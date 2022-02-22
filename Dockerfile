# Build step
FROM golang:1.17-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /hiper

# Minimal container used for running the binary
FROM scratch

COPY --from=build /hiper /hiper

ENTRYPOINT ["/hiper"]