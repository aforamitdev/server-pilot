package system

import (
	"github.com/aforamitdev/server-pilot/internal/protogen"
	"github.com/aforamitdev/server-pilot/pkg/logger"
)

type SystemInformer struct {
	protogen.UnimplementedStatusServiceServer
	log *logger.Logger
}

func NewSystemInformer(log *logger.Logger) *SystemInformer {
	return &SystemInformer{
		log: log,
	}
}
