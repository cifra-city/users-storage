/*
Cifra SSO REST API

SSO REST API for Cifra app

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Member type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Member{}

// Member struct for Member
type Member struct {
	Data MemberData `json:"data"`
}

type _Member Member

// NewMember instantiates a new Member object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMember(data MemberData) *Member {
	this := Member{}
	this.Data = data
	return &this
}

// NewMemberWithDefaults instantiates a new Member object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberWithDefaults() *Member {
	this := Member{}
	return &this
}

// GetData returns the Data field value
func (o *Member) GetData() MemberData {
	if o == nil {
		var ret MemberData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Member) GetDataOk() (*MemberData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Member) SetData(v MemberData) {
	o.Data = v
}

func (o Member) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Member) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *Member) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMember := _Member{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMember)

	if err != nil {
		return err
	}

	*o = Member(varMember)

	return err
}

type NullableMember struct {
	value *Member
	isSet bool
}

func (v NullableMember) Get() *Member {
	return v.value
}

func (v *NullableMember) Set(val *Member) {
	v.value = val
	v.isSet = true
}

func (v NullableMember) IsSet() bool {
	return v.isSet
}

func (v *NullableMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMember(val *Member) *NullableMember {
	return &NullableMember{value: val, isSet: true}
}

func (v NullableMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

