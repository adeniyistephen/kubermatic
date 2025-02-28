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

// ResetAlertmanagerReader is a Reader for the ResetAlertmanager structure.
type ResetAlertmanagerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetAlertmanagerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetAlertmanagerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewResetAlertmanagerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewResetAlertmanagerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewResetAlertmanagerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResetAlertmanagerOK creates a ResetAlertmanagerOK with default headers values
func NewResetAlertmanagerOK() *ResetAlertmanagerOK {
	return &ResetAlertmanagerOK{}
}

/* ResetAlertmanagerOK describes a response with status code 200, with default header values.

EmptyResponse is a empty response
*/
type ResetAlertmanagerOK struct {
}

func (o *ResetAlertmanagerOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] resetAlertmanagerOK ", 200)
}

func (o *ResetAlertmanagerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetAlertmanagerUnauthorized creates a ResetAlertmanagerUnauthorized with default headers values
func NewResetAlertmanagerUnauthorized() *ResetAlertmanagerUnauthorized {
	return &ResetAlertmanagerUnauthorized{}
}

/* ResetAlertmanagerUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ResetAlertmanagerUnauthorized struct {
}

func (o *ResetAlertmanagerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] resetAlertmanagerUnauthorized ", 401)
}

func (o *ResetAlertmanagerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetAlertmanagerForbidden creates a ResetAlertmanagerForbidden with default headers values
func NewResetAlertmanagerForbidden() *ResetAlertmanagerForbidden {
	return &ResetAlertmanagerForbidden{}
}

/* ResetAlertmanagerForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ResetAlertmanagerForbidden struct {
}

func (o *ResetAlertmanagerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] resetAlertmanagerForbidden ", 403)
}

func (o *ResetAlertmanagerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResetAlertmanagerDefault creates a ResetAlertmanagerDefault with default headers values
func NewResetAlertmanagerDefault(code int) *ResetAlertmanagerDefault {
	return &ResetAlertmanagerDefault{
		_statusCode: code,
	}
}

/* ResetAlertmanagerDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ResetAlertmanagerDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the reset alertmanager default response
func (o *ResetAlertmanagerDefault) Code() int {
	return o._statusCode
}

func (o *ResetAlertmanagerDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/projects/{project_id}/clusters/{cluster_id}/alertmanager/config][%d] resetAlertmanager default  %+v", o._statusCode, o.Payload)
}
func (o *ResetAlertmanagerDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ResetAlertmanagerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
