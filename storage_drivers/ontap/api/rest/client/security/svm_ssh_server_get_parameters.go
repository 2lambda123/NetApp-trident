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

// NewSvmSSHServerGetParams creates a new SvmSSHServerGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmSSHServerGetParams() *SvmSSHServerGetParams {
	return &SvmSSHServerGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmSSHServerGetParamsWithTimeout creates a new SvmSSHServerGetParams object
// with the ability to set a timeout on a request.
func NewSvmSSHServerGetParamsWithTimeout(timeout time.Duration) *SvmSSHServerGetParams {
	return &SvmSSHServerGetParams{
		timeout: timeout,
	}
}

// NewSvmSSHServerGetParamsWithContext creates a new SvmSSHServerGetParams object
// with the ability to set a context for a request.
func NewSvmSSHServerGetParamsWithContext(ctx context.Context) *SvmSSHServerGetParams {
	return &SvmSSHServerGetParams{
		Context: ctx,
	}
}

// NewSvmSSHServerGetParamsWithHTTPClient creates a new SvmSSHServerGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmSSHServerGetParamsWithHTTPClient(client *http.Client) *SvmSSHServerGetParams {
	return &SvmSSHServerGetParams{
		HTTPClient: client,
	}
}

/* SvmSSHServerGetParams contains all the parameters to send to the API endpoint
   for the svm ssh server get operation.

   Typically these are written to a http.Request.
*/
type SvmSSHServerGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* SvmUUID.

	   SVM UUID
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the svm ssh server get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerGetParams) WithDefaults() *SvmSSHServerGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm ssh server get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the svm ssh server get params
func (o *SvmSSHServerGetParams) WithTimeout(timeout time.Duration) *SvmSSHServerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm ssh server get params
func (o *SvmSSHServerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm ssh server get params
func (o *SvmSSHServerGetParams) WithContext(ctx context.Context) *SvmSSHServerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm ssh server get params
func (o *SvmSSHServerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm ssh server get params
func (o *SvmSSHServerGetParams) WithHTTPClient(client *http.Client) *SvmSSHServerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm ssh server get params
func (o *SvmSSHServerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the svm ssh server get params
func (o *SvmSSHServerGetParams) WithFieldsQueryParameter(fields []string) *SvmSSHServerGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the svm ssh server get params
func (o *SvmSSHServerGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithSVMUUIDPathParameter adds the svmUUID to the svm ssh server get params
func (o *SvmSSHServerGetParams) WithSVMUUIDPathParameter(svmUUID string) *SvmSSHServerGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the svm ssh server get params
func (o *SvmSSHServerGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmSSHServerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSvmSSHServerGet binds the parameter fields
func (o *SvmSSHServerGetParams) bindParamFields(formats strfmt.Registry) []string {
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
