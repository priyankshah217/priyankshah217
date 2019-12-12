package services

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"github.com/priyankshah217/model"
	"github.com/priyankshah217/producer-api/repository"
)

var dir, _ = os.Getwd()
var pactDir = fmt.Sprintf(os.Getenv("GOPATH")+"/src/github.com/priyankshah217/pact/", dir)
var logDir = fmt.Sprintf("%s/log", dir)
var port, _ = utils.GetFreePort()

func TestPactProvider(t *testing.T) {
	go startServer()
	pact := createPact()
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:    fmt.Sprintf("http://127.0.0.1:%d", port),
		Tags:               []string{"master"},
		FailIfNoPactsFound: false,
		// PactURLs: []string{filepath.FromSlash(fmt.Sprintf("%s/pact_consumer-pact_provider.json",
		// 	os.Getenv("GOPATH")+"/src/github.com/priyankshah217/pact/"))},
		BrokerURL:       fmt.Sprintf("%s://%s", os.Getenv("PACT_BROKER_PROTO"), os.Getenv("PACT_BROKER_URL")),
		BrokerUsername:  os.Getenv("PACT_BROKER_USERNAME"),
		BrokerPassword:  os.Getenv("PACT_BROKER_PASSWORD"),
		ProviderVersion: "1.0.0",
		StateHandlers:   stateHandlers,
		RequestFilter:   fixToken,
		// BeforeEach: func() error {
		// 	userRepository = sallyExists
		// 	return nil
		// },
	})
	if err != nil {
		t.Log("Pact test failed")
	}
}

var stateHandlers = types.StateHandlers{
	"User sally exists": func() error {
		userRepository = sallyExists
		return nil
	},
	"User sally does not exists": func() error {
		userRepository = sallyDoesNotExist
		return nil
	},
	"User is not authenticated": func() error {
		userRepository = sallyUnauthorized
		return nil
	},
}

func fixToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			r.Header.Set("Authorization", getAuthToken())
		}
		next.ServeHTTP(w, r)
	})
}

func startServer() {
	mux := GetHttpHandler()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Printf("API starting: port %d (%s)", port, ln.Addr())
	log.Printf("API terminating: %v", http.Serve(ln, mux))
}

// Provider States data sets
var sallyExists = &repository.UserRepository{
	Users: map[string]*model.User{
		"sally": {
			FirstName: "Jean-Marie",
			LastName:  "de La Beaujardière😀😍",
			Username:  "sally",
			Type:      "admin",
			ID:        10,
		},
	},
}

var sallyDoesNotExist = &repository.UserRepository{}

var sallyUnauthorized = &repository.UserRepository{
	Users: map[string]*model.User{
		"sally": {
			FirstName: "Jean-Marie",
			LastName:  "de La Beaujardière😀😍",
			Username:  "sally",
			Type:      "blocked",
			ID:        10,
		},
	},
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Provider:                 "PACT_PROVIDER",
		LogDir:                   logDir,
		PactDir:                  pactDir,
		DisableToolValidityCheck: true,
		LogLevel:                 "INFO",
	}
}
