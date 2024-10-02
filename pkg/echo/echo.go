package echo

import (
	"io"

	"github.com/banggibima/backend-agile/config"
	"github.com/labstack/echo/v4"
)

func Init(config *config.Config) (*echo.Echo, error) {
	e := echo.New()

	e.HideBanner = true
	e.Logger.SetOutput(io.Discard)

	return e, nil
}
