FROM golang:1.21 as build
ADD . /go/web5-components-test
WORKDIR /go/web5-components-test
RUN CGO_ENABLED=0 go build ./cmd/web5-components-tests

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/web5-components-test/web5-components-tests /web5-components-tests
CMD [ "/web5-components-tests", "-no-start" ]
