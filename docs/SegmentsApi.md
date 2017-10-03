# \SegmentsApi

All URIs are relative to *https://api.moosend.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddingCriteriaToSegments**](SegmentsApi.md#AddingCriteriaToSegments) | **Post** /lists/{MailingListID}/segments/{SegmentID}/criteria/add.{Format} | Adding criteria to segments
[**CreatingANewSegment**](SegmentsApi.md#CreatingANewSegment) | **Post** /lists/{MailingListID}/segments/create.{Format} | Creating a new segment
[**DeletingASegment**](SegmentsApi.md#DeletingASegment) | **Delete** /lists/{MailingListID}/segments/{SegmentID}/delete.{Format} | Deleting A Segment
[**GettingSegmentDetails**](SegmentsApi.md#GettingSegmentDetails) | **Get** /lists/{MailingListID}/segments/{SegmentID}/details.{Format} | Getting segment details
[**GettingSegmentSubscribers**](SegmentsApi.md#GettingSegmentSubscribers) | **Get** /lists/{MailingListID}/segments/{SegmentID}/members.{Format} | Getting segment subscribers
[**GettingSegments**](SegmentsApi.md#GettingSegments) | **Get** /lists/{MailingListID}/segments.{Format} | Getting segments
[**UpdatingASegment**](SegmentsApi.md#UpdatingASegment) | **Post** /lists/{MailingListID}/segments/{SegmentID}/update.{Format} | Updating a segment
[**UpdatingSegmentCriteria**](SegmentsApi.md#UpdatingSegmentCriteria) | **Post** /lists/{MailingListID}/segments/{SegmentID}/criteria/{CriteriaID}/update.{Format} | Updating segment criteria


# **AddingCriteriaToSegments**
> AddingCriteriaToSegmentsResponse AddingCriteriaToSegments($format, $mailingListID, $apikey, $segmentID, $body)

Adding criteria to segments

Adds a new criterion (a rule) to the specified segment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list the specified segment belongs. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **segmentID** | **string**| The ID of the segment to update. | 
 **body** | [**AddingCriteriaToSegmentsRequest**](AddingCriteriaToSegmentsRequest.md)|  | 

### Return type

[**AddingCriteriaToSegmentsResponse**](AddingCriteriaToSegmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatingANewSegment**
> CreatingANewSegmentResponse CreatingANewSegment($format, $mailingListID, $apikey, $body)

Creating a new segment

Creates a new empty segment (without criteria) for the given mailing list. You may specify the name of the segment and the way the criteria will match together.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list the specified segment belongs | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **body** | [**CreatingANewSegmentRequest**](CreatingANewSegmentRequest.md)|  | 

### Return type

[**CreatingANewSegmentResponse**](CreatingANewSegmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletingASegment**
> DeletingASegmentResponse DeletingASegment($format, $mailingListID, $apikey, $segmentID)

Deleting A Segment

Deletes a segment along with its criteria from the mailing list. The subscribers of the mailing list that the segment returned are not deleted or affected in any way.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list the specified segment belongs. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **segmentID** | **string**| The ID of the segment to update. | 

### Return type

[**DeletingASegmentResponse**](DeletingASegmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingSegmentDetails**
> GettingSegmentDetailsResponse GettingSegmentDetails($format, $mailingListID, $segmentID, $apikey)

Getting segment details

Gets detailed information on a specific segment and its criteria. However, it does not include the subscribers returned by the segment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list the specified segment belongs | 
 **segmentID** | **string**| The ID of the segment to fetch results for. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 

### Return type

[**GettingSegmentDetailsResponse**](GettingSegmentDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingSegmentSubscribers**
> GettingSegmentSubscribersResponse GettingSegmentSubscribers($format, $mailingListID, $segmentID, $apikey)

Getting segment subscribers

Gets a list of the subscribers that the specified segment returns according to its criteria. Because the results for this call could be quite big, paging information is required as input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list the specified segment belongs | 
 **segmentID** | **string**| The ID of the segment to fetch results for. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 

### Return type

[**GettingSegmentSubscribersResponse**](GettingSegmentSubscribersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingSegments**
> GettingSegmentsResponse GettingSegments($format, $mailingListID, $apikey)

Getting segments

Get a list of all segments with their criteria for the given mailing list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 

### Return type

[**GettingSegmentsResponse**](GettingSegmentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatingASegment**
> UpdatingASegmentResponse UpdatingASegment($format, $mailingListID, $apikey, $segmentID, $body)

Updating a segment

Updates the properties of an existing segment. You may update the name and match type of the segment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list the specified segment belongs | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **segmentID** | **string**| The ID of the segment to update. | 
 **body** | [**UpdatingASegmentRequest**](UpdatingASegmentRequest.md)|  | 

### Return type

[**UpdatingASegmentResponse**](UpdatingASegmentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatingSegmentCriteria**
> UpdatingSegmentCriteriaResponse UpdatingSegmentCriteria($format, $mailingListID, $apikey, $segmentID, $criteriaID, $body)

Updating segment criteria

Updates an existing criterion in the specified segment.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list the specified segment belongs. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **segmentID** | **string**| The ID of the segment to update. | 
 **criteriaID** | **float64**| The ID of the criterion to process. | 
 **body** | [**UpdatingSegmentCriteriaRequest**](UpdatingSegmentCriteriaRequest.md)|  | 

### Return type

[**UpdatingSegmentCriteriaResponse**](UpdatingSegmentCriteriaResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

