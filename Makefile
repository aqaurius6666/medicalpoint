dev-recreate:
	@docker-compose --project-name=medicalpoint-dev --env-file deploy/dev/.env -f deploy/dev/docker-compose.yaml up -d --build --force-recreate

#dev-migration-up:
#	NETWORK_NAME=medical-chain-server-dev docker-compose --project-name=medical-chain-dev -f deploy/dev/docker-compose-migration-tool.yaml up up
#
#dev-migration-down:
#	NETWORK_NAME=medical-chain-server-dev docker-compose --project-name=medical-chain-dev -f deploy/dev/docker-compose-migration-tool.yaml up down
build-and-push-image: build-image push-image

build-image:
	@docker build . --target=release -t aqaurius6666/medicalpoint-gateway:pre-release

push-image:
	@docker tag aqaurius6666/medicalpoint-gateway:pre-release aqaurius6666/medicalpoint-gateway${TAG}
	@docker push aqaurius6666/medicalpoint-gateway${TAG}

build:
	@go build -o ./dist/server ./src

debug:
	@/bin/sh -c "dlv --listen=0.0.0.0:2345 --headless=true --api-version=2 exec ./dist/server -- --disable-profiler --allow-kill serve"

serve:
	@./dist/server serve

dev:
	@./dist/server --log-format plain --log-level debug --disable-profiler --allow-kill serve

cleanDB:
	@./dist/server clean

seed:
	@./dist/server seed-data --clean
	@echo Hello

test:
	#go test ./src/cockroach/... -v -check.f "CockroachDbGraphTestSuite.*"
	@go test ./... -v

test-prepare-up:
	@docker exec  up -f deploy/dev/docker-compose.yaml main-cdb -d

test-prepare-down:
	@docker-compose down -f deploy/dev/docker-compose.yaml main-cdb

grpc-client:
	@grpc-client-cli localhost:${GRPC_PORT}

kill:
	@(echo '{}' | grpc-client-cli -service CommonService -method Kill localhost:${GRPC_PORT}) > /nil 2> /nil || return 0

logs:
	@docker logs -f medicalpoint-dev_gateway_1

proto2:
	@./scripts/genproto.sh

swagger: proto2
	@go generate ./src/services/swagger github.com/medicalpoint/gateway/src/services/swagger 

rebase: 
	@git pull --rebase origin dev

push:
	@git push origin HEAD:automatic-branch -f
