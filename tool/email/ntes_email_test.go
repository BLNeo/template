package email

import (
	"fmt"
	"testing"
)

func TestSendNTESEmail(t *testing.T) {
	in := &Config{
		Host:     "smtp.163.com",
		User:     "glgneo163@163.com",
		AuthCode: "QFCKEISRMWUFFJCV",
		Port:     25,
	}
	SetNTESEmailConfig(in)

	fmt.Println(emailConfig)

	emailTo := []string{"1072511881@qq.com"}
	emailCc := []string{}
	subject := "xxx图书馆验证码"
	body := "验证码：109721"
	fileName := ""
	if err := SendNTESEmail(emailTo, emailCc, subject, body, fileName); err != nil {
		t.Error(err)
	}

}
