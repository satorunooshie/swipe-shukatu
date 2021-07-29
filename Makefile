
ENV_DEV = $(shell cat .env.dev)
ROOT_PASS = Password!

.PHONY: up
up:
	docker-compose up --build -d

.PHONY: down
down:
	docker-compose down

.PHONY: ps
ps:
	docker-compose ps

.PHONY: test
test:
	docker-compose exec api richgo test ./... -v -count=1 -cover

.PHONY: gen
gen:
	go generate ./...

.PHONY: lint
lint:
	docker-compose -f docker-compose.dev.yml exec api golangci-lint run -v

.PHONY: test-lint
test-lint: test lint

# .PHONY: schema-diff
# schema-diff:
# 	docker-compose -f docker-compose.dev.yml exec api skeema diff dev -p$(ROOT_PASS) --allow-unsafe
#
# .PHONY: schema-migrate
# schema-migrate:
# 	docker-compose -f docker-compose.dev.yml exec api sh ./schema.sh $(ROOT_PASS)

.PHONY: deps
deps:
	cat tools.go | awk -F'"' '/_/ {print $$2}' | xargs -tI {} go install {}
