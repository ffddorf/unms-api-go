# \NmsApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NmsConnectionGet**](NmsApi.md#NmsConnectionGet) | **Get** /nms/connection | Return connect string for unms instance
[**NmsEnumsGet**](NmsApi.md#NmsEnumsGet) | **Get** /nms/enums | Return current codebooks
[**NmsMailserverGet**](NmsApi.md#NmsMailserverGet) | **Get** /nms/mailserver | Fetch mail server settings
[**NmsMailserverPut**](NmsApi.md#NmsMailserverPut) | **Put** /nms/mailserver | Update mail server settings
[**NmsMailserverTestPost**](NmsApi.md#NmsMailserverTestPost) | **Post** /nms/mailserver/test | test mail server settings by sending email
[**NmsMaintenanceBackupGet**](NmsApi.md#NmsMaintenanceBackupGet) | **Get** /nms/maintenance/backup | Download data backup
[**NmsMaintenanceBackupPut**](NmsApi.md#NmsMaintenanceBackupPut) | **Put** /nms/maintenance/backup | Upload data backup
[**NmsMaintenanceBackupRestoreGet**](NmsApi.md#NmsMaintenanceBackupRestoreGet) | **Get** /nms/maintenance/backup/restore | Restores uploaded backup file
[**NmsMaintenanceSupportinfoGet**](NmsApi.md#NmsMaintenanceSupportinfoGet) | **Get** /nms/maintenance/supportinfo | Download support info
[**NmsRefreshCertificatePut**](NmsApi.md#NmsRefreshCertificatePut) | **Put** /nms/refresh-certificate | Refresh NMS certificate
[**NmsServerConfigGet**](NmsApi.md#NmsServerConfigGet) | **Get** /nms/server-config | Return UNMS server configuration
[**NmsSettingsGet**](NmsApi.md#NmsSettingsGet) | **Get** /nms/settings | Get nms settings
[**NmsSettingsPut**](NmsApi.md#NmsSettingsPut) | **Put** /nms/settings | Updated nms settings
[**NmsSetupGet**](NmsApi.md#NmsSetupGet) | **Get** /nms/setup | Return status of UNMS configuration
[**NmsSetupPost**](NmsApi.md#NmsSetupPost) | **Post** /nms/setup | Set up the UNMS instance
[**NmsStatisticsGet**](NmsApi.md#NmsStatisticsGet) | **Get** /nms/statistics | Get UNMS network statistics
[**NmsSummaryGet**](NmsApi.md#NmsSummaryGet) | **Get** /nms/summary | Various badge-count like values, e.g. unread logs count
[**NmsUpdateGet**](NmsApi.md#NmsUpdateGet) | **Get** /nms/update | Get nms update status
[**NmsUpdatePut**](NmsApi.md#NmsUpdatePut) | **Put** /nms/update | Request NMS update


# **NmsConnectionGet**
> string NmsConnectionGet(ctx, xAuthToken)
Return connect string for unms instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

**string**

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsEnumsGet**
> Enums NmsEnumsGet(ctx, )
Return current codebooks

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Enums**](Enums.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsMailserverGet**
> MailServer NmsMailserverGet(ctx, xAuthToken)
Fetch mail server settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**MailServer**](MailServer.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsMailserverPut**
> MailServer NmsMailserverPut(ctx, xAuthToken, body)
Update mail server settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**MailServer**](MailServer.md)|  | 

### Return type

[**MailServer**](MailServer.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsMailserverTestPost**
> Status NmsMailserverTestPost(ctx, xAuthToken, body)
test mail server settings by sending email

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**Body8**](Body8.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsMaintenanceBackupGet**
> *os.File NmsMaintenanceBackupGet(ctx, xAuthToken, retention)
Download data backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **retention** | **float32**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsMaintenanceBackupPut**
> Status NmsMaintenanceBackupPut(ctx, xAuthToken, file)
Upload data backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **file** | ***os.File**| The uploaded backup file | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsMaintenanceBackupRestoreGet**
> Status NmsMaintenanceBackupRestoreGet(ctx, xAuthToken)
Restores uploaded backup file

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

# **NmsMaintenanceSupportinfoGet**
> *os.File NmsMaintenanceSupportinfoGet(ctx, xAuthToken)
Download support info

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsRefreshCertificatePut**
> Status NmsRefreshCertificatePut(ctx, xAuthToken)
Refresh NMS certificate

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

# **NmsServerConfigGet**
> InlineResponse2001 NmsServerConfigGet(ctx, )
Return UNMS server configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsSettingsGet**
> NmsSettings NmsSettingsGet(ctx, xAuthToken)
Get nms settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**NmsSettings**](NmsSettings.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsSettingsPut**
> NmsSettings NmsSettingsPut(ctx, xAuthToken, body)
Updated nms settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**NmsSettings**](NmsSettings.md)|  | 

### Return type

[**NmsSettings**](NmsSettings.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsSetupGet**
> InlineResponse200 NmsSetupGet(ctx, )
Return status of UNMS configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsSetupPost**
> Status NmsSetupPost(ctx, body)
Set up the UNMS instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**Body6**](Body6.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsStatisticsGet**
> NmsStatistics NmsStatisticsGet(ctx, xAuthToken, interval)
Get UNMS network statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **interval** | **string**|  | 

### Return type

[**NmsStatistics**](NMSStatistics.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsSummaryGet**
> NmsSummary NmsSummaryGet(ctx, xAuthToken, outagesTimestamp, logsTimestamp, logsLevel)
Various badge-count like values, e.g. unread logs count

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **outagesTimestamp** | **string**|  | 
  **logsTimestamp** | **string**|  | 
  **logsLevel** | [**[]string**](string.md)|  | 

### Return type

[**NmsSummary**](NmsSummary.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsUpdateGet**
> NmsUpdateStatus NmsUpdateGet(ctx, xAuthToken)
Get nms update status

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**NmsUpdateStatus**](NmsUpdateStatus.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NmsUpdatePut**
> Status NmsUpdatePut(ctx, xAuthToken, body)
Request NMS update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **body** | [**Body7**](Body7.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

