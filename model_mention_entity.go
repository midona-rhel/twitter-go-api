/*
Tweets and Users

API Reference — Labs v2

API version: 2.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package twitterapi

import (
	"encoding/json"
)

// MentionEntity struct for MentionEntity
type MentionEntity struct {
	// Index (zero-based) at which position this entity starts.
	Start int32 `json:"start"`
	// Index (zero-based) at which position this entity ends.
	End int32 `json:"end"`
	// The Twitter handle (screen name) of this user.
	Username string `json:"username"`
	// Unique identifier of this User. This is returned as a string in order to avoid complications with languages and tools that cannot handle large integers.
	Id string `json:"id"`
}

// NewMentionEntity instantiates a new MentionEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMentionEntity(start int32, end int32, username string, id string) *MentionEntity {
	this := MentionEntity{}
	this.Start = start
	this.End = end
	this.Username = username
	this.Id = id
	return &this
}

// NewMentionEntityWithDefaults instantiates a new MentionEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMentionEntityWithDefaults() *MentionEntity {
	this := MentionEntity{}
	return &this
}

// GetStart returns the Start field value
func (o *MentionEntity) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *MentionEntity) GetStartOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *MentionEntity) SetStart(v int32) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *MentionEntity) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *MentionEntity) GetEndOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *MentionEntity) SetEnd(v int32) {
	o.End = v
}

// GetUsername returns the Username field value
func (o *MentionEntity) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *MentionEntity) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *MentionEntity) SetUsername(v string) {
	o.Username = v
}

// GetId returns the Id field value
func (o *MentionEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MentionEntity) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MentionEntity) SetId(v string) {
	o.Id = v
}

func (o MentionEntity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if true {
		toSerialize["username"] = o.Username
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableMentionEntity struct {
	value *MentionEntity
	isSet bool
}

func (v NullableMentionEntity) Get() *MentionEntity {
	return v.value
}

func (v *NullableMentionEntity) Set(val *MentionEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableMentionEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableMentionEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMentionEntity(val *MentionEntity) *NullableMentionEntity {
	return &NullableMentionEntity{value: val, isSet: true}
}

func (v NullableMentionEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMentionEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

