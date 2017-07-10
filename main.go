package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/go-gomail/gomail"
)

func QQCNT() {
	m := gomail.NewMessage()
	m.SetHeader("From", "big brother <727266990@qq.com>")
	m.SetHeader("To", "1852951552@qq.com")
	//	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "请注意")
	m.SetBody("text/html", "明天下午在大会议室见面")
	f, err := os.Create("重要的文件呦.txt")
	if err != nil {
		panic(err)
	}

	defer func(name string) {
		time.Sleep(time.Second * 10)
		os.Remove(name)
	}(f.Name())

	defer f.Close()
	f.WriteString("这是一个不能说的秘密")
	h := make(map[string][]string, 0)
	h["Content-Type"] = []string{`application/octet-stream; charset=utf-8; name="` + f.Name() + `"`} //要设置这个，否则中文会乱码
	//	h["Content-Type"] = []string{`text/plain; name="` + f.Name() + `"`}
	fileSetting := gomail.SetHeader(h)
	m.Attach(f.Name(), fileSetting)
	d := gomail.NewDialer("smtp.qq.com", 587, "727266990@qq.com", "password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	fmt.Println("================")
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	fmt.Println("结束")
}

func main() {
	QQCNT()
}
