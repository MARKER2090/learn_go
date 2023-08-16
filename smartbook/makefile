.PHONY: AAA
docker:
	@rm smartbook || true
	@go build -tags=k8s -o smartbook .
	@docker rmi -f smartbook:v0.0.1
	@docker build -t smartbook:v0.0.1 .
