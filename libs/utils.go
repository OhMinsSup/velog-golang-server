package libs

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/OhMinsSup/story-server/ent"
	"github.com/gin-gonic/gin"
	"os"
)

type JSON = map[string]interface{}

// CreateHash 해시를 생성하는 함수
func CreateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr
}

// GetEnvWithKey process.env 에 key값과 일치하는 값을 가져온다
func GetEnvWithKey(key string) string {
	return os.Getenv(key)
}

// Repository database client
type Repository struct {
	Client *ent.Client
	DB     *sql.DB
}

// NewRepository - Repository client Create
func NewRepository(ctx *gin.Context) *Repository {
	client := ctx.MustGet("client").(*ent.Client)
	db := ctx.MustGet("db").(*sql.DB)

	return &Repository{
		Client: client,
		DB:     db,
	}
}

// GetClient - get entGo client
func (repository *Repository) GetClient() *ent.Client {
	return repository.Client
}

// GetClient - get sql client
func (repository *Repository) GetDB() *sql.DB {
	return repository.DB
}
