// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElasticsearchSystemSettings A subset of Elasticsearch settings. TIP: To define the complete set of Elasticsearch settings, use `ElasticsearchSystemSettings` with `user_settings_override*` and `user_settings*`.
//
// swagger:model ElasticsearchSystemSettings
type ElasticsearchSystemSettings struct {

	// If true (the default), then any write operation on an index that does not currently exist will create it. NOTES: (Corresponds to the parameter 'action.auto_create_index')
	AutoCreateIndex *bool `json:"auto_create_index,omitempty"`

	// (2.x only - to get the same result in 5.x template mappings must be used) Sets the default number of shards per index, defaulting to 1 if not specified. (Corresponds to the parameter 'index.number_of_shards' in 2.x, not supported in 5.x)
	DefaultShardsPerIndex int32 `json:"default_shards_per_index,omitempty"`

	// If true (default is false) then the index deletion API will not support wildcards or '_all'. NOTES: (Corresponds to the parameter 'action.destructive_requires_name')
	DestructiveRequiresName *bool `json:"destructive_requires_name,omitempty"`

	// Defaults to false on versions <= 7.2.0, true otherwise. If false, then the API commands to close indices are disabled. This is important because Elasticsearch does not snapshot or migrate close indices on versions under 7.2.0, therefore standard Elastic Cloud configuration operations will cause irretrievable loss of indices' data. NOTES: (Corresponds to the parameter 'cluster.indices.close.enable')
	EnableCloseIndex *bool `json:"enable_close_index,omitempty"`

	// The default interval at which monitoring information from the cluster if collected, if monitoring is enabled. NOTES: (Corresponds to the parameter 'marvel.agent.interval' in 2.x and 'xpack.monitoring.collection.interval' in 5.x)
	MonitoringCollectionInterval int32 `json:"monitoring_collection_interval,omitempty"`

	// The duration for which monitoring history is stored (format '(NUMBER)d' eg '3d' for 3 days). NOTES: ('Corresponds to the parameter xpack.monitoring.history.duration' in 5.x, defaults to '7d')
	MonitoringHistoryDuration string `json:"monitoring_history_duration,omitempty"`

	// Limits remote Elasticsearch clusters that can be used as the source for '_reindex' API commands
	ReindexWhitelist []string `json:"reindex_whitelist"`

	// scripting
	Scripting *ElasticsearchScriptingUserSettings `json:"scripting,omitempty"`

	// The trigger engine for Watcher, defaults to 'scheduler' - see the xpack documentation for more information. NOTES: (Corresponds to the parameter '(xpack.)watcher.trigger.schedule.engine', depending on version. Ignored from 6.x onwards.)
	WatcherTriggerEngine string `json:"watcher_trigger_engine,omitempty"`
}

// Validate validates this elasticsearch system settings
func (m *ElasticsearchSystemSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScripting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchSystemSettings) validateScripting(formats strfmt.Registry) error {
	if swag.IsZero(m.Scripting) { // not required
		return nil
	}

	if m.Scripting != nil {
		if err := m.Scripting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scripting")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this elasticsearch system settings based on the context it is used
func (m *ElasticsearchSystemSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScripting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchSystemSettings) contextValidateScripting(ctx context.Context, formats strfmt.Registry) error {

	if m.Scripting != nil {
		if err := m.Scripting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scripting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchSystemSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchSystemSettings) UnmarshalBinary(b []byte) error {
	var res ElasticsearchSystemSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
