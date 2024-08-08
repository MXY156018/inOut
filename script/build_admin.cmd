SET ROOT=%cd%
mkdir %ROOT%\..\bin\


:: 管理员核心模块 (API/RPC)
cd ..\service\admin && go build -ldflags="-w -s" -o mall_admin.exe && move mall_admin.exe %ROOT%\..\bin\  && cd %ROOT%
