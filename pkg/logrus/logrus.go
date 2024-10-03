package logrus

import (
	"os"

	"github.com/banggibima/agile-backend/config"
	"github.com/sirupsen/logrus"
)

func Init(config *config.Config) (*logrus.Logger, error) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)

	return logger, nil
}
