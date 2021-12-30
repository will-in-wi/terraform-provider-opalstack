# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MailuserCreate**](MailuserApi.md#MailuserCreate) | **Post** /api/v1/mailuser/create/ | 
[**MailuserDelete**](MailuserApi.md#MailuserDelete) | **Post** /api/v1/mailuser/delete/ | 
[**MailuserList**](MailuserApi.md#MailuserList) | **Get** /api/v1/mailuser/list/ | 
[**MailuserRead**](MailuserApi.md#MailuserRead) | **Get** /api/v1/mailuser/read/{uuid} | 
[**MailuserUpdate**](MailuserApi.md#MailuserUpdate) | **Post** /api/v1/mailuser/update/ | 
[**MailuserUpdatePublic**](MailuserApi.md#MailuserUpdatePublic) | **Post** /api/v1/mailuser/update_public/ | 

# **MailuserCreate**
> []MailUserResponse MailuserCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MailUserCreate**](MailUserCreate.md)|  | 

### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailuserDelete**
> MailuserDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MailUserDelete**](MailUserDelete.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailuserList**
> []MailUserResponse MailuserList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailuserRead**
> MailUserResponse MailuserRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailuserUpdate**
> []MailUserResponse MailuserUpdate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MailUserUpdate**](MailUserUpdate.md)|  | 

### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MailuserUpdatePublic**
> []MailUserResponse MailuserUpdatePublic(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MailUserUpdatePublic**](MailUserUpdatePublic.md)|  | 

### Return type

[**[]MailUserResponse**](MailUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

