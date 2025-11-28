package impl

import (
	"be/models/domains"
	"fmt"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioServImpl struct {
	Client     *twilio.RestClient
	FromNumber string
}

func NewTwilioServImpl(model domains.Twilio) *TwilioServImpl {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: model.AccountId,
		Password: model.AuthToken,
	})

	return &TwilioServImpl{
		Client:     client,
		FromNumber: model.FromNumber,
	}
}

func (serv *TwilioServImpl) SendOtpToPhoneNumber(phoneNumber, otpCode string) error {
	if phoneNumber == "" {
		return fmt.Errorf("phone number cannot be empty")
	}
	if otpCode == "" {
		return fmt.Errorf("OTP code cannot be empty")
	}

	body := fmt.Sprintf("XSELL Official - Your OTP is: %s. Valid for 5 minutes.", otpCode)

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(serv.FromNumber)
	params.SetBody(body)

	resp, err := serv.Client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("failed to send OTP to %s: %w", phoneNumber, err)
	}

	if resp.Sid != nil {
		fmt.Printf("OTP sent successfully. Message SID: %s\n", *resp.Sid)
	}

	return nil
}
