module protoExample

go 1.17

replace github.com/farhadmirr/proto v0.0.1 => ./proto

require github.com/farhadmirr/simple-gRPC-handler v0.0.0-20211227193921-c1a043f42399

require (
	github.com/farhadmirr/proto v0.0.1 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.43.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
