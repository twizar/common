GOLANG_CI_LINT_IMAGE=golangci/golangci-lint:latest-alpine
GOLANG_IMAGE=golang:1.17.2

go-lint:
	docker run -v ${PWD}:/app -w /app $(GOLANG_CI_LINT_IMAGE) golangci-lint run -v --timeout 600m --fix

up-dev-env:
	docker-compose -f ./build/dev/docker-compose.yml up -d
down-dev-env:
	docker-compose -f ./build/dev/docker-compose.yml down

test-call:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{ "path":"/teams", "httpMethod":"GET" }'

go-test:
	docker run \
		--env-file=test.env \
		-v ${PWD}:/app \
		-w /app $(GOLANG_IMAGE) \
		go test -race -cover -v -coverpkg=./... -coverprofile=cover.out ./...
		go tool cover -html=cover.out

generate-mocks:
	~/go/bin/mockgen -source=./pkg/client/aws_lambda.go -destination=./test/mock/aws_lambda_mock.go -package=mock
	~/go/bin/mockgen -source=./pkg/client/team.go -destination=./test/mock/team_mock.go -package=mock
