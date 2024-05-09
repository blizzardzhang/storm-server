package svc

import (
	"gorm.io/gorm"
	"storm-server/app/sys/cmd/rpc/internal/config"
	"storm-server/app/sys/model/sys"
	"storm-server/common/enter"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := enter.InitGorm(c.DB.DataSource, "test_", sys.Models)
	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
