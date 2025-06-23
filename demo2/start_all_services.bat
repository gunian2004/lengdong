@echo off
echo 正在启动所有服务...

REM 设置工作目录为脚本所在目录
cd /d %~dp0

REM 启动用户服务
start "user-service" /D services\user-service go run main.go

REM 启动后台管理服务
start "admin-service" /D services\admin-service go run main.go

REM 启动区块链服务
start "blockchain-service" /D services\blockchain-service go run main.go

REM 启动查询服务
start "query-service" /D services\query-service go run main.go

REM 启动工厂服务
start "factory-service" /D services\factory-service go run main.go

REM 启动经销商服务
start "saler-service" /D services\saler-service go run main.go

echo 所有服务已启动！
pause