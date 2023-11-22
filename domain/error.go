package domain

type ErrorFormat struct {
	Code    int
	Message string
}

type ResponseFormat struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
}

var (
	ErrorBadRequest         = ErrorFormat{Code: 400, Message: "bad request"}
	ErrorServer             = ErrorFormat{Code: 500, Message: "Server Error"}
	ErrUserAlreadyExist     = ErrorFormat{Code: 1001, Message: "User already exists"}
	ErrUserNotFound         = ErrorFormat{Code: 1002, Message: "User not found"}
	ErrRemainDatesNotEnough = ErrorFormat{Code: 1003, Message: "Remain dates not enough"}
)
