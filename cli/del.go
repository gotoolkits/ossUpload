package cli

func Del(objName string) {

	initOSS()
	err := bucket.DeleteObject(objName)

	errCheck(err, "del the object error!")

	log.Infoln("delete a object name:", objName)

}
