package checkpoint

import "testing"

func TestOpeningAnotherCheckpointRetainsEarlierState(t *testing.T) {
	coordinator := NewCoordinator()
	first, err := coordinator.Open("metrics", []string{"a", "b"})
	if err != nil { t.Fatalf("first open: %v", err) }
	second, err := coordinator.Open("metrics", []string{"c"})
	if err != nil { t.Fatalf("second open: %v", err) }
	if first.ID == second.ID { t.Fatalf("checkpoint IDs collided: %q", first.ID) }
	if _, err := coordinator.Snapshot(first.ID); err != nil { t.Fatalf("first checkpoint disappeared: %v", err) }
	first.Expected[0] = "changed"
	stored, err := coordinator.Snapshot(first.ID)
	if err != nil { t.Fatalf("snapshot: %v", err) }
	if stored.Expected[0] == "changed" { t.Fatalf("snapshot shares expected shards: %#v", stored.Expected) }
}
