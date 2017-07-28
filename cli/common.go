package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	bucket *oss.Bucket
	bkt    string
)

var log = logrus.New()

func errCheck(err error, s string) {

	if err != nil {
		log.Errorln(s, err)
		os.Exit(1)
	}

}

func initOSS() {

	var err error
	var accKeyID, accKeySec string
	var endpoint = "http://oss-cn-shenzhen.aliyuncs.com"

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/oss/")

	err = viper.ReadInConfig()
	errCheck(err, "Config file no find,init config failed!")

	accKeyID = viper.GetString("oss.accessKeyID")
	accKeySec = viper.GetString("oss.accessKeySecret")
	bkt = viper.GetString("oss.bucket")

	client, err := oss.New(endpoint, accKeyID, accKeySec)
	errCheck(err, "New oss instance error :")

	bucket, err = client.Bucket(bkt)
	errCheck(err, "Get Bucketerror :")

}

func putToOss(objKey, localPath string, index int, ulChan chan ulOss, errChan chan error) {

	err := bucket.PutObjectFromFile(objKey, localPath)

	if err != nil {
		errChan <- err
	} else {

		path := fmt.Sprintf("LocalPath:%s --> OssPath:%s", localPath, objKey)
		ulChan <- ulOss{path, index}
	}

}

func getFilelist(path string) []string {
	var paths []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		paths = append(paths, path)
		return nil
	})

	if err != nil {
		log.Errorf("filepath returned %v\n", err)
		return nil
	}

	return paths
}

func listObjs() {
	lsReg, err := bucket.ListObjects()
	errCheck(err, "Get Bucket List Objects error :")

	for i, v := range lsReg.Objects {
		fmt.Printf("%d %-60s \t %20v \t\t\t %8v\n", i, v.Key, v.LastModified, v.Size)
	}
}

func listObj(prefix string) {
	lsReg, err := bucket.ListObjects(oss.Prefix(prefix))
	errCheck(err, "Get Bucket List Objects error :")

	for i, v := range lsReg.Objects {
		fmt.Printf("%d %-60s \t %20v \t\t\t %8v\n", i, v.Key, v.LastModified, v.Size)
	}

}
