.PHONY: run
run:
	reflex -d fancy -s go run .

.PHONY: docker-build
docker-build:
	docker build -t tenderbuttons .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 tenderbuttons