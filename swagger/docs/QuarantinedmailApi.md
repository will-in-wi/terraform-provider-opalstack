# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QuarantinedmailDelete**](QuarantinedmailApi.md#QuarantinedmailDelete) | **Post** /api/v1/quarantinedmail/delete/ | 
[**QuarantinedmailList**](QuarantinedmailApi.md#QuarantinedmailList) | **Get** /api/v1/quarantinedmail/list/ | 
[**QuarantinedmailRead**](QuarantinedmailApi.md#QuarantinedmailRead) | **Get** /api/v1/quarantinedmail/read/{uuid} | 
[**QuarantinedmailSubmit**](QuarantinedmailApi.md#QuarantinedmailSubmit) | **Post** /api/v1/quarantinedmail/submit/ | 

# **QuarantinedmailDelete**
> QuarantinedmailDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]QuarantinedMailRead**](QuarantinedMailRead.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuarantinedmailList**
> []QuarantinedMailResponse QuarantinedmailList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]QuarantinedMailResponse**](QuarantinedMailResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuarantinedmailRead**
> QuarantinedMailResponse QuarantinedmailRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**QuarantinedMailResponse**](QuarantinedMailResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **QuarantinedmailSubmit**
> []QuarantinedMailResponse QuarantinedmailSubmit(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]QuarantinedMailUpdate**](QuarantinedMailUpdate.md)|  | 

### Return type

[**[]QuarantinedMailResponse**](QuarantinedMailResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

