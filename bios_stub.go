// +build !windows,!linux

package bios

func bios() (*BiosInfo, error) {
	return nil, newErrNotImplemented()
}
