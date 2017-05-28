FROM golang:alpine

# add user
RUN \
 adduser -s /bin/bash -D bot && \
 mkdir -p /bot && \
 chown -R bot:bot /bot

# build bot
ADD . /go/src/github.com/lastfreeacc/telehabrrssbot
RUN \
 apk add --update --no-progress git && \
 cd /go/src/github.com/lastfreeacc/telehabrrssbot && \
 go get && \
 go build -o /bot/telehabrrssbot && \
 apk del git && \
 rm -rf /go/src/* && \
 rm -rf /var/cache/apk/*

RUN \
 echo "#!/bin/sh" > /bot/exec.sh && \
 echo "/bot/telehabrrssbot &> /bot/out.log &" /bot/exec.sh && \
 chmod +x /bot/exec.sh

USER bot
WORKDIR /bot
ENTRYPOINT ["/bot/exec.sh"]