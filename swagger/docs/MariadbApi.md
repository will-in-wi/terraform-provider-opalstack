# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MariadbCreate**](MariadbApi.md#MariadbCreate) | **Post** /api/v1/mariadb/create/ | 
[**MariadbDelete**](MariadbApi.md#MariadbDelete) | **Post** /api/v1/mariadb/delete/ | 
[**MariadbList**](MariadbApi.md#MariadbList) | **Get** /api/v1/mariadb/list/ | 
[**MariadbRead**](MariadbApi.md#MariadbRead) | **Get** /api/v1/mariadb/read/{uuid} | 
[**MariadbUpdate**](MariadbApi.md#MariadbUpdate) | **Post** /api/v1/mariadb/update/ | 

# **MariadbCreate**
> []MariaDbResponse MariadbCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MariaDbCreate**](MariaDBCreate.md)|  | 

### Return type

[**[]MariaDbResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariadbDelete**
> MariadbDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MariaDbRead**](MariaDBRead.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariadbList**
> []MariaDbResponse MariadbList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]MariaDbResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariadbRead**
> MariaDbResponse MariadbRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**MariaDbResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MariadbUpdate**
> []MariaDbResponse MariadbUpdate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]MariaDbUpdate**](MariaDBUpdate.md)|  | 

### Return type

[**[]MariaDbResponse**](MariaDBResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

