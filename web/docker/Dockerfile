# Name the node stage "builder"
FROM node:15 AS build-a-bot-web
ENV CGO_ENABLED 0

# Copy the source code into the container
COPY ./web /build-a-bot 

# Set working directory
WORKDIR /build-a-bot 

# Deleting these if they are grabbed from my local
RUN rm -rf node_modules

# install node modules and build assets
RUN npm install 
RUN npm run build

# nginx state for serving content
FROM nginx:latest

# Set working directory to nginx asset directory
WORKDIR /usr/share/nginx/html

# Remove default nginx static assets
RUN rm -rf ./*

# Copy static assets from web stage
COPY --from=build-a-bot-web /build-a-bot/dist/. .

# Containers run nginx with global directives and daemon off
#ENTRYPOINT ["nginx", "-g", "daemon off;"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.title="build-a-bot-web" \
    org.opencontainers.image.authors="Sean Amoroso <amopitt@gmail.com>" \
    org.opencontainers.image.source="https://github.com/amopitt/sean-build-a-bot/web" \
    org.opencontainers.image.revision="${VCS_REF}" \
    org.opencontainers.image.vendor="AmoCoder"