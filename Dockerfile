FROM --platform=arm64 golang:alpine3.23

WORKDIR /

COPY / /

RUN go build

RUN chmod +x ./bicho

CMD ["./bicho", "api"]

EXPOSE 8080
