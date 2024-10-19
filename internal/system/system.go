package system

import (
	"context"

	apiv1 "github.com/aforamitdev/server-pilot/app/spilothq/gen/proto/api/v1"
	"github.com/aforamitdev/server-pilot/pkg/logger"
	"github.com/shirou/gopsutil/v3/host"
)

type SystemInformer struct {
	apiv1.UnimplementedSystemServicesServer
	log *logger.Logger
}

func NewSystemInformer(log *logger.Logger) *SystemInformer {
	return &SystemInformer{
		log: log,
	}
}

func (s *SystemInformer) GetStatus(ctx context.Context, req *apiv1.GetStatusRequest) (*apiv1.GetStatusResponse, error) {

	host, err := host.Info()
	if err != nil {
		return nil, err
	}

	rsp := apiv1.GetStatusResponse{System: string(host.OS)}
	return &rsp, nil

}
