package cli

import (
	"fmt"
)

func Del(objName string) {

	initOSS()
	err := bucket.DeleteObject(objName)

	errCheck(err, "del the object error!")

	fmt.Println("delete a object name:", objName)

}
