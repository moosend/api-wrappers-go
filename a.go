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

type A struct {

	// 
	ABVersion float64 `json:"ABVersion,omitempty"`

	// 
	CampaignDeliveredOn string `json:"CampaignDeliveredOn,omitempty"`

	// 
	CampaignID string `json:"CampaignID,omitempty"`

	// 
	CampaignIsArchived bool `json:"CampaignIsArchived,omitempty"`

	// 
	CampaignName string `json:"CampaignName,omitempty"`

	// 
	CampaignSubject string `json:"CampaignSubject,omitempty"`

	// 
	From string `json:"From,omitempty"`

	// 
	MailingLists []MailingLists134 `json:"MailingLists,omitempty"`

	// 
	Sent float64 `json:"Sent,omitempty"`

	// 
	To string `json:"To,omitempty"`

	// 
	TotalBounces float64 `json:"TotalBounces,omitempty"`

	// 
	TotalComplaints float64 `json:"TotalComplaints,omitempty"`

	// 
	TotalForwards float64 `json:"TotalForwards,omitempty"`

	// 
	TotalLinkClicks float64 `json:"TotalLinkClicks,omitempty"`

	// 
	TotalOpens float64 `json:"TotalOpens,omitempty"`

	// 
	TotalUnsubscribes float64 `json:"TotalUnsubscribes,omitempty"`

	// 
	UniqueForwards float64 `json:"UniqueForwards,omitempty"`

	// 
	UniqueLinkClicks float64 `json:"UniqueLinkClicks,omitempty"`

	// 
	UniqueOpens float64 `json:"UniqueOpens,omitempty"`
}
