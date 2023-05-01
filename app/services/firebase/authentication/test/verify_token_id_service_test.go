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

func TestVerifyTokenIDService(t *testing.T) {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	err := godotenv.Load("../../../../../.env")
	if err != nil {
		fmt.Println("can't load the env : ", err.Error())
		return
	}

	// defined tokenID
	tokenId := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjFlOTczZWUwZTE2ZjdlZWY0ZjkyMWQ1MGRjNjFkNzBiMmVmZWZjMTkiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiWWVyZW1pYSBDaHJpcyBTYXJhZ2kiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUdObXl4YmMzOS1PN2J4Sy1UMzY5bEtGV2JWTVI5WU1VcFd3c3JGRzNUbjVpZz1zOTYtYyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS90YXVrdS1kODUyNCIsImF1ZCI6InRhdWt1LWQ4NTI0IiwiYXV0aF90aW1lIjoxNjc5NjQxMzQ1LCJ1c2VyX2lkIjoiN1N6UXJmS3lYY1FCTERiNkdMblZTRUd1cUh1MSIsInN1YiI6IjdTelFyZkt5WGNRQkxEYjZHTG5WU0VHdXFIdTEiLCJpYXQiOjE2Nzk2NDEzNDUsImV4cCI6MTY3OTY0NDk0NSwiZW1haWwiOiJ5ZXJlbWlhOTk3QGdtYWlsLmNvbSIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJmaXJlYmFzZSI6eyJpZGVudGl0aWVzIjp7Imdvb2dsZS5jb20iOlsiMTE0MTg1OTAwMzQxMjk5NjAxMTMxIl0sImVtYWlsIjpbInllcmVtaWE5OTdAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.HsS51UTkvoUFCkJbVWA4yNGJhiS8egakPTjBnTkrBz-5dSWPqzpIHr7hBibjpWiFVgKj0BlJ9ozwKViwxXvLja42ilfrl5jl9OYyyND5mWC4MD1Q5-FMfU3Mb_18hKNC4dHOjRKeTDFVHMkxOyJm2V1NvblJm5IDwB50QUc8krVU0mB_uynq4s3AvRLjgD0wM0yUQe9FgjtggFAluchCEGHpLWazybG4EdkiXQ7oZwMKp2ti-FrGKHBzsYtZxH1UWxAhLToaD4beDM2ZDHAHwgUFwNvMOYvV1z91Rg1hJh4TT6F1BtzePWuHyxquK8t9zGqzcgEDqm67LABTI3efLw"

	// call service
	service := authentication.VerifyTokenIDService{
		TokenID: tokenId,
		Repo:    authentication2.AuthenticationRepository{},
	}.Do()

	if !service.Status {
		println("something went wrong ", service.Message)
		return
	}

	fmt.Println("oke doke", service.Data)

}
