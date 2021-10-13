FROM alpine
LABEL name="httpserver"
WORKDIR /opt/httpserver
COPY httpserver /opt/httpserver/
EXPOSE 8080
ENTRYPOINT  ["/opt/httpserver/httpserver"]
