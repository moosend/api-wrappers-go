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

type MailingList69 struct {

	// 
	ActiveMemberCount float64 `json:"ActiveMemberCount,omitempty"`

	// 
	BouncedMemberCount float64 `json:"BouncedMemberCount,omitempty"`

	// 
	CreatedBy string `json:"CreatedBy,omitempty"`

	// 
	CreatedOn string `json:"CreatedOn,omitempty"`

	// 
	CustomFieldsDefinition []interface{} `json:"CustomFieldsDefinition,omitempty"`

	// 
	ID string `json:"ID,omitempty"`

	// 
	ImportOperation string `json:"ImportOperation,omitempty"`

	// 
	Name string `json:"Name,omitempty"`

	// 
	RemovedMemberCount float64 `json:"RemovedMemberCount,omitempty"`

	// 
	Status float64 `json:"Status,omitempty"`

	// 
	UnsubscribedMemberCount float64 `json:"UnsubscribedMemberCount,omitempty"`

	// 
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// 
	UpdatedOn string `json:"UpdatedOn,omitempty"`
}