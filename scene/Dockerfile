FROM golang:1.17 as build
RUN mkdir -p /app/building
WORKDIR /app/building
ADD . /app/building
RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim
# Copy from docker build
COPY --from=build /app/building/dist/bin/discovery /app/bin/
# Copy from local build
#ADD  dist/ /app/
ENV  LOG_DIR    /app/log
EXPOSE 7171
WORKDIR /app/
CMD  /app/bin/discovery