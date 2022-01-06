package bios

import (
	"fmt"
	"runtime"
)

type BiosInfo struct {
	Version string
	Vendor  string
	Date    string
}

func BIOS() (*BiosInfo, error) {
	bi, err := bios()
	if err != nil {
		return nil, err
	} // if
	return &BiosInfo{
		Version: bi.Version,
		Vendor:  bi.Vendor,
		Date:    bi.Date,
	}, nil
}

func newErrNotImplemented() error {
	return ErrNotImplemented{runtime.GOOS, runtime.GOARCH}
}

type ErrNotImplemented struct {
	GOOS   string
	GOARCH string
}

func (e ErrNotImplemented) Error() string {
	return fmt.Sprintf("not implemented on %s-%s", e.GOOS, e.GOARCH)
}
