package bios

import (
	"github.com/StackExchange/wmi"
)

const wqlBIOS = "SELECT InstallDate, Manufacturer, Version FROM CIM_BIOSElement"

type win32BIOS struct {
	InstallDate  *string
	Manufacturer *string
	Version      *string
}

func bios() (*BiosInfo, error) {
	// Getting data from WMI
	var win32BIOSDescriptions []win32BIOS
	if err := wmi.Query(wqlBIOS, &win32BIOSDescriptions); err != nil {
		return nil, err
	}
	bi := BiosInfo{}
	if len(win32BIOSDescriptions) > 0 {
		bi.Vendor = *win32BIOSDescriptions[0].Manufacturer
		bi.Version = *win32BIOSDescriptions[0].Version
		bi.Date = *win32BIOSDescriptions[0].InstallDate
	}
	return &bi, nil
}
