package checkpoint

import (
	"errors"
	"testing"
)

func TestPublicErrorsAndCheckpointIDsRemainStable(t *testing.T) {
	coordinator := NewCoordinator()
	for name, want := range map[string]error{
		"invalid stream": ErrInvalidStream,
		"no shards":      ErrNoShards,
	} {
		var err error
		if name == "invalid stream" {
			_, err = coordinator.Open("", []string{"a"})
		} else {
			_, err = coordinator.Open("metrics", nil)
		}
		if !errors.Is(err, want) {
			t.Fatalf("%s error = %v, want errors.Is(_, %v)", name, err, want)
		}
	}
	first, err := coordinator.Open("metrics", []string{"a"})
	if err != nil { t.Fatalf("open: %v", err) }
	if _, err := coordinator.Snapshot("missing"); !errors.Is(err, ErrUnknownCheckpoint) {
		t.Fatalf("unknown checkpoint = %v", err)
	}
	if _, err := coordinator.Acknowledge(first.ID, "other", 1); !errors.Is(err, ErrUnexpectedShard) {
		t.Fatalf("unexpected shard = %v", err)
	}
	second, err := coordinator.Open("metrics", []string{"b"})
	if err != nil { t.Fatalf("second open: %v", err) }
	if second.ID != "metrics-2" { t.Fatalf("checkpoint ID = %q, want metrics-2", second.ID) }
}
