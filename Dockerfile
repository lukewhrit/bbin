FROM golang:1.15.5-alpine3.12.1

RUN mkdir /opt/bbin

COPY . /opt/bbin
WORKDIR /opt/bbin

# Download dependencies
RUN go mod download

# Build the binary
RUN go build --ldflags "-s -w" -o bin/bbin ./

# Run the generated binary
CMD ["/opt/bbin/bin/bbin"]
