# The GOTH stack
A simple stack with little dependencies

For a live version of this app you can go [here](https://goth-stats.click)
# Design Choices
This stack is an adaption of [GOTTH](https://github.com/TomDoesTech/GOTTH) with less dependencies. Here's a breakdown of each and why I used it
## net/http
I use the standard lib http because it's router is powerful enough for me so I don't have to have a dependency like [Chi](https://github.com/go-chi/chi). It has support for routing based on method and path and middleware which is everything I need.
## http/template
I use http/template instead of [Templ](https://github.com/a-h/templ), again because the standard library has every feature I need. I found that the added complexity of Templ just isn't worth the hassle for my apps
## tailwind
I use [tailwind](https://github.com/tailwindlabs/tailwindcss) to take the guesswork out of making the GUI look presentable. I am a backend dev by trade and am not very good at creating GUIs but tailwind with its sensible defaults makes it very easy, of course to minimize dependencies further it's also possible to write the CSS yourself but that isn't for me.
## htmx
As replacement for a JavaScript framework like React or Angular I use [htmx](https://github.com/bigskysoftware/htmx) as it offers a very minimal way for creating interactive GUIs. Its minimal approach of extending HTML and rendering HTML server-side and just swapping it in is IMHO a very nice way of creating interactivity with a very minimal set of functions.
## sqlc
For my database integration I use [sqlc](https://github.com/sqlc-dev/sqlc) to generate type safe Go code. I like the repository pattern over ORMs as it is simpler for me to write SQL queries instead of using a framework to generate queries on the fly.
## migrate
For automatically migrating my database i use [migrate](https://github.com/golang-migrate/migrate) as it is very simple and works well together with sqlc
## pgx/v5
For my database connection I use [pgx](https://github.com/jackc/pgx) as it was recommended for use with sqlc.
## envconfig
For configuration I use environment variables and [envconfig](https://github.com/kelseyhightower/envconfig) provides a simple interface to load environment variables into structs
# Supporting tools
This are the tools I use to deploy the resulting app
## Makefile
For building I use a Makefile so that I can alias commonly used commands in a simple way without cluttering my zsh config.
## Docker
For deploying the GOTH app I use Docker stacks with NGINX Proxy Manager handling the reverse proxy.