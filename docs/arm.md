## Compile to ARM
docker build --platform linux/arm64/v8  -f Dockerfile.arm -t zalbiraw/go-api-test-service:arm .
docker push zalbiraw/go-api-test-service:arm