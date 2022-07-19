FROM golang:1.17 as build

ENV http_proxy "http://172.17.0.1:7890"
ENV HTTP_PROXY "http://172.17.0.1:7890"
ENV https_proxy "http://172.17.0.1:7890"
ENV HTTPS_PROXY "http://172.17.0.1:7890"

WORKDIR /build
COPY . .

RUN make build

FROM debian:bullseye-slim

COPY --from=build /build/flex-gpu-device-plugin /usr/bin/flex-gpu-device-plugin

ENTRYPOINT ["flex-gpu-device-plugin"]
