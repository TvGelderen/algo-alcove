FROM node:21.6-bookworm AS node
WORKDIR /app
COPY . .
RUN npm install
RUN npx tailwindcss -i assets/css/app.css -o assets/css/dist/style.css

FROM golang:1.22.0-bookworm AS builder
WORKDIR /app
COPY --from=node /app .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/assets assets
COPY .env .

EXPOSE 3000

CMD [ "/app/main" ]
