package system

import (
	"context"
	"os/exec"

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

func (s *SystemInformer) GetStatus(ctx context.Context, statusRequest *protogen.StatusRequest) (*protogen.StatusResponse, error) {

	cmd := exec.Command("lsb_release", "", "--all")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	statusResp := protogen.StatusResponse{
		Status: string(output),
	}
	return &statusResp, nil

}
