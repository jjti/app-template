FROM golang:1.20.5
WORKDIR /go/src/github.com/jjtimmons/sblast/
COPY . ./
RUN CGO_ENABLED=0 go build -o sblast -buildvcs=false .
CMD ["./sblast"]
EXPOSE 8080
