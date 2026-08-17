package checkpoint

import (
	"errors"
	"testing"
)

func TestOpenPreservesErrorAndCheckpointLifecycleContracts(t *testing.T) {
	coordinator := NewCoordinator()
	if _, err := coordinator.Open("", []string{"a"}); !errors.Is(err, ErrInvalidStream) { t.Fatalf("empty stream = %v", err) }
	first, err := coordinator.Open("metrics", []string{"a"})
	if err != nil { t.Fatalf("first open: %v", err) }
	second, err := coordinator.Open("metrics", []string{"b"})
	if err != nil { t.Fatalf("second open: %v", err) }
	if first.ID == second.ID || second.ID != "metrics-2" { t.Fatalf("ids = %q, %q", first.ID, second.ID) }
	if _, err := coordinator.Snapshot(first.ID); err != nil { t.Fatalf("first checkpoint disappeared: %v", err) }
}
