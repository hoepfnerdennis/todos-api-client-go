# Todo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewTodo

`func NewTodo(title string, ) *Todo`

NewTodo instantiates a new Todo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTodoWithDefaults

`func NewTodoWithDefaults() *Todo`

NewTodoWithDefaults instantiates a new Todo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Todo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Todo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Todo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Todo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *Todo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Todo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Todo) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Todo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Todo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Todo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Todo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


