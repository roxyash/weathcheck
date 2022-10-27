package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"weathcheck/internal/types"

	"github.com/gin-gonic/gin"

	"github.com/sirupsen/logrus"
)

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, types.ErrorResponse{ErrorMessage: message})
}

func CheckStatus(statusCode int, bodyText []byte) error {
	if statusCode != http.StatusOK {
		var data interface{}
		err := json.Unmarshal(bodyText, &data)
		if err != nil {
			return err
		}
		newData, _ := data.(map[string]interface{})
		
		return errors.New(newData["error"].(string))
	}

	return nil
}
