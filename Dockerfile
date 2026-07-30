FROM ubuntu:resolute

COPY loopback /opt/loopback

ENTRYPOINT ["/opt/loopback"]
