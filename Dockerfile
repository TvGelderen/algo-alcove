FROM node:21.6-bookworm AS node
WORKDIR /app
COPY . .
RUN npm install
RUN make init-tailwind

FROM golang:1.22.0-bookworm AS builder
WORKDIR /app
COPY --from=node /app .
RUN make init-go
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/assets assets

EXPOSE 3000

CMD [ "/app/main" ]
