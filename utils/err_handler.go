package utils

import "gochiapp/model"

func ErrorResponseWeb(err error, code int) {

	if err != nil {
		panic(model.ResponseFailWeb{
			Status:     "Failed",
			StatusCode: code,
			Error:      err.Error(),
		})
	}
}
