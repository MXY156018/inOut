package types

type OutRecordReq struct {
	Records string `json:"records"`
}

type GetRecordsReq struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Sort      string `json:"sort"`
}
type InRecordReq struct {
	Type     string  `json:"type"`
	TypeId   int32   `json:"type_id"`
	Weight   float32 `json:"weight"`
	Price    float32 `json:"price"`
	Sum      float32 `json:"sum"`
	IsSettle int8    `json:"is_settle"`
	Settle   float32 `json:"settle"`
}

type TypeDetailReq struct {
	Id        int8   `json:"id"`
	Page      int    `json:"page"`
	PageSize  int    `json:"pageSize"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Settle    int    `json:"settle"`
}
