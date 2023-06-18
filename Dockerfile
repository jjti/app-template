FROM golang:1.20.5
WORKDIR /
COPY . ./
RUN CGO_ENABLED=0 go build -o app -buildvcs=false ./cmd
CMD ["./app"]
EXPOSE 8080
