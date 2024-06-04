FROM golang:1.22.1-alpine AS builder

WORKDIR /build

COPY go.mod .

COPY . .

RUN go build -o cron_service ./cmd/cron/gocron.go

FROM ubuntu:20.04

ENV APP_DIR /build

WORKDIR $APP_DIR

RUN apt-get update \
    && groupadd -r web \
    && useradd -d $APP_DIR -r -g web web \
    && chown web:web -R $APP_DIR \
    && apt-get install -y netcat-traditional \
    && apt-get install -y acl \
    && apt-get install -y ca-certificates

COPY --from=builder /build/cron_service $APP_DIR/cron_service
COPY --from=builder /build/deploy/scripts/prod-cron-service-start.sh $APP_DIR/prod-cron-service-start.sh
COPY --from=builder /build/deploy/env/.env.prod $APP_DIR//deploy/env/.env.prod

RUN setfacl -R -m u:web:rwx $APP_DIR/prod-cron-service-start.sh

USER web

ENTRYPOINT ["bash", "prod-cron-service-start.sh"]