FROM powerman/dockerize:0.19.0 AS dockerize
FROM golang:1.23-alpine AS build

WORKDIR /src

COPY go.mod go.sum /src/

RUN go mod download

COPY . /src/

RUN go build -a -tags musl -o /app/server cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=dockerize /usr/local/bin/dockerize /usr/sbin/
COPY --from=build /app/ /app/
COPY --from=build /src/schema/ /app/schema/
COPY --from=build /src/docker/config_template.yml /app/etc/config/config_template.yml

CMD [ "dockerize", "-template", "/app/etc/config/config_template.yml:/app/etc/config/config.yml", "/app/server" ]