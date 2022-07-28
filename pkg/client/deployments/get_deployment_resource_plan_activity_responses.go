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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetDeploymentResourcePlanActivityReader is a Reader for the GetDeploymentResourcePlanActivity structure.
type GetDeploymentResourcePlanActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentResourcePlanActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentResourcePlanActivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentResourcePlanActivityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentResourcePlanActivityOK creates a GetDeploymentResourcePlanActivityOK with default headers values
func NewGetDeploymentResourcePlanActivityOK() *GetDeploymentResourcePlanActivityOK {
	return &GetDeploymentResourcePlanActivityOK{}
}

/* GetDeploymentResourcePlanActivityOK describes a response with status code 200, with default header values.

Returning the single plan activity for the specified resource
*/
type GetDeploymentResourcePlanActivityOK struct {
	Payload *models.CommonClusterPlanInfo
}

func (o *GetDeploymentResourcePlanActivityOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/activity/{attempt_id}][%d] getDeploymentResourcePlanActivityOK  %+v", 200, o.Payload)
}
func (o *GetDeploymentResourcePlanActivityOK) GetPayload() *models.CommonClusterPlanInfo {
	return o.Payload
}

func (o *GetDeploymentResourcePlanActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonClusterPlanInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentResourcePlanActivityNotFound creates a GetDeploymentResourcePlanActivityNotFound with default headers values
func NewGetDeploymentResourcePlanActivityNotFound() *GetDeploymentResourcePlanActivityNotFound {
	return &GetDeploymentResourcePlanActivityNotFound{}
}

/* GetDeploymentResourcePlanActivityNotFound describes a response with status code 404, with default header values.

 * The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
* The Resource specified by {ref_id} cannot be found. (code: `deployments.deployment_resource_not_found`)
*/
type GetDeploymentResourcePlanActivityNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentResourcePlanActivityNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/{resource_kind}/{ref_id}/activity/{attempt_id}][%d] getDeploymentResourcePlanActivityNotFound  %+v", 404, o.Payload)
}
func (o *GetDeploymentResourcePlanActivityNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentResourcePlanActivityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
