SET ROOT=%cd%
mkdir %ROOT%\..\bin\

cd ..\service\upload\api && go build -ldflags="-w -s" -o mall_upload.exe && move mall_upload.exe %ROOT%\..\bin\  && cd %ROOT%
