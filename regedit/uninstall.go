package regedit

import (
	"golang.org/x/sys/windows/registry"
	"golang.org/x/text/encoding/korean"
	"log"
)

const (
	HKEY_LOCAL_MACHINE = "HKEY_LOCAL_MACHINE"
	UNINSTALL_PATH1    = `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`
	UNINSTALL_PATH2    = `SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`
)

func GetUninstall() ([]string, error) {
	const si = len(HKEY_LOCAL_MACHINE) + 1

	decoder := korean.EUCKR.NewDecoder()
	paths := make([]string, 0)
	if data, err := command(paths, HKEY_LOCAL_MACHINE+"\\"+UNINSTALL_PATH1, si, decoder); err != nil {
		return nil, err
	} else {
		paths = data
	}
	if data, err := command(paths, HKEY_LOCAL_MACHINE+"\\"+UNINSTALL_PATH2, si, decoder); err != nil {
		return nil, err
	} else {
		paths = data
	}

	for _, path := range paths {
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE)
		if err != nil {
			return nil, err
		} else {
			err := k.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return paths, nil
}
