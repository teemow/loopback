FROM ubuntu:wily-20160706

COPY loopback /opt/loopback

ENTRYPOINT ["/opt/loopback"]
