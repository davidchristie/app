FROM golang:1.16 AS app-builder

WORKDIR /usr/src/app

ENV CGO_ENABLED=0

COPY go.mod go.sum ./

RUN go mod download

COPY services/app services/app

RUN go build ./services/app

FROM node:16.5 as web-builder

WORKDIR /usr/src/app

COPY package-lock.json package.json ./

COPY ./clients/web/package.json ./clients/web/

RUN npm ci

COPY ./clients/web ./clients/web

RUN npm run build

FROM alpine:3.14 as runtime

ENV DATABASE_MIGRATIONS=file://migrations

WORKDIR /usr/src/app

COPY --from=app-builder /usr/src/app/app .
COPY --from=app-builder /usr/src/app/services/app/migrations ./migrations/

COPY --from=web-builder /usr/src/app/clients/web/build ./public

ENTRYPOINT [ "./app" ]
