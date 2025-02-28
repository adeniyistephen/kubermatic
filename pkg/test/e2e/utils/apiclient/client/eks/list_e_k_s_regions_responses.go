// Code generated by go-swagger; DO NOT EDIT.

package eks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListEKSRegionsReader is a Reader for the ListEKSRegions structure.
type ListEKSRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSRegionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSRegionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSRegionsOK creates a ListEKSRegionsOK with default headers values
func NewListEKSRegionsOK() *ListEKSRegionsOK {
	return &ListEKSRegionsOK{}
}

/* ListEKSRegionsOK describes a response with status code 200, with default header values.

EKSRegions
*/
type ListEKSRegionsOK struct {
	Payload []models.EKSRegions
}

func (o *ListEKSRegionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/regions][%d] listEKSRegionsOK  %+v", 200, o.Payload)
}
func (o *ListEKSRegionsOK) GetPayload() []models.EKSRegions {
	return o.Payload
}

func (o *ListEKSRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSRegionsUnauthorized creates a ListEKSRegionsUnauthorized with default headers values
func NewListEKSRegionsUnauthorized() *ListEKSRegionsUnauthorized {
	return &ListEKSRegionsUnauthorized{}
}

/* ListEKSRegionsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSRegionsUnauthorized struct {
}

func (o *ListEKSRegionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/regions][%d] listEKSRegionsUnauthorized ", 401)
}

func (o *ListEKSRegionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSRegionsForbidden creates a ListEKSRegionsForbidden with default headers values
func NewListEKSRegionsForbidden() *ListEKSRegionsForbidden {
	return &ListEKSRegionsForbidden{}
}

/* ListEKSRegionsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSRegionsForbidden struct {
}

func (o *ListEKSRegionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/regions][%d] listEKSRegionsForbidden ", 403)
}

func (o *ListEKSRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSRegionsDefault creates a ListEKSRegionsDefault with default headers values
func NewListEKSRegionsDefault(code int) *ListEKSRegionsDefault {
	return &ListEKSRegionsDefault{
		_statusCode: code,
	}
}

/* ListEKSRegionsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSRegionsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s regions default response
func (o *ListEKSRegionsDefault) Code() int {
	return o._statusCode
}

func (o *ListEKSRegionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/providers/eks/regions][%d] listEKSRegions default  %+v", o._statusCode, o.Payload)
}
func (o *ListEKSRegionsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSRegionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
