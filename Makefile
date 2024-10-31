IMAGE_NAME=petherin/spacetickets
IMAGE_TAG=latest

start: ##@App Start app in Docker
	docker compose up -d

stop: ##@App Stop app in Docker
	docker compose down -v --remove-orphans
	docker compose rm -v -f -s

logs: ##@App Tails logs from app in Docker
	docker compose logs -f --tail 100 spacetickets-api

build: ##@App Builds the local Dockerfile for both arm64 and amd64 architectures
	docker buildx build --platform linux/amd64,linux/arm64 -t $(IMAGE_NAME):$(IMAGE_TAG) .

push: ##@App Push image to Docker, provided you're logged in as the author
	docker push $(IMAGE_NAME):$(IMAGE_TAG)

test: ##@App Run all unit tests
	go test -v ./...

db: ##@Database Opens terminal in database container
	docker exec -it spacetickets-db psql -U postgres -d example

tables: ##@Database List database tables
	docker exec -it spacetickets-db psql -U postgres -d example -c "\dt public.*"

desc-bookings: ##@Database Describe bookings database table
	docker exec -it spacetickets-db psql -U postgres -d example -c "\d bookings"

desc-launchpads: ##@Database Describe launchpads database table
	docker exec -it spacetickets-db psql -U postgres -d example -c "\d launchpads"

desc-destinations: ##@Database Describe destinations database table
	docker exec -it spacetickets-db psql -U postgres -d example -c "\d destinations"

# Color settings for the making the help information look pretty
 GREEN  := $(shell tput -Txterm setaf 2)
 WHITE  := $(shell tput -Txterm setaf 7)
 YELLOW := $(shell tput -Txterm setaf 3)
 RESET  := $(shell tput -Txterm sgr0)

 # Add the following 'help' target to your Makefile
 # And add help text after each target name starting with '\#\#'
 # A category can be added with @category
 # link: https://gist.github.com/prwhite/8168133#gistcomment-1727513
 HELP_FUN = \
 	%help; \
 	while(<>) { push @{$$help{$$2 // 'options'}}, [$$1, $$3] if /^([a-zA-Z\-]+)\s*:.*\#\#(?:@([a-zA-Z\-]+))?\s(.*)$$/ }; \
 	print "usage: make [target]\n\n"; \
 	for (sort keys %help) { \
 	print "${WHITE}$$_:${RESET}\n"; \
 	for (@{$$help{$$_}}) { \
 	$$sep = " " x (32 - length $$_->[0]); \
 	print "  ${YELLOW}$$_->[0]${RESET}$$sep${GREEN}$$_->[1]${RESET}\n"; \
 	}; \
 	print "\n"; }

help: ##@other Show this help.
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)

toc: ##@other Generate the Table of Contents in README.md
	docker run --platform=linux/amd64 -v $(PWD):/code -it node:lts-alpine3.17 /bin/sh -c 'npx --yes markdown-toc -i /code/README.md'

# Running just the `make` command will now print out the help information
# instead of printing the first command in the file
.DEFAULT_GOAL := help
.PHONY: db
