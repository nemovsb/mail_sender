OS=$(shell uname -o)
PROJECTNAME=$(shell basename "$(PWD)")

# COLORED OUTPUT
GREEN=
LGREEN=
YELLOW=
ORANGE=
NC=# No Color

ifeq (${OS}, GNU/Linux)
GREEN=\033[1;32m
LGREEN=\033[0;32m
YELLOW=\033[1;33m
ORANGE=\033[0;33m
NC=\033[0m # No Color
endif


.PHONY: build test  

## build: build project
build:
	@echo
	go build
	@echo ">${GREEN} complete${NC}"
	@echo

## test: run unit tests
test:
	@echo ">${YELLOW} Running unit tests...${NC}"
	go test -v ./...
	@echo ">${GREEN} All tests passed${NC}"

## build for docker-compose
GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

ifeq ($(OS), Msys)
GOOS=windows
endif

BUILD_DIR=build/package
BUILDVARS=GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED}
DOCKER_BUILDVARS=GOOS=linux GOARCH=${GOARCH} CGO_ENABLED=${CGO_ENABLED}

build-for-compose:
	@echo
	${DOCKER_BUILDVARS} go build -o mail_sender -ldflags "-s -w"
	@echo ">${GREEN} complete${NC}"
	@echo


COMPOSE_TEST_FILE=docker-compose.yml
COMPOSE_TEST_CMD=docker-compose --project-name dev_${PROJECTNAME} --file ${COMPOSE_TEST_FILE}
COMPOSE_TEST_PULL_CMD=${COMPOSE_TEST_CMD} pull

.PHONY: compose-test-up
compose-test-up: build-for-compose
	@echo ">${YELLOW} Raise the whole project from docker-compose.yml...${NC}"
	${COMPOSE_TEST_PULL_CMD}
	${COMPOSE_TEST_CMD} up --build --detach
	@echo ">${GREEN} Project raised${NC}"

## compose-test-down: destroy everything raised from docker-compose.yml
.PHONY: compose-test-down
compose-test-down:
	@echo ">${YELLOW} Destroying everything raised from docker-compose.yml...${NC}"
	${COMPOSE_TEST_CMD} down --remove-orphans
	@echo ">${GREEN} Everything destroyed${NC}"

## compose-test-down-with-volumes: destroy everything raised from docker-compose.yml with volumes
.PHONY: compose-test-down-with-volumes
compose-test-down-with-volumes:
	@echo ">${YELLOW} Destroying everything raised from docker-compose.yml with volumes...${NC}"
	${COMPOSE_TEST_CMD} down --remove-orphans --volumes
	@echo ">${GREEN} Everything destroyed${NC}"


## go-mod-verify: clean and verify go modules
go-mod-verify:
	@echo ">${YELLOW} Fixing modules...${NC}"
	@echo "  >${ORANGE} Adding missing and removing unused modules...${NC}"
	go mod tidy
	@echo "  >${ORANGE} Verifying dependencies have expected content...${NC}"
	go mod verify
	@echo ">${GREEN} Modules fixed${NC}"



