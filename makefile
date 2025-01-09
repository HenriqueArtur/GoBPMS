.PHONY: db-start db-stop test

test: db-start
	go test ./src/... -v

watch-test:
	find ./src -type f -name '*.go' | entr -r make test

db-start:
	docker compose up -d

db-stop:
	docker compose down
