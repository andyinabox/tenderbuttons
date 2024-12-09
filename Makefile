.PHONY: run
run:
	go run .

.PHONY: watch
watch:
	reflex -d fancy -s make run
