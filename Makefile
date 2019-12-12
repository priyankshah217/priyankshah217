TEST?=./...

include ./make/config.mk

install:
	@if [ ! -d pact/bin ]; then\
		echo "--- 🛠 Installing Pact CLI dependencies";\
		curl -fsSL https://raw.githubusercontent.com/pact-foundation/pact-ruby-standalone/master/install.sh | bash;\
    fi

run-consumer:
	@go run consumer/client/cmd/main.go

run-provider:
	@go run provider/cmd/usersvc/main.go

deploy-consumer: install
	@echo "--- ✅ Checking if we can deploy consumer"
	@pact-broker can-i-deploy \
		--pacticipant $(CONSUMER_NAME) \
		--broker-base-url ${PACT_BROKER_PROTO}://$(PACT_BROKER_URL) \
		--broker-username $(PACT_BROKER_USERNAME) \
		--broker-password $(PACT_BROKER_PASSWORD) \
		--latest

deploy-provider: install
	@echo "--- ✅ Checking if we can deploy provider"
	@pact-broker can-i-deploy \
		--pacticipant $(PROVIDER_NAME) \
		--broker-base-url ${PACT_BROKER_PROTO}://$(PACT_BROKER_URL) \
		--broker-username $(PACT_BROKER_USERNAME) \
		--broker-password $(PACT_BROKER_PASSWORD) \
		--latest

publish: install
	@echo "--- 📝 Publishing Pacts"
	go run consumer-api/pact/publish.go
	@echo
	@echo "Pact contract publishing complete!"
	@echo
	@echo "Head over to $(PACT_BROKER_PROTO)://$(PACT_BROKER_URL) and login with"
	@echo "=> Username: $(PACT_BROKER_USERNAME)"
	@echo "=> Password: $(PACT_BROKER_PASSWORD)"
	@echo "to see your published contracts.	"


unit:
	@echo "--- 🔨Running Unit tests "
	go test -tags=unit -count=1 github.com/priyankshah217/consumer-api/client -run 'TestClientUnit'

consumer: install
	@echo "--- 🔨Running Consumer Pact tests "
	go test -tags=integration -count=1 github.com/priyankshah217/consumer-api/client -run 'TestClientPact'

provider: install
	@echo "--- 🔨Running Provider Pact tests "
	go test -count=1 -tags=integration github.com/priyankshah217/producer-api/services -run "TestPactProvider"

.PHONY: install unit consumer provider run-provider run-consumer