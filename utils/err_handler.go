package utils

import "gochiapp/model"

func ErrorResponseWeb(err error, code int) {

	var status = "Failed"

	if code == 500 {
		status = "Error"
	}

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     status,
			StatusCode: code,
			Error:      err.Error(),
		})
	}
}
