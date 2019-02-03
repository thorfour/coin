FROM scratch
RUN update-ca-certificates
EXPOSE 80
COPY bin/server/coin /
CMD ["/coin"]
