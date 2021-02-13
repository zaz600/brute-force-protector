# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -v -o bp-server ./cmd/bp-server/...

# final stage
FROM alpine
WORKDIR /app
EXPOSE 50051
COPY --from=build-env /src/bp-server /app/
ENTRYPOINT ./bp-server