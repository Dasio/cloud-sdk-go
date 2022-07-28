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
	"github.com/go-openapi/validate"
)

// IntegrationsServer Holds diagnostics for an Integrations Server resource
//
// swagger:model IntegrationsServer
type IntegrationsServer struct {

	// The backend plan as JSON
	// Required: true
	BackendPlan interface{} `json:"backend_plan"`

	// The human readable name (defaults to the generated cluster id if not specified)
	// Required: true
	DisplayName *string `json:"display_name"`

	// The user-specified id of the Elasticsearch Cluster that this will link to
	// Required: true
	ElasticsearchClusterRefID *string `json:"elasticsearch_cluster_ref_id"`

	// A locally-unique user-specified id
	// Required: true
	RefID *string `json:"ref_id"`
}

// Validate validates this integrations server
func (m *IntegrationsServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackendPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElasticsearchClusterRefID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRefID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationsServer) validateBackendPlan(formats strfmt.Registry) error {

	if m.BackendPlan == nil {
		return errors.Required("backend_plan", "body", nil)
	}

	return nil
}

func (m *IntegrationsServer) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsServer) validateElasticsearchClusterRefID(formats strfmt.Registry) error {

	if err := validate.Required("elasticsearch_cluster_ref_id", "body", m.ElasticsearchClusterRefID); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationsServer) validateRefID(formats strfmt.Registry) error {

	if err := validate.Required("ref_id", "body", m.RefID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this integrations server based on context it is used
func (m *IntegrationsServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationsServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationsServer) UnmarshalBinary(b []byte) error {
	var res IntegrationsServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
