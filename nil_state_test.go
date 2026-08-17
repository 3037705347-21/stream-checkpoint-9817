package checkpoint

import "testing"

func TestCoordinatorInitializesAllCheckpointState(t *testing.T) {
	coordinator := NewCoordinator()
	opened, err := coordinator.Open("metrics", []string{"a", "b"})
	if err != nil { t.Fatalf("open: %v", err) }
	if opened.ID == "" || len(opened.Expected) != 2 { t.Fatalf("opened snapshot = %#v", opened) }
	if _, err := coordinator.Acknowledge(opened.ID, "a", 1); err != nil { t.Fatalf("first acknowledge: %v", err) }
	completed, err := coordinator.Acknowledge(opened.ID, "b", 2)
	if err != nil { t.Fatalf("second acknowledge: %v", err) }
	batch, ok := Plan(completed)
	if !ok || len(batch.Markers) != 2 { t.Fatalf("batch = %#v, ready = %v", batch, ok) }
}
