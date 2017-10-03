# \MailingListsApi

All URIs are relative to *https://api.moosend.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatingACustomField**](MailingListsApi.md#CreatingACustomField) | **Post** /lists/{MailingListID}/customfields/create.{Format} | Creating a custom field
[**CreatingAMailingList**](MailingListsApi.md#CreatingAMailingList) | **Post** /lists/create.{Format} | Creating a mailing list
[**DeletingAMailingList**](MailingListsApi.md#DeletingAMailingList) | **Delete** /lists/{MailingListID}/delete.{Format} | Deleting a mailing list
[**GettingAllActiveMailingLists**](MailingListsApi.md#GettingAllActiveMailingLists) | **Get** /lists.{Format} | Getting all active mailing lists
[**GettingAllActiveMailingListsWithPaging**](MailingListsApi.md#GettingAllActiveMailingListsWithPaging) | **Get** /lists/{Page}/{PageSize}.{Format} | Getting all active mailing lists with paging
[**GettingMailingListDetails**](MailingListsApi.md#GettingMailingListDetails) | **Get** /lists/{MailingListID}/details.{Format} | Getting mailing list details
[**RemovingACustomField**](MailingListsApi.md#RemovingACustomField) | **Delete** /lists/{MailingListID}/customfields/{CustomFieldID}/delete.{Format} | Removing a custom field
[**UpdatingACustomField**](MailingListsApi.md#UpdatingACustomField) | **Post** /lists/{MailingListID}/customfields/{CustomFieldID}/update.{Format} | Updating a custom field
[**UpdatingAMailingList**](MailingListsApi.md#UpdatingAMailingList) | **Post** /lists/{MailingListID}/update.{Format} | Updating a mailing list


# **CreatingACustomField**
> CreatingACustomFieldResponse CreatingACustomField($format, $apikey, $mailingListID, $body)

Creating a custom field

Creates a new custom field in the specified mailing list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list where the custom field will belong. | 
 **body** | [**CreatingACustomFieldRequest**](CreatingACustomFieldRequest.md)|  | 

### Return type

[**CreatingACustomFieldResponse**](CreatingACustomFieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatingAMailingList**
> CreatingAMailingListResponse CreatingAMailingList($format, $apikey, $body)

Creating a mailing list

Creates a new empty mailing list in your account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **body** | [**CreatingAMailingListRequest**](CreatingAMailingListRequest.md)|  | 

### Return type

[**CreatingAMailingListResponse**](CreatingAMailingListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletingAMailingList**
> DeletingAMailingListResponse DeletingAMailingList($format, $apikey, $mailingListID)

Deleting a mailing list

Deletes a mailing list from your account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to be deleted. | 

### Return type

[**DeletingAMailingListResponse**](DeletingAMailingListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingAllActiveMailingLists**
> GettingAllActiveMailingListsResponse GettingAllActiveMailingLists($format, $apikey, $withStatistics, $shortBy, $sortMethod)

Getting all active mailing lists

Gets a list of your active mailing lists in your account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **withStatistics** | **string**| Specifies whether to fetch statistics for the subscribers or not. If omitted, results will be returned with statistics by default. | [optional] 
 **shortBy** | **string**| The name of the campaign property to sort results by. If not specified, results will be sorted by the CreatedOn property | [optional] 
 **sortMethod** | **string**| The method to sort results: ASC for ascending, DESC for descending. If not specified, &#x60;ASC&#x60; will be assumed | [optional] 

### Return type

[**GettingAllActiveMailingListsResponse**](GettingAllActiveMailingListsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingAllActiveMailingListsWithPaging**
> GettingAllActiveMailingListsWithPagingResponse GettingAllActiveMailingListsWithPaging($format, $apikey, $page, $pageSize, $shortBy, $sortMethod)

Getting all active mailing lists with paging

Gets a list of your active mailing lists in your account. Because the results for this call could be quite big, paging information is required as input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **page** | **float64**| The page that you want to get. | 
 **pageSize** | **float64**| Lists Per Page. | 
 **shortBy** | **string**| The name of the campaign property to sort results by. If not specified, results will be sorted by the CreatedOn property | [optional] 
 **sortMethod** | **string**| The method to sort results: ASC for ascending, DESC for descending. If not specified, &#x60;ASC&#x60; will be assumed | [optional] 

### Return type

[**GettingAllActiveMailingListsWithPagingResponse**](GettingAllActiveMailingListsWithPagingResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingMailingListDetails**
> GettingMailingListDetailsResponse GettingMailingListDetails($format, $mailingListID, $apikey, $withStatistics)

Getting mailing list details

Gets details for a given mailing list. You may include subscriber statistics in your results or not. Any segments existing for the requested mailing list will not be included in the results.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list to be returned. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **withStatistics** | **string**| Specifies whether to fetch statistics for the subscribers or not. If omitted, results will be returned with statistics by default. | [optional] 

### Return type

[**GettingMailingListDetailsResponse**](GettingMailingListDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovingACustomField**
> RemovingACustomFieldResponse RemovingACustomField($format, $customFieldID, $apikey, $mailingListID)

Removing a custom field

Removes a custom field definition from the specified mailing list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **customFieldID** | **string**| The ID of the custom field to be deleted. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list where the custom field belongs. | 

### Return type

[**RemovingACustomFieldResponse**](RemovingACustomFieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatingACustomField**
> UpdatingACustomFieldResponse UpdatingACustomField($format, $customFieldID, $apikey, $mailingListID, $body)

Updating a custom field

Updates the properties of an existing custom field in the specified mailing list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **customFieldID** | **string**| The ID of the custom field to be updated. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list where the custom field belongs. | 
 **body** | [**UpdatingACustomFieldRequest**](UpdatingACustomFieldRequest.md)|  | 

### Return type

[**UpdatingACustomFieldResponse**](UpdatingACustomFieldResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatingAMailingList**
> UpdatingAMailingListResponse UpdatingAMailingList($format, $apikey, $mailingListID, $body)

Updating a mailing list

Updates the properties of an existing mailing list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to be updated. | 
 **body** | [**UpdatingAMailingListRequest**](UpdatingAMailingListRequest.md)|  | 

### Return type

[**UpdatingAMailingListResponse**](UpdatingAMailingListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

