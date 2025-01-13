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

// checks if the MemberUpdateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MemberUpdateDataAttributes{}

// MemberUpdateDataAttributes struct for MemberUpdateDataAttributes
type MemberUpdateDataAttributes struct {
	// Team id
	TeamId string `json:"team_id"`
	// User id
	UserId string `json:"user_id"`
	// User role
	Role *string `json:"role,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
}

type _MemberUpdateDataAttributes MemberUpdateDataAttributes

// NewMemberUpdateDataAttributes instantiates a new MemberUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMemberUpdateDataAttributes(teamId string, userId string) *MemberUpdateDataAttributes {
	this := MemberUpdateDataAttributes{}
	this.TeamId = teamId
	this.UserId = userId
	return &this
}

// NewMemberUpdateDataAttributesWithDefaults instantiates a new MemberUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMemberUpdateDataAttributesWithDefaults() *MemberUpdateDataAttributes {
	this := MemberUpdateDataAttributes{}
	return &this
}

// GetTeamId returns the TeamId field value
func (o *MemberUpdateDataAttributes) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *MemberUpdateDataAttributes) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *MemberUpdateDataAttributes) SetTeamId(v string) {
	o.TeamId = v
}

// GetUserId returns the UserId field value
func (o *MemberUpdateDataAttributes) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *MemberUpdateDataAttributes) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *MemberUpdateDataAttributes) SetUserId(v string) {
	o.UserId = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *MemberUpdateDataAttributes) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberUpdateDataAttributes) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *MemberUpdateDataAttributes) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *MemberUpdateDataAttributes) SetRole(v string) {
	o.Role = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MemberUpdateDataAttributes) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MemberUpdateDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MemberUpdateDataAttributes) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MemberUpdateDataAttributes) SetDescription(v string) {
	o.Description = &v
}

func (o MemberUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MemberUpdateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["team_id"] = o.TeamId
	toSerialize["user_id"] = o.UserId
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *MemberUpdateDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"team_id",
		"user_id",
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

	varMemberUpdateDataAttributes := _MemberUpdateDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMemberUpdateDataAttributes)

	if err != nil {
		return err
	}

	*o = MemberUpdateDataAttributes(varMemberUpdateDataAttributes)

	return err
}

type NullableMemberUpdateDataAttributes struct {
	value *MemberUpdateDataAttributes
	isSet bool
}

func (v NullableMemberUpdateDataAttributes) Get() *MemberUpdateDataAttributes {
	return v.value
}

func (v *NullableMemberUpdateDataAttributes) Set(val *MemberUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableMemberUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableMemberUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMemberUpdateDataAttributes(val *MemberUpdateDataAttributes) *NullableMemberUpdateDataAttributes {
	return &NullableMemberUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableMemberUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMemberUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


