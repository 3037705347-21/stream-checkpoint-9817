package checkpoint

// Batch groups markers into a deterministic commit unit.
type Batch struct {
	Checkpoint CheckpointID
	Markers    []Marker
}

// Plan returns a commit batch only once every expected shard has acknowledged.
func Plan(snapshot Snapshot) (Batch, bool) {
	if !snapshot.Complete {
		return Batch{}, false
	}
	return Batch{Checkpoint: snapshot.ID, Markers: cloneMarkers(snapshot.Markers)}, true
}
