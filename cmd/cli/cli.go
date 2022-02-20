package main

import (
	"fmt"
	"logClient/services"
	"logClient/types/auth"
)

func main() {
	// res, err := logit.GetUser(
	// 	// &idea.Idea{
	// 	// 	Name: "ok",
	// 	// },
	// 	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Im9rQGdtYWlsLmNvbSIsImV4cCI6MTY0MTA1MTY0MSwiaWQiOiI2MWNmMThiOTNkN2JjNmZjNmIxMWYwOWIifQ.uNfzoZY8_XTDiar-ZuclTd5GXWuvRkfoAtQsMn2uLfM",
	// )
	// if err != nil {
	// 	fmt.Println("err: ", err.Error())

	// }
	// fmt.Println("response: ", string(res))
	ls := services.LogitService{}
	output := make(chan *auth.AuthResponse)
	exit := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("output=>", <-output)

		}
		exit <- true
	}()
	// for i := 0; i < 3; i++ {
	ls.GetToken(&auth.AuthRequest{Email: "ok@gmail.com", Password: "password"}, output, exit)

	// }
}
