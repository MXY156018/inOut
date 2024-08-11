package logic

import (
	"context"
	"mall-admin/api/internal/svc"
	"mall-admin/model"
	"mall-admin/types"
	"mall-pkg/api"
	"mall-pkg/utils"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type In struct {
	ctx  context.Context
	sCtx *svc.ServiceContext
}

func NewIn(ctx context.Context, svcCtx *svc.ServiceContext) In {
	return In{
		ctx:  ctx,
		sCtx: svcCtx,
	}
}

func (l *In) AddInRecord(req *types.InRecordReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	if req.Sum == 0 {
		resp.Code = api.Error_Server
		resp.Msg = "请输入金额"
		return resp
	}
	if req.Type == "" {
		resp.Code = api.Error_Server
		resp.Msg = "请输入类型"
		return resp
	}
	var data model.InRecords
	var typeData model.InTypes
	if req.TypeId == 0 {
		typeData.Type = req.Type
		err := l.sCtx.DB.Save(&typeData).Error
		if err != nil {
			l.sCtx.Log.Error("添加失败", zap.Error(err))
			resp.Code = api.Error_Server
			resp.Msg = "添加失败"
			return resp
		}
		data.TypeId = typeData.Id
	} else {
		data.TypeId = req.TypeId
	}

	var date = time.Now()
	data.Date = date
	data.IsSettle = req.IsSettle
	data.Price = req.Price
	data.Settle = req.Settle
	data.Sum = req.Sum
	data.Type = req.Type
	data.Weight = req.Weight

	err := l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Create(&data).Error
		if merr != nil {
			return merr
		}
		merr = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "date"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"money": gorm.Expr("money+?", req.Settle),
			}),
		}).Create(&model.InDay{
			Date:  date,
			Money: float32(req.Settle),
		}).Error
		if merr != nil {
			return merr
		}
		merr = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "date"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"money": gorm.Expr("money+?", req.Settle),
			}),
		}).Create(&model.InMonth{
			Date:  date.Format("2006-01"),
			Money: float32(req.Settle),
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
func (l *In) GetInRecords(req *types.GetRecordsReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	offset := req.PageSize * (req.Page - 1)
	limit := req.PageSize
	var data []model.InRecords
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
	if req.Sort != "" {
		db = db.Order(req.Sort)
	} else {
		db = db.Order("date desc")
	}
	err := db.Count(&count).Offset(offset).Limit(limit).Find(&data).Error
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
func (l *In) DelInRecords(req *api.IDReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	var count int64
	var data model.InRecords
	l.sCtx.DB.Model(&model.InRecords{}).Where("id = ?", req.ID).Count(&count).Find(&data)
	if count == 0 {
		resp.Code = api.Error_Server
		resp.Msg = "该记录不存在"
		return resp
	}
	dayDate := data.Date.Format("2006-01-02")
	month := data.Date.Format("2006-01")
	err := l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Where("id = ?", req.ID).Delete(&model.InRecords{}).Error
		if merr != nil {
			return merr
		}
		merr = tx.Model(&model.OutDay{}).Where("date = ?", dayDate).Update("money", gorm.Expr("money-?", data.Settle)).Error
		if merr != nil {
			return merr
		}
		merr = tx.Model(&model.OutMonth{}).Where("date = ?", month).Update("money", gorm.Expr("money-?", data.Settle)).Error
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

func (l *In) InGetType() *api.BaseResp {
	var resp = &api.BaseResp{}
	var data []model.InTypes
	err := l.sCtx.DB.Model(&data).Find(&data).Error
	if err != nil {
		resp.Code = api.Error_Server
		resp.Msg = "查询失败"
		return resp
	}
	resp.Code = api.Error_OK
	resp.Data = data
	return resp
}
func (l *In) InGetTypeDetail(req *types.TypeDetailReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	var data []model.InRecords
	var count int64
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := l.sCtx.DB.Model(&data).Where("type_id=?", req.Id)
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
	if req.Settle != -1 {
		db = db.Where("is_settle = ?", req.Settle)
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

func (l *In) InGetById(req *api.IDReq) *api.BaseResp {
	var resp = &api.BaseResp{}
	if req.ID == 0 {
		resp.Code = api.Error_Server
		resp.Msg = "id不能为空"
		return resp
	}
	var data model.InRecords
	err := l.sCtx.DB.Model(&data).Where("id = ?", req.ID).Find(&data).Error
	if err != nil {
		l.sCtx.Log.Error("查询失败", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = "查询失败"
	}
	resp.Code = api.Error_OK
	resp.Data = data
	return resp
}
func (l *In) InUpdateById(req *model.InRecords) *api.BaseResp {
	var resp = &api.BaseResp{}
	if req.Id == 0 {
		resp.Code = api.Error_Server
		resp.Msg = "id不能为空"
		return resp
	}
	var data model.InRecords
	err := l.sCtx.DB.Model(&data).Where("id = ?", req.Id).Find(&data).Error
	if err != nil {
		l.sCtx.Log.Error("查询失败", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = "查询失败"
	}
	if req.Settle > data.Sum {
		resp.Code = api.Error_Server
		resp.Msg = "结算金额不能大于总金额"
		return resp
	}
	update := req.Settle - data.Settle
	if req.IsSettle == 1 && req.Settle != data.Sum {
		//结算金额发生变化
		resp.Code = api.Error_Server
		resp.Msg = "结算金额错误"
		return resp
	} else {
		if req.IsSettle == 2 && req.Settle > data.Sum-data.Settle {
			resp.Code = api.Error_Server
			resp.Msg = "结算金额错误"
			return resp
		} else if req.IsSettle == 2 && req.Settle == data.Sum {
			resp.Code = api.Error_Server
			resp.Msg = "应选已结账"
			return resp
		}
	}
	err = l.sCtx.DB.Transaction(func(tx *gorm.DB) error {
		merr := tx.Model(&data).Where("id = ?", req.Id).Updates(req).Error
		if err != nil {
			return merr
		}
		month := data.Date.Format("2006-01")
		merr = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "date"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"money": gorm.Expr("money+?", update),
			}),
		}).Create(&model.InDay{
			Date:  data.Date,
			Money: float32(req.Settle),
		}).Error
		if err != nil {
			return merr
		}
		merr = tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "date"}},
			DoUpdates: clause.Assignments(map[string]interface{}{
				"money": gorm.Expr("money+?", update),
			}),
		}).Create(&model.InMonth{
			Date:  month,
			Money: float32(req.Settle),
		}).Error
		if err != nil {
			return merr
		}
		return nil
	})

	if err != nil {
		l.sCtx.Log.Error("更新失败", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = "更新失败"
	}
	resp.Code = api.Error_OK
	resp.Msg = "更新成功"
	return resp
}
func (l *In) GetInSum(req *types.Date) *api.BaseResp {
	var resp = &api.BaseResp{}
	var data []model.InRecords
	db := l.sCtx.DB.Model(&data)
	if req.StartTime == "" && req.EndTime == "" {
		start, end := utils.GetTodayDateTime(time.Now())
		db = db.Where("date >= ? and date<?", start, end)
	} else {
		startTime, err := utils.StringToTime(req.StartTime)
		if err != nil {
			resp.Code = api.Error_Server
			resp.Msg = "时间格式错误"
			return resp
		}
		endTime, err := utils.StringToTime(req.EndTime)
		if err != nil {
			resp.Code = api.Error_Server
			resp.Msg = "时间格式错误"
			return resp
		}
		db = db.Where("date >= ? && date<?", startTime, endTime)
	}
	err := db.Find(&data).Error
	if err != nil {
		l.sCtx.Log.Error("查询失败", zap.Error(err))
		resp.Code = api.Error_Server
		resp.Msg = "查询失败"
		return resp
	}
	var total, unsettle, finish, unfinish, unfinish_settle = 0.0, 0.0, 0.0, 0.0, 0.0

	for _, v := range data {

		total += float64(v.Sum)
		if v.IsSettle == 0 {
			unsettle += float64(v.Sum)
		} else if v.IsSettle == 1 {
			finish += float64(v.Settle)
		} else if v.IsSettle == 2 {
			unfinish += float64(v.Sum - v.Settle)
			unfinish_settle += float64(v.Settle)
		}

	}
	resp.Code = api.Error_OK
	resp.Data = types.InSumResp{
		Total:          total,
		UnSettle:       unsettle,
		Finish:         finish,
		UnFinish:       unfinish,
		UnFinishSettle: unfinish_settle,
	}
	return resp
}
