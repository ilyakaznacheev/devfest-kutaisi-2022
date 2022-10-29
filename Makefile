build:
	docker build . -t gdg-kutaisi

run:
	docker run --rm -p 8888:8080 gdg-kutaisi