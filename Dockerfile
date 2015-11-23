
FROM scratch
ENTRYPOINT ["/gohttprelay"]

# Add the binary
ADD gohttprelay /
EXPOSE 8080
