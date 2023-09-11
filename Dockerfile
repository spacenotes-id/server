FROM golang:1.19.13-alpine3.18 AS builder

WORKDIR /src

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod download && \
  go build -ldflags="-s -w" -o spacenotes

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder --chown=nonroot:nonroot /src/spacenotes /app/spacenotes
COPY --from=builder --chown=nonroot:nonroot /src/database /app/database
COPY --from=builder --chown=nonroot:nonroot /src/docs /app/docs

ENTRYPOINT [ "/app/spacenotes" ]
