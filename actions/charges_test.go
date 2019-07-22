package actions

import (
	"gopkg.in/h2non/gock.v1"
)

func (as *ActionSuite) Test_CreateChargeHandler() {
	defer gock.Off()

	reqBody := createChargeRequest{
		AmountInCents: 9991,
	}

	gock.New("https://api.stripe.com/").
		Post("/v1/charges").
		Reply(201).
		JSON(map[string]interface{}{
			"id":     "ch_xxxxxxxx",
			"amount": 9991,
		})

	res := as.JSON("/charges").Post(reqBody)
	resBody := createChargeResponse{}
	res.Bind(&resBody)
	as.Equal(201, res.Code)
	as.Equal(resBody.AmountInCents, reqBody.AmountInCents)
	as.NotNil(resBody.ID)

}
