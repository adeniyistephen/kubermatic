// Code generated by go-swagger; DO NOT EDIT.

package ipampool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// CreateIPAMPoolReader is a Reader for the CreateIPAMPool structure.
type CreateIPAMPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIPAMPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateIPAMPoolCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateIPAMPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateIPAMPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateIPAMPoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIPAMPoolCreated creates a CreateIPAMPoolCreated with default headers values
func NewCreateIPAMPoolCreated() *CreateIPAMPoolCreated {
	return &CreateIPAMPoolCreated{}
}

/* CreateIPAMPoolCreated describes a response with status code 201, with default header values.

EmptyResponse is a empty response
*/
type CreateIPAMPoolCreated struct {
}

func (o *CreateIPAMPoolCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/ipampools][%d] createIpAMPoolCreated ", 201)
}

func (o *CreateIPAMPoolCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateIPAMPoolUnauthorized creates a CreateIPAMPoolUnauthorized with default headers values
func NewCreateIPAMPoolUnauthorized() *CreateIPAMPoolUnauthorized {
	return &CreateIPAMPoolUnauthorized{}
}

/* CreateIPAMPoolUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type CreateIPAMPoolUnauthorized struct {
}

func (o *CreateIPAMPoolUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/ipampools][%d] createIpAMPoolUnauthorized ", 401)
}

func (o *CreateIPAMPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateIPAMPoolForbidden creates a CreateIPAMPoolForbidden with default headers values
func NewCreateIPAMPoolForbidden() *CreateIPAMPoolForbidden {
	return &CreateIPAMPoolForbidden{}
}

/* CreateIPAMPoolForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type CreateIPAMPoolForbidden struct {
}

func (o *CreateIPAMPoolForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/ipampools][%d] createIpAMPoolForbidden ", 403)
}

func (o *CreateIPAMPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateIPAMPoolDefault creates a CreateIPAMPoolDefault with default headers values
func NewCreateIPAMPoolDefault(code int) *CreateIPAMPoolDefault {
	return &CreateIPAMPoolDefault{
		_statusCode: code,
	}
}

/* CreateIPAMPoolDefault describes a response with status code -1, with default header values.

errorResponse
*/
type CreateIPAMPoolDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create IP a m pool default response
func (o *CreateIPAMPoolDefault) Code() int {
	return o._statusCode
}

func (o *CreateIPAMPoolDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/seeds/{seed_name}/ipampools][%d] createIPAMPool default  %+v", o._statusCode, o.Payload)
}
func (o *CreateIPAMPoolDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateIPAMPoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
