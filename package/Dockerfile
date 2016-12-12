FROM ubuntu:16.04
COPY ssh-pub-key-mgr /usr/bin/

# File obtained from curl -o ca-certificates.crt https://curl.haxx.se/ca/cacert.pem
COPY ca-certificates.crt /etc/ssl/certs/
CMD ["ssh-pub-key-mgr"]
