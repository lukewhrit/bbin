FROM golang:1.15.5-alpine3.12

RUN mkdir /opt/bbin

COPY . /opt/bbin
WORKDIR /opt/bbin

# GCC is required for cgo support in DataDog/zstd
RUN apk add --no-cache build-base

# Download dependencies
RUN go mod download

# Build the binary
RUN go build --ldflags "-s -w" -o bin/bbin ./

# Run the generated binary
CMD ["/opt/bbin/bin/bbin"]
