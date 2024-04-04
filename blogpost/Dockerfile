FROM golang:latest AS builder

COPY . /go/src/blogpost
WORKDIR /go/src/blogpost
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /go/bin/blogpost ./cmd/blogpost

FROM scratch
COPY --from=builder /go/bin/blogpost /go/bin/blogpost
ENTRYPOINT ["/go/bin/blogpost"]