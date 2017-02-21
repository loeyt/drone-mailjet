FROM scratch
ADD drone-mailjet /
ENTRYPOINT ["/drone-mailjet"]
