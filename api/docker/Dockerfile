FROM golang:1.16 as build-a-bot-api
ENV CGO_ENABLED 0
ARG VCS_REF

# Copy the source code into the container
COPY . /build-a-bot 

WORKDIR /build-a-bot/api
RUN go build -ldflags "-X main.build=${VCS_REF}"

# Run the Go Binary in Alpine.
FROM alpine:3.13
ARG BUILD_DATE
ARG VCS_REF
COPY --from=build-a-bot-api /build-a-bot/api/. /service/build-a-bot-api/.
WORKDIR /service/build-a-bot-api
CMD ["./build-a-bot-api"]

LABEL org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.title="build-a-bot-api" \
    org.opencontainers.image.authors="Sean Amoroso <amopitt@gmail.com>" \
    org.opencontainers.image.source="https://github.com/amopitt/sean-build-a-bot/api" \
    org.opencontainers.image.revision="${VCS_REF}" \
    org.opencontainers.image.vendor="AmoCoder"