# UserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User ID | 
**Type** | **string** |  | 
**Attributes** | [**UserDataAttributes**](UserDataAttributes.md) |  | 

## Methods

### NewUserData

`func NewUserData(id string, type_ string, attributes UserDataAttributes, ) *UserData`

NewUserData instantiates a new UserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataWithDefaults

`func NewUserDataWithDefaults() *UserData`

NewUserDataWithDefaults instantiates a new UserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UserData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *UserData) GetAttributes() UserDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UserData) GetAttributesOk() (*UserDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UserData) SetAttributes(v UserDataAttributes)`

SetAttributes sets Attributes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


