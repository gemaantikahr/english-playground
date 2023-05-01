package test

import (
	"english_playground/app/base"
	"english_playground/app/components"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestStoreData(t *testing.T) {

	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		println("error load env")
		return
	}

	client := base.InitRedis()
	sessionStorage := components.SessionStorageInit(client)

	// defined data
	var data AutoGenerated
	data.Type = "CTC"
	data.GroupName = "REGULAR"
	data.ServiceName = "JNE"

	if err := sessionStorage.PutCache("couriers", "jne_ctc_regular", data, 1000); err != nil {
		println("something went wrong !", err.Error())
		return
	}

	println("success")
	return

}

func TestGetData(t *testing.T) {

	errEnv := godotenv.Load("../.env")
	if errEnv != nil {
		println("error load env")
		return
	}

	client := base.InitRedis()
	sessionStorage := components.SessionStorageInit(client)
	var data AutoGenerated
	if err := sessionStorage.GetCache("couriers", "jne_ctc_regular", &data); err != nil {
		println("something went wrong ", err.Error())
		return
	}

	fmt.Println("the data", data)
	return
}

type AutoGenerated struct {
	ServiceName string `json:"service_name"`
	GroupName   string `json:"group_name"`
	Type        string `json:"type"`
}
