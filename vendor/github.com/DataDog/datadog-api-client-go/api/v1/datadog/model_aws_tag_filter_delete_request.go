/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// AWSTagFilterDeleteRequest The objects used to delete an AWS tag filter entry.
type AWSTagFilterDeleteRequest struct {
	// The unique identifier of your AWS account.
	AwsAccountIdentifier *string       `json:"aws_account_identifier,omitempty"`
	Namespace            *AWSNamespace `json:"namespace,omitempty"`
}

// NewAWSTagFilterDeleteRequest instantiates a new AWSTagFilterDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSTagFilterDeleteRequest() *AWSTagFilterDeleteRequest {
	this := AWSTagFilterDeleteRequest{}
	return &this
}

// NewAWSTagFilterDeleteRequestWithDefaults instantiates a new AWSTagFilterDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSTagFilterDeleteRequestWithDefaults() *AWSTagFilterDeleteRequest {
	this := AWSTagFilterDeleteRequest{}
	return &this
}

// GetAwsAccountIdentifier returns the AwsAccountIdentifier field value if set, zero value otherwise.
func (o *AWSTagFilterDeleteRequest) GetAwsAccountIdentifier() string {
	if o == nil || o.AwsAccountIdentifier == nil {
		var ret string
		return ret
	}
	return *o.AwsAccountIdentifier
}

// GetAwsAccountIdentifierOk returns a tuple with the AwsAccountIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterDeleteRequest) GetAwsAccountIdentifierOk() (*string, bool) {
	if o == nil || o.AwsAccountIdentifier == nil {
		return nil, false
	}
	return o.AwsAccountIdentifier, true
}

// HasAwsAccountIdentifier returns a boolean if a field has been set.
func (o *AWSTagFilterDeleteRequest) HasAwsAccountIdentifier() bool {
	if o != nil && o.AwsAccountIdentifier != nil {
		return true
	}

	return false
}

// SetAwsAccountIdentifier gets a reference to the given string and assigns it to the AwsAccountIdentifier field.
func (o *AWSTagFilterDeleteRequest) SetAwsAccountIdentifier(v string) {
	o.AwsAccountIdentifier = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AWSTagFilterDeleteRequest) GetNamespace() AWSNamespace {
	if o == nil || o.Namespace == nil {
		var ret AWSNamespace
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSTagFilterDeleteRequest) GetNamespaceOk() (*AWSNamespace, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AWSTagFilterDeleteRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given AWSNamespace and assigns it to the Namespace field.
func (o *AWSTagFilterDeleteRequest) SetNamespace(v AWSNamespace) {
	o.Namespace = &v
}

func (o AWSTagFilterDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AwsAccountIdentifier != nil {
		toSerialize["aws_account_identifier"] = o.AwsAccountIdentifier
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableAWSTagFilterDeleteRequest struct {
	value *AWSTagFilterDeleteRequest
	isSet bool
}

func (v NullableAWSTagFilterDeleteRequest) Get() *AWSTagFilterDeleteRequest {
	return v.value
}

func (v *NullableAWSTagFilterDeleteRequest) Set(val *AWSTagFilterDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSTagFilterDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSTagFilterDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSTagFilterDeleteRequest(val *AWSTagFilterDeleteRequest) *NullableAWSTagFilterDeleteRequest {
	return &NullableAWSTagFilterDeleteRequest{value: val, isSet: true}
}

func (v NullableAWSTagFilterDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSTagFilterDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
