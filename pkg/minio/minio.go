package minio

import (
	"github.com/banggibima/agile-backend/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func Client(config *config.Config) (*minio.Client, error) {
	endpoint := config.Minio.Endpoint
	accessKeyID := config.Minio.AccessKeyID
	secretAccessKey := config.Minio.SecretAccessKey
	useSSL := config.Minio.UseSSL

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
