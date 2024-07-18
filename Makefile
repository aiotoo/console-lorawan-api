.PHONY: go swagger js

all:
	docker-compose up

go:
	docker-compose run --rm console-lorawan-api-go

swagger:
	docker-compose run --rm console-lorawan-api-swagger
