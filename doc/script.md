goctl api go --dir . --style=goZero --api 

goctl rpc protoc library.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./ --style=goZero

goctl model mysql datasource -url="root:root@tcp(localhost:3306)/mathless" -table="project" -dir=./ --cache=true --style=goZero