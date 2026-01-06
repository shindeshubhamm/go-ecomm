.PHONY: db migrate

db:
	. ./.env && docker compose up postgres -d


migrate:
	. ./.env && goose up