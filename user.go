package intistelecom

type Balance struct {
	Currency string  `json:"currency"`
	Amount   float32 `json:"amount"`
}

type User struct {
	UserName string `json:"username"`
	Id       int32  `json:"id"`
}

func GetBalance() (Balance, error) {
	cl := getClient()
	result := Balance{}
	response, err := cl.Get("/user/balance")
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}

func GetMe() (User, error) {
	cl := getClient()
	result := User{}
	response, err := cl.Get("/user/me")
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}
