FROM node:23.5-bookworm AS node
WORKDIR /app
COPY . .
RUN npm install
RUN make init-tailwind

FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY --from=node /app .
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go get
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/assets assets

EXPOSE 5000

CMD [ "/app/main" ]
