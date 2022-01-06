package bios

import "testing"

func TestBIOS(t *testing.T) {
	bi, err := BIOS()
	if err != nil {
		t.Fatalf("failed to get bios information: %q", err)
	}
	if bi.Date == "" && bi.Vendor == "" && bi.Version == "" {
		t.Fatal("bios information is empty")
	}
}
