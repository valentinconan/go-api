FROM gcr.io/distroless/base-debian10
ARG VERSION
ENV APPLICATION_VERSION=$VERSION

WORKDIR /

COPY bin/go-api-amd64-linux /go-api-amd64-linux

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-api-amd64-linux"]
