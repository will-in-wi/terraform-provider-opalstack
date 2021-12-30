# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariauserCreate**](MariauserApi.md#MariauserCreate) | **Post** /api/v1/mariauser/create/ | 
[**MariauserDelete**](MariauserApi.md#MariauserDelete) | **Post** /api/v1/mariauser/delete/ | 
[**MariauserList**](MariauserApi.md#MariauserList) | **Get** /api/v1/mariauser/list/ | 
[**MariauserRead**](MariauserApi.md#MariauserRead) | **Get** /api/v1/mariauser/read/{uuid} | 
[**MariauserUpdate**](MariauserApi.md#MariauserUpdate) | **Post** /api/v1/mariauser/update/ | 

# **MariauserCreate**
> []MariaUserResponse MariauserCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MariaUserCreate**](MariaUserCreate.md)|  | 

### Return type

[**[]MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariauserDelete**
> MariauserDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MariaUserRead**](MariaUserRead.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariauserList**
> []MariaUserResponse MariauserList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariauserRead**
> MariaUserResponse MariauserRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariauserUpdate**
> []MariaUserResponse MariauserUpdate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MariaUserUpdate**](MariaUserUpdate.md)|  | 

### Return type

[**[]MariaUserResponse**](MariaUserResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

