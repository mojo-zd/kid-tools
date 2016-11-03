package main

import (
	"flag"
	"fmt"
	"kid-tools/config"
	"math/rand"
	"time"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	glog.CopyStandardLogTo("INFO")

	certFile := flag.Arg(0)
	keyFile := flag.Arg(1)

	if len(certFile) == 0 || len(keyFile) == 0 {
		fmt.Println("请输入公钥和私钥的文件路径!")
		return
	}
	pk, _, err := config.LoadCertAndKey(certFile, keyFile)
	if err != nil {
		fmt.Println("生成公钥和私钥文件出错 %s", err)
		return
	}
	fmt.Println(pk.KeyID())
}
