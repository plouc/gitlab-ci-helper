.PHONY: install update test coverage_backend run format build build_checksums

SHA1     = $(shell git rev-parse HEAD)
GO_PKG   = ./,./commands,./integrations/flowdock,./integrations/hipchat
GO_FILES = $(shell find $(GO_PROJECTS_PATHS) -maxdepth 1 -type f -name "*.go")
OS       = $(shell uname)

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  HELP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

.DEFAULT: help

# COLORS
RED    = $(shell printf "\33[31m")
GREEN  = $(shell printf "\33[32m")
WHITE  = $(shell printf "\33[37m")
YELLOW = $(shell printf "\33[33m")
RESET  = $(shell printf "\33[0m")

# Add the following 'help' target to your Makefile
# And add help text after each target name starting with '\#\#'
# A category can be added with @category
HELP_SCRIPT = \
    %help; \
    while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-\%_]+)\s*:.*\#\#(?:@([a-zA-Z\-\%]+))?\s(.*)$$/ }; \
    print "usage: make [target]\n\n"; \
    for (sort keys %help) { \
    print "${WHITE}$$_:${RESET}\n"; \
    for (@{$$help{$$_}}) { \
    $$sep = " " x (32 - length $$_->[0]); \
    print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
    }; \
    print "\n"; }

help: ##prints help
	@perl -e '${HELP_SCRIPT}' ${MAKEFILE_LIST}

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  SETUP
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

install: ##@setup install dependencies
	@echo "${YELLOW}Installing dependencies${RESET}"
	go list -f '{{range .Imports}}{{.}} {{end}}' ./... | xargs go get -v ${install_args}
	go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | xargs go get -v ${install_args}
	go get -v ${install_args} github.com/wadey/gocovmerge
	@echo "${GREEN}✔ successfully installed dependencies${RESET}\n"

update: ##@setup update dependencies
	@${MAKE} install install_args=-u

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  TEST
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

test: ##@test run tests and cs tools
	@echo "${YELLOW}Running tests${RESET}"
	go test -v ./cli/...
	go vet ./cli/...
	gofmt -l -s -e ./cli
	@exit `gofmt -l -s -e ./cli | wc -l`
	@echo "${GREEN}✔ tests successfully passed${RESET}\n"

coverage_backend: ##@test run coverage tests
	mkdir -p build/coverage
	rm -rf build/coverage/*.cov
	go test -v -timeout 60s -coverpkg $(GO_PKG) -covermode count -coverprofile=build/coverage/main.cov ./
	go test -v -timeout 60s -coverpkg $(GO_PKG) -covermode count -coverprofile=build/coverage/commands.cov ./commands
	go test -v -timeout 60s -coverpkg $(GO_PKG) -covermode count -coverprofile=build/coverage/integration_flowdock.cov ./integrations/flowdock
	go test -v -timeout 60s -coverpkg $(GO_PKG) -covermode count -coverprofile=build/coverage/integration_hipchat.cov ./integrations/hipchat
	gocovmerge build/coverage/* > build/gitlabcihelper.coverage
	go tool cover -html=./build/gitlabcihelper.coverage -o build/gitlabcihelper.html

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  MISC
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

run: ##@misc run the command
	go run cli/main.go

format: ##@misc format the code and generate commands.md file
	@echo "${YELLOW}Formatting code${RESET}"
	gofmt -l -w -s .
	go fix ./...
	#go run cli/main.go dump:readme > commands.md
	@echo "${GREEN}✔ code was formatted as expected${RESET}\n"

#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-
#
#  BUILD
#
#=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-

build: ##@build build binaries
	@echo "${YELLOW}Buiding CLI build binaries${RESET}"
	GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.RefLog=$(SHA1) -s -w" -o build/darwin-amd64-gitlab-ci-helper cli/main.go
	GOOS=darwin GOARCH=386   go build -ldflags "-X main.RefLog=$(SHA1) -s -w" -o build/darwin-386-gitlab-ci-helper   cli/main.go
	GOOS=linux  GOARCH=amd64 go build -ldflags "-X main.RefLog=$(SHA1) -s -w" -o build/linux-amd64-gitlab-ci-helper  cli/main.go
	GOOS=linux  GOARCH=386   go build -ldflags "-X main.RefLog=$(SHA1) -s -w" -o build/linux-386-gitlab-ci-helper    cli/main.go
	GOOS=linux  GOARCH=arm   go build -ldflags "-X main.RefLog=$(SHA1) -s -w" -o build/linux-arm-gitlab-ci-helper    cli/main.go
	GOOS=linux  GOARCH=arm64 go build -ldflags "-X main.RefLog=$(SHA1) -s -w" -o build/linux-arm64-gitlab-ci-helper  cli/main.go
    #ifneq ("", "$(shell which docker)")
	#	docker run --rm -v $(shell pwd):/usr/src/myapp -v $(GOPATH):/usr/src/myapp/vendor -w /usr/src/myapp -e "GOPATH=/usr/src/myapp/vendor:/go" -e GOOS=linux -e GOARCH=amd64 golang:1.9-alpine go build -ldflags "-X main.RefLog=$(SHA1) -s -w" -o build/alpine-amd64-gitlab-ci-helper cli/main.go
    #endif
	@echo "${GREEN}✔ successfully built CLI binaries${RESET}\n"

build_checksums: ##@build generate checksums for CLI binaries
	@echo "${YELLOW}Generating CLI build checksums${RESET}"
	@rm -f build/checksums.txt

    # for OSX users where you have md5 instead of md5sum
    ifeq (${OS}, Darwin)
        # md5 output has the following format:
        #
        # MD5 (darwin-amd64-glc) = 8eb317789e5d08e1c800cc469c20325a
        #
        # that's why sed and awk are used to cleanup
		@cd build && ls . | grep gitlab-ci-helper \
            | xargs md5 \
            | awk '{ printf("%s\n%s\n\n", $$2, $$4) }' \
            | sed 's/[()]//g' \
            >> checksums.txt
    else
        # md5sum output has the following format:
        #
        # 8eb317789e5d08e1c800cc469c20325a darwin-amd64-glc
        #
        # that's why awk is used to cleanup
		@cd build && ls . | grep gitlab-ci-helper \
            | xargs md5sum \
            | awk '{ printf("%s\n%s\n\n", $$2, $$1) }' \
            >> checksums.txt
    endif
	@echo "${GREEN}✔ successfully generated CLI build checksums to ${WHITE}build/checksums.txt${RESET}\n"
