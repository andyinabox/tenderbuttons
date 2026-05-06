.PHONY: watch
watch:
	reflex -d none -G "assets/*" -s make run-dev

.PHONY: run-dev
run-dev:
	go run . -v -a

.PHONY: run-https
run-https: .cert/localhost.crt
	go run . -v -m https-ss

IMAGE := code.andydayton.com/andy/tenderbuttons

.PHONY: docker-build
docker-build:
	docker build -t $(IMAGE):dev .

.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 $(IMAGE):dev

.PHONY: release
release:
	@test -n "$(VERSION)" || (echo "usage: make release VERSION=v1.0.0" && exit 1)
	./sh/release.sh $(VERSION)

.PHONY: stats
stats:
	go run ./cmd/stats/main.go 3 35 1000


# non-phonies

.cert/localhost.crt:
	./script/cert.sh