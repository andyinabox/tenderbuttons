.PHONY: run
run:
	go run .

.PHONY: watch
watch:
	reflex -d fancy -s make run

.PHONY: docker-build
docker-build:
	docker build -t tenderbuttons .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 tenderbuttons