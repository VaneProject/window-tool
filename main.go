package main

import (
	"window-tool/regedit"
)

func errorLog(err error) {
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
}

func main() {
	_, err := regedit.GetUninstall()
	errorLog(err)

	//k, err := window-tool.OpenKey(window-tool.LOCAL_MACHINE, UNINSTALL1, window-tool.ENUMERATE_SUB_KEYS|window-tool.QUERY_VALUE)
	//errorLog(err)
	//defer errorLog(k.Close())
	//
	//info, err := k.ReadSubKeyNames(-1)
	//errorLog(err)
	//fmt.Println(info)

	//keyNames, err := k.ReadSubKeyNames(0)
	//errorLog(err)
	//for _, name := range keyNames {
	//	fmt.Println(name)
	//}
}
