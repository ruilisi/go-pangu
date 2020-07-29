FROM alpine:3.10
WORKDIR /app
ADD invoice .
COPY conf/application.yml conf/application
RUN chmod +x invoice
CMD ["/app/invoice"]
