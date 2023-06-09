package repositories

import (
	"encoding/json"
	"english_playground/app/models"
	"english_playground/app/services/http_example_service/base"
	"io"
)

type HttpExampleInterface interface {
	GetPost() models.Posts
}

type HttpExampleRepository struct{}

func (HttpExampleRepository) GetPost() models.Posts {
	call, err := base.ExampleApi{}.Get("/posts").Call()
	if err != nil {
		return nil
	}
	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return nil
	}
	var data models.Posts
	errUm := json.Unmarshal(result, &data)
	if errUm != nil {
		return nil
	}
	return data
}
