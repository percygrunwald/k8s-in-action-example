ARG GO_VERSION=1.18

FROM golang:${GO_VERSION} AS build
ENV GO_BUILD="CGO_ENABLED=0 go build -installsuffix 'static' -o /go/bin/ /go/src/cmd"
WORKDIR /go/src
COPY . /go/src/
RUN sh -c "$GO_BUILD/k8s-in-action-example"

FROM gcr.io/distroless/static AS final
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /go/bin /app
ENTRYPOINT ["/app/k8s-in-action-example"]
