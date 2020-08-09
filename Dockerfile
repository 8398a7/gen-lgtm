FROM scratch
COPY gen-lgtm /
ENTRYPOINT ["/gen-lgtm"]
