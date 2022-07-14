FROM golang:1.17 as build

ENV http_proxy "http://172.17.0.1:7890"
ENV HTTP_PROXY "http://172.17.0.1:7890"
ENV https_proxy "http://172.17.0.1:7890"
ENV HTTPS_PROXY "http://172.17.0.1:7890"

WORKDIR /build
COPY . .

RUN make build

FROM alpine:3.15

COPY --from=build /build/gpu-share-device-plugin /usr/bin/gpu-share-device-plugin

ENTRYPOINT ["gpu-share-device-plugin"]
