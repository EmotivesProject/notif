SLEEP_TIME=0

.PHONY: lint test integration

lint:
	golangci-lint run

test:
	go test -v ./...

integration:
	sleep $(SLEEP_TIME)
	go test -v -tags=integration ./...