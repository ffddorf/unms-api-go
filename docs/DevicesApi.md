# \DevicesApi

All URIs are relative to *https://&lt;unms-deployment&gt;/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesAircubesIdGet**](DevicesApi.md#DevicesAircubesIdGet) | **Get** /devices/aircubes/{id} | Return AirCube detail
[**DevicesAirmaxesIdGet**](DevicesApi.md#DevicesAirmaxesIdGet) | **Get** /devices/airmaxes/{id} | Return AirMax detail
[**DevicesApsProfilesGet**](DevicesApi.md#DevicesApsProfilesGet) | **Get** /devices/aps/profiles | List of all access points connection profiles
[**DevicesBackupsPost**](DevicesApi.md#DevicesBackupsPost) | **Post** /devices/backups | Create and download a new multi device backup
[**DevicesDeviceIdBackupsBackupIdApplyPost**](DevicesApi.md#DevicesDeviceIdBackupsBackupIdApplyPost) | **Post** /devices/{deviceId}/backups/{backupId}/apply | Update device configuration from backup file
[**DevicesDeviceIdBackupsBackupIdDelete**](DevicesApi.md#DevicesDeviceIdBackupsBackupIdDelete) | **Delete** /devices/{deviceId}/backups/{backupId} | Delete backup
[**DevicesDeviceIdBackupsBackupIdGet**](DevicesApi.md#DevicesDeviceIdBackupsBackupIdGet) | **Get** /devices/{deviceId}/backups/{backupId} | Return device configuration backup file
[**DevicesDeviceIdInterfacesGet**](DevicesApi.md#DevicesDeviceIdInterfacesGet) | **Get** /devices/{deviceId}/interfaces | Return list of device interfaces
[**DevicesDeviceIdInterfacesInterfaceNameBlockPost**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNameBlockPost) | **Post** /devices/{deviceId}/interfaces/{interfaceName}/block | Block interface
[**DevicesDeviceIdInterfacesInterfaceNameDelete**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNameDelete) | **Delete** /devices/{deviceId}/interfaces/{interfaceName} | Delete interface
[**DevicesDeviceIdInterfacesInterfaceNameGet**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNameGet) | **Get** /devices/{deviceId}/interfaces/{interfaceName} | Get interface configuration
[**DevicesDeviceIdInterfacesInterfaceNameOspfDelete**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNameOspfDelete) | **Delete** /devices/{deviceId}/interfaces/{interfaceName}/ospf | Unset interface OSPF config
[**DevicesDeviceIdInterfacesInterfaceNameOspfPut**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNameOspfPut) | **Put** /devices/{deviceId}/interfaces/{interfaceName}/ospf | Set interface OSPF config
[**DevicesDeviceIdInterfacesInterfaceNamePut**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNamePut) | **Put** /devices/{deviceId}/interfaces/{interfaceName} | Update interface configuration
[**DevicesDeviceIdInterfacesInterfaceNameResetstatsPost**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNameResetstatsPost) | **Post** /devices/{deviceId}/interfaces/{interfaceName}/resetstats | Reset interface statistics
[**DevicesDeviceIdInterfacesInterfaceNameUnblockPost**](DevicesApi.md#DevicesDeviceIdInterfacesInterfaceNameUnblockPost) | **Post** /devices/{deviceId}/interfaces/{interfaceName}/unblock | Unblock interface
[**DevicesDeviceIdInterfacesPppoePost**](DevicesApi.md#DevicesDeviceIdInterfacesPppoePost) | **Post** /devices/{deviceId}/interfaces/pppoe | Create new pppoe interface
[**DevicesDeviceIdInterfacesVlanPost**](DevicesApi.md#DevicesDeviceIdInterfacesVlanPost) | **Post** /devices/{deviceId}/interfaces/vlan | Create new vlan interface
[**DevicesDeviceIdServicesGet**](DevicesApi.md#DevicesDeviceIdServicesGet) | **Get** /devices/{deviceId}/services | Device services
[**DevicesDeviceIdServicesPut**](DevicesApi.md#DevicesDeviceIdServicesPut) | **Put** /devices/{deviceId}/services | Update device services
[**DevicesDeviceIdSystemGet**](DevicesApi.md#DevicesDeviceIdSystemGet) | **Get** /devices/{deviceId}/system | Device system configuration
[**DevicesDeviceIdSystemPut**](DevicesApi.md#DevicesDeviceIdSystemPut) | **Put** /devices/{deviceId}/system | Update device system configuration
[**DevicesDeviceIdSystemUnmsGet**](DevicesApi.md#DevicesDeviceIdSystemUnmsGet) | **Get** /devices/{deviceId}/system/unms | Device specific UNMS settings
[**DevicesDeviceIdSystemUnmsPut**](DevicesApi.md#DevicesDeviceIdSystemUnmsPut) | **Put** /devices/{deviceId}/system/unms | Update device specific unms settings
[**DevicesEroutersIdDelete**](DevicesApi.md#DevicesEroutersIdDelete) | **Delete** /devices/erouters/{id} | 
[**DevicesEroutersIdDhcpLeasesGet**](DevicesApi.md#DevicesEroutersIdDhcpLeasesGet) | **Get** /devices/erouters/{id}/dhcp/leases | DHCP IP address leases
[**DevicesEroutersIdDhcpLeasesPost**](DevicesApi.md#DevicesEroutersIdDhcpLeasesPost) | **Post** /devices/erouters/{id}/dhcp/leases | Create DHCP IP address lease
[**DevicesEroutersIdDhcpLeasesServerNameLeaseIdDelete**](DevicesApi.md#DevicesEroutersIdDhcpLeasesServerNameLeaseIdDelete) | **Delete** /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId} | Delete DHCP IP lease
[**DevicesEroutersIdDhcpLeasesServerNameLeaseIdPut**](DevicesApi.md#DevicesEroutersIdDhcpLeasesServerNameLeaseIdPut) | **Put** /devices/erouters/{id}/dhcp/leases/{serverName}/{leaseId} | Update DHCP IP lease
[**DevicesEroutersIdDhcpServersGet**](DevicesApi.md#DevicesEroutersIdDhcpServersGet) | **Get** /devices/erouters/{id}/dhcp/servers | Device DHCP servers
[**DevicesEroutersIdDhcpServersPost**](DevicesApi.md#DevicesEroutersIdDhcpServersPost) | **Post** /devices/erouters/{id}/dhcp/servers | Create new DHCP server
[**DevicesEroutersIdDhcpServersServerNameBlockPost**](DevicesApi.md#DevicesEroutersIdDhcpServersServerNameBlockPost) | **Post** /devices/erouters/{id}/dhcp/servers/{serverName}/block | Block DHCP server
[**DevicesEroutersIdDhcpServersServerNameDelete**](DevicesApi.md#DevicesEroutersIdDhcpServersServerNameDelete) | **Delete** /devices/erouters/{id}/dhcp/servers/{serverName} | Delete DHCP server
[**DevicesEroutersIdDhcpServersServerNameGet**](DevicesApi.md#DevicesEroutersIdDhcpServersServerNameGet) | **Get** /devices/erouters/{id}/dhcp/servers/{serverName} | Get DHCP server configuration
[**DevicesEroutersIdDhcpServersServerNamePut**](DevicesApi.md#DevicesEroutersIdDhcpServersServerNamePut) | **Put** /devices/erouters/{id}/dhcp/servers/{serverName} | Update DHCP server configuration
[**DevicesEroutersIdDhcpServersServerNameUnblockPost**](DevicesApi.md#DevicesEroutersIdDhcpServersServerNameUnblockPost) | **Post** /devices/erouters/{id}/dhcp/servers/{serverName}/unblock | Unblock dhcp server
[**DevicesEroutersIdGet**](DevicesApi.md#DevicesEroutersIdGet) | **Get** /devices/erouters/{id} | Return erouter detail
[**DevicesEroutersIdRouterOspfAreasAreaDelete**](DevicesApi.md#DevicesEroutersIdRouterOspfAreasAreaDelete) | **Delete** /devices/erouters/{id}/router/ospf/areas/{area} | Delete OSPF area
[**DevicesEroutersIdRouterOspfAreasAreaPut**](DevicesApi.md#DevicesEroutersIdRouterOspfAreasAreaPut) | **Put** /devices/erouters/{id}/router/ospf/areas/{area} | Update OSPF area
[**DevicesEroutersIdRouterOspfAreasGet**](DevicesApi.md#DevicesEroutersIdRouterOspfAreasGet) | **Get** /devices/erouters/{id}/router/ospf/areas | All OSPF areas
[**DevicesEroutersIdRouterOspfAreasPost**](DevicesApi.md#DevicesEroutersIdRouterOspfAreasPost) | **Post** /devices/erouters/{id}/router/ospf/areas | Create new OSPF area
[**DevicesEroutersIdRouterOspfGet**](DevicesApi.md#DevicesEroutersIdRouterOspfGet) | **Get** /devices/erouters/{id}/router/ospf | Get OSPF configuration
[**DevicesEroutersIdRouterOspfPut**](DevicesApi.md#DevicesEroutersIdRouterOspfPut) | **Put** /devices/erouters/{id}/router/ospf | Update OSPF configuration
[**DevicesEroutersIdRouterRoutesBlockPost**](DevicesApi.md#DevicesEroutersIdRouterRoutesBlockPost) | **Post** /devices/erouters/{id}/router/routes/block | Edit static route
[**DevicesEroutersIdRouterRoutesDeletePost**](DevicesApi.md#DevicesEroutersIdRouterRoutesDeletePost) | **Post** /devices/erouters/{id}/router/routes/delete | Delete route
[**DevicesEroutersIdRouterRoutesGet**](DevicesApi.md#DevicesEroutersIdRouterRoutesGet) | **Get** /devices/erouters/{id}/router/routes | All routes
[**DevicesEroutersIdRouterRoutesPost**](DevicesApi.md#DevicesEroutersIdRouterRoutesPost) | **Post** /devices/erouters/{id}/router/routes | Create new static route
[**DevicesEroutersIdRouterRoutesPut**](DevicesApi.md#DevicesEroutersIdRouterRoutesPut) | **Put** /devices/erouters/{id}/router/routes | Edit static route
[**DevicesEroutersIdRouterRoutesUnblockPost**](DevicesApi.md#DevicesEroutersIdRouterRoutesUnblockPost) | **Post** /devices/erouters/{id}/router/routes/unblock | Edit static route
[**DevicesGet**](DevicesApi.md#DevicesGet) | **Get** /devices | List of all devices in network
[**DevicesIdAuthorizePost**](DevicesApi.md#DevicesIdAuthorizePost) | **Post** /devices/{id}/authorize | Authorize device
[**DevicesIdBackupsGet**](DevicesApi.md#DevicesIdBackupsGet) | **Get** /devices/{id}/backups | Return list of device backups
[**DevicesIdBackupsPost**](DevicesApi.md#DevicesIdBackupsPost) | **Post** /devices/{id}/backups | Create new device backup
[**DevicesIdBackupsPut**](DevicesApi.md#DevicesIdBackupsPut) | **Put** /devices/{id}/backups | Upload data backup
[**DevicesIdDelete**](DevicesApi.md#DevicesIdDelete) | **Delete** /devices/{id} | Delete device
[**DevicesIdGet**](DevicesApi.md#DevicesIdGet) | **Get** /devices/{id} | Device status overview
[**DevicesIdLocatePost**](DevicesApi.md#DevicesIdLocatePost) | **Post** /devices/{id}/locate | Start device locating
[**DevicesIdRefreshPost**](DevicesApi.md#DevicesIdRefreshPost) | **Post** /devices/{id}/refresh | Refresh device communication, remove macAesKey
[**DevicesIdResetPost**](DevicesApi.md#DevicesIdResetPost) | **Post** /devices/{id}/reset | Reset device statistics, action clear-traffic-analysis
[**DevicesIdRestartPost**](DevicesApi.md#DevicesIdRestartPost) | **Post** /devices/{id}/restart | Restart device, device action reboot
[**DevicesIdStatisticsGet**](DevicesApi.md#DevicesIdStatisticsGet) | **Get** /devices/{id}/statistics | Return device statistics
[**DevicesIdUnlocatePost**](DevicesApi.md#DevicesIdUnlocatePost) | **Post** /devices/{id}/unlocate | Stop device locating
[**DevicesMacsGet**](DevicesApi.md#DevicesMacsGet) | **Get** /devices/macs/ | Return a list of all devices&#39; MAC address (except for ONUs)
[**DevicesOltsGet**](DevicesApi.md#DevicesOltsGet) | **Get** /devices/olts | List of all OLTs in network
[**DevicesOltsIdDelete**](DevicesApi.md#DevicesOltsIdDelete) | **Delete** /devices/olts/{id} | Delete OLT form UNMS
[**DevicesOltsIdGet**](DevicesApi.md#DevicesOltsIdGet) | **Get** /devices/olts/{id} | Return OLT detail
[**DevicesOltsIdProfilesGet**](DevicesApi.md#DevicesOltsIdProfilesGet) | **Get** /devices/olts/{id}/profiles | List of all OLT ONU profiles
[**DevicesOltsIdProfilesPost**](DevicesApi.md#DevicesOltsIdProfilesPost) | **Post** /devices/olts/{id}/profiles | Create OLT ONU profile
[**DevicesOltsIdProfilesProfileDelete**](DevicesApi.md#DevicesOltsIdProfilesProfileDelete) | **Delete** /devices/olts/{id}/profiles/{profile} | Delete OLT form UNMS
[**DevicesOltsIdProfilesProfilePut**](DevicesApi.md#DevicesOltsIdProfilesProfilePut) | **Put** /devices/olts/{id}/profiles/{profile} | Update OLT ONU profile
[**DevicesOnusGet**](DevicesApi.md#DevicesOnusGet) | **Get** /devices/onus | List of all onus in network
[**DevicesOnusIdBlockPost**](DevicesApi.md#DevicesOnusIdBlockPost) | **Post** /devices/onus/{id}/block | Delete ONU GPON password and set status to blocked
[**DevicesOnusIdClientsGet**](DevicesApi.md#DevicesOnusIdClientsGet) | **Get** /devices/onus/{id}/clients | Get all clients connected to ONU
[**DevicesOnusIdDelete**](DevicesApi.md#DevicesOnusIdDelete) | **Delete** /devices/onus/{id} | Delete ONU device - call OLT action DELETE_ONU (former forget action)
[**DevicesOnusIdGet**](DevicesApi.md#DevicesOnusIdGet) | **Get** /devices/onus/{id} | Return ONU detail
[**DevicesOnusIdResetstatsPost**](DevicesApi.md#DevicesOnusIdResetstatsPost) | **Post** /devices/onus/{id}/resetstats | Reset onu data statistics
[**DevicesOnusIdUnblockPost**](DevicesApi.md#DevicesOnusIdUnblockPost) | **Post** /devices/onus/{id}/unblock | Unblock ONU device - device action UNBLOCK_ONU
[**DevicesOnusIdUpgradePost**](DevicesApi.md#DevicesOnusIdUpgradePost) | **Post** /devices/onus/{id}/upgrade | Upgrade device
[**DevicesSsidsGet**](DevicesApi.md#DevicesSsidsGet) | **Get** /devices/ssids | Get devices wireless configuration


# **DevicesAircubesIdGet**
> AirCubeStatusDetail DevicesAircubesIdGet(ctx, xAuthToken, id)
Return AirCube detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**AirCubeStatusDetail**](AirCubeStatusDetail.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesAirmaxesIdGet**
> AirMaxStatusDetail DevicesAirmaxesIdGet(ctx, xAuthToken, id)
Return AirMax detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**AirMaxStatusDetail**](AirMaxStatusDetail.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesApsProfilesGet**
> []DeviceApsProfile DevicesApsProfilesGet(ctx, xAuthToken)
List of all access points connection profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**[]DeviceApsProfile**](DeviceAPSProfile.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesBackupsPost**
> *os.File DevicesBackupsPost(ctx, xAuthToken, payload)
Create and download a new multi device backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **payload** | [**Payload4**](Payload4.md)|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdBackupsBackupIdApplyPost**
> Status DevicesDeviceIdBackupsBackupIdApplyPost(ctx, xAuthToken, deviceId, backupId)
Update device configuration from backup file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **backupId** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdBackupsBackupIdDelete**
> Status DevicesDeviceIdBackupsBackupIdDelete(ctx, xAuthToken, deviceId, backupId)
Delete backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **backupId** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdBackupsBackupIdGet**
> *os.File DevicesDeviceIdBackupsBackupIdGet(ctx, xAuthToken, deviceId, backupId)
Return device configuration backup file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **backupId** | **string**|  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesGet**
> []InterfaceOverview DevicesDeviceIdInterfacesGet(ctx, xAuthToken, deviceId)
Return list of device interfaces

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 

### Return type

[**[]InterfaceOverview**](InterfaceOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNameBlockPost**
> Status DevicesDeviceIdInterfacesInterfaceNameBlockPost(ctx, xAuthToken, deviceId, interfaceName)
Block interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNameDelete**
> Status DevicesDeviceIdInterfacesInterfaceNameDelete(ctx, xAuthToken, deviceId, interfaceName)
Delete interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNameGet**
> InterfaceConfig DevicesDeviceIdInterfacesInterfaceNameGet(ctx, xAuthToken, deviceId, interfaceName)
Get interface configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 

### Return type

[**InterfaceConfig**](InterfaceConfig.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNameOspfDelete**
> Status DevicesDeviceIdInterfacesInterfaceNameOspfDelete(ctx, xAuthToken, deviceId, interfaceName)
Unset interface OSPF config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNameOspfPut**
> Status DevicesDeviceIdInterfacesInterfaceNameOspfPut(ctx, xAuthToken, deviceId, interfaceName, optional)
Set interface OSPF config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string**|  | 
 **deviceId** | **string**|  | 
 **interfaceName** | **string**|  | 
 **body** | [**OspfInterface**](OspfInterface.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNamePut**
> InterfaceConfig DevicesDeviceIdInterfacesInterfaceNamePut(ctx, xAuthToken, deviceId, interfaceName, body)
Update interface configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 
  **body** | [**InterfaceConfig**](InterfaceConfig.md)|  | 

### Return type

[**InterfaceConfig**](InterfaceConfig.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNameResetstatsPost**
> Status DevicesDeviceIdInterfacesInterfaceNameResetstatsPost(ctx, xAuthToken, deviceId, interfaceName)
Reset interface statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesInterfaceNameUnblockPost**
> Status DevicesDeviceIdInterfacesInterfaceNameUnblockPost(ctx, xAuthToken, deviceId, interfaceName)
Unblock interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **deviceId** | **string**|  | 
  **interfaceName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesPppoePost**
> InterfaceConfig DevicesDeviceIdInterfacesPppoePost(ctx, xAuthToken, deviceId, body)
Create new pppoe interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **body** | [**Body10**](Body10.md)|  | 

### Return type

[**InterfaceConfig**](InterfaceConfig.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdInterfacesVlanPost**
> InterfaceConfig DevicesDeviceIdInterfacesVlanPost(ctx, xAuthToken, deviceId, body)
Create new vlan interface

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **body** | [**Body9**](Body9.md)|  | 

### Return type

[**InterfaceConfig**](InterfaceConfig.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdServicesGet**
> []Services DevicesDeviceIdServicesGet(ctx, xAuthToken, deviceId)
Device services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 

### Return type

[**[]Services**](Services.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdServicesPut**
> Services DevicesDeviceIdServicesPut(ctx, xAuthToken, deviceId, body)
Update device services

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **body** | [**Services**](Services.md)|  | 

### Return type

[**Services**](Services.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdSystemGet**
> []System DevicesDeviceIdSystemGet(ctx, xAuthToken, deviceId)
Device system configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 

### Return type

[**[]System**](System.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdSystemPut**
> System DevicesDeviceIdSystemPut(ctx, xAuthToken, deviceId, body)
Update device system configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **body** | [**System**](System.md)|  | 

### Return type

[**System**](System.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdSystemUnmsGet**
> UnmsSettings DevicesDeviceIdSystemUnmsGet(ctx, xAuthToken, deviceId)
Device specific UNMS settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 

### Return type

[**UnmsSettings**](UnmsSettings.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesDeviceIdSystemUnmsPut**
> UnmsSettings DevicesDeviceIdSystemUnmsPut(ctx, xAuthToken, deviceId, body)
Update device specific unms settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **deviceId** | **string**|  | 
  **body** | [**UnmsSettings**](UnmsSettings.md)|  | 

### Return type

[**UnmsSettings**](UnmsSettings.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDelete**
> Status DevicesEroutersIdDelete(ctx, xAuthToken, id)


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

# **DevicesEroutersIdDhcpLeasesGet**
> []DhcpLease DevicesEroutersIdDhcpLeasesGet(ctx, xAuthToken, id)
DHCP IP address leases

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]DhcpLease**](DHCPLease.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpLeasesPost**
> DhcpLease DevicesEroutersIdDhcpLeasesPost(ctx, xAuthToken, id, body)
Create DHCP IP address lease

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**Body11**](Body11.md)|  | 

### Return type

[**DhcpLease**](DHCPLease.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpLeasesServerNameLeaseIdDelete**
> Status DevicesEroutersIdDhcpLeasesServerNameLeaseIdDelete(ctx, xAuthToken, id, serverName, leaseId)
Delete DHCP IP lease

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **serverName** | **string**|  | 
  **leaseId** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpLeasesServerNameLeaseIdPut**
> DhcpLease DevicesEroutersIdDhcpLeasesServerNameLeaseIdPut(ctx, xAuthToken, id, serverName, leaseId, body)
Update DHCP IP lease

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **serverName** | **string**|  | 
  **leaseId** | **string**|  | 
  **body** | [**DhcpLease**](DhcpLease.md)|  | 

### Return type

[**DhcpLease**](DHCPLease.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpServersGet**
> []DhcpServerConfiguration DevicesEroutersIdDhcpServersGet(ctx, xAuthToken, id)
Device DHCP servers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]DhcpServerConfiguration**](DHCPServerConfiguration.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpServersPost**
> DhcpServerConfiguration DevicesEroutersIdDhcpServersPost(ctx, xAuthToken, id, body)
Create new DHCP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**DhcpServerConfiguration**](DhcpServerConfiguration.md)|  | 

### Return type

[**DhcpServerConfiguration**](DHCPServerConfiguration.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpServersServerNameBlockPost**
> Status DevicesEroutersIdDhcpServersServerNameBlockPost(ctx, xAuthToken, id, serverName)
Block DHCP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **id** | **string**|  | 
  **serverName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpServersServerNameDelete**
> Status DevicesEroutersIdDhcpServersServerNameDelete(ctx, xAuthToken, id, serverName)
Delete DHCP server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **serverName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpServersServerNameGet**
> DhcpServerConfiguration DevicesEroutersIdDhcpServersServerNameGet(ctx, xAuthToken, id, serverName)
Get DHCP server configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **serverName** | **string**|  | 

### Return type

[**DhcpServerConfiguration**](DHCPServerConfiguration.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpServersServerNamePut**
> DhcpServerConfiguration DevicesEroutersIdDhcpServersServerNamePut(ctx, xAuthToken, id, serverName, body)
Update DHCP server configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **serverName** | **string**|  | 
  **body** | [**DhcpServerConfiguration**](DhcpServerConfiguration.md)|  | 

### Return type

[**DhcpServerConfiguration**](DHCPServerConfiguration.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdDhcpServersServerNameUnblockPost**
> Status DevicesEroutersIdDhcpServersServerNameUnblockPost(ctx, xAuthToken, id, serverName)
Unblock dhcp server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **id** | **string**|  | 
  **serverName** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdGet**
> ErouterStatusDetail DevicesEroutersIdGet(ctx, xAuthToken, id)
Return erouter detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**ErouterStatusDetail**](ErouterStatusDetail.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterOspfAreasAreaDelete**
> Status DevicesEroutersIdRouterOspfAreasAreaDelete(ctx, xAuthToken, id, area)
Delete OSPF area

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **area** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterOspfAreasAreaPut**
> OspfArea DevicesEroutersIdRouterOspfAreasAreaPut(ctx, xAuthToken, id, area, body)
Update OSPF area

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **area** | **string**|  | 
  **body** | [**OspfArea**](OspfArea.md)|  | 

### Return type

[**OspfArea**](OSPFArea.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterOspfAreasGet**
> []OspfArea DevicesEroutersIdRouterOspfAreasGet(ctx, xAuthToken, id)
All OSPF areas

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]OspfArea**](OSPFArea.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterOspfAreasPost**
> OspfArea DevicesEroutersIdRouterOspfAreasPost(ctx, xAuthToken, id, body)
Create new OSPF area

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**OspfArea**](OspfArea.md)|  | 

### Return type

[**OspfArea**](OSPFArea.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterOspfGet**
> OspfConfig DevicesEroutersIdRouterOspfGet(ctx, xAuthToken, id)
Get OSPF configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**OspfConfig**](OSPFConfig.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterOspfPut**
> OspfConfig DevicesEroutersIdRouterOspfPut(ctx, xAuthToken, id, body)
Update OSPF configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**OspfConfig**](OspfConfig.md)|  | 

### Return type

[**OspfConfig**](OSPFConfig.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterRoutesBlockPost**
> RouteOverview DevicesEroutersIdRouterRoutesBlockPost(ctx, xAuthToken, id, body)
Edit static route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**StaticRoute**](StaticRoute.md)|  | 

### Return type

[**RouteOverview**](RouteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterRoutesDeletePost**
> Status DevicesEroutersIdRouterRoutesDeletePost(ctx, xAuthToken, id, body)
Delete route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**StaticRoute**](StaticRoute.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterRoutesGet**
> []RouteOverview DevicesEroutersIdRouterRoutesGet(ctx, xAuthToken, id)
All routes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]RouteOverview**](RouteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterRoutesPost**
> RouteOverview DevicesEroutersIdRouterRoutesPost(ctx, xAuthToken, id, body)
Create new static route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**StaticRoute**](StaticRoute.md)|  | 

### Return type

[**RouteOverview**](RouteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterRoutesPut**
> RouteOverview DevicesEroutersIdRouterRoutesPut(ctx, xAuthToken, id, body)
Edit static route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**StaticRoute**](StaticRoute.md)|  | 

### Return type

[**RouteOverview**](RouteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesEroutersIdRouterRoutesUnblockPost**
> RouteOverview DevicesEroutersIdRouterRoutesUnblockPost(ctx, xAuthToken, id, body)
Edit static route

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**StaticRoute**](StaticRoute.md)|  | 

### Return type

[**RouteOverview**](RouteOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesGet**
> []DeviceStatusOverview DevicesGet(ctx, xAuthToken, optional)
List of all devices in network

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
 **siteId** | **string**|  | 
 **parentId** | **string**|  | 

### Return type

[**[]DeviceStatusOverview**](DeviceStatusOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesIdAuthorizePost**
> Status DevicesIdAuthorizePost(ctx, xAuthToken, id, payload)
Authorize device

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **payload** | [**Payload3**](Payload3.md)|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesIdBackupsGet**
> []DeviceBackup DevicesIdBackupsGet(ctx, xAuthToken, id)
Return list of device backups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]DeviceBackup**](DeviceBackup.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesIdBackupsPost**
> DeviceBackup DevicesIdBackupsPost(ctx, xAuthToken, id)
Create new device backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**DeviceBackup**](DeviceBackup.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesIdBackupsPut**
> []DeviceBackup DevicesIdBackupsPut(ctx, xAuthToken, id, file)
Upload data backup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **file** | ***os.File**| The uploaded backup file or files | 

### Return type

[**[]DeviceBackup**](DeviceBackup.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesIdDelete**
> Status DevicesIdDelete(ctx, xAuthToken, id)
Delete device

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

# **DevicesIdGet**
> DeviceStatusOverview DevicesIdGet(ctx, xAuthToken, id)
Device status overview

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**DeviceStatusOverview**](DeviceStatusOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesIdLocatePost**
> Status DevicesIdLocatePost(ctx, xAuthToken, id)
Start device locating

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

# **DevicesIdRefreshPost**
> Status DevicesIdRefreshPost(ctx, xAuthToken, id)
Refresh device communication, remove macAesKey

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

# **DevicesIdResetPost**
> Status DevicesIdResetPost(ctx, xAuthToken, id)
Reset device statistics, action clear-traffic-analysis

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

# **DevicesIdRestartPost**
> Status DevicesIdRestartPost(ctx, xAuthToken, id)
Restart device, device action reboot

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

# **DevicesIdStatisticsGet**
> DeviceStatistics DevicesIdStatisticsGet(ctx, xAuthToken, id, interval)
Return device statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **interval** | **string**|  | 

### Return type

[**DeviceStatistics**](DeviceStatistics.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesIdUnlocatePost**
> Status DevicesIdUnlocatePost(ctx, xAuthToken, id)
Stop device locating

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

# **DevicesMacsGet**
> []string DevicesMacsGet(ctx, xAuthToken)
Return a list of all devices' MAC address (except for ONUs)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

**[]string**

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOltsGet**
> []DeviceStatusOverview DevicesOltsGet(ctx, xAuthToken, optional)
List of all OLTs in network

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
 **siteId** | **string**|  | 

### Return type

[**[]DeviceStatusOverview**](DeviceStatusOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOltsIdDelete**
> Status DevicesOltsIdDelete(ctx, xAuthToken, id)
Delete OLT form UNMS

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

# **DevicesOltsIdGet**
> OltStatusDetail DevicesOltsIdGet(ctx, xAuthToken, id)
Return OLT detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**OltStatusDetail**](OLTStatusDetail.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOltsIdProfilesGet**
> []OltonuProfile DevicesOltsIdProfilesGet(ctx, xAuthToken, id)
List of all OLT ONU profiles

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]OltonuProfile**](OLTONUProfile.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOltsIdProfilesPost**
> OltonuProfile DevicesOltsIdProfilesPost(ctx, xAuthToken, id, body)
Create OLT ONU profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **body** | [**OltonuProfile**](OltonuProfile.md)|  | 

### Return type

[**OltonuProfile**](OLTONUProfile.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOltsIdProfilesProfileDelete**
> Status DevicesOltsIdProfilesProfileDelete(ctx, xAuthToken, id, profile)
Delete OLT form UNMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **profile** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOltsIdProfilesProfilePut**
> OltonuProfile DevicesOltsIdProfilesProfilePut(ctx, xAuthToken, id, profile, body)
Update OLT ONU profile

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 
  **profile** | **string**|  | 
  **body** | [**OltonuProfile**](OltonuProfile.md)|  | 

### Return type

[**OltonuProfile**](OLTONUProfile.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOnusGet**
> []OnuStatusOverview DevicesOnusGet(ctx, xAuthToken, optional)
List of all onus in network

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
 **siteId** | **string**|  | 
 **parentId** | **string**|  | 

### Return type

[**[]OnuStatusOverview**](ONUStatusOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOnusIdBlockPost**
> Status DevicesOnusIdBlockPost(ctx, xAuthToken, id)
Delete ONU GPON password and set status to blocked

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

# **DevicesOnusIdClientsGet**
> []OnuClient DevicesOnusIdClientsGet(ctx, xAuthToken, id)
Get all clients connected to ONU

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**[]OnuClient**](ONUClient.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOnusIdDelete**
> Status DevicesOnusIdDelete(ctx, xAuthToken, id)
Delete ONU device - call OLT action DELETE_ONU (former forget action)

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

# **DevicesOnusIdGet**
> OnuStatusOverview DevicesOnusIdGet(ctx, xAuthToken, id)
Return ONU detail

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 
  **id** | **string**|  | 

### Return type

[**OnuStatusOverview**](ONUStatusOverview.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOnusIdResetstatsPost**
> Status DevicesOnusIdResetstatsPost(ctx, xAuthToken, id)
Reset onu data statistics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**|  | 
  **id** | **string**|  | 

### Return type

[**Status**](Status.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DevicesOnusIdUnblockPost**
> Status DevicesOnusIdUnblockPost(ctx, xAuthToken, id)
Unblock ONU device - device action UNBLOCK_ONU

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

# **DevicesOnusIdUpgradePost**
> Status DevicesOnusIdUpgradePost(ctx, xAuthToken, id)
Upgrade device

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

# **DevicesSsidsGet**
> []InlineResponse2002 DevicesSsidsGet(ctx, xAuthToken)
Get devices wireless configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **xAuthToken** | **string**| User authorization token | 

### Return type

[**[]InlineResponse2002**](inline_response_200_2.md)

### Authorization

[UserSecurity](../README.md#UserSecurity)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

