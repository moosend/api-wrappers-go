/* 
 * Moosend API
 *
 * TODO: Add a description
 *
 * OpenAPI spec version: 1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package Moosend.Wrappers.GOWrapper

type Context84 struct {

	// 
	ABCampaignData string `json:"ABCampaignData,omitempty"`

	// 
	ConfirmationTo string `json:"ConfirmationTo,omitempty"`

	// 
	CreatedOn string `json:"CreatedOn,omitempty"`

	// 
	DeliveredOn string `json:"DeliveredOn,omitempty"`

	// 
	FormatType float64 `json:"FormatType,omitempty"`

	// 
	HTMLContent string `json:"HTMLContent,omitempty"`

	// 
	ID string `json:"ID,omitempty"`

	// 
	IsTransactional bool `json:"IsTransactional,omitempty"`

	// 
	MailingLists []MailingList85 `json:"MailingLists,omitempty"`

	// 
	Name string `json:"Name,omitempty"`

	// 
	PlainContent string `json:"PlainContent,omitempty"`

	ReplyToEmail ReplyToEmail `json:"ReplyToEmail,omitempty"`

	// 
	ScheduledFor string `json:"ScheduledFor,omitempty"`

	Sender Sender `json:"Sender,omitempty"`

	// 
	Status float64 `json:"Status,omitempty"`

	// 
	Subject string `json:"Subject,omitempty"`

	// 
	Timezone string `json:"Timezone,omitempty"`

	// 
	UpdatedOn string `json:"UpdatedOn,omitempty"`

	// 
	WebLocation string `json:"WebLocation,omitempty"`
}