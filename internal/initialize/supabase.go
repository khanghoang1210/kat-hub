package initialize

import (
	"os"
	storage_go "github.com/supabase-community/storage-go"
)

var StorageClient *storage_go.Client

func CreateSession() {
	projUrl := os.Getenv("SUPABASE_PROJECT_URL")
	secretKey := os.Getenv("SUPABASE_SECRET_KEY")
	//var err error
	StorageClient =  storage_go.NewClient(projUrl, secretKey, nil)
}
