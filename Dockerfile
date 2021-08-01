FROM golang:1.16 AS server-builder

WORKDIR /usr/src/app

ENV CGO_ENABLED=0

COPY go.mod go.sum ./

RUN go mod download

COPY services/server services/server

RUN go build ./services/server

FROM node:16.5 as web-builder

WORKDIR /usr/src/app

COPY package-lock.json package.json ./

COPY ./clients/web/package.json ./clients/web/

RUN npm ci

COPY ./clients/web ./clients/web

RUN npm run build

FROM alpine:3.14 as runtime

WORKDIR /usr/src/app

COPY --from=server-builder /usr/src/app/server .

COPY --from=web-builder /usr/src/app/clients/web/build ./public

ENTRYPOINT [ "./server" ]
