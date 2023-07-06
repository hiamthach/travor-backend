//go:build tools
// +build tools

package tools

import (
	_ "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	_ "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
