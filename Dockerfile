ARG IMAGE_ARCH=amd64

# this is needed to change the file permissions on an arm64 build
# since arm64 chown and chmod won't run on an amd64 build host
FROM alpine:3 as permer
ARG GOARCH=amd64

COPY ./OPATH/mockster-*.linux-${GOARCH} /usr/local/bin/mockster
RUN chown nobody:nogroup /usr/local/bin/mockster && chmod +x /usr/local/bin/mockster

FROM ${IMAGE_ARCH}/alpine:3
LABEL maintainer "The Trickster Authors <trickster-developers@googlegroups.com>"

# expects that you are in in $src/github.com/tricksterproxy/mockster
# and have already ran "make release" for binaries to reside in OPATH
COPY  --from=permer /usr/local/bin/mockster /usr/local/bin/mockster
COPY LICENSE                                /LICENSE
COPY NOTICE                                 /NOTICE

USER nobody
EXPOSE 8482
ENTRYPOINT ["/usr/local/bin/mockster"]
