package aws

import (
	"github.com/OhMinsSup/story-server/libs"
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

// StorageRepository Aws Storage
type StorageRepository struct {
	sess *session.Session
}

// Initialize aws sdk project server init function
func Initialize() *session.Session {
	AccessKeyID = libs.GetEnvWithKey("AWS_ACCESS_KEY_ID")
	SecretAccessKey = libs.GetEnvWithKey("AWS_SECRET_ACCESS_KEY")
	Region = libs.GetEnvWithKey("AWS_REGION")

	// aws sdk session create
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

// Inject server context save session info
func Inject(sess *session.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	}
}

// NewStorageRepository storage create
func NewStorageRepository(sess *session.Session) *StorageRepository {
	return &StorageRepository{
		sess: sess,
	}
}

// GetS3PresignedUrl Generate S3 Presigned Url
func (s *StorageRepository) GetS3PresignedUrl(bucket, key string, expiration time.Duration) (string, error) {
	// Create S3 service client
	s3Client := s3.New(s.sess)

	// Construct a GetObjectRequest request
	req, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
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
