FROM ubuntu:wily

COPY loopback /opt/loopback

ENTRYPOINT ["/opt/loopback"]
