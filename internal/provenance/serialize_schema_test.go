package provenance

import (
	"strings"
	"testing"
)

func TestUnmarshalRejectsAnUnsupportedSchema(t *testing.T) {
	data, err := Marshal(&Snapshot{Schema: "9.9.9"})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	_, err = Unmarshal(data)
	if err == nil || !strings.Contains(err.Error(), "unsupported schema") {
		t.Fatalf("want an unsupported-schema error, got %v", err)
	}
}
