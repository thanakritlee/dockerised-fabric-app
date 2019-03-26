# GoLang build stage
FROM golang:1.12 as builder

WORKDIR /app

# Final stage
FROM scratch

WORKDIR /app