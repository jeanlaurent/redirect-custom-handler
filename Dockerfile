FROM golang:1.7.1-alpine as gobuilder
WORKDIR /go/src/github.com/jeanlaurent/redirect
COPY . ./
RUN go build -o redirect ./

FROM alpine:3.6
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=gobuilder /go/src/github.com/jeanlaurent/redirect /app/
ENTRYPOINT ["./redirect"]
