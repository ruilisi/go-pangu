FROM alpine:3.10
WORKDIR /app
ADD invoice .
COPY application.yml .
RUN chmod +x invoice
CMD ["/app/invoice"]
