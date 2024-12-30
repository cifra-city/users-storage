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

// checks if the UserDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataAttributes{}

// UserDataAttributes struct for UserDataAttributes
type UserDataAttributes struct {
	// User ID
	Id string `json:"id"`
	// Username
	Username string `json:"username"`
	// User title
	Title string `json:"title"`
	// User status
	Status string `json:"status"`
	// User avatar
	Avatar string `json:"avatar"`
	// User bio
	Bio string `json:"bio"`
}

type _UserDataAttributes UserDataAttributes

// NewUserDataAttributes instantiates a new UserDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataAttributes(id string, username string, title string, status string, avatar string, bio string) *UserDataAttributes {
	this := UserDataAttributes{}
	this.Id = id
	this.Username = username
	this.Title = title
	this.Status = status
	this.Avatar = avatar
	this.Bio = bio
	return &this
}

// NewUserDataAttributesWithDefaults instantiates a new UserDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataAttributesWithDefaults() *UserDataAttributes {
	this := UserDataAttributes{}
	return &this
}

// GetId returns the Id field value
func (o *UserDataAttributes) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserDataAttributes) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserDataAttributes) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *UserDataAttributes) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserDataAttributes) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserDataAttributes) SetUsername(v string) {
	o.Username = v
}

// GetTitle returns the Title field value
func (o *UserDataAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *UserDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *UserDataAttributes) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *UserDataAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserDataAttributes) SetStatus(v string) {
	o.Status = v
}

// GetAvatar returns the Avatar field value
func (o *UserDataAttributes) GetAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Avatar
}

// GetAvatarOk returns a tuple with the Avatar field value
// and a boolean to check if the value has been set.
func (o *UserDataAttributes) GetAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Avatar, true
}

// SetAvatar sets field value
func (o *UserDataAttributes) SetAvatar(v string) {
	o.Avatar = v
}

// GetBio returns the Bio field value
func (o *UserDataAttributes) GetBio() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Bio
}

// GetBioOk returns a tuple with the Bio field value
// and a boolean to check if the value has been set.
func (o *UserDataAttributes) GetBioOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bio, true
}

// SetBio sets field value
func (o *UserDataAttributes) SetBio(v string) {
	o.Bio = v
}

func (o UserDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	toSerialize["title"] = o.Title
	toSerialize["status"] = o.Status
	toSerialize["avatar"] = o.Avatar
	toSerialize["bio"] = o.Bio
	return toSerialize, nil
}

func (o *UserDataAttributes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"title",
		"status",
		"avatar",
		"bio",
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

	varUserDataAttributes := _UserDataAttributes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserDataAttributes)

	if err != nil {
		return err
	}

	*o = UserDataAttributes(varUserDataAttributes)

	return err
}

type NullableUserDataAttributes struct {
	value *UserDataAttributes
	isSet bool
}

func (v NullableUserDataAttributes) Get() *UserDataAttributes {
	return v.value
}

func (v *NullableUserDataAttributes) Set(val *UserDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataAttributes(val *UserDataAttributes) *NullableUserDataAttributes {
	return &NullableUserDataAttributes{value: val, isSet: true}
}

func (v NullableUserDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

