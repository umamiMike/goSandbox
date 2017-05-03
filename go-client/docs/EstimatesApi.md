# \EstimatesApi

All URIs are relative to *https://api.uber.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimatesPriceGet**](EstimatesApi.md#EstimatesPriceGet) | **Get** /estimates/price | Price Estimates
[**EstimatesTimeGet**](EstimatesApi.md#EstimatesTimeGet) | **Get** /estimates/time | Time Estimates


# **EstimatesPriceGet**
> []PriceEstimate EstimatesPriceGet($startLatitude, $startLongitude, $endLatitude, $endLongitude)

Price Estimates

The Price Estimates endpoint returns an estimated price range for each product offered at a given location. The price estimate is provided as a formatted string with the full price range and the localized currency symbol.<br><br>The response also includes low and high estimates, and the [ISO 4217](http://en.wikipedia.org/wiki/ISO_4217) currency code for situations requiring currency conversion. When surge is active for a particular product, its surge_multiplier will be greater than 1, but the price estimate already factors in this multiplier. 


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startLatitude** | **float64**| Latitude component of start location. | 
 **startLongitude** | **float64**| Longitude component of start location. | 
 **endLatitude** | **float64**| Latitude component of end location. | 
 **endLongitude** | **float64**| Longitude component of end location. | 

### Return type

[**[]PriceEstimate**](PriceEstimate.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EstimatesTimeGet**
> []Product EstimatesTimeGet($startLatitude, $startLongitude, $customerUuid, $productId)

Time Estimates

The Time Estimates endpoint returns ETAs for all products offered at a given location, with the responses expressed as integers in seconds. We recommend that this endpoint be called every minute to provide the most accurate, up-to-date ETAs.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startLatitude** | **float64**| Latitude component of start location. | 
 **startLongitude** | **float64**| Longitude component of start location. | 
 **customerUuid** | **string**| Unique customer identifier to be used for experience customization. | [optional] 
 **productId** | **string**| Unique identifier representing a specific product for a given latitude &amp; longitude. | [optional] 

### Return type

[**[]Product**](Product.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

