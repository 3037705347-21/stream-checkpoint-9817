package checkpoint

// Coordinator owns independent checkpoint lifecycles.
type Coordinator struct {
	registry *registry
}

// NewCoordinator returns a coordinator with no checkpoints.
func NewCoordinator() *Coordinator {
	return &Coordinator{registry: newRegistry()}
}

// Open starts a checkpoint for every supplied shard.
func (c *Coordinator) Open(stream string, shards []string) (Snapshot, error) {
	if err := validateOpen(stream, shards); err != nil {
		return Snapshot{}, err
	}
	return c.registry.create(stream, shards), nil
}

// Acknowledge records a shard's latest durable marker.
func (c *Coordinator) Acknowledge(id CheckpointID, shard string, sequence int64) (Snapshot, error) {
	marker := Marker{Shard: shard, Sequence: sequence}
	if err := validateMarker(marker); err != nil {
		return Snapshot{}, err
	}
	return c.registry.acknowledge(id, marker)
}

// Snapshot returns the current state without exposing mutable storage.
func (c *Coordinator) Snapshot(id CheckpointID) (Snapshot, error) {
	return c.registry.get(id)
}
