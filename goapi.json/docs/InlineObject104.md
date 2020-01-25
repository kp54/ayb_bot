# InlineObject104

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithFiles** | **bool** | ファイルが添付された投稿に限定するか否か | [optional] 
**MediaOnly** | **bool** | ファイルが添付された投稿に限定するか否か (このパラメータは廃止予定です。代わりに withFiles を使ってください。) | [optional] 
**FileType** | **[]string** | 指定された種類のファイルが添付された投稿のみを取得します | [optional] 
**ExcludeNsfw** | **bool** | true にすると、NSFW指定されたファイルを除外します(fileTypeが指定されている場合のみ有効) | [optional] 
**Limit** | **float32** |  | [optional] [default to 10]
**SinceId** | **string** |  | [optional] 
**UntilId** | **string** |  | [optional] 
**SinceDate** | **float32** |  | [optional] 
**UntilDate** | **float32** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

