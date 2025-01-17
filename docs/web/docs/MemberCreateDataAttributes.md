# MemberCreateDataAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User ID | 
**TeamId** | **string** | Team ID | 
**Role** | **string** | User role | 
**Description** | Pointer to **string** | Description | [optional] 

## Methods

### NewMemberCreateDataAttributes

`func NewMemberCreateDataAttributes(userId string, teamId string, role string, ) *MemberCreateDataAttributes`

NewMemberCreateDataAttributes instantiates a new MemberCreateDataAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberCreateDataAttributesWithDefaults

`func NewMemberCreateDataAttributesWithDefaults() *MemberCreateDataAttributes`

NewMemberCreateDataAttributesWithDefaults instantiates a new MemberCreateDataAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *MemberCreateDataAttributes) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MemberCreateDataAttributes) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MemberCreateDataAttributes) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTeamId

`func (o *MemberCreateDataAttributes) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MemberCreateDataAttributes) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MemberCreateDataAttributes) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetRole

`func (o *MemberCreateDataAttributes) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MemberCreateDataAttributes) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MemberCreateDataAttributes) SetRole(v string)`

SetRole sets Role field to given value.


### GetDescription

`func (o *MemberCreateDataAttributes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MemberCreateDataAttributes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MemberCreateDataAttributes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MemberCreateDataAttributes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


