run:
	@npx tailwindcss -i assets/css/app.css -o assets/css/dist/style.css
	@templ generate view
	@go build -o tmp/main main.go
	@./tmp/main

templ:
	@templ generate --watch --proxy="http://localhost:3000"

tailwind:
	@npx tailwindcss -i assets/css/app.css -o assets/css/dist/style.css --watch

init-tailwind:
	@npx tailwindcss -i assets/css/app.css -o assets/css/dist/style.css

init-go:
	@go install github.com/a-h/templ/cmd/templ@latest
	@templ generate
	@go get
