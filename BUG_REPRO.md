# 修复前故障复现（Docker）

## 项目与标准命令

模块为 `example.com/stream-checkpoint`，验证命令为 `go test -count=20 ./...`。

## 环境构建与编译

使用当前平台的 `golang:1.22` 镜像构建成功，容器内 `go build ./...` 成功。

## 故障触发步骤

在容器工作目录执行 `go test -count=20 ./...`。

## 实际错误输出

测试退出码为 1，第一次打开带分片的检查点即触发 `panic: assignment to entry in nil map`。

## 期望行为

新协调器应能创建检查点、记录分片确认并生成完整提交批次，不应出现 nil 状态异常。
