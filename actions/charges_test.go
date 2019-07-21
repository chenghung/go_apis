package actions

func (as *ActionSuite) Test_CreateChargeHandler() {
	reqBody := createChargeRequest{
		AmountInCents: 9991,
	}

	res := as.JSON("/charges").Post(reqBody)
	resBody := createChargeResponse{}
	res.Bind(&resBody)
	as.Equal(201, res.Code)
	as.Equal(resBody.AmountInCents, reqBody.AmountInCents)
	as.NotNil(resBody.ID)

}
