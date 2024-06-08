package system

import (
	"context"

	system_proto "github.com/aforamitdev/server-pilot/internal/proto/system"
)

type NewSystemInformer struct {
	system_proto.UnimplementedInformerServer
}

func (i NewSystemInformer) GetSystem(ctx context.Context, req *system_proto.SystemRequest) (*system_proto.SystemResponse, error) {
	return &system_proto.SystemResponse{
		Name: "Amti rai",
	}, nil
}
