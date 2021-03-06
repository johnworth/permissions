FROM discoenv/golang-base:master

ENV CONF_TEMPLATE=/go/src/github.com/cyverse-de/permissions/permissions.yaml.tmpl
ENV CONF_FILENAME=permissions.yaml
ENV PROGRAM=permissions

COPY . /go/src/github.com/cyverse-de/permissions

RUN git clone https://github.com/swagger-api/swagger-ui.git /tmp/swagger-ui \
    && cd /tmp/swagger-ui \
    && git checkout v2.1.4 \
    && mkdir -p /docs \
    && cp -pr dist/* /docs/ \
    && cd / \
    && rm -rf /tmp/swagger-ui \
    && cp /go/src/github.com/cyverse-de/permissions/index.html /docs/index.html

RUN go get github.com/constabulary/gb/...

RUN cd /go/src/github.com/cyverse-de/permissions && \
    gb build && \
    cp bin/permissions-server /bin/permissions

WORKDIR /
EXPOSE 60000

ARG git_commit=unknown
ARG version="2.9.0"

LABEL org.cyverse.git-ref="$git_commit"
LABEL org.cyverse.version="$version"
