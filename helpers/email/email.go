package email

import (
	"bytes"
	"gopkg.in/gomail.v2"
	template2 "html/template"
)

type BindData struct {
	Keyword string
	Url     string
}

type Config struct {
	Host     string
	Port     string
	Addr     string
	Username string
	Password string
}

type EmailSender struct {
	conf     *Config
	template string
}

type Sender interface {
	Send(to string) error
	ParseTemplate(filepath string, data interface{}) error
}

// Send 이메일을 보내는 메소드
func (e *EmailSender) Send(to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.conf.Addr)
	//m.SetHeader("To", strings.Join(to[:], ","))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Velog 인증")
	m.SetBody("text/html", e.template)

	d := gomail.NewDialer("smtp.gmail.com", 587, e.conf.Username, e.conf.Password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// ParseTemplate html 템플릿을 읽고 리턴한다
func (e *EmailSender) ParseTemplate(filepath string, data interface{}) error {
	template, errTemp := template2.ParseFiles(filepath)
	if errTemp != nil {
		return errTemp
	}

	buf := new(bytes.Buffer)
	if err := template.Execute(buf, data); err != nil {
		return err
	}

	e.template = buf.String()
	return nil
}

// NewEmailSender 이메일 생성 및 보내는 구조체를 정의
func NewEmailSender(conf *Config, template string) Sender {
	return &EmailSender{conf: conf, template: template}
}

// SetupEmailCredentials 설정값 정의
func SetupEmailCredentials(host, port, senderAddr, username, password string) Config {
	return Config{
		host, port, senderAddr, username, password,
	}
}
