module github.com/jhump/grpctunnel

go 1.22.3

toolchain go1.24.1

require (
	github.com/fullstorydev/grpchan v1.1.1
	github.com/siderolabs/grpc-proxy v0.5.1
	github.com/stretchr/testify v1.10.0
	golang.org/x/sync v0.10.0
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/siderolabs/grpc-proxy v0.5.1 => github.com/cbws/grpc-proxy v0.0.0-20250311113103-e3a676b98e4c
