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

// ListEKSInstanceTypesNoCredentialsReader is a Reader for the ListEKSInstanceTypesNoCredentials structure.
type ListEKSInstanceTypesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListEKSInstanceTypesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListEKSInstanceTypesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListEKSInstanceTypesNoCredentialsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListEKSInstanceTypesNoCredentialsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListEKSInstanceTypesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListEKSInstanceTypesNoCredentialsOK creates a ListEKSInstanceTypesNoCredentialsOK with default headers values
func NewListEKSInstanceTypesNoCredentialsOK() *ListEKSInstanceTypesNoCredentialsOK {
	return &ListEKSInstanceTypesNoCredentialsOK{}
}

/* ListEKSInstanceTypesNoCredentialsOK describes a response with status code 200, with default header values.

EKSInstanceTypes
*/
type ListEKSInstanceTypesNoCredentialsOK struct {
	Payload models.EKSInstanceTypes
}

func (o *ListEKSInstanceTypesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/instancetypes][%d] listEKSInstanceTypesNoCredentialsOK  %+v", 200, o.Payload)
}
func (o *ListEKSInstanceTypesNoCredentialsOK) GetPayload() models.EKSInstanceTypes {
	return o.Payload
}

func (o *ListEKSInstanceTypesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListEKSInstanceTypesNoCredentialsUnauthorized creates a ListEKSInstanceTypesNoCredentialsUnauthorized with default headers values
func NewListEKSInstanceTypesNoCredentialsUnauthorized() *ListEKSInstanceTypesNoCredentialsUnauthorized {
	return &ListEKSInstanceTypesNoCredentialsUnauthorized{}
}

/* ListEKSInstanceTypesNoCredentialsUnauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type ListEKSInstanceTypesNoCredentialsUnauthorized struct {
}

func (o *ListEKSInstanceTypesNoCredentialsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/instancetypes][%d] listEKSInstanceTypesNoCredentialsUnauthorized ", 401)
}

func (o *ListEKSInstanceTypesNoCredentialsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSInstanceTypesNoCredentialsForbidden creates a ListEKSInstanceTypesNoCredentialsForbidden with default headers values
func NewListEKSInstanceTypesNoCredentialsForbidden() *ListEKSInstanceTypesNoCredentialsForbidden {
	return &ListEKSInstanceTypesNoCredentialsForbidden{}
}

/* ListEKSInstanceTypesNoCredentialsForbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type ListEKSInstanceTypesNoCredentialsForbidden struct {
}

func (o *ListEKSInstanceTypesNoCredentialsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/instancetypes][%d] listEKSInstanceTypesNoCredentialsForbidden ", 403)
}

func (o *ListEKSInstanceTypesNoCredentialsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListEKSInstanceTypesNoCredentialsDefault creates a ListEKSInstanceTypesNoCredentialsDefault with default headers values
func NewListEKSInstanceTypesNoCredentialsDefault(code int) *ListEKSInstanceTypesNoCredentialsDefault {
	return &ListEKSInstanceTypesNoCredentialsDefault{
		_statusCode: code,
	}
}

/* ListEKSInstanceTypesNoCredentialsDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListEKSInstanceTypesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list e k s instance types no credentials default response
func (o *ListEKSInstanceTypesNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListEKSInstanceTypesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v2/projects/{project_id}/kubernetes/clusters/{cluster_id}/providers/eks/instancetypes][%d] listEKSInstanceTypesNoCredentials default  %+v", o._statusCode, o.Payload)
}
func (o *ListEKSInstanceTypesNoCredentialsDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListEKSInstanceTypesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
