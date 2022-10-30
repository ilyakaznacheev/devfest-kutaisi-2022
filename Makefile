build:
	@docker build . -t gdg-kutaisi

run:
	@docker run --rm -p 8888:8080 \
	-e API_ADDRESS=0.0.0.0:8080 \
	-e DB_ADDRESS=host.docker.internal:6379 \
	-e DB_COLLECTION=wines \
	gdg-kutaisi

start: build run

up:
	@docker compose up

down:
	@docker compose down
	@docker compose rm