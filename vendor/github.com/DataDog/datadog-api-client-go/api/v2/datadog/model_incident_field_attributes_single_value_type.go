/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// IncidentFieldAttributesSingleValueType Type of the single value field definitions.
type IncidentFieldAttributesSingleValueType string

// List of IncidentFieldAttributesSingleValueType
const (
	INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_DROPDOWN IncidentFieldAttributesSingleValueType = "dropdown"
	INCIDENTFIELDATTRIBUTESSINGLEVALUETYPE_TEXTBOX  IncidentFieldAttributesSingleValueType = "textbox"
)

func (v *IncidentFieldAttributesSingleValueType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IncidentFieldAttributesSingleValueType(value)
	for _, existing := range []IncidentFieldAttributesSingleValueType{"dropdown", "textbox"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IncidentFieldAttributesSingleValueType", value)
}

// Ptr returns reference to IncidentFieldAttributesSingleValueType value
func (v IncidentFieldAttributesSingleValueType) Ptr() *IncidentFieldAttributesSingleValueType {
	return &v
}

type NullableIncidentFieldAttributesSingleValueType struct {
	value *IncidentFieldAttributesSingleValueType
	isSet bool
}

func (v NullableIncidentFieldAttributesSingleValueType) Get() *IncidentFieldAttributesSingleValueType {
	return v.value
}

func (v *NullableIncidentFieldAttributesSingleValueType) Set(val *IncidentFieldAttributesSingleValueType) {
	v.value = val
	v.isSet = true
}

func (v NullableIncidentFieldAttributesSingleValueType) IsSet() bool {
	return v.isSet
}

func (v *NullableIncidentFieldAttributesSingleValueType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncidentFieldAttributesSingleValueType(val *IncidentFieldAttributesSingleValueType) *NullableIncidentFieldAttributesSingleValueType {
	return &NullableIncidentFieldAttributesSingleValueType{value: val, isSet: true}
}

func (v NullableIncidentFieldAttributesSingleValueType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncidentFieldAttributesSingleValueType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
