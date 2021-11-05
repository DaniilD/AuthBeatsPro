include .env

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

.PHONY: script
script:
	@docker-compose exec app bash -c "$(exec)"

.PHONY: build
build: ## сборка образов и приложения
	@echo "================================================================================"
	@printf "\033[36mПересборка образов\033[0m\n"
	@if [ ! -f .env ] ; then  printf "\033[31m.env file not found\033[0m\n" ; exit 1 ; fi
	@DOCKER_BUILDKIT=1 docker build \
		--pull \
		--build-arg=UID=${DOCKER_UID} \
		--build-arg=GID=${DOCKER_GID} \
		-t "${SERVICE_NAME}-${LOCAL_BASE_DOMAIN_SLUG}" \
		.
	@echo "================================================================================"
	@printf "\033[36mЗапуск контейнеров\033[0m\n"
	@docker-compose up -d
	@echo "================================================================================"
	@printf "\033[36mСборка приложения\033[0m\n"
	@$(MAKE) script exec="./scripts/build.sh"
	@$(MAKE) migrate

.PHONY: migrate
migrate: ## накатывание миграций
	@echo "================================================================================"
	@printf "\033[36mВыполнение миграций\033[0m\n"
	@@$(MAKE) script exec="./scripts/migrate.sh"


#migrate -path migrations -database "mysql://root:root@tcp(localhost:3306)/test" -verbose up
#migrate create -ext sql -dir migrations <name_table>