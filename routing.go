package intistelecom

type Routing struct {
	Currency string  `json:"currency"`
	Mcc      string  `json:"mcc"`
	Mnc      string  `json:"mnc"`
	Price    float32 `json:"price"`
}

func Cost(phoneNumber string) (Routing, error) {
	cl := getClient()
	result := Routing{}
	response, err := cl.Get("/routing/" + phoneNumber)
	if err != nil {
		return result, err
	}

	err = cl.JSONDecoder(response, &result)
	return result, err
}
