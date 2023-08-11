# Intistele.com

###### With the intistele.com, you can access and organise SMS mailouts from your own software via API
###### [Intistele.com documentation](https://api-go2.intistele.com/docs/)


### Environment variables:

INTISTELECOM_USERNAME - Your [Username](https://go2.intistele.com/#/api/open-api)

INTISTELECOM_API_KEY - Your [API key](https://go2.intistele.com/#/api/open-api)

### Quickstart

```bash
package main

import sms "github.com/pixel365/go-intistelecom"

func main() {
	msg := sms.MessageBody{
		Destination:    "17779990011",
		Originator:     "test",
		Text:           "some text",
		TimeToSend:     "",
		CallbackURL:    "https://some-callback-url.com",
		ValidityPeriod: 0,
		UseLocaltime:   false,
	}
	sms.Send(&msg)
}
```
