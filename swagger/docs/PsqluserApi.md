# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PsqluserCreate**](PsqluserApi.md#PsqluserCreate) | **Post** /api/v1/psqluser/create/ | 
[**PsqluserDelete**](PsqluserApi.md#PsqluserDelete) | **Post** /api/v1/psqluser/delete/ | 
[**PsqluserList**](PsqluserApi.md#PsqluserList) | **Get** /api/v1/psqluser/list/ | 
[**PsqluserRead**](PsqluserApi.md#PsqluserRead) | **Get** /api/v1/psqluser/read/{uuid} | 
[**PsqluserUpdate**](PsqluserApi.md#PsqluserUpdate) | **Post** /api/v1/psqluser/update/ | 

# **PsqluserCreate**
> []PsqlUserResponse PsqluserCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PsqlUserCreate**](PsqlUserCreate.md)|  | 

### Return type

[**[]PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqluserDelete**
> PsqluserDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PsqlUserRead**](PsqlUserRead.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqluserList**
> []PsqlUserResponse PsqluserList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqluserRead**
> PsqlUserResponse PsqluserRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqluserUpdate**
> []PsqlUserResponse PsqluserUpdate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PsqlUserUpdate**](PsqlUserUpdate.md)|  | 

### Return type

[**[]PsqlUserResponse**](PsqlUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

