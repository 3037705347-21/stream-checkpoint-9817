package checkpoint

import (
	"errors"
	"reflect"
	"testing"
)

func TestCheckpointLifecycleProducesSortedBatch(t *testing.T) {
	coordinator := NewCoordinator()
	opened, err := coordinator.Open("metrics", []string{"edge-b", "edge-a"})
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if opened.Complete {
		t.Fatal("new checkpoint was complete")
	}
	if _, err := coordinator.Acknowledge(opened.ID, "edge-b", 11); err != nil {
		t.Fatalf("acknowledge edge-b: %v", err)
	}
	completed, err := coordinator.Acknowledge(opened.ID, "edge-a", 7)
	if err != nil {
		t.Fatalf("acknowledge edge-a: %v", err)
	}
	batch, ok := Plan(completed)
	if !ok {
		t.Fatal("complete checkpoint did not produce a batch")
	}
	want := []Marker{{Shard: "edge-a", Sequence: 7}, {Shard: "edge-b", Sequence: 11}}
	if !reflect.DeepEqual(batch.Markers, want) {
		t.Fatalf("markers = %#v, want %#v", batch.Markers, want)
	}
}

func TestSnapshotDoesNotShareState(t *testing.T) {
	coordinator := NewCoordinator()
	opened, err := coordinator.Open("audit", []string{"a", "b"})
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	opened.Expected[0] = "mutated"
	snapshot, err := coordinator.Snapshot(opened.ID)
	if err != nil {
		t.Fatalf("snapshot: %v", err)
	}
	if snapshot.Expected[0] != "a" {
		t.Fatalf("snapshot exposed shared state: %#v", snapshot.Expected)
	}
}

func TestValidationAndMonotonicMarkers(t *testing.T) {
	coordinator := NewCoordinator()
	if _, err := coordinator.Open("", []string{"a"}); !errors.Is(err, ErrInvalidStream) {
		t.Fatalf("empty stream error = %v", err)
	}
	opened, err := coordinator.Open("audit", []string{"a"})
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	if _, err := coordinator.Acknowledge(opened.ID, "a", 4); err != nil {
		t.Fatalf("first ack: %v", err)
	}
	if _, err := coordinator.Acknowledge(opened.ID, "a", 3); !errors.Is(err, ErrSequenceRegression) {
		t.Fatalf("regression error = %v", err)
	}
}
