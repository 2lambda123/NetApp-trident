// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewS3AuditGetParams creates a new S3AuditGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3AuditGetParams() *S3AuditGetParams {
	return &S3AuditGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3AuditGetParamsWithTimeout creates a new S3AuditGetParams object
// with the ability to set a timeout on a request.
func NewS3AuditGetParamsWithTimeout(timeout time.Duration) *S3AuditGetParams {
	return &S3AuditGetParams{
		timeout: timeout,
	}
}

// NewS3AuditGetParamsWithContext creates a new S3AuditGetParams object
// with the ability to set a context for a request.
func NewS3AuditGetParamsWithContext(ctx context.Context) *S3AuditGetParams {
	return &S3AuditGetParams{
		Context: ctx,
	}
}

// NewS3AuditGetParamsWithHTTPClient creates a new S3AuditGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3AuditGetParamsWithHTTPClient(client *http.Client) *S3AuditGetParams {
	return &S3AuditGetParams{
		HTTPClient: client,
	}
}

/* S3AuditGetParams contains all the parameters to send to the API endpoint
   for the s3 audit get operation.

   Typically these are written to a http.Request.
*/
type S3AuditGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditGetParams) WithDefaults() *S3AuditGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the s3 audit get params
func (o *S3AuditGetParams) WithTimeout(timeout time.Duration) *S3AuditGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 audit get params
func (o *S3AuditGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 audit get params
func (o *S3AuditGetParams) WithContext(ctx context.Context) *S3AuditGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 audit get params
func (o *S3AuditGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 audit get params
func (o *S3AuditGetParams) WithHTTPClient(client *http.Client) *S3AuditGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 audit get params
func (o *S3AuditGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the s3 audit get params
func (o *S3AuditGetParams) WithFieldsQueryParameter(fields []string) *S3AuditGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the s3 audit get params
func (o *S3AuditGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithSVMUUIDPathParameter adds the svmUUID to the s3 audit get params
func (o *S3AuditGetParams) WithSVMUUIDPathParameter(svmUUID string) *S3AuditGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the s3 audit get params
func (o *S3AuditGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3AuditGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamS3AuditGet binds the parameter fields
func (o *S3AuditGetParams) bindParamFields(formats strfmt.Registry) []string {
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
