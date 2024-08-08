#!/bin/bash

echo '启动服务'

cd ../bin

rs=`ps -ef | grep etcd | grep -v grep | wc -l`
if [[ "$rs" == "0" ]];then
    echo "未启动 etcd,启动 etcd"
    nohup etcd > etcd.out 2>&1 &
fi

sleep 3
ps -aux | grep etcd

echo '启动管理员服务'
nohup ./mall_admin > mall_admin.out 2>&1 &
sleep 3

echo '启动用户服务'
nohup ./mall_user > mall_user.out 2>&1 &
sleep 3

echo '启动商品服务'
nohup ./mall_product > mall_product.out 2>&1 &
sleep 3

echo '启动订单服务'
nohup ./mall_order > mall_order.out 2>&1 &
sleep 3

echo '启动消息服务'
nohup ./mall_message > mall_message.out 2>&1 &
sleep 3

ps -aux | grep mall_admin
ps -aux | grep mall_product
ps -aux | grep mall_user
ps -aux | grep mall_order
ps -aux | grep mall_message

cd -