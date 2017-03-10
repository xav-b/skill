# Go.mk
# vim:ft=make

PLATFORM ?= "windows linux darwin"

all: $(BINARY)

container-up: ## start a pending container workspace
	docker run -d --name $(PROJECT) \
		-v $(PWD):$(GO_PROJECTS)/$(PROJECT) \
		-w $(GO_PROJECTS)/$(PROJECT) $(CONTAINER) sleep infinity

container-shell: ## start a shell in the workspace container
	docker exec -it $(PROJECT) bash

# TODO simple make build
# TODO only build on changes
crossbuild: $(SOURCES) container
	docker run -it --rm \
  	-v $(PWD):$(GO_PROJECTS)/$(PROJECT) \
  	-w $(GO_PROJECTS)/$(PROJECT) $(PROJECT) wgo cross

# usage:
# PLATFORM=darwin \
# COMMENT="Buggy MVP completed" \
# RELEASE_OPTS="-prerelease -b 'Buggy MVP completed'" \
# make release
release: crossbuild ## git tag and publish a release on Github
ifndef COMMENT
	$(error no tag description provided)
endif
	git tag -a $(VERSION) -m '$(COMMENT)'
	git push --tags
	ghr $(RELEASE_OPTS) v$(VERSION) $(BUILD_PATH)/$(VERSION)/

$(BINARY): $(SOURCES) ## compile project
	go build -v -ldflags ${LDFLAGS} -o ${BINARY}

.PHONY: install.tools
install-tools: ## install development tools
	# code coverage
	go get github.com/axw/gocov/gocov
	# cross-compilation
	go get github.com/mitchellh/gox
	# github release publication
	go get github.com/tcnksm/ghr
	# code linting
	# FIXME make circleci to fail
	go get github.com/alecthomas/gometalinter && \
		gometalinter --install --update

install-hack: install.tools ## install dev tools and project deps
	go get ./...

install: ## compile and globally install the cli
	go install -ldflags ${LDFLAGS}

lint:
	docker exec -it $(PROJECT) gometalinter --exclude=vendor ./...

tests: lint
	docker exec -it $(PROJECT) go test $(TESTARGS)

.PHONY: godoc
godoc: ## run Go doc server
	godoc -http=0.0.0.0:6060

# FIXME need sudo because of docker
.PHONY: clean
clean: ## remove build artifacts
	@test -d $(BUILD_PATH) && rm -rf $(BUILD_PATH)
	@test -f $(BINARY) && rm -rf $(BINARY)
