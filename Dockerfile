#############################################
# STEP 1: Build optimized executable binary #
#############################################
FROM golang:1.17-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /hiper

#######################################################
# STEP 2: Build a minimal image which runs the binary #
#######################################################
FROM alpine

COPY --from=build /hiper /hiper

ENTRYPOINT ["/hiper"]