# 修复前故障复现（Docker）

## 项目与标准命令

模块为 `example.com/stream-checkpoint`，验证命令为 `go test -count=20 ./...`。

## 环境构建与编译

使用当前平台的 `golang:1.22` 镜像构建成功，容器内 `go build ./...` 成功。

## 故障触发步骤

在容器工作目录执行 `go test -count=20 ./...`。

## 实际错误输出

测试退出码为 1；空流校验返回 `open rejected: stream rejected: stream must not be empty`，并且错误链校验失败，检查点编号为 `metrics-3` 而非 `metrics-2`。

## 期望行为

调用方应能使用 `errors.Is` 识别公开错误，连续创建的检查点编号应连续递增。
