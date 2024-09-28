generate-api:
	oapi-codegen -generate types,server,spec -package oapi -o api/oapi/oapi.gen.go docs/api/openapi.yaml

run-api:
	cd api && docker compose up

stop-api:
	cd api && docker compose down