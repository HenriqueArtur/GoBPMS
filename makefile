.PHONY: db-start db-stop

db-start:
	docker compose up -d

db-stop:
	docker compose down
