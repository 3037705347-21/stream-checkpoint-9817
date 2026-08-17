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
	markers := cloneMarkers(snapshot.Markers)
	for left, right := 0, len(markers)-1; left < right; left, right = left+1, right-1 {
		markers[left], markers[right] = markers[right], markers[left]
	}
	return Batch{Checkpoint: snapshot.ID, Markers: markers}, true
}
