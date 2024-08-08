#!/bin/bash

echo "编译所有服务"
rootDir=`pwd`

echo ${rootDir}

binDir=${rootDir}/../bin
mkdir -p ${binDir}

serviceDir=${rootDir}/../service

echo ${serviceDir}


cd ../service/admin && go build -ldflags="-w -s" -o mall_admin && mv mall_admin ${binDir}   && cd ${rootDir}

cd ../service/product && go build -ldflags="-w -s" -o mall_product && mv mall_product ${binDir}   && cd ${rootDir}

cd ../service/user && go build -ldflags="-w -s" -o mall_user && mv mall_user ${binDir}   && cd ${rootDir}

cd ../service/order && go build -ldflags="-w -s" -o mall_order && mv mall_order ${binDir}   && cd ${rootDir}

cd ../service/message && go build -ldflags="-w -s" -o mall_message && mv mall_message ${binDir}   && cd ${rootDir}
