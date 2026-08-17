package checkpoint

// CheckpointID identifies a single coordination round.
type CheckpointID string

// Marker records the sequence that a shard has durably processed.
type Marker struct {
	Shard    string
	Sequence int64
}

// Snapshot is the public, immutable view of a checkpoint.
type Snapshot struct {
	ID       CheckpointID
	Stream   string
	Expected []string
	Markers  []Marker
	Complete bool
}

func cloneStrings(values []string) []string {
	cloned := append([]string(nil), values...)
	return cloned
}

func cloneMarkers(values []Marker) []Marker {
	cloned := append([]Marker(nil), values...)
	return cloned
}
