// Code generated by go-swagger; DO NOT EDIT.

package nutanix

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListNutanixSubnetsParams creates a new ListNutanixSubnetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNutanixSubnetsParams() *ListNutanixSubnetsParams {
	return &ListNutanixSubnetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNutanixSubnetsParamsWithTimeout creates a new ListNutanixSubnetsParams object
// with the ability to set a timeout on a request.
func NewListNutanixSubnetsParamsWithTimeout(timeout time.Duration) *ListNutanixSubnetsParams {
	return &ListNutanixSubnetsParams{
		timeout: timeout,
	}
}

// NewListNutanixSubnetsParamsWithContext creates a new ListNutanixSubnetsParams object
// with the ability to set a context for a request.
func NewListNutanixSubnetsParamsWithContext(ctx context.Context) *ListNutanixSubnetsParams {
	return &ListNutanixSubnetsParams{
		Context: ctx,
	}
}

// NewListNutanixSubnetsParamsWithHTTPClient creates a new ListNutanixSubnetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNutanixSubnetsParamsWithHTTPClient(client *http.Client) *ListNutanixSubnetsParams {
	return &ListNutanixSubnetsParams{
		HTTPClient: client,
	}
}

/* ListNutanixSubnetsParams contains all the parameters to send to the API endpoint
   for the list nutanix subnets operation.

   Typically these are written to a http.Request.
*/
type ListNutanixSubnetsParams struct {

	// Credential.
	Credential *string

	// NutanixCluster.
	NutanixCluster string

	// NutanixPassword.
	NutanixPassword *string

	/* NutanixProject.

	   Project query parameter. Can be omitted to query subnets without project scope
	*/
	NutanixProject *string

	// NutanixProxyURL.
	NutanixProxyURL *string

	// NutanixUsername.
	NutanixUsername *string

	/* Dc.

	   KKP Datacenter to use for endpoint
	*/
	DC string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list nutanix subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNutanixSubnetsParams) WithDefaults() *ListNutanixSubnetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list nutanix subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNutanixSubnetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithTimeout(timeout time.Duration) *ListNutanixSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithContext(ctx context.Context) *ListNutanixSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithHTTPClient(client *http.Client) *ListNutanixSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithCredential(credential *string) *ListNutanixSubnetsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithNutanixCluster adds the nutanixCluster to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithNutanixCluster(nutanixCluster string) *ListNutanixSubnetsParams {
	o.SetNutanixCluster(nutanixCluster)
	return o
}

// SetNutanixCluster adds the nutanixCluster to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetNutanixCluster(nutanixCluster string) {
	o.NutanixCluster = nutanixCluster
}

// WithNutanixPassword adds the nutanixPassword to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithNutanixPassword(nutanixPassword *string) *ListNutanixSubnetsParams {
	o.SetNutanixPassword(nutanixPassword)
	return o
}

// SetNutanixPassword adds the nutanixPassword to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetNutanixPassword(nutanixPassword *string) {
	o.NutanixPassword = nutanixPassword
}

// WithNutanixProject adds the nutanixProject to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithNutanixProject(nutanixProject *string) *ListNutanixSubnetsParams {
	o.SetNutanixProject(nutanixProject)
	return o
}

// SetNutanixProject adds the nutanixProject to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetNutanixProject(nutanixProject *string) {
	o.NutanixProject = nutanixProject
}

// WithNutanixProxyURL adds the nutanixProxyURL to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithNutanixProxyURL(nutanixProxyURL *string) *ListNutanixSubnetsParams {
	o.SetNutanixProxyURL(nutanixProxyURL)
	return o
}

// SetNutanixProxyURL adds the nutanixProxyUrl to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetNutanixProxyURL(nutanixProxyURL *string) {
	o.NutanixProxyURL = nutanixProxyURL
}

// WithNutanixUsername adds the nutanixUsername to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithNutanixUsername(nutanixUsername *string) *ListNutanixSubnetsParams {
	o.SetNutanixUsername(nutanixUsername)
	return o
}

// SetNutanixUsername adds the nutanixUsername to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetNutanixUsername(nutanixUsername *string) {
	o.NutanixUsername = nutanixUsername
}

// WithDC adds the dc to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) WithDC(dc string) *ListNutanixSubnetsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list nutanix subnets params
func (o *ListNutanixSubnetsParams) SetDC(dc string) {
	o.DC = dc
}

// WriteToRequest writes these params to a swagger request
func (o *ListNutanixSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	// header param NutanixCluster
	if err := r.SetHeaderParam("NutanixCluster", o.NutanixCluster); err != nil {
		return err
	}

	if o.NutanixPassword != nil {

		// header param NutanixPassword
		if err := r.SetHeaderParam("NutanixPassword", *o.NutanixPassword); err != nil {
			return err
		}
	}

	if o.NutanixProject != nil {

		// header param NutanixProject
		if err := r.SetHeaderParam("NutanixProject", *o.NutanixProject); err != nil {
			return err
		}
	}

	if o.NutanixProxyURL != nil {

		// header param NutanixProxyURL
		if err := r.SetHeaderParam("NutanixProxyURL", *o.NutanixProxyURL); err != nil {
			return err
		}
	}

	if o.NutanixUsername != nil {

		// header param NutanixUsername
		if err := r.SetHeaderParam("NutanixUsername", *o.NutanixUsername); err != nil {
			return err
		}
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
