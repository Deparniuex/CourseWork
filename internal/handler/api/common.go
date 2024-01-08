package api

type ID struct {
	Value int64 `json:"-" uri:"id,min=0" binding:"required" example:"21"`
}

type Username struct {
	Value string `uri:"username" binding:"required,min=5,max=50" example:"username"`
}

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
