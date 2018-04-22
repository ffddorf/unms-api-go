# \UserApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGet**](UserApi.md#UserGet) | **Get** /user | Get the authenticated user
[**UserLoginPost**](UserApi.md#UserLoginPost) | **Post** /user/login | Login
[**UserLoginTotpauthPost**](UserApi.md#UserLoginTotpauthPost) | **Post** /user/login/totpauth | Two Factor Authentication login step 2
[**UserLogoutPost**](UserApi.md#UserLogoutPost) | **Post** /user/logout | Logout
[**UserPasswordRequestresetPost**](UserApi.md#UserPasswordRequestresetPost) | **Post** /user/password/requestreset | Login
[**UserPasswordResetPost**](UserApi.md#UserPasswordResetPost) | **Post** /user/password/reset | Login
[**UserProfileGet**](UserApi.md#UserProfileGet) | **Get** /user/profile | Get the authenticated user profile
[**UserProfilePut**](UserApi.md#UserProfilePut) | **Put** /user/profile | Updates authenticated user
[**UserPut**](UserApi.md#UserPut) | **Put** /user | Updates authenticated user
[**UserTotpauthGet**](UserApi.md#UserTotpauthGet) | **Get** /user/totpauth | Gets new information for two factor authentication
[**UserTotpauthPut**](UserApi.md#UserTotpauthPut) | **Put** /user/totpauth | Sets two factor authentication for user


# **UserGet**
> User UserGet(ctx, xAuthToken)
Get the authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**User**](User.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLoginPost**
> User UserLoginPost(ctx, body)
Login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**Login**](Login.md)|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLoginTotpauthPost**
> User UserLoginTotpauthPost(ctx, body)
Two Factor Authentication login step 2

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**Body1**](Body1.md)|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogoutPost**
> Status UserLogoutPost(ctx, xAuthToken)
Logout

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

# **UserPasswordRequestresetPost**
> Status UserPasswordRequestresetPost(ctx, body)
Login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**Body3**](Body3.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPasswordResetPost**
> Status UserPasswordResetPost(ctx, body)
Login

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**Body4**](Body4.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserProfileGet**
> UserProfile UserProfileGet(ctx, xAuthToken)
Get the authenticated user profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**UserProfile**](UserProfile.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserProfilePut**
> UserProfile UserProfilePut(ctx, xAuthToken, body)
Updates authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**UserProfile**](UserProfile.md)|  | 

### Return type

[**UserProfile**](UserProfile.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPut**
> User UserPut(ctx, xAuthToken, body)
Updates authenticated user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**Body**](Body.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserTotpauthGet**
> TwoFactorSecret UserTotpauthGet(ctx, xAuthToken)
Gets new information for two factor authentication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**TwoFactorSecret**](TwoFactorSecret.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserTotpauthPut**
> User UserTotpauthPut(ctx, xAuthToken, body)
Sets two factor authentication for user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**Body2**](Body2.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

