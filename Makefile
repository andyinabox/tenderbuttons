.PHONY: watch
watch:
	reflex -d none -s make run

.PHONY: run
run:
	go run . -v

.PHONY: run-https
run-https: .cert/localhost.crt
	go run . -v -m https-ss

.PHONY: docker-build
docker-build:
	docker build -t andyinabox/tenderbuttons .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 andyinabox/tenderbuttons

.PHONY: docker-push
docker-push:
	docker push andyinabox/tenderbuttons

# non-phonies

.cert/localhost.crt:
	./script/cert.sh