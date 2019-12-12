package client

import (
	"fmt"
	"os"
	"testing"

	"net/http"
	"net/url"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/priyankshah217/model"
	"github.com/stretchr/testify/assert"
)

var commonHeaders = dsl.MapMatcher{
	"Content-Type":         term("application/json; charset=utf-8", `application\/json`),
	"X-Api-Correlation-Id": dsl.Like("100"),
}

var headersWithToken = dsl.MapMatcher{
	"Authorization": dsl.Like("Bearer 2019-01-01"),
}

var u *url.URL
var client *Client

func TestMain(m *testing.M) {
	var exitCode int

	// Setup Pact and related test stuff
	setup()

	// Run all the tests
	exitCode = m.Run()

	// Shutdown the Mock Service and Write pact files to disk
	if err := pact.WritePact(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pact.Teardown()
	os.Exit(exitCode)
}

func TestClientPact_GetUser(t *testing.T) {
	t.Run("the user exists", func(t *testing.T) {
		id := 10

		pact.
			AddInteraction().
			Given("User sally exists").
			UponReceiving("A request to login with user 'sally'").
			WithRequest(request{
				Method:  "GET",
				Path:    term("/user/10", "/user/[0-9]+"),
				Headers: headersWithToken,
			}).
			WillRespondWith(dsl.Response{
				Status:  http.StatusOK,
				Body:    dsl.Match(model.User{}),
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			user, err := client.WithToken("2019-01-01").GetUser(id)

			// Assert basic fact
			if user.ID != id {
				return fmt.Errorf("wanted user with ID %d but got %d", id, user.ID)
			}

			return err
		})

		if err != nil {
			t.Fatalf("Error on Verify: %v", err)
		}
	})
	t.Run("the user does not exist", func(t *testing.T) {

		pact.
			AddInteraction().
			Given("User sally does not exists").
			UponReceiving("A request to login with user 'sally'").
			WithRequest(request{
				Method:  "GET",
				Path:    term("/user/10", "/user/[0-9]+"),
				Headers: headersWithToken,
			}).
			WillRespondWith(dsl.Response{
				Status:  http.StatusNotFound,
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			_, err := client.WithToken("2019-01-01").GetUser(10)
			return err
		})
		assert.Equal(t, model.ErrNotFound, err)
	})
	t.Run("not authenticated user", func(t *testing.T) {
		pact.
			AddInteraction().
			Given("User is not authenticated").
			UponReceiving("A request to login with user 'sally'").
			WithRequest(request{
				Method: "GET",
				Path:   term("/user/10", "/user/[0-9]+"),
			}).
			WillRespondWith(dsl.Response{
				Status:  http.StatusUnauthorized,
				Headers: commonHeaders,
			})

		err := pact.Verify(func() error {
			_, err := client.WithToken("").GetUser(10)
			return err
		})

		assert.Equal(t, model.ErrUnauthorized, err)
	})

}

// Common test data
var pact dsl.Pact

// Aliases
var term = dsl.Term

type request = dsl.Request

func setup() {
	pact = createPact()

	// Proactively start service to get access to the port
	pact.Setup(true)

	u, _ = url.Parse(fmt.Sprintf("http://localhost:%d", pact.Server.Port))

	client = &Client{
		BaseURL: u,
	}

}

func createPact() dsl.Pact {
	return dsl.Pact{
		Consumer:                 "PACT_CONSUMER",
		Provider:                 "PACT_PROVIDER",
		LogDir:                   "./log/",
		PactDir:                  os.Getenv("GOPATH") + "/src/github.com/priyankshah217/pact",
		LogLevel:                 "INFO",
		DisableToolValidityCheck: true,
	}
}
