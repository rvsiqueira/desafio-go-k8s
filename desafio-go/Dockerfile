FROM golang:alpine as golang
WORKDIR /go/src/app
COPY src/main.go .
# Static build required so that we can safely copy the binary over.
RUN CGO_ENABLED=0 go install
RUN rm main.go

FROM scratch
# the test program:
COPY --from=golang /go/bin/app /app
ENTRYPOINT ["/app"]
EXPOSE 8080