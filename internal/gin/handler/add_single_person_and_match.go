package handler

import (
	"log"
	"tinder/domain"
	"tinder/internal/provider"

	"github.com/gin-gonic/gin"
)

type AddSinglePersonAndMatchRequest struct {
	Name        string `json:"name" binding:"required" example:"Jason"`
	Height      uint8  `json:"height" binding:"required,gt=0,lt=250" example:"165"`
	Gender      uint8  `json:"gender" binding:"required,gt=0,lt=3" example:"1"`
	RemainDates uint8  `json:"wanted_dates" binding:"required,gt=0,lt=250" example:"10"`
}
type SinglePersonAndMatchResponse struct {
	Users []SinglePersonAndMatchResponseUser `json:"users"`
}
type SinglePersonAndMatchResponseUser struct {
	Name        string        `json:"name" example:"Jason"`
	Height      uint8         `json:"height" example:"165"`
	Gender      domain.Gender `json:"gender" example:"1"`
	RemainDates uint8         `json:"remainDates" example:"10"`
}

// @Summary 加入用戶
// @Description 將新使用者新增至匹配系統並為新使用者尋找任何可能的符合項目。
// @Tags tinder
// @Accept json
// @Produce json
// @Param Request body AddSinglePersonAndMatchRequest true "<ul><li>gender: 性別 男=1, 女=2</li></ul>"
// @Success      200  {object}  domain.ResponseFormat{data=SinglePersonAndMatchResponse}
// @Failure      500  {object}  domain.ErrorFormat
// @Router /people/add_and_match [post]
func AddSinglePersonAndMatch(c *gin.Context) {
	var req AddSinglePersonAndMatchRequest
	if err := c.BindJSON(&req); err != nil {
		log.Println("AddSinglePersonAndMatch error: ", err)
		Failed(c, domain.ErrorBadRequest, err.Error())
		return
	}

	newUser := domain.User{
		Name:        req.Name,
		Height:      req.Height,
		Gender:      domain.Gender(req.Gender),
		RemainDates: req.RemainDates,
	}

	svc, err := provider.NewService()
	if err != nil {
		log.Println("AddSinglePersonAndMatch error: ", err)
		Failed(c, domain.ErrorServer, err.Error())
		return
	}

	matches, errFormat := svc.AddSinglePersonAndMatch(&newUser)
	if errFormat != nil {
		log.Println("AddSinglePersonAndMatch error: ", errFormat.Message)
		Failed(c, domain.ErrorServer, errFormat.Message)
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
