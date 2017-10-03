# \CampaignsApi

All URIs are relative to *https://api.moosend.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ABTestCampaignSummary**](CampaignsApi.md#ABTestCampaignSummary) | **Get** /campaigns/{CampaignID}/view_ab_summary.{Format} | AB Test Campaign Summary
[**ActivityByLocation**](CampaignsApi.md#ActivityByLocation) | **Get** /campaigns/{CampaignID}/stats/countries.{Format} | Activity By Location
[**CampaignSummary**](CampaignsApi.md#CampaignSummary) | **Get** /campaigns/{CampaignID}/view_summary.{Format} | Campaign Summary
[**CloningAnExistingCampaign**](CampaignsApi.md#CloningAnExistingCampaign) | **Post** /campaigns/{CampaignID}/clone.{Format} | Cloning An Existing Campaign
[**CreatingADraftCampaign**](CampaignsApi.md#CreatingADraftCampaign) | **Post** /campaigns/create.{Format} | Creating A Draft Campaign
[**DeletingACampaign**](CampaignsApi.md#DeletingACampaign) | **Delete** /campaigns/{CampaignID}/delete.{Format} | Deleting A Campaign
[**GetAllCampaigns**](CampaignsApi.md#GetAllCampaigns) | **Get** /campaigns.{Format} | Get All Campaigns
[**GetCampaignStatisticsWithPagingFiltered**](CampaignsApi.md#GetCampaignStatisticsWithPagingFiltered) | **Get** /campaigns/{CampaignID}/stats/{Type}.{Format} | Get Campaign Statistics With Paging &amp; Filtered
[**GetCampaignsByPage**](CampaignsApi.md#GetCampaignsByPage) | **Get** /campaigns/{Page}.{Format} | Get Campaigns By Page
[**GetCampaignsByPageAndPagesize**](CampaignsApi.md#GetCampaignsByPageAndPagesize) | **Get** /campaigns/{Page}/{PageSize}.{Format} | Get Campaigns By Page And Pagesize
[**GettingAllYourSenders**](CampaignsApi.md#GettingAllYourSenders) | **Get** /senders/find_all.{Format} | Getting All Your Senders
[**GettingCampaignDetails**](CampaignsApi.md#GettingCampaignDetails) | **Get** /campaigns/{CampaignID}/view.{Format} | Getting Campaign Details
[**GettingSenderDetails**](CampaignsApi.md#GettingSenderDetails) | **Get** /senders/find_one.{Format} | Getting Sender Details
[**LinkActivity**](CampaignsApi.md#LinkActivity) | **Get** /campaigns/{CampaignID}/stats/links.{Format} | Link Activity
[**SchedulingACampaign**](CampaignsApi.md#SchedulingACampaign) | **Post** /campaigns/{CampaignID}/schedule.{Format} | Scheduling A Campaign
[**SendingACampaign**](CampaignsApi.md#SendingACampaign) | **Post** /campaigns/{CampaignID}/send.{Format} | Sending a campaign
[**TestingACampaign**](CampaignsApi.md#TestingACampaign) | **Post** /campaigns/{CampaignID}/send_test.{Format} | Testing a campaign
[**UnschedulingACampaign**](CampaignsApi.md#UnschedulingACampaign) | **Post** /campaigns/{CampaignID}/unschedule.{Format} | Unscheduling a campaign
[**UpdatingADraftCampaign**](CampaignsApi.md#UpdatingADraftCampaign) | **Post** /campaigns/{CampaignID}/update.{Format} | Updating A Draft Campaign


# **ABTestCampaignSummary**
> AbTestCampaignSummaryResponse ABTestCampaignSummary($format, $apikey, $campaignID)

AB Test Campaign Summary

Provides a basic summary of the results for a sent AB test campaign, separately for each version (A and B), such as the number of recipients, opens, clicks, bounces, unsubscribes, forwards etc to date.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the requested AB test campaign | 

### Return type

[**AbTestCampaignSummaryResponse**](AbTestCampaignSummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivityByLocation**
> ActivityByLocationResponse ActivityByLocation($format, $apikey, $campaignID)

Activity By Location

Returns a detailed report of your campaign opens (unique and total) by country.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the requested campaign | 

### Return type

[**ActivityByLocationResponse**](ActivityByLocationResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CampaignSummary**
> CampaignSummaryResponse CampaignSummary($format, $apikey, $campaignID)

Campaign Summary

Provides a basic summary of the results for any sent campaign such as the number of recipients, opens, clicks, bounces, unsubscribes, forwards etc. to date.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the requested campaign | 

### Return type

[**CampaignSummaryResponse**](CampaignSummaryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloningAnExistingCampaign**
> CloningAnExistingCampaignResponse CloningAnExistingCampaign($format, $campaignID, $apikey)

Cloning An Existing Campaign

Creates an exact copy of an existing campaign. The new campaign is created as a draft.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **campaignID** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 

### Return type

[**CloningAnExistingCampaignResponse**](CloningAnExistingCampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatingADraftCampaign**
> CreatingADraftCampaignResponse CreatingADraftCampaign($format, $apikey, $body)

Creating A Draft Campaign

Creates a new campaign in your account. This method does not send the campaign, but rather creates it as a draft, ready for sending or testing.  You can choose to send either a regular campaign or an AB split campaign. Campaign content must be specified from a web location.  Ignore ***(A/B Split Campaign Option)*** if you want to create a regular campaign.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **body** | [**CreatingADraftCampaignRequest**](CreatingADraftCampaignRequest.md)|  | 

### Return type

[**CreatingADraftCampaignResponse**](CreatingADraftCampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletingACampaign**
> DeletingACampaignResponse DeletingACampaign($format, $apikey, $campaignID)

Deleting A Campaign

Deletes a campaign from your account, draft or even sent.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the draft campaign to update. | 

### Return type

[**DeletingACampaignResponse**](DeletingACampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllCampaigns**
> GetAllCampaignsResponse GetAllCampaigns($format, $apikey)

Get All Campaigns

Returns a list of all campaigns in your account with detailed information. Because the results for this call could be quite big, paging information is required as input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 

### Return type

[**GetAllCampaignsResponse**](GetAllCampaignsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignStatisticsWithPagingFiltered**
> GetCampaignStatisticsWithPagingFilteredResponse GetCampaignStatisticsWithPagingFiltered($format, $apikey, $campaignID, $type_, $page, $pageSize, $from, $to)

Get Campaign Statistics With Paging & Filtered

Returns a detailed list of statistics for a given campaign based on activity such as emails sent, opened, bounced, link clicked, etc.  Because the results for this call could be quite big, paging information is required as input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the requested AB test campaign | 
 **type_** | **string**| The type of the activity to display results for. This must be one of the following values : * Sent : To get information about when and to which recipients the campaign was sent. * Opened : To get information about who opened the campaign. * LinkClicked : To get information about who clicked on which link. * Forward : To get information about who forwarded the campaign using the relevant link on the email body and when. * Unsubscribed : To get information about who unsubscribed from the campaign by clicking on the unsubscribe link and when. * Bounced : To get information about which email recipients failed to receive the campaign. If not specified, the value Sent will be used by default. | 
 **page** | **string**| The page number to display results for. If not specified, the first page will be returned. | [optional] 
 **pageSize** | **string**| The maximum number of results per page. This must be a positive integer up to 100. If not specified, 50 results per page will be returned.  If a value greater than 100 is specified, it will be treated as 100. | [optional] 
 **from** | **string**| A date value that specifies since when to start returning results. If omitted, results will be returned from the moment the campaign was sent. | [optional] 
 **to** | **string**| A date value that specifies up to when to return results. If omitted, results will be returned up to date. | [optional] 

### Return type

[**GetCampaignStatisticsWithPagingFilteredResponse**](GetCampaignStatisticsWithPaging&amp;FilteredResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsByPage**
> GetCampaignsByPageResponse GetCampaignsByPage($format, $apikey, $page)

Get Campaigns By Page

Returns a list of all campaigns in your account with detailed information. Because the results for this call could be quite big, paging information is required as input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **page** | **float64**| The page number to display results for. | 

### Return type

[**GetCampaignsByPageResponse**](GetCampaignsByPageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsByPageAndPagesize**
> GetCampaignsByPageAndPagesizeResponse GetCampaignsByPageAndPagesize($format, $apikey, $page, $pageSize, $shortBy, $sortMethod)

Get Campaigns By Page And Pagesize

Returns a list of all campaigns in your account with detailed information. Because the results for this call could be quite big, paging information is required as input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **page** | **float64**| The page number to display results for. | 
 **pageSize** | **float64**| The maximum number of results per page.  This must be a positive integer up to 100. If not specified, 50 results per page will be returned.  If a value greater than 100 is specified, it will be treated as 100. | 
 **shortBy** | **string**| The name of the campaign property to sort results by. If not specified, results will be sorted by the CreatedOn property | [optional] 
 **sortMethod** | **string**| The method to sort results: ASC for ascending, DESC for descending. If not specified, &#x60;ASC&#x60; will be assumed | [optional] 

### Return type

[**GetCampaignsByPageAndPagesizeResponse**](GetCampaignsByPageAndPagesizeResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingAllYourSenders**
> GettingAllYourSendersResponse GettingAllYourSenders($format, $apikey)

Getting All Your Senders

Gets a list of your active senders in your account. You may specify any email address of these senders when sending a campaign.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 

### Return type

[**GettingAllYourSendersResponse**](GettingAllYourSendersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingCampaignDetails**
> GettingCampaignDetailsResponse GettingCampaignDetails($format, $apikey, $campaignID)

Getting Campaign Details

Returns a complete set of properties that describe the requested campaign in detail. No statistics are included in the result.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the requested campaign | 

### Return type

[**GettingCampaignDetailsResponse**](GettingCampaignDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingSenderDetails**
> GettingSenderDetailsResponse GettingSenderDetails($format, $email, $apikey)

Getting Sender Details

Returns basic information for the specified sender identified by its email address.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **email** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 

### Return type

[**GettingSenderDetailsResponse**](GettingSenderDetailsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkActivity**
> LinkActivityResponse LinkActivity($format, $apikey, $campaignID)

Link Activity

Returns a list with your campaign links and how many clicks have been made by your recipients, either unique or total.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the requested campaign | 

### Return type

[**LinkActivityResponse**](LinkActivityResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SchedulingACampaign**
> SchedulingACampaignResponse SchedulingACampaign($format, $apikey, $campaignID, $body)

Scheduling A Campaign

Assigns a scheduled date and time at which the campaign will be delivered.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the campaign to be scheduled | 
 **body** | [**SchedulingACampaignRequest**](SchedulingACampaignRequest.md)|  | 

### Return type

[**SchedulingACampaignResponse**](SchedulingACampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendingACampaign**
> SendingACampaignResponse SendingACampaign($format, $apikey, $campaignID)

Sending a campaign

Sends an existing draft campaign to all recipients specified in its mailing list. The campaign is sent immediatelly.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the draft campaign to be sent. | 

### Return type

[**SendingACampaignResponse**](SendingACampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TestingACampaign**
> TestingACampaignResponse TestingACampaign($format, $apikey, $campaignID, $body)

Testing a campaign

Sends a test email of a draft campaign to a list of email addresses you specify for previewing.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the draft campaign to be tested. | 
 **body** | [**TestingACampaignRequest**](TestingACampaignRequest.md)|  | 

### Return type

[**TestingACampaignResponse**](TestingACampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnschedulingACampaign**
> UnschedulingACampaignResponse UnschedulingACampaign($format, $apikey, $campaignID)

Unscheduling a campaign

Removes a previously defined scheduled date and time from a campaign, so that it will be delivered immediately if already queued or when sent.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the campaign to be scheduled | 

### Return type

[**UnschedulingACampaignResponse**](UnschedulingACampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatingADraftCampaign**
> UpdatingADraftCampaignResponse UpdatingADraftCampaign($format, $apikey, $campaignID, $body)

Updating A Draft Campaign

Updates properties of an existing draft A/B campaign in your account. Non-draft campaigns cannot be updated. Ignore ***(A/B Split Campaign Option)*** if you want to create a regular campaign.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **campaignID** | **string**| The ID of the draft campaign to update. | 
 **body** | [**UpdatingADraftCampaignRequest**](UpdatingADraftCampaignRequest.md)|  | 

### Return type

[**UpdatingADraftCampaignResponse**](UpdatingADraftCampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

