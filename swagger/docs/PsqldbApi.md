# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PsqldbCreate**](PsqldbApi.md#PsqldbCreate) | **Post** /api/v1/psqldb/create/ | 
[**PsqldbDelete**](PsqldbApi.md#PsqldbDelete) | **Post** /api/v1/psqldb/delete/ | 
[**PsqldbList**](PsqldbApi.md#PsqldbList) | **Get** /api/v1/psqldb/list/ | 
[**PsqldbRead**](PsqldbApi.md#PsqldbRead) | **Get** /api/v1/psqldb/read/{uuid} | 
[**PsqldbUpdate**](PsqldbApi.md#PsqldbUpdate) | **Post** /api/v1/psqldb/update/ | 

# **PsqldbCreate**
> []PsqlDbResponse PsqldbCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PsqlDbCreate**](PsqlDBCreate.md)|  | 

### Return type

[**[]PsqlDbResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqldbDelete**
> PsqldbDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PsqlDbRead**](PsqlDBRead.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqldbList**
> []PsqlDbResponse PsqldbList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PsqlDbResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqldbRead**
> PsqlDbResponse PsqldbRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**PsqlDbResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PsqldbUpdate**
> []PsqlDbResponse PsqldbUpdate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]PsqlDbUpdate**](PsqlDBUpdate.md)|  | 

### Return type

[**[]PsqlDbResponse**](PsqlDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

