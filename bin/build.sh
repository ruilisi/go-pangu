GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -tags release -ldflags '-w -s' -o invoice
docker build -t invoice-server .
docker tag invoice-server ccr.ccs.tencentyun.com/ruilisi/invoice-server
docker push ccr.ccs.tencentyun.com/ruilisi/invoice-server
