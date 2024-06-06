// Package nats_connect_shared
/*
Copyright 6/6/24 STY Holdings Inc

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the “Software”), to deal in
the Software without restriction, including without limitation the rights to use,
copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
OTHER DEALINGS IN THE SOFTWARE.

*/
package nats_connect_shared

// SendGrid - Twilio

type SendEmailRequest struct {
	BodyPlain          string           `json:"body_plain,omitempty"`
	BodyHTML           string           `json:"body_html,omitempty"`
	EmailToRecipient   []emailRecipient `json:"email_to_recipient"`
	SaaSKey            string           `json:"saas_key"`
	SenderEmailAddress string           `json:"sender_email_address,omitempty"`
	SenderName         string           `json:"sender_name,omitempty"`
	Subject            string           `json:"subject,omitempty"`
}

type emailRecipient struct {
	name    string
	address string
}

// Stripe

type ListPaymentMethodRequest struct {
	SaaSKey string `json:"saas_key"`
}

type ListPaymentIntentRequest struct {
	CustomerId    string `json:"customer_id,omitempty"`
	Limit         int64  `json:"limit,omitempty"`
	SaaSKey       string `json:"saas_key"`
	StartingAfter string `json:"starting_after,omitempty"`
}

type PaymentIntentRequest struct {
	Amount                  float64 `json:"amount"`
	AutomaticPaymentMethods bool    `json:"automatic_payment_methods,omitempty"`
	Currency                string  `json:"currency"`
	Description             string  `json:"description,omitempty"`
	ReceiptEmail            string  `json:"receipt_email"`
	ReturnURL               string  `json:"return_url,omitempty"`
	SaaSKey                 string  `json:"saas_key"`
	// Confirm            bool     `json:"confirm,omitempty"`
	// PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
}

type CancelPaymentIntentRequest struct {
	CancellationReason string `json:"cancellation_reason"`
	PaymentIntentId    string `json:"id"`
	SaaSKey            string `json:"saas_key"`
}

type ConfirmPaymentIntentRequest struct {
	CaptureMethod   string `json:"capture_method,omitempty"`
	PaymentIntentId string `json:"id"`
	PaymentMethod   string `json:"payment_method,omitempty"`
	ReceiptEmail    string `json:"receipt_email,omitempty"`
	ReturnURL       string `json:"return_url,omitempty,omitempty"`
	SaaSKey         string `json:"saas_key"`
}

// Synadia Cloud

type GetPersonalAccessTokenRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TokenId string `json:"token_id"`
}

type GetPrometheusMetricsRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type GetSystemLimitsRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type GetSystemRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type GetTeamRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type GetTeamLimitsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type GetVersionRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
}

type ListAccountsRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type ListInfoAppUserTeamRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type ListNATSUsersRequest struct {
	SaaSKey   string `json:"saas_key"`
	BaseURL   string `json:"base_url"`
	AccountId string `json:"account_id"`
}

type ListPersonalAccessTokensRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type ListSystemsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}

type ListSystemAccountInfoRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type ListSystemServerInfoRequest struct {
	SaaSKey  string `json:"saas_key"`
	BaseURL  string `json:"base_url"`
	SystemId string `json:"system_id"`
}

type ListTeamsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
}

type ListTeamServerAccountsRequest struct {
	SaaSKey string `json:"saas_key"`
	BaseURL string `json:"base_url"`
	TeamId  string `json:"team_id"`
}
