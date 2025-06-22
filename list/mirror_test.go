package list

import (
	"testing"
)

func TestMirrorList(t *testing.T) {
	for _, source := range MirrorSources {
		if source.HomeURL == "" {
			t.Error("Unexpected empty HomeUrl for MirrorSources", "Name",
				source.Name)
		}
	}

	for _, site := range MirrorSites {
		if site.HomeURL == "" {
			t.Error("Unexpected empty HomeUrl for MirrorSites", "Name",
				site.Name)
		}
	}
}
