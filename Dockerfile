FROM golang:1.12-stretch
LABEL maintainer="Borys Shchipanskyi <borys_shchipanskyi@bose.com>"

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/BoseCorp/monetization/
WORKDIR /go/src/github.com/BoseCorp/monetization/

RUN \
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o service ./

EXPOSE 8080

# Run the service command when the container starts.
ENTRYPOINT ["/service"]