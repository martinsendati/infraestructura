FROM golang:alpine3.15

WORKDIR /src/
COPY app.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/app

ENV PORT=8080
EXPOSE 8080/tcp
ENTRYPOINT ["/bin/app"]
