FROM alpine:3.10
WORKDIR /app
ADD invoice .
RUN chmod +x invoice
CMD ["/app/invoice"]
