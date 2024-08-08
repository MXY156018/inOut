SET ROOT=%cd%
mkdir %ROOT%\..\bin\


:: 商品服务 (API/RPC)
cd ..\service\product && go build -ldflags="-w -s" -o mall_product.exe && move mall_product.exe %ROOT%\..\bin\  && cd %ROOT%
