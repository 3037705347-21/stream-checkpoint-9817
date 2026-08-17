# stream-checkpoint__004 Docker 交付说明

## 项目概览
- Stream Checkpoint is an in-memory Go library for coordinating durable progress
- Go module: `example.com/stream-checkpoint`

## 标准命令

```bash
go build ./...
go test ./...
```

## Docker 构建

```bash
./build_benzhi_docker.sh stream-checkpoint__004-benzhi linux/amd64
docker run --rm -it stream-checkpoint__004-benzhi bash
```

## 环境

- 基础镜像: `golang:1.22`
- 依赖在镜像构建阶段预下载，容器内可直接执行 Go 构建和测试命令。
- 代码目录: `/app`
