#!/bin/bash


echo "停止所有服务"

killall mall_admin
# killall mall_product
# killall mall_user
# killall mall_order
# killall mall_message

# sleep 3

ps -aux | grep mall_admin
# ps -aux | grep mall_product
# ps -aux | grep mall_user
# ps -aux | grep mall_order
# ps -aux | grep mall_message
