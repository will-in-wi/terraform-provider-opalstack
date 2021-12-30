# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OsuserCreate**](OsuserApi.md#OsuserCreate) | **Post** /api/v1/osuser/create/ | 
[**OsuserDelete**](OsuserApi.md#OsuserDelete) | **Post** /api/v1/osuser/delete/ | 
[**OsuserList**](OsuserApi.md#OsuserList) | **Get** /api/v1/osuser/list/ | 
[**OsuserRead**](OsuserApi.md#OsuserRead) | **Get** /api/v1/osuser/read/{uuid} | 
[**OsuserUpdate**](OsuserApi.md#OsuserUpdate) | **Post** /api/v1/osuser/update/ | 

# **OsuserCreate**
> []OsUserResponse OsuserCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]OsUserCreate**](OSUserCreate.md)|  | 

### Return type

[**[]OsUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OsuserDelete**
> OsuserDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]OsUserRead**](OSUserRead.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OsuserList**
> []OsUserResponse OsuserList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]OsUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OsuserRead**
> OsUserResponse OsuserRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**OsUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OsuserUpdate**
> []OsUserResponse OsuserUpdate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]OsUserUpdate**](OSUserUpdate.md)|  | 

### Return type

[**[]OsUserResponse**](OSUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

