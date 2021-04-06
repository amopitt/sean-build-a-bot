# Used by `image`, `push` & `deploy` targets, override as required
IMAGE_REG ?= localhost:32000
IMAGE_REPO ?= build-a-bot-web
IMAGE_TAG ?= latest

image_web: ## 🔨 Build container image from Dockerfile 
	docker build . --no-cache --file web/docker/Dockerfile \
	--tag $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

image_api: ## 🔨 Build container image from Dockerfile 
	docker build . --no-cache --file api/docker/Dockerfile \
	--tag $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

push:  ## 📤 Push container image to registry 
	docker push $(IMAGE_REG)/$(IMAGE_REPO):$(IMAGE_TAG)

## k8s
run_api:
	kubectl apply -f k8s/backend/file.configmap.yml
	kubectl apply -f k8s/backend/file.deployment.yml
	kubectl apply -f k8s/backend/file.service.nodeport.yml

run_web:
	kubectl apply -f k8s/frontend/file.configmap.yml
	kubectl apply -f k8s/frontend/file.deployment.yml
	kubectl apply -f k8s/frontend/file.service.nodeport.yml

teardown_web:
	kubectl delete service frontend
	kubectl delete deployment frontend
	kubectl delete configmap frontend-app-settings

teardown_api:
	kubectl delete service backend
	kubectl delete deployment backend
	kubectl delete configmap backend-app-settings