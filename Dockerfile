# NodeJS build stage
FROM node:10.15.3-alpine as clientbuilder

WORKDIR /app

COPY . .

RUN cd web/client && \
    npm install && \
    npm run build

# GoLang build stage
FROM golang:1.12 as builder

WORKDIR /app

COPY . .

RUN cd web/ && \
    go mod download

RUN cd web/ && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o web

# Final stage
FROM scratch

WORKDIR /app

COPY --from=clientbuilder /app/web/client/build /app/web/client/build
COPY --from=builder /app/web/web /app/web/web
COPY --from=builder /app/web/config.yaml /app/web/config.yaml
COPY --from=builder /app/chaincode /app/chaincode
COPY --from=builder /app/fabric-network /app/fabric-network