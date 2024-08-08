package privilege

// import (
// 	"context"
// 	"fmt"
// 	"mall-pkg/api"
// 	"mall-pkg/service/cache"
// )

// // 检查管理员API调用这是否是商户
// //
// // ctx go-zero api context
// //
// // 0, nil 不是商户
// //
// // 非0,nil 是商户
// func IsMerchant(ctx context.Context) (int, error) {
// 	authorityId := ctx.Value(api.Middle_Header_AuthorityId).(string)
// 	isMerchat := authorityId == api.Merchant_AuthorityId
// 	if !isMerchat {
// 		return 0, nil
// 	}

// 	uid := ctx.Value(api.Context_Key_UID).(int)
// 	id, err := cache.Ctx.MerchantId.Get(uid)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if id <= 0 {
// 		return 0, fmt.Errorf("找不到管理员 %d 对应的商户ID", uid)
// 	}
// 	return id, nil
// }
