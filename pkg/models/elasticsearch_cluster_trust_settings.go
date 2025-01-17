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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ElasticsearchClusterTrustSettings Configuration of trust with other clusters.
//
// swagger:model ElasticsearchClusterTrustSettings
type ElasticsearchClusterTrustSettings struct {

	// The list of trust relationships with different accounts
	Accounts []*AccountTrustRelationship `json:"accounts"`

	// The list of trust relationships where the certificate is bundled with the trust setting. Allows configuring trust for clusters running outside of an Elastic Cloud managed environment or in an Elastic Cloud environment without an environment level trust established.
	Direct []*DirectTrustRelationship `json:"direct"`

	// The list of trust relationships with external entities
	External []*ExternalTrustRelationship `json:"external"`
}

// Validate validates this elasticsearch cluster trust settings
func (m *ElasticsearchClusterTrustSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterTrustSettings) validateAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.Accounts) { // not required
		return nil
	}

	for i := 0; i < len(m.Accounts); i++ {
		if swag.IsZero(m.Accounts[i]) { // not required
			continue
		}

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterTrustSettings) validateDirect(formats strfmt.Registry) error {
	if swag.IsZero(m.Direct) { // not required
		return nil
	}

	for i := 0; i < len(m.Direct); i++ {
		if swag.IsZero(m.Direct[i]) { // not required
			continue
		}

		if m.Direct[i] != nil {
			if err := m.Direct[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("direct" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterTrustSettings) validateExternal(formats strfmt.Registry) error {
	if swag.IsZero(m.External) { // not required
		return nil
	}

	for i := 0; i < len(m.External); i++ {
		if swag.IsZero(m.External[i]) { // not required
			continue
		}

		if m.External[i] != nil {
			if err := m.External[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this elasticsearch cluster trust settings based on the context it is used
func (m *ElasticsearchClusterTrustSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDirect(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterTrustSettings) contextValidateAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Accounts); i++ {

		if m.Accounts[i] != nil {
			if err := m.Accounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterTrustSettings) contextValidateDirect(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Direct); i++ {

		if m.Direct[i] != nil {
			if err := m.Direct[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("direct" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterTrustSettings) contextValidateExternal(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.External); i++ {

		if m.External[i] != nil {
			if err := m.External[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("external" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterTrustSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterTrustSettings) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterTrustSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
