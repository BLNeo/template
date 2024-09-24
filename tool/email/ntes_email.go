package email

import (
	"gopkg.in/gomail.v2"
	"os"
	"sync"
)

var (
	once        sync.Once
	emailConfig *Config
)

func SetNTESEmailConfig(conf *Config) {
	once.Do(func() {
		emailConfig = conf
	})
}

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	AuthCode string `yaml:"authCode"`
	Port     int    `yaml:"port"`
}

func SendNTESEmail(emailTo []string, emailCc []string, subject, body, fileName string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", emailConfig.User) // 发起人
	m.SetHeader("To", emailTo...)         //发送给多个用户
	if len(emailCc) > 0 {
		m.SetHeader("Cc", emailCc...) //抄送给多个用户
	}
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文
	// 有带附件
	if fileName != "" {
		defer os.Remove(fileName)
		m.Attach(fileName)
	}

	d := gomail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.User, emailConfig.AuthCode)
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil
}
