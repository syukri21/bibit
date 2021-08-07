package middlewares

import (
	"bibit-test/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusBadRequest, models.GenericRes{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Error:   err,
		})
	}

	c.Logger().Error(report)
	_ = c.JSON(report.Code, models.GenericRes{
		Code:    report.Code,
		Message: report.Message.(string),
		Error:   report,
	})
}
