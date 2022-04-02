## How to build containers

docker build -t zalbiraw/go-api-test-service:latest .

### ARM
docker build --platform linux/arm64/v8 --build-arg var_name=arm64 -t zalbiraw/go-api-test-service:arm .
