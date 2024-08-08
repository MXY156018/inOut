SET ROOT=%cd%

cd %ROOT%\..\bin\

::start cmd /k
start cmd /C .\mall_admin.exe
timeout  /nobreak /t 1
::start cmd /C .\mall_upload.exe
::timeout  /nobreak /t 1

::start cmd /C .\mall_user.exe
::timeout  /nobreak /t 3
::
::start cmd /C .\mall_product.exe
::timeout  /nobreak /t 2
::
::
::start cmd /C .\mall_order.exe
::timeout  /nobreak /t 3
::
::start cmd /C .\mall_message.exe
::timeout  /nobreak /t 3


cd %ROOT%