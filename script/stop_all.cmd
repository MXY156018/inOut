:: admin 服务
taskkill /fi "imagename  eq mall_admin.exe" /f
:: 上传服务
:: taskkill /fi "imagename  eq mall_upload.exe" /f
:: 商品服务
taskkill /fi "imagename  eq mall_product.exe" /f
:: 用户服务
taskkill /fi "imagename  eq mall_user.exe" /f
:: 订单服务
taskkill /fi "imagename  eq mall_order.exe" /f
:: 消息服务
taskkill /fi "imagename  eq mall_message.exe" /f

