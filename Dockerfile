FROM ubuntu:noble

COPY loopback /opt/loopback

ENTRYPOINT ["/opt/loopback"]
