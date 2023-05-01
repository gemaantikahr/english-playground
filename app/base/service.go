package base

import "english_playground/app/response"

type ServiceInterface interface {
	Do() response.ServiceResponse
}
