FROM golang:latest as build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/grandcolline/gke-demo-app
COPY . .
RUN go install


FROM gcr.io/distroless/base

COPY --from=build /go/bin/gke-demo-app /gke-demo-app
ENV PORT=8080
EXPOSE 8080

CMD ["/gke-demo-app"]
