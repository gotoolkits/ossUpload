package cli

import (
	"fmt"
	"strings"
	"time"
)

type ulOss struct {
	path  string
	index int
}

func UploadToOSS(res, des string) {

	initOSS()

	fpList := getFilelist(res)

	ulChan := make(chan ulOss, 100)
	errChan := make(chan error)

	fmt.Println("Get local file queue size :", len(fpList))

	for i, lp := range fpList {

		sufPath := strings.Split(lp, res)

		keyName := fmt.Sprintf("%s%s", des, sufPath[1])

		// add go routine upload the files
		go putToOss(keyName, lp, i, ulChan, errChan)

	}

loop1:
	for {
		select {
		case <-time.After(6 * time.Second):
			break loop1
		case errc := <-errChan:
			log.Errorln("Upload error happen:", errc)
		case ulc := <-ulChan:
			log.Infof("%d: %s Upload sucessful!", ulc.index+1, ulc.path)
		}
	}
}
