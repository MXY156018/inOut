SET ROOT=%cd%
mkdir %ROOT%\..\bin\



:: 管理员核心模块 (API/RPC)
cd ..\service\admin && go build -ldflags="-w -s" -o mall_admin.exe && move mall_admin.exe %ROOT%\..\bin\  && cd %ROOT%
::cd ..\service\product && go build -ldflags="-w -s" -o mall_product.exe && move mall_product.exe %ROOT%\..\bin\  && cd %ROOT%
::cd ..\service\upload\api && go build -ldflags="-w -s" -o mall_upload.exe && move mall_upload.exe %ROOT%\..\bin\  && cd %ROOT%
::cd ..\service\user && go build -ldflags="-w -s" -o mall_user.exe && move mall_user.exe %ROOT%\..\bin\  && cd %ROOT%
::cd ..\service\order && go build -ldflags="-w -s" -o mall_order.exe && move mall_order.exe %ROOT%\..\bin\  && cd %ROOT%
::cd ..\service\message && go build -ldflags="-w -s" -o mall_message.exe && move mall_message.exe %ROOT%\..\bin\  && cd %ROOT%

cd %ROOT%