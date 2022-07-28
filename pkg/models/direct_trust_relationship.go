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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectTrustRelationship The trust relationship with entities trusted directly having their certificate bundled together with the trust settings.
//
// swagger:model DirectTrustRelationship
type DirectTrustRelationship struct {

	// A list of node names trusted in addition to those deducible from trust_allowlist and scope id. Allows trusting nodes that don't have a scoped name at the cost of maintaining the list. Mandatory if scope id is not defined. Wildcards are not allowed.
	AdditionalNodeNames []string `json:"additional_node_names"`

	// The public ca certificate(s) to trust. Only one is required, but it is possible to specify multiple certificates in order to facilitate key rotation.
	// Required: true
	Certificates []*TrustedCertificate `json:"certificates"`

	// a human readable name of the trust relationship
	// Required: true
	Name *string `json:"name"`

	// A lowercase alphanumerical string of max 32 characters. Usually an organization id or an environment id, but could really be any suitable suffix for clusters using the CA certificate of this trust. Required unless trust_all is false and trust_allowlist is empty.
	// Example: abc123
	ScopeID string `json:"scope_id,omitempty"`

	// If true, scope_id is required and the `trust_allowlist` is ignored and all clusters matching the scope id will be trusted.
	// Required: true
	TrustAll *bool `json:"trust_all"`

	// The list of clusters with matching scope to trust. Only used when `trust_all` is false. Providing one or more clusters makes scope_id mandatory.
	TrustAllowlist []string `json:"trust_allowlist"`

	// The type can either be ESS, ECE or generic. If none is specified, then generic is assumed.
	// Enum: [ECE ESS generic]
	Type string `json:"type,omitempty"`

	// Auto generated identifier for this trust, allows distinguishing between update vs remove and add.
	UID string `json:"uid,omitempty"`
}

// Validate validates this direct trust relationship
func (m *DirectTrustRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrustAll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectTrustRelationship) validateCertificates(formats strfmt.Registry) error {

	if err := validate.Required("certificates", "body", m.Certificates); err != nil {
		return err
	}

	for i := 0; i < len(m.Certificates); i++ {
		if swag.IsZero(m.Certificates[i]) { // not required
			continue
		}

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DirectTrustRelationship) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DirectTrustRelationship) validateTrustAll(formats strfmt.Registry) error {

	if err := validate.Required("trust_all", "body", m.TrustAll); err != nil {
		return err
	}

	return nil
}

var directTrustRelationshipTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ECE","ESS","generic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directTrustRelationshipTypeTypePropEnum = append(directTrustRelationshipTypeTypePropEnum, v)
	}
}

const (

	// DirectTrustRelationshipTypeECE captures enum value "ECE"
	DirectTrustRelationshipTypeECE string = "ECE"

	// DirectTrustRelationshipTypeESS captures enum value "ESS"
	DirectTrustRelationshipTypeESS string = "ESS"

	// DirectTrustRelationshipTypeGeneric captures enum value "generic"
	DirectTrustRelationshipTypeGeneric string = "generic"
)

// prop value enum
func (m *DirectTrustRelationship) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, directTrustRelationshipTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DirectTrustRelationship) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this direct trust relationship based on the context it is used
func (m *DirectTrustRelationship) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectTrustRelationship) contextValidateCertificates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Certificates); i++ {

		if m.Certificates[i] != nil {
			if err := m.Certificates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("certificates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectTrustRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectTrustRelationship) UnmarshalBinary(b []byte) error {
	var res DirectTrustRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
