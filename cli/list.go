package cli

import (
	"fmt"
)

func List(pfx string) {
	initOSS()

	fmt.Println("List all object in bucket:", bkt)
	if pfx == "" {
		listObjs()
	} else {
		listObj(pfx)
	}

}
