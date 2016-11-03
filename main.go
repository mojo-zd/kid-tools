package main

import (
	"flag"
	"fmt"
	"kid-tools/config"
	"math/rand"
	"time"

	"strings"

	"github.com/golang/glog"
)

const (
	CERT_SUFFIX = ".crt"
	KEY_SUFFIX  = ".key"
)

var (
	certfile, keyfile string
)

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	glog.CopyStandardLogTo("INFO")

	if len(flag.Args()) == 0 {
		fmt.Println("请输入公钥和私钥的文件路径!")
		return
	}

	for _, ele := range flag.Args() {
		if strings.HasSuffix(ele, CERT_SUFFIX) {
			certfile = ele
		} else if strings.HasSuffix(ele, KEY_SUFFIX) {
			keyfile = ele
		}
	}

	if len(certfile) == 0 || len(keyfile) == 0 {
		fmt.Println("请输入公钥和私钥的文件路径!")
		return
	}

	pk, _, err := config.LoadCertAndKey(certfile, keyfile)
	if err != nil {
		fmt.Println("生成公钥和私钥文件出错 %s", err)
		return
	}

	fmt.Println(pk.KeyID())
}
