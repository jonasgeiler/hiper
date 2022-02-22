#############################################
# STEP 1: Build optimized executable binary #
#############################################
FROM golang:1.17-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /hiper

#######################################################
# STEP 2: Build a minimal image which runs the binary #
#######################################################
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /hiper /hiper

ENTRYPOINT ["/hiper"]