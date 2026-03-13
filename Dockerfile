FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
ARG APM_TAG=""
RUN CGO_ENABLED=0 go build -tags "${APM_TAG}" -ldflags="-s -w -X main.version=${VERSION}" -o /express-botx .

FROM alpine:3.21
RUN apk add --no-cache ca-certificates
COPY --from=build /express-botx /usr/local/bin/express-botx
ENTRYPOINT ["express-botx"]
