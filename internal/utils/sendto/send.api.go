package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	MessageBody string `json:"messageBody"`
	Subject     string `json:"subject"`
	Attachment  string `json:"attachment"`
}

func SendEmailToJavaByAPI(otp string, email string, purpose string) error {
	// URL API
	postURL := "http://localhost:8080/api/v1/send_text"

	// Data JSON
	mailRequest := MailRequest{
		ToEmail:     email,
		MessageBody: "OTP is " + otp,
		Subject:     "Verify OTP" + purpose,
		Attachment:  "path/to/email",
	}

	// Convert struct to JSON
	requestBody, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	fmt.Sprintln("Response status: ", resp.Status)

	return nil
}
