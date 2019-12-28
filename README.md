# grpc-gateway-example 

# Step 1:
Get know about gRPC Gateway over there: [gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway)

# Step 2:
Install some tools related gRPC Gateway

# Step 3:
Pull this example & run 
```bash
# Ensure export GO111MODULE=on
go mod vendor
```

# Step 4:
If you want to update protobuf definition
```bash
./scripts/pb-gen.sh
``` 

# Step 5:
Start server gRPC first: server/main.go
Then start server gRPC Gateway main.go

```bash
# Call a HTTP Request to Gateway & check result
curl -X GET \
  http://127.0.0.1:8080/get \
  -H 'Content-Type: application/json' \
  -d '{
    "value": "ekuuuu"
}'
```

```bash
# Call a HTTP Request to Gateway to get swagger content
curl -X GET \
  http://127.0.0.1:8080/swagger/template.swagger.json
```
