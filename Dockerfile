FROM golang:latest AS builder

RUN go get github.com/codegangsta/negroni
RUN go get github.com/kamase14/kadai

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/kamase14/kadai
COPY . .
RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/kamase14/kadai /app

CMD /app/main 80

