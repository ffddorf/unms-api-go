# RouteOverview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | K - kernel route, C - connected, S - static, R - RIP, O - OSPF, I - ISIS, B - BGP | [optional] [default to null]
**Enabled** | **bool** | nullable; bool only for static routes | [optional] [default to null]
**Destination** | **string** | cidr | [optional] [default to null]
**Gateway** | **string** | ip address | [optional] [default to null]
**GatewayStatus** | **string** |  | [optional] [default to null]
**Interface_** | **string** | interface name, for example eth0, eth1 etc. | [optional] [default to null]
**Distance** | **float32** |  | [optional] [default to null]
**Fib** | **bool** | true if route type includes *, for example S&gt;* | [optional] [default to null]
**Selected** | **bool** | true if route type includes &gt;, for example S&gt;* | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


