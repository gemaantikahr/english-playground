package test

import (
	authentication2 "english_playground/app/repositories/firebase/authentication"
	"english_playground/app/services/firebase/authentication"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func TestExecuteVerifyTokenServiceTest(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	service := authentication.ExecuteVerifyTokenService{
		TokenId: "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFlOTczZWUwZTE2ZjdlZWY0ZjkyMWQ1MGRjNjFkNzBiMmVmZWZjMTkiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiWWVyZW1pYSBDaHJpcyBTYXJhZ2kiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUdObXl4YmMzOS1PN2J4Sy1UMzY5bEtGV2JWTVI5WU1VcFd3c3JGRzNUbjVpZz1zOTYtYyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS90YXVrdS1kODUyNCIsImF1ZCI6InRhdWt1LWQ4NTI0IiwiYXV0aF90aW1lIjoxNjc5NjQ2MzI2LCJ1c2VyX2lkIjoiN1N6UXJmS3lYY1FCTERiNkdMblZTRUd1cUh1MSIsInN1YiI6IjdTelFyZkt5WGNRQkxEYjZHTG5WU0VHdXFIdTEiLCJpYXQiOjE2Nzk2NDYzMjYsImV4cCI6MTY3OTY0OTkyNiwiZW1haWwiOiJ5ZXJlbWlhOTk3QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7Imdvb2dsZS5jb20iOlsiMTE0MTg1OTAwMzQxMjk5NjAxMTMxIl0sImVtYWlsIjpbInllcmVtaWE5OTdAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.k_lYF4yOqzybxWLskrNy-DrCQD5NCDJ9pUdCn0J2Dea6xSuJq6aug_AMSZago7tVIwTh8Op6txxHouG347eJ7yfibxCriLYUeFnkp3tT1qmPRwaT2qU55ioN8DsklTpKQ5CfCAKXXpfr6_A1_1TivT6QzwVf-kfsAxzM6AS0oV4fC4nB_Gn3tJuH9pIIWOw20AIODi7wSET4wfzb0rUxF2ffLX2L4HPGM5h1OWnRrHQCmw8RXL6vd-7rSHvX4IuKjXnzA57tE6er48d7GLzuQaiJpYAQI7Ckhal8o1l09VKtSbRDVbtjY8Fcyl5Pt15-8TICmW53VWXIUlCgPUjEnA",
		Repo:    authentication2.AuthenticationRepository{},
	}.Do()

	if !service.Status {
		println("something went wrong ! ", service.Message)
		return
	}

	fmt.Println("success ", service.Data)
}
