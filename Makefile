CSS_INPUT = assets/src/style.css
CSS_OUTPUT = assets/dist/style.css
PROXY_PORT = 8080
APP_PORT = 3000

css-watch:
	npx tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT) --watch=always

templ-watch:
	templ generate -watch \
		-proxy="http://localhost:$(APP_PORT)" \
		-proxyport="$(PROXY_PORT)" \
		--open-browser=false \
		-cmd="go run ."

dev:
	parallel -j 2 --line-buffer ::: \
		'make css-watch' \
		'make templ-watch'

build:
	templ generate 
	npx tailwindcss -i $(CSS_INPUT) -o $(CSS_OUTPUT)
	go build -tags embed -o ./tmp/build