package handler

import (
	"log"
	"strconv"
	"tinder/domain"
	"tinder/internal/provider"

	"github.com/gin-gonic/gin"
)

// @Summary 配對n組
// @Description 尋找最多N個可能符合的單身人士，其中N是請求參數。
// @Tags tinder
// @Accept json
// @Produce json
// @Param        userName   path      string  true  "用戶名(帳號)"
// @Param        number   path      string  true  "配對數量"
// @Success      200  {object}  domain.ResponseFormat{data=SinglePersonAndMatchResponse}
// @Failure      500  {object}  domain.ErrorFormat
// @Router /people/{userName}/query_single_person/{number} [post]
func QuerySinglePerson(c *gin.Context) {
	number := c.Param("number")
	userName := c.Param("userName")

	n, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Println("strconv.ParseUint error: ", err)
		Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	svc, err := provider.NewService()
	if err != nil {
		log.Println("QuerySinglePerson error: ", err)
		Failed(c, domain.ErrorServer, err.Error())
		return
	}

	matches, errFormat := svc.QuerySinglePerson(userName, int(n))
	if errFormat != nil {
		log.Println("QuerySinglePerson error: ", errFormat.Message)
		Failed(c, *errFormat, "")
		return
	}

	output := []SinglePersonAndMatchResponseUser{}
	for _, match := range matches {
		output = append(output, SinglePersonAndMatchResponseUser{
			Name:        match.Name,
			Height:      match.Height,
			Gender:      match.Gender,
			RemainDates: match.RemainDates,
		})
	}

	Success(c, SinglePersonAndMatchResponse{Users: output})
}
