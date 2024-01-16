package common

func NewSuccessResponse() DefaultResponse {
	return DefaultResponse{
		Code:    200,
		Message: "Success",
	}
}

func NewCreatedSuccessResponse(insertedId string) CreatedSuccessResponse {
	return CreatedSuccessResponse{
		Code:    200,
		Message: "Success",
		Payload: insertedId,
	}
}
