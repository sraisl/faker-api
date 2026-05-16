FROM golang:1.26.3-trixie AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/faker-api .

FROM debian:trixie-slim AS runtime

RUN apt-get update \
	&& apt-get install -y --no-install-recommends ca-certificates \
	&& rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=build /out/faker-api /app/faker-api

EXPOSE 8888

ENTRYPOINT ["/app/faker-api"]
