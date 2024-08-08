package api

// 商户黑名单(冻结)
import (
	"net/http"

	"encoding/json"
	"mall-pkg/service/cache"
)

// 商户中黑名单中间件
//
// 从 redis 读取是否在黑名单
func MerchantBlackMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uid, ok := ctx.Value(Context_Key_UID).(int)
		if !ok {
			res := BaseResp{
				Code:   Error_InvalidToken,
				Msg:    "无效token",
				Reload: true,
			}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
		// 查找是否在很名单
		isBlack, err := cache.Ctx.Black.GetMerchant(uid)
		if err != nil {
			res := BaseResp{
				Code:   Error_InvalidToken,
				Msg:    err.Error(),
				Reload: true,
			}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}

		if isBlack {
			res := BaseResp{
				Code:   Error_InvalidToken,
				Msg:    "无效token",
				Reload: true,
			}
			resp, _ := json.Marshal(res)
			w.Write(resp)
			return
		}
		next(w, r)
	}
}
