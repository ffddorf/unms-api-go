# \LogsApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogsGet**](LogsApi.md#LogsGet) | **Get** /logs | List of all log items (former alerts/events)
[**LogsUnreadGet**](LogsApi.md#LogsUnreadGet) | **Get** /logs/unread | Number of unread logItems for logged-in user


# **LogsGet**
> InlineResponse2005 LogsGet(ctx, xAuthToken, count, page, optional)
List of all log items (former alerts/events)

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
 **level** | **string**|  | 
 **period** | **float32**| unix timestamp in milliseconds | 
 **query** | **string**| search pattern | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsUnreadGet**
> InlineResponse2004 LogsUnreadGet(ctx, xAuthToken, timestamp, optional)
Number of unread logItems for logged-in user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **timestamp** | **float32**| unix timestamp in ms | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string**| User authorization token | 
 **timestamp** | **float32**| unix timestamp in ms | 
 **level** | **string**|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

