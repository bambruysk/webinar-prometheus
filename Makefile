.PHONY: build_blog
build_blog:
	docker build -t blogpost ./blogpost

.PHONY: run
run:
	docker compose -f ./prom/docker-compose.yml -p prom up --force-recreate -d --build
