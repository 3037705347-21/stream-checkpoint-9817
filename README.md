# Stream Checkpoint

Stream Checkpoint is an in-memory Go library for coordinating durable progress
markers in stream processing workers. It models checkpoint creation, shard
acknowledgements, deterministic batch planning, and immutable status snapshots.

## Layout

- `checkpoint.go`: public coordinator and checkpoint lifecycle.
- `model.go`: identifiers, marker values, and snapshots.
- `registry.go`: checkpoint storage and acknowledgement tracking.
- `planner.go`: deterministic work-batch planning.
- `validate.go`: input and lifecycle validation.

## Use

The package is intended to be embedded by stream workers:

```go
coordinator := checkpoint.NewCoordinator()
checkpoint, err := coordinator.Open("deploy-42", []string{"edge-a", "edge-b"})
if err != nil { /* handle input */ }
_, err = coordinator.Acknowledge(checkpoint.ID, "edge-a", 120)
```

Run the library checks with:

```text
go build ./...
go test ./...
```
