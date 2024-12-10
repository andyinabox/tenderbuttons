.PHONY: watch
watch:
	reflex -d none -s make run

.PHONY: run
run:
	go run . -v

.PHONY: docker-build
docker-build:
	docker build -t andyinabox/tenderbuttons .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 andyinabox/tenderbuttons

.PHONY: docker-push
docker-push:
	docker push andyinabox/tenderbuttons