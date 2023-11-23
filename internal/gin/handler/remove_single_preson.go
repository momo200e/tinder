package handler

import (
	"log"
	"tinder/domain"
	"tinder/internal/provider"

	"github.com/gin-gonic/gin"
)

// @Summary 刪除用戶
// @Description 從匹配系統中刪除一個用戶，使該用戶無法再被匹配。
// @Tags tinder
// @Accept json
// @Produce json
// @Param        userName   path      string  true  "用戶名(帳號)"
// @Success 200 {object} domain.ResponseFormat "ok"
// @Failure      500  {object}  domain.ErrorFormat
// @Router /people/{userName} [delete]
func RemoveSinglePerson(c *gin.Context) {
	userName := c.Param("userName")

	svc, err := provider.NewService()
	if err != nil {
		log.Println("RemoveSinglePerson error: ", err)
		Failed(c, domain.ErrorServer, err.Error())
		return
	}

	errFormat := svc.RemoveSinglePerson(userName)
	if errFormat != nil {
		log.Println("RemoveSinglePerson error: ", errFormat.Message)
		Failed(c, *errFormat, "")
		return
	}

	Success(c, nil)
}
