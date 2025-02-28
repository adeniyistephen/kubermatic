// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// RevokeClusterAdminTokenV2Reader is a Reader for the RevokeClusterAdminTokenV2 structure.
type RevokeClusterAdminTokenV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeClusterAdminTokenV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeClusterAdminTokenV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRevokeClusterAdminTokenV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeClusterAdminTokenV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRevokeClusterAdminTokenV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeClusterAdminTokenV2OK creates a RevokeClusterAdminTokenV2OK with default headers values
func NewRevokeClusterAdminTokenV2OK() *RevokeClusterAdminTokenV2OK {
	return &RevokeClusterAdminTokenV2OK{}
}

/* RevokeClusterAdminTokenV2OK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenV2OK struct {
}

func (o *RevokeClusterAdminTokenV2OK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2OK ", 200)
}

func (o *RevokeClusterAdminTokenV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenV2Unauthorized creates a RevokeClusterAdminTokenV2Unauthorized with default headers values
func NewRevokeClusterAdminTokenV2Unauthorized() *RevokeClusterAdminTokenV2Unauthorized {
	return &RevokeClusterAdminTokenV2Unauthorized{}
}

/* RevokeClusterAdminTokenV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenV2Unauthorized struct {
}

func (o *RevokeClusterAdminTokenV2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2Unauthorized ", 401)
}

func (o *RevokeClusterAdminTokenV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenV2Forbidden creates a RevokeClusterAdminTokenV2Forbidden with default headers values
func NewRevokeClusterAdminTokenV2Forbidden() *RevokeClusterAdminTokenV2Forbidden {
	return &RevokeClusterAdminTokenV2Forbidden{}
}

/* RevokeClusterAdminTokenV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type RevokeClusterAdminTokenV2Forbidden struct {
}

func (o *RevokeClusterAdminTokenV2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2Forbidden ", 403)
}

func (o *RevokeClusterAdminTokenV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRevokeClusterAdminTokenV2Default creates a RevokeClusterAdminTokenV2Default with default headers values
func NewRevokeClusterAdminTokenV2Default(code int) *RevokeClusterAdminTokenV2Default {
	return &RevokeClusterAdminTokenV2Default{
		_statusCode: code,
	}
}

/* RevokeClusterAdminTokenV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type RevokeClusterAdminTokenV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the revoke cluster admin token v2 default response
func (o *RevokeClusterAdminTokenV2Default) Code() int {
	return o._statusCode
}

func (o *RevokeClusterAdminTokenV2Default) Error() string {
	return fmt.Sprintf("[PUT /api/v2/projects/{project_id}/clusters/{cluster_id}/token][%d] revokeClusterAdminTokenV2 default  %+v", o._statusCode, o.Payload)
}
func (o *RevokeClusterAdminTokenV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *RevokeClusterAdminTokenV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
