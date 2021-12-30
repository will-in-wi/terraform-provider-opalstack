# SiteUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [default to null]
**Name** | **string** |  | [optional] [default to null]
**Ip4** | **string** |  | [optional] [default to null]
**Ip6** | **string** |  | [optional] [default to null]
**Domains** | **[]string** |  | [optional] [default to null]
**Routes** | [**[]Route**](Route.md) |  | [optional] [default to null]
**Cert** | **string** |  | [optional] [default to null]
**Redirect** | **bool** | Automatically redirect to https:// | [optional] [default to null]
**GenerateLe** | **bool** | Automatically issue Lets Encrypt certificate for the domains on this site? | [optional] [default to null]
**LeHttpChallengeTokens** | **[]string** |  | [optional] [default to null]
**Disabled** | **bool** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

