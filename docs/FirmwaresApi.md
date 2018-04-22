# \FirmwaresApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FirmwaresDeletePost**](FirmwaresApi.md#FirmwaresDeletePost) | **Post** /firmwares/delete | Batch firmware delete
[**FirmwaresGet**](FirmwaresApi.md#FirmwaresGet) | **Get** /firmwares | Fetch fetch available firmwares
[**FirmwaresPost**](FirmwaresApi.md#FirmwaresPost) | **Post** /firmwares | Upload new firmware image


# **FirmwaresDeletePost**
> []Firmware FirmwaresDeletePost(ctx, xAuthToken, payload)
Batch firmware delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **payload** | [**[]Payload8**](payload_8.md)|  | 

### Return type

[**[]Firmware**](Firmware.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirmwaresGet**
> []Firmware FirmwaresGet(ctx, xAuthToken)
Fetch fetch available firmwares

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**[]Firmware**](Firmware.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FirmwaresPost**
> Firmware FirmwaresPost(ctx, xAuthToken, file)
Upload new firmware image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **file** | ***os.File**| The uploaded firmware file or multiple files | 

### Return type

[**Firmware**](Firmware.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

