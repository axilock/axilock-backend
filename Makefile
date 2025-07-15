run: build
	./bin/axilock-backend

build:
	@go build -ldflags="-s -w" -o ./bin/axilock-backend .

sqlc:
	@sqlc generate

createdb:
	docker exec -it postgres:17 createdb -u root axilockdb

dropdb:
	docker exec -it postgres:17 drobdb axilockdb

migratecreate:
	migrate create -ext sql -dir migrations -seq ${NAME}

remotemigratedown:
	migrate -path migrations -database "postgresql://admin:secure@localhost:5431/axilockdb?sslmode=disable" down 1

remotemigrateup:
	migrate -path migrations -database "postgresql://admin:secure@localhost:5431/axilockdb?sslmode=disable" up 1

migrateup:
	migrate -path migrations -database "postgresql://admin:secure@localhost:5432/axilockdb?sslmode=disable" up

migrateup1:
	migrate -path migrations -database "postgresql://admin:secure@localhost:5432/axilockdb?sslmode=disable" up 1

migratedown:
	migrate -path migrations -database "postgresql://admin:secure@localhost:5432/axilockdb?sslmode=disable" down

migratedown1:
	migrate -path migrations -database "postgresql://admin:secure@localhost:5432/axilockdb?sslmode=disable" down 1

app: proto
	docker compose up --build

update:
	git submodule update --recursive  --remote    

rebuild:
	git submodule update --recursive  --remote    
	docker compose up -d --build backend

mock:
	@go install github.com/vektra/mockery/v2@v2.53.0
	@mockery 

test:
	go test -v -cover ./...

lint:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.6
	@golangci-lint run --config .golangci.yaml

live:
	@go install github.com/air-verse/air@latest
	@air

evans:
	evans --host localhost -r repl -p 8090

proto:
	git submodule update --recursive  --remote

swagger:
	@protoc -I ./protos/proto --openapiv2_out ./swagger/ --openapiv2_opt allow_merge=true --openapiv2_opt merge_file_name=api  \
    	protos/proto/*.proto

swagger-ui:
	@swagger-cli bundle -o ./swagger/swagger.json -t json \
	 ./swagger/rpc_integrations.swagger.json ./swagger/rpc_userservice.swagger.json
	@docker run -p 80:8080 \
    -e SWAGGER_JSON=/swagger/swagger.json \
    -v ./swagger/:/swagger \
    swaggerapi/swagger-ui

send-to-air: proto
	 rsync -avz --exclude ".git" --exclude .env --exclude "bin" --exclude "tmp" --delete . axilock:/home/ubuntu/axilock
	 ssh axilock "cd /home/ubuntu/axilock && docker compose up -d --build backend"

dev:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/vektra/mockery/v2@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

.PHONY: run sqlc createdb drobdb migratedown migratedown live proto migratecreate evans test app rebuild lint swagger swagger-ui
