package list

import (
	"testing"
)

func TestMirrorList(t *testing.T) {
	for _, source := range MirrorSources {
		if source.HomeUrl == "" {
			t.Error("Unexpected empty HomeUrl for MirrorSources", "Name",
				source.Name)
		}
	}
	for _, source := range MirrorSources {
		if source.HomeUrl == "" {
			t.Error("Unexpected empty HomeUrl for MirrorSources", "Name",
				source.Name)
		}
	}
}
