// Code generated by go-swagger; DO NOT EDIT.

package security

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
	"github.com/go-openapi/swag"
)

// NewClusterNisGetParams creates a new ClusterNisGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterNisGetParams() *ClusterNisGetParams {
	return &ClusterNisGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterNisGetParamsWithTimeout creates a new ClusterNisGetParams object
// with the ability to set a timeout on a request.
func NewClusterNisGetParamsWithTimeout(timeout time.Duration) *ClusterNisGetParams {
	return &ClusterNisGetParams{
		timeout: timeout,
	}
}

// NewClusterNisGetParamsWithContext creates a new ClusterNisGetParams object
// with the ability to set a context for a request.
func NewClusterNisGetParamsWithContext(ctx context.Context) *ClusterNisGetParams {
	return &ClusterNisGetParams{
		Context: ctx,
	}
}

// NewClusterNisGetParamsWithHTTPClient creates a new ClusterNisGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterNisGetParamsWithHTTPClient(client *http.Client) *ClusterNisGetParams {
	return &ClusterNisGetParams{
		HTTPClient: client,
	}
}

/* ClusterNisGetParams contains all the parameters to send to the API endpoint
   for the cluster nis get operation.

   Typically these are written to a http.Request.
*/
type ClusterNisGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster nis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNisGetParams) WithDefaults() *ClusterNisGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster nis get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNisGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster nis get params
func (o *ClusterNisGetParams) WithTimeout(timeout time.Duration) *ClusterNisGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster nis get params
func (o *ClusterNisGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster nis get params
func (o *ClusterNisGetParams) WithContext(ctx context.Context) *ClusterNisGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster nis get params
func (o *ClusterNisGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster nis get params
func (o *ClusterNisGetParams) WithHTTPClient(client *http.Client) *ClusterNisGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster nis get params
func (o *ClusterNisGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the cluster nis get params
func (o *ClusterNisGetParams) WithFieldsQueryParameter(fields []string) *ClusterNisGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cluster nis get params
func (o *ClusterNisGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterNisGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamClusterNisGet binds the parameter fields
func (o *ClusterNisGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
