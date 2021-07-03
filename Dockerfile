FROM golang:1.16.5 AS builder



WORKDIR $GOPATH/src/github.com/mchirico/gosubshell

# Copy the entire project and build it

COPY . $GOPATH/src/github.com/mchirico/gosubshell
RUN go get -v -t -d .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o /bin/project


# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /bin/project .
USER 65532:65532

ENTRYPOINT ["/project"]



