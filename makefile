.SILENT:

sw:
	go run websocket/server/server.go

cw:
	go run websocket/client/client.go

sg:
	go run grpc/server/server.go

cg:
	go run grpc/client/client.go

e:
	go run experiment/experiment.go

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc\proto\p.proto