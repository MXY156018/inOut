package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	"mall-admin/model"
	"mall-admin/types"
	"mall-pkg/api"
	"mall-pkg/utils"
	"strconv"
	"strings"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Out struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewOut(ctx context.Context, svcCtx *svc.ServiceContext) Out {
	return Out{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

func (l *Out) AddOutRecord(req *types.OutRecordReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	if req.Records == "" {
		resp.Code = api.Error_Server
		resp.Msg = "请输入金额"
		return resp
	}
	records := strings.Split(req.Records, ",")

	var data []model.OutRecords
	var sum float64
	var date = time.Now()
	for _, v := range records {
		if v == "" {
			continue
		}
		amount, _ := strconv.ParseFloat(v, 32)
		sum += amount
		data = append(data, model.OutRecords{
			Money:  float32(amount),
			Date:   date,
			Remark: "",
		})
	}

	err := l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Create(&data).Error
		if merr != nil {
			return merr
		}
		merr = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "date"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"money": gorm.Expr("money+?", sum),
			}),
		}).Create(&model.OutDay{
			Date:  date,
			Money: float32(sum),
		}).Error
		if merr != nil {
			return merr
		}
		merr = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "date"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"money": gorm.Expr("money+?", sum),
			}),
		}).Create(&model.OutMonth{
			Date:  date.Format("2006-01"),
			Money: float32(sum),
		}).Error
		if merr != nil {
			return merr
		}
		return nil
	})
	if err != nil {
		l.sCtx.Log.Error("添加失败", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = "添加失败"
		return resp
	}
	resp.Code = api.Error_OK
	resp.Msg = "添加成功"
	return resp
}
func (l *Out) GetOutRecords(req *types.GetRecordsReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	offset := req.PageSize * (req.Page - 1)
	limit := req.PageSize
	var data []model.OutRecords
	var count int64

	db := l.sCtx.DB.Model(&data)

	if req.StartTime != "" && req.EndTime != "" {
		var startTime, endTime time.Time
		startTime, err := utils.StringToTime(req.StartTime)
		if err != nil {
			resp.Code = api.Error_Server
			resp.Msg = "时间格式错误"
			return resp
		}
		endTime, err = utils.StringToTime(req.EndTime)
		if err != nil {
			resp.Code = api.Error_Server
			resp.Msg = "时间格式错误"
			return resp
		}

		db = db.Where("date >= ? && date<?", startTime, endTime)
	}

	err := db.Count(&count).Offset(offset).Limit(limit).Order("date desc").Find(&data).Error
	if err != nil {
		l.sCtx.Log.Error("查询失败", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = "查询失败"
		return resp
	}
	resp.Code = api.Error_OK
	resp.Data = api.PageResp{
		Total:    count,
		List:     data,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	return resp
}
func (l *Out) DelOutRecords(req *api.IDReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	var count int64
	var data model.OutRecords
	l.sCtx.DB.Model(&model.OutRecords{}).Where("id = ?", req.ID).Count(&count).Find(&data)
	if count == 0 {
		resp.Code = api.Error_Server
		resp.Msg = "该记录不存在"
		return resp
	}
	dayDate := data.Date.Format("2006-01-02")
	month := data.Date.Format("2006-01")
	err := l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Where("id = ?", req.ID).Delete(&model.OutRecords{}).Error
		if merr != nil {
			return merr
		}
		merr = tx.Model(&model.OutDay{}).Where("date = ?", dayDate).Update("money", gorm.Expr("money-?", data.Money)).Error
		if merr != nil {
			return merr
		}
		merr = tx.Model(&model.OutMonth{}).Where("date = ?", month).Update("money", gorm.Expr("money-?", data.Money)).Error
		if merr != nil {
			return merr
		}
		return nil
	})
	if err != nil {
		l.sCtx.Log.Error("删除失败", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = "删除失败"
	}
	resp.Code = api.Error_OK
	resp.Msg = "删除成功"
	return resp
}
