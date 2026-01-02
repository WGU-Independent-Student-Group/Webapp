package data

import "testing"

func TestDepIDUnique(t *testing.T) {
	metadata, _, _, _, err := Load("mrds.csv")
	if err != nil {
		t.Fatal(err)
	}

	exists := make(map[int64]struct{})
	for _, m := range metadata {
		if _, ok := exists[m.MetadataID]; ok {
			t.Fatalf("duplicate DepID found: %d", m.MetadataID)
		}
		exists[m.MetadataID] = struct{}{}
	}
}
