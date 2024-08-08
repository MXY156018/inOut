package system

import (
	"context"

	"mall-admin/cmd/initdb/logic"
	sysModel "mall-admin/model"

	"gorm.io/gorm"
)

const initOrderAuthorityBtn = initOrderMenuAuthority + 1

type initAuthorityBtn struct{}

// auto run
// func init() {
// 	logic.RegisterInit(initOrderAuthorityBtn, &initAuthorityBtn{})
// }

func (i *initAuthorityBtn) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, logic.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysAuthorityBtn{})
}

func (i *initAuthorityBtn) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysAuthorityBtn{})
}

func (i initAuthorityBtn) InitializerName() string {
	var entity sysModel.SysAuthorityBtn
	return entity.TableName()
}

func (i *initAuthorityBtn) InitializeData(ctx context.Context) (context.Context, error) {
	_, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, logic.ErrMissingDBContext
	}

	return ctx, nil
}

func (i *initAuthorityBtn) DataInserted(ctx context.Context) bool {

	return true
}
