package response

type Response struct {
	Statuscode int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`

}