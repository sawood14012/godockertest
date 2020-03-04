From golang:latest
ADD . /go/src/github.com/sawood14012/godockertest
# Build the contact_registry command inside the container.
RUN go install github.com/sawood14012/godockertest
# Run the contact_registry command when the container starts.
ENTRYPOINT /go/bin/godockertest
# http server listens on port 8080.
EXPOSE 8080