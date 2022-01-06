package bios

import (
	"os"
	"path/filepath"
	"strings"
)

func bios() (*BiosInfo, error) {
	bi := BiosInfo{}
	bi.Version = bestEffortReadLinuxDMI("bios_version")
	bi.Vendor = bestEffortReadLinuxDMI("bios_vendor")
	bi.Date = bestEffortReadLinuxDMI("bios_date")
	return &bi, nil
}

func bestEffortReadLinuxDMI(value string) string {
	path := filepath.Join("/", "sys", "class", "dmi", "id", value)
	b, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(b))
}
