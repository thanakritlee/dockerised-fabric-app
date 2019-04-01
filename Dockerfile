# GoLang build stage
FROM golang:1.12 as builder

WORKDIR /app

COPY . .


# Final stage
FROM scratch

WORKDIR /app