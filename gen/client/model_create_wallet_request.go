/*
Coinbase Platform API

This is the OpenAPI 3.0 specification for the Coinbase Platform APIs, used in conjunction with the Coinbase Platform SDKs.

API version: 0.0.1-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateWalletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateWalletRequest{}

// CreateWalletRequest struct for CreateWalletRequest
type CreateWalletRequest struct {
	Wallet CreateWalletRequestWallet `json:"wallet"`
}

type _CreateWalletRequest CreateWalletRequest

// NewCreateWalletRequest instantiates a new CreateWalletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWalletRequest(wallet CreateWalletRequestWallet) *CreateWalletRequest {
	this := CreateWalletRequest{}
	this.Wallet = wallet
	return &this
}

// NewCreateWalletRequestWithDefaults instantiates a new CreateWalletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWalletRequestWithDefaults() *CreateWalletRequest {
	this := CreateWalletRequest{}
	return &this
}

// GetWallet returns the Wallet field value
func (o *CreateWalletRequest) GetWallet() CreateWalletRequestWallet {
	if o == nil {
		var ret CreateWalletRequestWallet
		return ret
	}

	return o.Wallet
}

// GetWalletOk returns a tuple with the Wallet field value
// and a boolean to check if the value has been set.
func (o *CreateWalletRequest) GetWalletOk() (*CreateWalletRequestWallet, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wallet, true
}

// SetWallet sets field value
func (o *CreateWalletRequest) SetWallet(v CreateWalletRequestWallet) {
	o.Wallet = v
}

func (o CreateWalletRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWalletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["wallet"] = o.Wallet
	return toSerialize, nil
}

func (o *CreateWalletRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"wallet",
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

	varCreateWalletRequest := _CreateWalletRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateWalletRequest)

	if err != nil {
		return err
	}

	*o = CreateWalletRequest(varCreateWalletRequest)

	return err
}

type NullableCreateWalletRequest struct {
	value *CreateWalletRequest
	isSet bool
}

func (v NullableCreateWalletRequest) Get() *CreateWalletRequest {
	return v.value
}

func (v *NullableCreateWalletRequest) Set(val *CreateWalletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWalletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWalletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWalletRequest(val *CreateWalletRequest) *NullableCreateWalletRequest {
	return &NullableCreateWalletRequest{value: val, isSet: true}
}

func (v NullableCreateWalletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWalletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


