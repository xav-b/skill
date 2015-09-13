# Makefile
# vim:ft=make

PROJECT=`basename $(PWD)`
GIT_USER=`git config --get user.name`
GO_PROJECTS="/go/src/github.com/$(GIT_USER)"

all: build

build:
	docker build --rm -t $(PROJECT) .
	docker run -it --rm \
  	-v $(PWD):$(GO_PROJECTS)/$(PROJECT) \
  	-w $(GO_PROJECTS)/$(PROJECT) $(PROJECT) wgo cross

install: build
	# TODO detect platform
	cp ./bin/skill-darwin-amd64 /usr/local/bin

lint:
	docker exec -it $(PROJECT) gometalinter --exclude=vendor ./...

clean:
	docker stop $(PROJECT) && docker rm -v $(PROJECT)
	rm -r ./bin
