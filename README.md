# Api

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
