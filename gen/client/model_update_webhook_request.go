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

// checks if the UpdateWebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateWebhookRequest{}

// UpdateWebhookRequest struct for UpdateWebhookRequest
type UpdateWebhookRequest struct {
	// The ID of the blockchain network
	NetworkId *string `json:"network_id,omitempty"`
	EventType WebhookEventType `json:"event_type"`
	// Webhook will monitor all events that matches any one of the event filters.
	EventFilters []WebhookEventFilter `json:"event_filters"`
	// The Webhook uri that updates to
	NotificationUri string `json:"notification_uri"`
}

type _UpdateWebhookRequest UpdateWebhookRequest

// NewUpdateWebhookRequest instantiates a new UpdateWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWebhookRequest(eventType WebhookEventType, eventFilters []WebhookEventFilter, notificationUri string) *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	this.EventType = eventType
	this.EventFilters = eventFilters
	this.NotificationUri = notificationUri
	return &this
}

// NewUpdateWebhookRequestWithDefaults instantiates a new UpdateWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWebhookRequestWithDefaults() *UpdateWebhookRequest {
	this := UpdateWebhookRequest{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *UpdateWebhookRequest) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *UpdateWebhookRequest) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *UpdateWebhookRequest) SetNetworkId(v string) {
	o.NetworkId = &v
}

// GetEventType returns the EventType field value
func (o *UpdateWebhookRequest) GetEventType() WebhookEventType {
	if o == nil {
		var ret WebhookEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetEventTypeOk() (*WebhookEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *UpdateWebhookRequest) SetEventType(v WebhookEventType) {
	o.EventType = v
}

// GetEventFilters returns the EventFilters field value
func (o *UpdateWebhookRequest) GetEventFilters() []WebhookEventFilter {
	if o == nil {
		var ret []WebhookEventFilter
		return ret
	}

	return o.EventFilters
}

// GetEventFiltersOk returns a tuple with the EventFilters field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetEventFiltersOk() ([]WebhookEventFilter, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventFilters, true
}

// SetEventFilters sets field value
func (o *UpdateWebhookRequest) SetEventFilters(v []WebhookEventFilter) {
	o.EventFilters = v
}

// GetNotificationUri returns the NotificationUri field value
func (o *UpdateWebhookRequest) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *UpdateWebhookRequest) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *UpdateWebhookRequest) SetNotificationUri(v string) {
	o.NotificationUri = v
}

func (o UpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateWebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["network_id"] = o.NetworkId
	}
	toSerialize["event_type"] = o.EventType
	toSerialize["event_filters"] = o.EventFilters
	toSerialize["notification_uri"] = o.NotificationUri
	return toSerialize, nil
}

func (o *UpdateWebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"event_type",
		"event_filters",
		"notification_uri",
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

	varUpdateWebhookRequest := _UpdateWebhookRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateWebhookRequest)

	if err != nil {
		return err
	}

	*o = UpdateWebhookRequest(varUpdateWebhookRequest)

	return err
}

type NullableUpdateWebhookRequest struct {
	value *UpdateWebhookRequest
	isSet bool
}

func (v NullableUpdateWebhookRequest) Get() *UpdateWebhookRequest {
	return v.value
}

func (v *NullableUpdateWebhookRequest) Set(val *UpdateWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWebhookRequest(val *UpdateWebhookRequest) *NullableUpdateWebhookRequest {
	return &NullableUpdateWebhookRequest{value: val, isSet: true}
}

func (v NullableUpdateWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


