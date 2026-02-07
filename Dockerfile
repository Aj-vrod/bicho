# --- Builder
FROM golang:1.24 AS builder

WORKDIR /

COPY . .
RUN make build
# --- End of Builder

FROM ubuntu:24.04
COPY --from=builder out/bicho /bin/bicho

RUN chmod a+x bin/bicho

ENTRYPOINT [ "bin/bicho", "api" ]
