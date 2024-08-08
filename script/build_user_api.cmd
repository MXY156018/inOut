SET ROOT=%cd%
mkdir %ROOT%\..\bin\


cd ..\service\user\api && go build -ldflags="-w -s" -o user_api.exe && move user_api.exe %ROOT%\..\bin\  && cd %ROOT%
