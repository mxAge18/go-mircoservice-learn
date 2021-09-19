protoc -I . --grpc-gateway_out ../serices/product \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=../services/product/
--grpc-gateway_opt standalone=true \
product.proto

protoc -I . \
--go_out ../services/product --go_opt paths=source_relative \
--go-grpc_out ../services/product --go-grpc_opt paths=source_relative \
product.proto

protoc -I . --grpc-gateway_out ../services/product \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
product.proto