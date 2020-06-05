# s3-local-test

## usage

```shell script
cp .env.example .env
vim .env
docker-compose up -d
aws s3 mb --endpoint-url http://localhost:9001 s3://testkun
curl -X PUT --upload-file test.txt $(go run main.go)
aws s3 ls --endpoint-url http://localhost:9001 s3://testkun/test/
```