package main

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type Users struct {
	Email    string `json:"email" from:"email"`
	Password string `json:"password" from:"password"`
}

var ListUser []Users

func main() {

}
