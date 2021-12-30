# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsrecordCreate**](DnsrecordApi.md#DnsrecordCreate) | **Post** /api/v1/dnsrecord/create/ | 
[**DnsrecordDelete**](DnsrecordApi.md#DnsrecordDelete) | **Post** /api/v1/dnsrecord/delete/ | 
[**DnsrecordList**](DnsrecordApi.md#DnsrecordList) | **Get** /api/v1/dnsrecord/list/ | 
[**DnsrecordRead**](DnsrecordApi.md#DnsrecordRead) | **Get** /api/v1/dnsrecord/read/{uuid} | 
[**DnsrecordUpdate**](DnsrecordApi.md#DnsrecordUpdate) | **Post** /api/v1/dnsrecord/update/ | 

# **DnsrecordCreate**
> []DnsRecordResponse DnsrecordCreate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DnsRecordCreate**](DNSRecordCreate.md)|  | 

### Return type

[**[]DnsRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsrecordDelete**
> DnsrecordDelete(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DnsRecordRead**](DNSRecordRead.md)|  | 

### Return type

 (empty response body)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsrecordList**
> []DnsRecordResponse DnsrecordList(ctx, )


### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DnsRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsrecordRead**
> DnsRecordResponse DnsrecordRead(ctx, uuid)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uuid** | [**string**](.md)|  | 

### Return type

[**DnsRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DnsrecordUpdate**
> []DnsRecordResponse DnsrecordUpdate(ctx, body)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DnsRecordUpdate**](DNSRecordUpdate.md)|  | 

### Return type

[**[]DnsRecordResponse**](DNSRecordResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

