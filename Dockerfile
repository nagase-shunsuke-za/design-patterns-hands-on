FROM node:21 as node
FROM golang:1.22

COPY --from=node /usr/local/bin/node /usr/local/bin/
COPY --from=node /usr/local/lib/node_modules/ /usr/local/lib/node_modules/

RUN ln -s /usr/local/bin/node /usr/local/bin/nodejs \
        && ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npm \
        && ln -s /usr/local/lib/node_modules/npm/bin/npm-cli.js /usr/local/bin/npx


ENV PATH $PATH:/go/src/node_modules/.bin

WORKDIR /go/src
