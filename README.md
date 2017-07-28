# ossUpload
`使用阿里OSS-SDK基础上开发的简单工具，可用于OSS的文件或目录上传/List/Del.`

## 安装

 >  依赖go环境，请确保go环境安装正常
    
     1) go get github.com/gotoolkits/ossUpload
     2) cd $GOPATH/github.com/gotoolkits/
        go install ossUpload.go

## 命令操作

 > Help 帮助

     ./ossUpload --help
     aliyun OSS Upload file tool, 
    You Can Upload/List/Del objects in the bucket.

     Usage:
       ossUpload [flags]
       ossUpload [command]

     Available Commands:
       del         Del a object
       help        Help about any command
       list        List the object in bucket

     Flags:
       -d, --destination path string   Specify the oss directories
       -h, --help                      help for ossUpload
       -o, --object name string        Specify the object name or prefix  in  bucket
       -r, --resource path string      Specify the resource files or directories

     Use "ossUpload [command] --help" for more information about a command.

 >  上传目录(必须参数: `-r 本地目录 -d 对象路径`   )

       ./ossUpload -r /tmp/aaa/ -d samples/

>  查询（`可选`参数: `-o 对象名称前缀`）

        1) //查询Bucket下所以objects
           ./ossUpload list
        2) //过滤查询，指定关键字前缀
          ./ossUpload list -o prefix_


>  删除（必须参数: `-o 对象名称`）

        ./ossUpload del -o samples/tt.txt


## 配置文件

>  默认查找路径：
>         1.   当前目录config.json (与执行文件同目录) `优先`
>         2.   /etc/oss/config.json

    {
      "oss":{
        "endPoint":"http://oss-cn-shenzhen.aliyuncs.com",
        "bucket":"********",
        "accessKeyID":"*************",
        "accessKeySecret":"******************************"
       }
    }

## Todo-Finish
 >   1) goroutine 实现多线程并行上传 ［已完成］
 >   2) 日志标准化 ［已完成］
 
     ERRO[0000] Upload error happen: open /tmp/com.apple.launchd.4gkZ1Fghle/org.macosforge.xquartz:0: operation not supported on socket 
     ERRO[0000] Upload error happen: open /tmp/com.apple.launchd.Aar8VHjSaL/Listeners: operation not supported on socket 
     ERRO[0000] Upload error happen: open /tmp/com.apple.launchd.IN2K6pG99U/Render: operation not supported on socket 
     INFO[0000] 26: LocalPath:/tmp/wifi-OBZk1X.log --> OssPath:yuch/wifi-OBZk1X.log Upload sucessful! 
     INFO[0000] 31: LocalPath:/tmp/wifi-XgtVzJ.log --> OssPath:yuch/wifi-XgtVzJ.log Upload sucessful! 
     INFO[0000] 5: LocalPath:/tmp/aaa/bb/cc.log --> OssPath:yuch/aaa/bb/cc.log Upload sucessful! 
     INFO[0003] 37: LocalPath:/tmp/wifi-gICNAD.log --> OssPath:yuch/wifi-gICNAD.log Upload sucessful! 
     INFO[0002] 19: LocalPath:/tmp/wifi-1EPROn.log --> OssPath:yuch/wifi-1EPROn.log Upload sucessful! 
     INFO[0009] 20: LocalPath:/tmp/wifi-4vVbMA.log --> OssPath:yuch/wifi-4vVbMA.log Upload sucessful! 
     INFO[0009] 21: LocalPath:/tmp/wifi-AWCdIW.log --> OssPath:yuch/wifi-AWCdIW.log Upload sucessful! 

