package responses

type SuccessResponseReturn struct {
	status  string `json:"status"`
}

var SuccessResponse = SuccessResponseReturn{
	status: "success",
}