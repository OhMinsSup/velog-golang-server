package storage

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

var (
	AccessKeyID     string
	SecretAccessKey string
	Region          string
)

func ConnectionByAws() *session.Session {
	AccessKeyID = helpers.GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey = helpers.GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	Region = helpers.GetEnvWithKey("AWS_REGION")

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(Region),
			Credentials: credentials.NewStaticCredentials(
				AccessKeyID,
				SecretAccessKey,
				"", // a token will be created when the session it's used.
			),
		})

	if err != nil {
		panic(err)
	}
	return sess
}

func Inject(sess *session.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	}
}
