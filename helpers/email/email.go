package email

import (
	"context"
	"github.com/OhMinsSup/story-server/helpers"
	mailgun2 "github.com/mailgun/mailgun-go/v4"
	"log"
	"time"
)

// 이메일 템플릿에 필요한 데이터
type EmailAuthTemplate struct {
	Keyword string `json:"keyword"`
	Url     string `json:"url"`
}

// mailgun에서 등록된 템플릿 이용해서 메일을 보낸다
func SendTemplateMessage(email string, template EmailAuthTemplate) (string, error) {
	apiKey := helpers.GetEnvWithKey("MAILGUN_API_KEY")
	domain := helpers.GetEnvWithKey("MAILGUN_DOMAIN_NAME")

	mailgun := mailgun2.NewMailgun(domain, apiKey)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	// Give time for template to show up in the system.
	time.Sleep(time.Second * 1)

	// Create a new message with template
	mail := mailgun.NewMessage("Veloss form Velog <mailgun@"+domain+">", "이메일 인증", "")
	mail.SetTemplate("velog-email")
	mail.AddRecipient(email)
	mail.AddVariable("keyword", template.Keyword)
	mail.AddVariable("url", template.Url)

	_, id, err := mailgun.Send(ctx, mail)
	if err != nil {
		panic(err)
	}

	log.Printf("Queued: %s", id)
	return id, err
}
