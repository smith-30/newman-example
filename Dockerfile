FROM docker:stable-dind

RUN apk --update --no-cache add bash docker-compose
# COPY --from=docker/compose:1.25.0-alpine /usr/local/bin/docker-compose /usr/local/bin/
# RUN chmod +x /usr/local/bin/docker-compose

WORKDIR /go/src/github.com/smith-30/cidind
COPY ./ ./

RUN docker-compose -v
ENV DOCKER_TLS_CERTDIR ""


ENTRYPOINT ["/bin/sh"]