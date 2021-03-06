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

type Paging struct {

	// 
	CurrentPage float64 `json:"CurrentPage,omitempty"`

	// 
	PageSize float64 `json:"PageSize,omitempty"`

	// 
	SortExpression string `json:"SortExpression,omitempty"`

	// 
	SortIsAscending bool `json:"SortIsAscending,omitempty"`

	// 
	TotalPageCount float64 `json:"TotalPageCount,omitempty"`

	// 
	TotalResults float64 `json:"TotalResults,omitempty"`
}
