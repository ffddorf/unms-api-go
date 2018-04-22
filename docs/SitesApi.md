# \SitesApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SitesGet**](SitesApi.md#SitesGet) | **Get** /sites | List of all sites in network
[**SitesIdDelete**](SitesApi.md#SitesIdDelete) | **Delete** /sites/{id} | Delete site and unauthorize all connected devices
[**SitesIdGet**](SitesApi.md#SitesIdGet) | **Get** /sites/{id} | Return a site&#39;s detail
[**SitesIdImagesGet**](SitesApi.md#SitesIdImagesGet) | **Get** /sites/{id}/images | Return all site images sorted by image order
[**SitesIdImagesPost**](SitesApi.md#SitesIdImagesPost) | **Post** /sites/{id}/images | Upload new image and create image thumbnail
[**SitesIdPut**](SitesApi.md#SitesIdPut) | **Put** /sites/{id} | Update site
[**SitesPost**](SitesApi.md#SitesPost) | **Post** /sites | Create new site
[**SitesSiteIdImagesImageIdDelete**](SitesApi.md#SitesSiteIdImagesImageIdDelete) | **Delete** /sites/{siteId}/images/{imageId} | Delete image
[**SitesSiteIdImagesImageIdGet**](SitesApi.md#SitesSiteIdImagesImageIdGet) | **Get** /sites/{siteId}/images/{imageId} | Return image file
[**SitesSiteIdImagesImageIdPatch**](SitesApi.md#SitesSiteIdImagesImageIdPatch) | **Patch** /sites/{siteId}/images/{imageId} | Update image
[**SitesSiteIdImagesImageIdReorderPost**](SitesApi.md#SitesSiteIdImagesImageIdReorderPost) | **Post** /sites/{siteId}/images/{imageId}/reorder | Change image order
[**SitesSiteIdImagesImageIdRotateleftPost**](SitesApi.md#SitesSiteIdImagesImageIdRotateleftPost) | **Post** /sites/{siteId}/images/{imageId}/rotateleft | Rotate the image 90 degrees to left
[**SitesSiteIdImagesImageIdRotaterightPost**](SitesApi.md#SitesSiteIdImagesImageIdRotaterightPost) | **Post** /sites/{siteId}/images/{imageId}/rotateright | Rotate the image 90 degrees to right


# **SitesGet**
> []SiteOverview SitesGet(ctx, xAuthToken, optional)
List of all sites in network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string**| User authorization token | 
 **id** | **string**|  | 

### Return type

[**[]SiteOverview**](SiteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesIdDelete**
> Status SitesIdDelete(ctx, xAuthToken, id)
Delete site and unauthorize all connected devices

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

# **SitesIdGet**
> SiteOverview SitesIdGet(ctx, xAuthToken, id)
Return a site's detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**SiteOverview**](SiteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesIdImagesGet**
> []ImageOverview SitesIdImagesGet(ctx, xAuthToken, id)
Return all site images sorted by image order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]ImageOverview**](ImageOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesIdImagesPost**
> Status SitesIdImagesPost(ctx, xAuthToken, id, file)
Upload new image and create image thumbnail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **file** | ***os.File**| The uploaded image or multiple images | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesIdPut**
> SiteOverview SitesIdPut(ctx, xAuthToken, id, payload)
Update site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **payload** | [**SiteOverview**](SiteOverview.md)|  | 

### Return type

[**SiteOverview**](SiteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesPost**
> SiteOverview SitesPost(ctx, xAuthToken, payload)
Create new site

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **payload** | [**Payload**](Payload.md)|  | 

### Return type

[**SiteOverview**](SiteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesSiteIdImagesImageIdDelete**
> Status SitesSiteIdImagesImageIdDelete(ctx, xAuthToken, imageId, siteId)
Delete image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **imageId** | **string**|  | 
  **siteId** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesSiteIdImagesImageIdGet**
> *os.File SitesSiteIdImagesImageIdGet(ctx, xAuthToken, siteId, imageId)
Return image file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **siteId** | **string**|  | 
  **imageId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesSiteIdImagesImageIdPatch**
> ImageOverview SitesSiteIdImagesImageIdPatch(ctx, xAuthToken, siteId, imageId, payload)
Update image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **siteId** | **string**|  | 
  **imageId** | **string**|  | 
  **payload** | [**Payload1**](Payload1.md)|  | 

### Return type

[**ImageOverview**](ImageOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesSiteIdImagesImageIdReorderPost**
> Status SitesSiteIdImagesImageIdReorderPost(ctx, xAuthToken, siteId, imageId, payload)
Change image order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **siteId** | **string**|  | 
  **imageId** | **string**|  | 
  **payload** | [**Payload2**](Payload2.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesSiteIdImagesImageIdRotateleftPost**
> Status SitesSiteIdImagesImageIdRotateleftPost(ctx, xAuthToken, siteId, imageId)
Rotate the image 90 degrees to left

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **siteId** | **string**|  | 
  **imageId** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SitesSiteIdImagesImageIdRotaterightPost**
> Status SitesSiteIdImagesImageIdRotaterightPost(ctx, xAuthToken, siteId, imageId)
Rotate the image 90 degrees to right

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **siteId** | **string**|  | 
  **imageId** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

