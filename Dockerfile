FROM golang:1.15.5-alpine3.12

RUN mkdir /opt/sojourner

COPY . /opt/sojourner
WORKDIR /opt/sojourner

# GCC is required for cgo support in DataDog/zstd
RUN apk add --no-cache build-base

# Download dependencies
RUN go mod download

# Build the binary
RUN go build --ldflags "-s -w" -o bin/sojourner ./

# Run the generated binary
CMD ["/opt/sojourner/bin/sojourner"]
