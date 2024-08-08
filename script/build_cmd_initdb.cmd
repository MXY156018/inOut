SET ROOT=%cd%
mkdir %ROOT%\..\bin\

cd ..\service\admin\cmd\initdb && go build -ldflags="-w -s" -o initdb.exe && move initdb.exe %ROOT%\..\bin\  && cd %ROOT%
