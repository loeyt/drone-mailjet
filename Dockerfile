FROM centurylink/ca-certs
ADD drone-mailjet /
ENTRYPOINT ["/drone-mailjet"]
