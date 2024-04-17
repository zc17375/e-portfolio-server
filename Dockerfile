FROM golang:1.21-alpine as builder

WORKDIR /build/myapp

# COPY go.mod, go.sum and download the dependencies
COPY go.* .
RUN go mod download

# COPY All things inside the project and build
COPY . .
COPY config.yaml /build/myapp/config.yaml
RUN go build -o /project/go-docker/build/myapp .

FROM scratch
COPY --from=builder /project/go-docker/build/myapp /project/go-docker/build/myapp
COPY --from=builder /build/myapp/config.yaml /config.yaml


EXPOSE 8888
ENTRYPOINT [ "/project/go-docker/build/myapp" ]