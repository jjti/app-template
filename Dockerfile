FROM golang:1.20.5
WORKDIR /
COPY . ./
RUN CGO_ENABLED=0 go build -o sblast -buildvcs=false .
CMD ["./sblast"]
EXPOSE 8080
