// Code generated by go-swagger; DO NOT EDIT.

package applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// GetApplicationInstallationReader is a Reader for the GetApplicationInstallation structure.
type GetApplicationInstallationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApplicationInstallationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApplicationInstallationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetApplicationInstallationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetApplicationInstallationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetApplicationInstallationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetApplicationInstallationOK creates a GetApplicationInstallationOK with default headers values
func NewGetApplicationInstallationOK() *GetApplicationInstallationOK {
	return &GetApplicationInstallationOK{}
}

/* GetApplicationInstallationOK describes a response with status code 200, with default header values.

ApplicationInstallation
*/
type GetApplicationInstallationOK struct {
	Payload *models.ApplicationInstallation
}

func (o *GetApplicationInstallationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] getApplicationInstallationOK  %+v", 200, o.Payload)
}
func (o *GetApplicationInstallationOK) GetPayload() *models.ApplicationInstallation {
	return o.Payload
}

func (o *GetApplicationInstallationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplicationInstallation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApplicationInstallationUnauthorized creates a GetApplicationInstallationUnauthorized with default headers values
func NewGetApplicationInstallationUnauthorized() *GetApplicationInstallationUnauthorized {
	return &GetApplicationInstallationUnauthorized{}
}

/* GetApplicationInstallationUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type GetApplicationInstallationUnauthorized struct {
}

func (o *GetApplicationInstallationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] getApplicationInstallationUnauthorized ", 401)
}

func (o *GetApplicationInstallationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationInstallationForbidden creates a GetApplicationInstallationForbidden with default headers values
func NewGetApplicationInstallationForbidden() *GetApplicationInstallationForbidden {
	return &GetApplicationInstallationForbidden{}
}

/* GetApplicationInstallationForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type GetApplicationInstallationForbidden struct {
}

func (o *GetApplicationInstallationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] getApplicationInstallationForbidden ", 403)
}

func (o *GetApplicationInstallationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetApplicationInstallationDefault creates a GetApplicationInstallationDefault with default headers values
func NewGetApplicationInstallationDefault(code int) *GetApplicationInstallationDefault {
	return &GetApplicationInstallationDefault{
		_statusCode: code,
	}
}

/* GetApplicationInstallationDefault describes a response with status code -1, with default header values.

errorResponse
*/
type GetApplicationInstallationDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get application installation default response
func (o *GetApplicationInstallationDefault) Code() int {
	return o._statusCode
}

func (o *GetApplicationInstallationDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/clusters/{cluster_id}/applicationinstallations/{namespace}/{appinstall_name}][%d] getApplicationInstallation default  %+v", o._statusCode, o.Payload)
}
func (o *GetApplicationInstallationDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetApplicationInstallationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
