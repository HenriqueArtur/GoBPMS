.PHONY: db-start db-stop test watch-test test-ci

test: db-start
	go test ./src/... -v

watch-test:
	find ./src -type f -name '*.go' | entr -r make test

test-ci:
	go test ./src/... -v

db-start:
	docker compose up -d

db-stop:
	docker compose down
