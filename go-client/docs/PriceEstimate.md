# PriceEstimate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Unique identifier representing a specific product for a given latitude &amp; longitude. For example, uberX in San Francisco will have a different product_id than uberX in Los Angeles | [optional] [default to null]
**CurrencyCode** | **string** | [ISO 4217](http://en.wikipedia.org/wiki/ISO_4217) currency code. | [optional] [default to null]
**DisplayName** | **string** | Display name of product. | [optional] [default to null]
**Estimate** | **string** | Formatted string of estimate in local currency of the start location. Estimate could be a range, a single number (flat rate) or \&quot;Metered\&quot; for TAXI. | [optional] [default to null]
**LowEstimate** | **float32** | Lower bound of the estimated price. | [optional] [default to null]
**HighEstimate** | **float32** | Upper bound of the estimated price. | [optional] [default to null]
**SurgeMultiplier** | **float32** | Expected surge multiplier. Surge is active if surge_multiplier is greater than 1. Price estimate already factors in the surge multiplier. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


