ARG IMAGE_ARCH=amd64

FROM ${IMAGE_ARCH}/alpine:3
LABEL maintainer "The Trickster Authors <trickster-developers@googlegroups.com>"

ARG GOARCH=amd64
# expects that you are in in $src/github.com/tricksterproxy/mockster
# and have already ran "make release" for binaries to reside in OPATH
COPY ./OPATH/mockster-*.linux-${GOARCH} /usr/local/bin/mockster
COPY LICENSE                                /LICENSE
COPY NOTICE                                 /NOTICE

RUN chown nobody:nogroup /usr/local/bin/mockster && chmod +x /usr/local/bin/mockster

USER nobody
EXPOSE 8482
ENTRYPOINT ["/usr/local/bin/mockster"]
