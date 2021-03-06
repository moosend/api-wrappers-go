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

type Segment struct {

	// 
	CreatedBy string `json:"CreatedBy,omitempty"`

	// 
	CreatedOn string `json:"CreatedOn,omitempty"`

	// 
	Criteria []Criterion `json:"Criteria,omitempty"`

	// 
	Description string `json:"Description,omitempty"`

	// 
	FetchType float64 `json:"FetchType,omitempty"`

	// 
	FetchValue float64 `json:"FetchValue,omitempty"`

	// 
	ID float64 `json:"ID,omitempty"`

	// 
	MatchType float64 `json:"MatchType,omitempty"`

	// 
	Name string `json:"Name,omitempty"`

	// 
	UpdatedBy string `json:"UpdatedBy,omitempty"`

	// 
	UpdatedOn string `json:"UpdatedOn,omitempty"`
}
