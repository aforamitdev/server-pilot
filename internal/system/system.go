package system

import (
	"context"

	system_proto "github.com/aforamitdev/server-pilot/internal/proto/system"
	"github.com/aforamitdev/server-pilot/pkg/logger"
)

type SystemInformer struct {
	system_proto.UnimplementedInformerServer
	log *logger.Logger
}

func NewSystemInformer(log *logger.Logger) *SystemInformer {
	return &SystemInformer{
		log: log,
	}
}

func (s SystemInformer) GetSystem(ctx context.Context, req *system_proto.SystemRequest) (*system_proto.SystemResponse, error) {
	return &system_proto.SystemResponse{
		Name: "Amti rai",
	}, nil
}
