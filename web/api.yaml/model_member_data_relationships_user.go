/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
)

// checks if the MemberDataRelationshipsUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberDataRelationshipsUser{}

// MemberDataRelationshipsUser struct for MemberDataRelationshipsUser
type MemberDataRelationshipsUser struct {
	Data *MemberDataRelationshipsUserData `json:"data,omitempty"`
}

// NewMemberDataRelationshipsUser instantiates a new MemberDataRelationshipsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberDataRelationshipsUser() *MemberDataRelationshipsUser {
	this := MemberDataRelationshipsUser{}
	return &this
}

// NewMemberDataRelationshipsUserWithDefaults instantiates a new MemberDataRelationshipsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberDataRelationshipsUserWithDefaults() *MemberDataRelationshipsUser {
	this := MemberDataRelationshipsUser{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MemberDataRelationshipsUser) GetData() MemberDataRelationshipsUserData {
	if o == nil || IsNil(o.Data) {
		var ret MemberDataRelationshipsUserData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberDataRelationshipsUser) GetDataOk() (*MemberDataRelationshipsUserData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MemberDataRelationshipsUser) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MemberDataRelationshipsUserData and assigns it to the Data field.
func (o *MemberDataRelationshipsUser) SetData(v MemberDataRelationshipsUserData) {
	o.Data = &v
}

func (o MemberDataRelationshipsUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberDataRelationshipsUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableMemberDataRelationshipsUser struct {
	value *MemberDataRelationshipsUser
	isSet bool
}

func (v NullableMemberDataRelationshipsUser) Get() *MemberDataRelationshipsUser {
	return v.value
}

func (v *NullableMemberDataRelationshipsUser) Set(val *MemberDataRelationshipsUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberDataRelationshipsUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberDataRelationshipsUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberDataRelationshipsUser(val *MemberDataRelationshipsUser) *NullableMemberDataRelationshipsUser {
	return &NullableMemberDataRelationshipsUser{value: val, isSet: true}
}

func (v NullableMemberDataRelationshipsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberDataRelationshipsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

