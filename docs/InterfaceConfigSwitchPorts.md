# InterfaceConfigSwitchPorts

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface_** | [***InterfaceIdentification**](InterfaceIdentification.md) |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Pvid** | **float32** | Value range is 1 - 4087. Pvid (port vlan id), when a packet comes to eth0 from outside network with no \&quot;vlan tag\&quot;, it will be tagged with pvid and treated as a packet with vlan tag. | [optional] [default to null]
**Vid** | **[]float32** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


