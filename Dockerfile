# syntax=docker/dockerfile:1
# Dev

# Step 0
FROM --platform=$BUILDPLATFORM tonistiigi/xx AS xx

# Step 1 - Build
FROM --platform=$BUILDPLATFORM golang:alpine AS build

COPY --from=xx / /

ARG TARGETPLATFORM
ENV CGO_ENABLED 0

WORKDIR /api
COPY go.mod .
COPY go.sum .
RUN xx-go mod download

COPY . .

RUN xx-go build -o bering-travel cmd/main.go

# Step 2 - prepare
FROM alpine:latest

WORKDIR /app

COPY /external ./external
COPY .env.dev .
COPY --from=build /api/bering-travel .

EXPOSE 7772

# Config
ENV PROJECT_ABS_PATH "/app"

ENTRYPOINT ["./bering-travel", "-mode=dev"]
