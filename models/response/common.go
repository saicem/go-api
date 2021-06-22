package response

type Response struct {
	Status  ApiCode     `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
