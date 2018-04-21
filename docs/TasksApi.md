# \TasksApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksGet**](TasksApi.md#TasksGet) | **Get** /tasks | List all tasks
[**TasksIdDelete**](TasksApi.md#TasksIdDelete) | **Delete** /tasks/{id} | Cancel a task
[**TasksIdGet**](TasksApi.md#TasksIdGet) | **Get** /tasks/{id} | Returns a mass upgrade task inner task items
[**TasksInProgressGet**](TasksApi.md#TasksInProgressGet) | **Get** /tasks/in-progress | Return number of tasks in progress


# **TasksGet**
> InlineResponse2003 TasksGet(ctx, xAuthToken, count, page, optional)
List all tasks

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
 **status** | **string**|  | 
 **period** | **float32**| unix timestamp in milliseconds | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TasksIdDelete**
> Status TasksIdDelete(ctx, xAuthToken, id)
Cancel a task

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TasksIdGet**
> InlineResponse2003 TasksIdGet(ctx, xAuthToken, id, count, page)
Returns a mass upgrade task inner task items

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **count** | **float32**|  | 
  **page** | **float32**|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TasksInProgressGet**
> InlineResponse2004 TasksInProgressGet(ctx, xAuthToken)
Return number of tasks in progress

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

