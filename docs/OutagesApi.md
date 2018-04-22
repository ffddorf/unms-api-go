# \OutagesApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutagesGet**](OutagesApi.md#OutagesGet) | **Get** /outages | List of all network outages for last month
[**OutagesUnreadGet**](OutagesApi.md#OutagesUnreadGet) | **Get** /outages/unread | Number of unread outageItems for logged-in user


# **OutagesGet**
> InlineResponse2006 OutagesGet(ctx, xAuthToken, count, page, optional)
List of all network outages for last month

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **count** | **float32**|  | 
  **page** | **float32**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string**| User authorization token | 
 **count** | **float32**|  | 
 **page** | **float32**|  | 
 **siteId** | **string**|  | 
 **deviceId** | **string**|  | 
 **type_** | **string**|  | 
 **period** | **float32**| unix timestamp in milliseconds | 
 **query** | **string**| search pattern | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OutagesUnreadGet**
> InlineResponse2004 OutagesUnreadGet(ctx, xAuthToken, timestamp)
Number of unread outageItems for logged-in user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **timestamp** | **float32**| unix timestamp in ms | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

