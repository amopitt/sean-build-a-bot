# Used by `image`, `push` & `deploy` targets, override as required
IMAGE_REG ?= localhost:32000
IMAGE_REPO ?= build-a-bot-web
IMAGE_TAG ?= latest

image: ## 🔨 Build container image from Dockerfile 
	docker build . --file web/docker/Dockerfile \
	--tag $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

push:  ## 📤 Push container image to registry 
	docker push $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

## k8s
run_web:
	kubectl apply -f k8s/frontend/file.deployment.yml
	kubectl apply -f k8s/frontend/file.service.nodeport.yml

teardown_web:
	kubectl delete service frontend
	kubectl delete deployment frontend