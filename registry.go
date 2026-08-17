package checkpoint

import "sync"

type storedCheckpoint struct {
	id       CheckpointID
	stream   string
	expected map[string]struct{}
	markers  map[string]int64
}

func (stored *storedCheckpoint) snapshot() Snapshot {
	expected := make([]string, 0, len(stored.expected))
	for shard := range stored.expected {
		expected = append(expected, shard)
	}
	markers := make([]Marker, 0, len(stored.markers))
	for shard, sequence := range stored.markers {
		markers = append(markers, Marker{Shard: shard, Sequence: sequence})
	}
	return Snapshot{
		ID:       stored.id,
		Stream:   stored.stream,
		Expected: cloneStrings(expected),
		Markers:  cloneMarkers(markers),
		Complete: len(stored.expected) == len(stored.markers),
	}
}

type registry struct {
	mu          sync.RWMutex
	checkpoints map[CheckpointID]*storedCheckpoint
	next        uint64
}

func newRegistry() *registry {
	return &registry{checkpoints: make(map[CheckpointID]*storedCheckpoint)}
}

func (r *registry) create(stream string, shards []string) Snapshot {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.next++
	id := CheckpointID(stream + "-" + itoa(r.next))
	expected := make(map[string]struct{}, len(shards))
	for _, shard := range shards {
		expected[shard] = struct{}{}
	}
	stored := &storedCheckpoint{id: id, stream: stream, expected: expected, markers: make(map[string]int64)}
	r.checkpoints[id] = stored
	return stored.snapshot()
}

func (r *registry) acknowledge(id CheckpointID, marker Marker) (Snapshot, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	stored, ok := r.checkpoints[id]
	if !ok {
		return Snapshot{}, ErrUnknownCheckpoint
	}
	if _, ok := stored.expected[marker.Shard]; !ok {
		return Snapshot{}, ErrUnexpectedShard
	}
	if previous, ok := stored.markers[marker.Shard]; ok && marker.Sequence < previous {
		return Snapshot{}, ErrSequenceRegression
	}
	stored.markers[marker.Shard] = marker.Sequence
	return stored.snapshot(), nil
}

func (r *registry) get(id CheckpointID) (Snapshot, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	stored, ok := r.checkpoints[id]
	if !ok {
		return Snapshot{}, ErrUnknownCheckpoint
	}
	return stored.snapshot(), nil
}
