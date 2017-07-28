package cli

import (
	"fmt"
	"strings"
)

func UploadToOSS(res, des string) {

	initOSS()

	fpList := getFilelist(res)

	fmt.Println("Get local file queue size :", len(fpList))

	for _, lp := range fpList {

		sufPath := strings.Split(lp, res)

		keyName := fmt.Sprintf("%s%s", des, sufPath[1])
		fmt.Println(lp, res, keyName)

		putToOss(keyName, lp)

		fmt.Println("Upload to OSS:", keyName)

	}

}
