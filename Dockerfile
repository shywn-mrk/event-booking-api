# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-event-book-api"
LABEL REPO="https://github.com/shywn-mrk/event-book-api"

ENV PROJPATH=/go/src/github.com/shywn-mrk/event-book-api

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/shywn-mrk/event-book-api
WORKDIR /go/src/github.com/shywn-mrk/event-book-api

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/shywn-mrk/event-book-api"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/event-book-api/bin

WORKDIR /opt/event-book-api/bin

COPY --from=build-stage /go/src/github.com/shywn-mrk/event-book-api/bin/event-book-api /opt/event-book-api/bin/
RUN chmod +x /opt/event-book-api/bin/event-book-api

# Create appuser
RUN adduser -D -g '' event-book-api
USER event-book-api

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/event-book-api/bin/event-book-api"]
