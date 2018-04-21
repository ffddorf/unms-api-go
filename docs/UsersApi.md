# \UsersApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGet**](UsersApi.md#UsersGet) | **Get** /users | Return list of users
[**UsersIdDelete**](UsersApi.md#UsersIdDelete) | **Delete** /users/{id} | Delete user
[**UsersIdPut**](UsersApi.md#UsersIdPut) | **Put** /users/{id} | Updates user
[**UsersIdReinvitePost**](UsersApi.md#UsersIdReinvitePost) | **Post** /users/{id}/reinvite | Reinvites user by email
[**UsersPost**](UsersApi.md#UsersPost) | **Post** /users | Creates new user


# **UsersGet**
> []User UsersGet(ctx, xAuthToken)
Return list of users

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**[]User**](User.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdDelete**
> Status UsersIdDelete(ctx, xAuthToken, id)
Delete user

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

# **UsersIdPut**
> User UsersIdPut(ctx, xAuthToken, id, body)
Updates user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**User**](User.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersIdReinvitePost**
> Status UsersIdReinvitePost(ctx, xAuthToken, id)
Reinvites user by email

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

# **UsersPost**
> User UsersPost(ctx, xAuthToken, body)
Creates new user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**Body5**](Body5.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

