# build stage
FROM golang:1.18.8-alpine3.15 as builder

ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

COPY . /src

WORKDIR /src

# RUN apk add git
RUN apk add make

RUN go mod download

RUN make build-app

# Run stage
FROM golang:1.18.8-alpine3.15 AS app

COPY . /app

WORKDIR /app

COPY --from=builder /src/mst /app

# App
EXPOSE 8000

CMD [ "/app/mst" ]