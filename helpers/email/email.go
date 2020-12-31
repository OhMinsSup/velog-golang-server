package email

import (
	"context"
	"github.com/OhMinsSup/story-server/helpers"
	mailgun2 "github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

// AuthTemplate 이메일 템플릿에 필요한 데이터
type AuthTemplate struct {
	Subject  string `json:"subject"`
	Template string `json:"template"`
	Keyword  string `json:"keyword"`
	Url      string `json:"url"`
}

// SendTemplateMessage mailgun에서 등록된 템플릿 이용해서 메일을 보낸다
func SendTemplateMessage(email string, template AuthTemplate) (string, error) {
	apiKey := helpers.GetEnvWithKey("MAILGUN_API_KEY")
	domain := helpers.GetEnvWithKey("MAILGUN_DOMAIN_NAME")

	mailgun := mailgun2.NewMailgun(domain, apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Give time for template to show up in the system.
	time.Sleep(time.Second * 1)

	// Create a new message with template
	mail := mailgun.NewMessage("Veloss form Velog <mailgun@"+domain+">", template.Subject, "")

	// 템플릿 타입 선택
	mail.SetTemplate(template.Template)
	// 받는 사람
	mail.AddRecipient(email)
	// 템플릿에 필요한 데이터 바인딩
	mail.AddVariable("keyword", template.Keyword)
	mail.AddVariable("url", template.Url)

	// 이메일 발송
	_, id, err := mailgun.Send(ctx, mail)
	if err != nil {
		return "", err
	}

	log.Println("Queued::" + id)
	return id, err
}
