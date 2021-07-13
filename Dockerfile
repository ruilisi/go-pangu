FROM alpine:3.10
WORKDIR /app
ADD dist/go-pangu-amd64-release-darwin go-pangu
COPY application.yml .
RUN chmod +x go-pangu
CMD ["/app/go-pangu"]
