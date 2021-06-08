# Api

Bump it up sir

## Docker Image

### run from sean-build-a-bot root folder

```
docker build -f api/docker/Dockerfile -t build-a-bot-api:1.0 --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`  --build-arg VCS_REF=`git rev-parse HEAD` .


docker build -f web/docker/Dockerfile -t build-a-bot-web:1.0 --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`  --build-arg VCS_REF=`git rev-parse HEAD` .
```

### compose up

```
docker-compose -f docker/compose.yaml up
```

## Makefile

Make file variables and default values, pass these in when calling make, e.g. make image IMAGE_REPO=blah/foo

### Build images and push them to microk8s local registry and run both containers and pods

```
# backend api
sudo make image_api IMAGE_REPO=build-a-bot-api
sudo make push IMAGE_REPO=build-a-bot-api
sudo make run_api
# remove
sudo make teardown_api

# frontend web (IMAGE_REPO defaults to build-a-bot-web)
sudo make image_web IMAGE_REPO=build-a-bot-web
sudo make push IMAGE_REPO=build-a-bot-web
sudo make run_web
# remove frontend web
sudo make teardown_web

```

## Logs

```
sudo kubectl logs deployment/backend
sudo kubectl logs deployment/frontend

```
