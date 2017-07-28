package main

import (
	"fmt"
	"os"

	"github.com/gotoolkits/aliyunOssUpload/cli"
	"github.com/spf13/cobra"
)

var res, des, obj string
var RootCmd = &cobra.Command{
	Use:   "ossUpload",
	Short: "aliyun OSS Upload file tool",
	Long: `aliyun OSS Upload file tool, 
You Can Upload/List/Del objects in the bucket.`,
	Run: func(cmd *cobra.Command, args []string) {
		if res == "" && des == "" {
			cmd.Usage()
		} else {

			cli.UploadToOSS(res, des)

		}

	},
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List the object in bucket",
	Long: `aliyun OSS Upload file tool,
	List the object in bucket. Can setting filter object prefix arg : " -o  xy/ "`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cli.List(obj)
	},
}

var DelCmd = &cobra.Command{
	Use:   "del",
	Short: "Del a object",
	Long: `aliyun OSS Upload file tool,
del a object in bucket. Need to set arg: " -o xy/test.file "`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff
		if obj == "" {
			cmd.Usage()
		}

		cli.Del(obj)
	},
}

func main() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&res, "resource path", "r", "", "Specify the resource files or directories")
	RootCmd.PersistentFlags().StringVarP(&des, "destination path", "d", "", "Specify the oss directories")
	RootCmd.PersistentFlags().StringVarP(&obj, "object name", "o", "", "Specify the object name or prefix  in  bucket")
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(DelCmd)
}
