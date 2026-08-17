# 修复前故障复现（Docker）

## 项目与标准命令

模块为 `example.com/stream-checkpoint`，验证命令为 `go test -count=20 ./...`。

## 环境构建与编译

使用当前平台的 `golang:1.22` 镜像构建成功，容器内 `go build ./...` 成功。

## 故障触发步骤

在容器工作目录执行 `go test -count=20 ./...`。

## 实际错误输出

测试退出码为 1；空流输入得到 `at least one shard is required`，与调用方期望的输入错误类别不一致。

## 期望行为

输入错误应保持准确类别，连续打开的检查点应有不同且连续的标识，并保留已有状态。
