.PHONY: tailwind-watch
tailwind-watch:
	tailwindcss -i ./static/input.css -o ./static/styles.min.css --watch --minify

.PHONY: tailwind-build
tailwind-build: 
	tailwindcss -i ./static/input.css -o ./static/styles.min.css --minify

.PHONY: dev
dev:
	air

.PHONY: build
build:
	docker buildx build --platform=linux/amd64,linux/arm64 -t hydrojohnny/goth-stats:latest .
