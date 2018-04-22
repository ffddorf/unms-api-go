# \DiscoveryApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoveryConnectPost**](DiscoveryApi.md#DiscoveryConnectPost) | **Post** /discovery/connect | Batch operation - connects given devices to UNMS
[**DiscoveryCredentialsPost**](DiscoveryApi.md#DiscoveryCredentialsPost) | **Post** /discovery/credentials | Batch operation - set new credentials of given devices
[**DiscoveryDelete**](DiscoveryApi.md#DiscoveryDelete) | **Delete** /discovery | Delete and/or terminate the discovery process
[**DiscoveryGet**](DiscoveryApi.md#DiscoveryGet) | **Get** /discovery | Fetch result of the last scan
[**DiscoveryStartPost**](DiscoveryApi.md#DiscoveryStartPost) | **Post** /discovery/start | Start new discovery run and terminate any ongoing for current user
[**DiscoveryStopPost**](DiscoveryApi.md#DiscoveryStopPost) | **Post** /discovery/stop | Cancel ongoing discovery task, discovery result is not deleted


# **DiscoveryConnectPost**
> Status DiscoveryConnectPost(ctx, xAuthToken, payload)
Batch operation - connects given devices to UNMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **payload** | [**Payload7**](Payload7.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoveryCredentialsPost**
> Status DiscoveryCredentialsPost(ctx, xAuthToken, payload)
Batch operation - set new credentials of given devices

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **payload** | [**Payload6**](Payload6.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoveryDelete**
> Status DiscoveryDelete(ctx, xAuthToken)
Delete and/or terminate the discovery process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoveryGet**
> DiscoveryResult DiscoveryGet(ctx, xAuthToken)
Fetch result of the last scan

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**DiscoveryResult**](DiscoveryResult.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoveryStartPost**
> DiscoveryResult DiscoveryStartPost(ctx, xAuthToken, payload)
Start new discovery run and terminate any ongoing for current user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **payload** | [**Payload5**](Payload5.md)|  | 

### Return type

[**DiscoveryResult**](DiscoveryResult.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoveryStopPost**
> DiscoveryResult DiscoveryStopPost(ctx, xAuthToken)
Cancel ongoing discovery task, discovery result is not deleted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**DiscoveryResult**](DiscoveryResult.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

