# \SubscribersApi

All URIs are relative to *https://api.moosend.com/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddingMultipleSubscribers**](SubscribersApi.md#AddingMultipleSubscribers) | **Post** /subscribers/{MailingListID}/subscribe_many.{Format} | Adding multiple subscribers
[**AddingSubscribers**](SubscribersApi.md#AddingSubscribers) | **Post** /subscribers/{MailingListID}/subscribe.{Format} | Adding subscribers
[**GetSubscriberByEmailAddress**](SubscribersApi.md#GetSubscriberByEmailAddress) | **Get** /subscribers/{MailingListID}/view.{Format} | Get subscriber by email address
[**GetSubscriberById**](SubscribersApi.md#GetSubscriberById) | **Get** /subscribers/{MailingListID}/find/{SubscriberID}.{Format} | Get subscriber by id
[**GettingSubscribers**](SubscribersApi.md#GettingSubscribers) | **Get** /lists/{MailingListID}/subscribers/{Status}.{Format} | Getting subscribers
[**RemovingASubscriber**](SubscribersApi.md#RemovingASubscriber) | **Post** /subscribers/{MailingListID}/remove.{Format} | Removing a subscriber
[**RemovingMultipleSubscribers**](SubscribersApi.md#RemovingMultipleSubscribers) | **Post** /subscribers/{MailingListID}/remove_many.{Format} | Removing multiple subscribers
[**UnsubscribingSubscribersFromAccount**](SubscribersApi.md#UnsubscribingSubscribersFromAccount) | **Post** /subscribers/unsubscribe.{Format} | Unsubscribing subscribers from account
[**UnsubscribingSubscribersFromMailingList**](SubscribersApi.md#UnsubscribingSubscribersFromMailingList) | **Post** /subscribers/{MailingListID}/unsubscribe.{Format} | Unsubscribing subscribers from mailing list
[**UnsubscribingSubscribersFromMailingListAndASpecifiedCampaign**](SubscribersApi.md#UnsubscribingSubscribersFromMailingListAndASpecifiedCampaign) | **Post** /subscribers/{MailingListID}/{CampaignID}/unsubscribe.{Format} | Unsubscribing subscribers from mailing list and a specified campaign
[**UpdatingASubscriber**](SubscribersApi.md#UpdatingASubscriber) | **Post** /subscribers/{MailingListID}/update/{SubscriberID}.{Format} | Updating a subscriber


# **AddingMultipleSubscribers**
> AddingMultipleSubscribersResponse AddingMultipleSubscribers($format, $apikey, $mailingListID, $body)

Adding multiple subscribers

This method allows you to add multiple subscribers in a mailing list with a single call. If some subscribers already exist with the given email addresses, they will be updated. If you try to add a subscriber with an invalid email address, this attempt will be ignored, as the process will skip to the next subscriber automatically.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to add subscribers to. | 
 **body** | [**AddingMultipleSubscribersRequest**](AddingMultipleSubscribersRequest.md)|  | 

### Return type

[**AddingMultipleSubscribersResponse**](AddingMultipleSubscribersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddingSubscribers**
> AddingSubscribersResponse AddingSubscribers($format, $mailingListID, $apikey, $body)

Adding subscribers

Adds a new subscriber to the specified mailing list. If there is already a subscriber with the specified email address in the list, an update will be performed instead.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list to add the new member. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **body** | [**AddingSubscribersRequest**](AddingSubscribersRequest.md)|  | 

### Return type

[**AddingSubscribersResponse**](AddingSubscribersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriberByEmailAddress**
> GetSubscriberByEmailAddressResponse GetSubscriberByEmailAddress($format, $apikey, $mailingListID, $email)

Get subscriber by email address

Searches for a subscriber with the specified email address in the specified mailing list and returns detailed information such as id, name, date created, date unsubscribed, status and custom fields


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list where the subscriber belongs. | 
 **email** | **string**| The email of the subscriber. | 

### Return type

[**GetSubscriberByEmailAddressResponse**](GetSubscriberByEmailAddressResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubscriberById**
> GetSubscriberByIdResponse GetSubscriberById($format, $apikey, $mailingListID, $subscriberID)

Get subscriber by id

Searches for a subscriber with the specified unique id in the specified mailing list and returns detailed information such as email, name, date created, date unsubscribed, status and custom fields.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to search the subscriber in. | 
 **subscriberID** | **string**| The id of the subscriber being searched. | 

### Return type

[**GetSubscriberByIdResponse**](GetSubscriberByIdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GettingSubscribers**
> GettingSubscribersResponse GettingSubscribers($format, $mailingListID, $apikey, $status, $page, $pageSize)

Getting subscribers

Gets a list of all subscribers in a given mailing list. You may filter the list by setting a date to fetch those subscribed since then and/or by their status. Because the results for this call could be quite big, paging information is required as input.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **mailingListID** | **string**| The ID of the mailing list where the subscribers belong. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **status** | **string**| Specifies what type of subscriber statistics results will be returned. | 
 **page** | **float64**| Specifies the page of subscriber statistics results will be returned. | [optional] 
 **pageSize** | **float64**| Specifies the page size of subscriber statistics results will be returned. | [optional] 

### Return type

[**GettingSubscribersResponse**](GettingSubscribersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovingASubscriber**
> RemovingASubscriberResponse RemovingASubscriber($format, $apikey, $mailingListID, $body)

Removing a subscriber

Removes a subscriber from the specified mailing list permanently (without moving to the suppression list).


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to remove the subscriber from. | 
 **body** | [**RemovingASubscriberRequest**](RemovingASubscriberRequest.md)|  | 

### Return type

[**RemovingASubscriberResponse**](RemovingASubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemovingMultipleSubscribers**
> RemovingMultipleSubscribersResponse RemovingMultipleSubscribers($format, $apikey, $mailingListID, $body)

Removing multiple subscribers

Removes a list of subscribers from the specified mailing list permanently (without putting them in the suppression list). Any invalid email addresses specified will be ignored.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to remove the subscribers from. | 
 **body** | [**RemovingMultipleSubscribersRequest**](RemovingMultipleSubscribersRequest.md)|  | 

### Return type

[**RemovingMultipleSubscribersResponse**](RemovingMultipleSubscribersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribingSubscribersFromAccount**
> UnsubscribingSubscribersFromAccountResponse UnsubscribingSubscribersFromAccount($format, $apikey, $body)

Unsubscribing subscribers from account

Unsubscribes a subscriber from the account.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **body** | [**UnsubscribingSubscribersFromAccountRequest**](UnsubscribingSubscribersFromAccountRequest.md)|  | 

### Return type

[**UnsubscribingSubscribersFromAccountResponse**](UnsubscribingSubscribersFromAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribingSubscribersFromMailingList**
> UnsubscribingSubscribersFromMailingListResponse UnsubscribingSubscribersFromMailingList($format, $apikey, $mailingListID, $body)

Unsubscribing subscribers from mailing list

Unsubscribes a subscriber from the specified mailing list. The subscriber is not deleted, but moved to the suppression list.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to add subscribers to. | 
 **body** | [**UnsubscribingSubscribersFromMailingListRequest**](UnsubscribingSubscribersFromMailingListRequest.md)|  | 

### Return type

[**UnsubscribingSubscribersFromMailingListResponse**](UnsubscribingSubscribersFromMailingListResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribingSubscribersFromMailingListAndASpecifiedCampaign**
> UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignResponse UnsubscribingSubscribersFromMailingListAndASpecifiedCampaign($format, $campaignID, $apikey, $mailingListID, $body)

Unsubscribing subscribers from mailing list and a specified campaign

Unsubscribes a subscriber from the specified mailing list and the specified campaign. The subscriber is not deleted, but moved to the suppression list.  This call will take into account the setting you have in \"suppression list and unsubscribe settings\" and will remove the subscriber from all other mailing lists or not accordingly.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **campaignID** | **string**| The ID of the campaign that was sent to the specific mailing list. | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list to remove the subscriber from. | 
 **body** | [**UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignRequest**](UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignRequest.md)|  | 

### Return type

[**UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignResponse**](UnsubscribingSubscribersFromMailingListAndASpecifiedCampaignResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatingASubscriber**
> UpdatingASubscriberResponse UpdatingASubscriber($format, $apikey, $mailingListID, $subscriberID, $body)

Updating a subscriber

Updates a subscriber in the specified mailing list. You can even update the subscribers email, if he has not unsubscribed.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**|  | 
 **apikey** | **string**| You may find your API Key or generate a new one in your account settings. | 
 **mailingListID** | **string**| The ID of the mailing list that contains the subscriber | 
 **subscriberID** | **string**| The id of the subscriber to be updated | 
 **body** | [**UpdatingASubscriberRequest**](UpdatingASubscriberRequest.md)|  | 

### Return type

[**UpdatingASubscriberResponse**](UpdatingASubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

