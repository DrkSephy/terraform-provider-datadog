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

// EventTimelineWidgetDefinitionType Type of the event timeline widget.
type EventTimelineWidgetDefinitionType string

// List of EventTimelineWidgetDefinitionType
const (
	EVENTTIMELINEWIDGETDEFINITIONTYPE_EVENT_TIMELINE EventTimelineWidgetDefinitionType = "event_timeline"
)

func (v *EventTimelineWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTimelineWidgetDefinitionType(value)
	for _, existing := range []EventTimelineWidgetDefinitionType{"event_timeline"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTimelineWidgetDefinitionType", value)
}

// Ptr returns reference to EventTimelineWidgetDefinitionType value
func (v EventTimelineWidgetDefinitionType) Ptr() *EventTimelineWidgetDefinitionType {
	return &v
}

type NullableEventTimelineWidgetDefinitionType struct {
	value *EventTimelineWidgetDefinitionType
	isSet bool
}

func (v NullableEventTimelineWidgetDefinitionType) Get() *EventTimelineWidgetDefinitionType {
	return v.value
}

func (v *NullableEventTimelineWidgetDefinitionType) Set(val *EventTimelineWidgetDefinitionType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTimelineWidgetDefinitionType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTimelineWidgetDefinitionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTimelineWidgetDefinitionType(val *EventTimelineWidgetDefinitionType) *NullableEventTimelineWidgetDefinitionType {
	return &NullableEventTimelineWidgetDefinitionType{value: val, isSet: true}
}

func (v NullableEventTimelineWidgetDefinitionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTimelineWidgetDefinitionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
