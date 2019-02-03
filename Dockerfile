FROM scratch
EXPOSE 80
COPY ca-certificates.crt /etc/ssl/certs/
COPY ./bin/server/coin /
CMD ["/coin"]
