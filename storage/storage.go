package storage

import (
	"github.com/OhMinsSup/story-server/helpers"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"time"
)

var (
	AccessKeyID     string
	SecretAccessKey string
	Region          string
)

func Initialize() *session.Session {
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

type StorageRepository struct {
	sess *session.Session
}

func NewStorageRepository(sess *session.Session) *StorageRepository {
	return &StorageRepository{
		sess: sess,
	}
}

func (s *StorageRepository) GetS3PresignedUrl(bucket, key string, expiration time.Duration) (string, error) {
	// Create S3 service client
	svc := s3.New(s.sess)

	// Construct a GetObjectRequest request
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	// presignedUrl with expiration time
	presignedUrl, err := req.Presign(expiration * time.Minute)

	// Check if it can be signed or not
	if err != nil {
		return "", err
	}

	// Return the presigned url
	return presignedUrl, nil
}
