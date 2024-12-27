FROM alpine:3.21 AS builder

WORKDIR /src

COPY . /src

RUN apk add \
        go \
        make \
        curl

RUN curl -L https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 \
        -o /usr/local/bin/tailwindcss \
        && chmod +x /usr/local/bin/tailwindcss

RUN go mod download
RUN make tailwind-build

RUN CGO_ENABLED=0 go build -o /src/build/server main.go

FROM scratch

WORKDIR /app
COPY --from=builder /src/build/ /app
COPY --from=builder /src/static /app/static

ENTRYPOINT [ "/app/server" ]
