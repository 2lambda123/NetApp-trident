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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewLocalCifsUserCreateParams creates a new LocalCifsUserCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalCifsUserCreateParams() *LocalCifsUserCreateParams {
	return &LocalCifsUserCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalCifsUserCreateParamsWithTimeout creates a new LocalCifsUserCreateParams object
// with the ability to set a timeout on a request.
func NewLocalCifsUserCreateParamsWithTimeout(timeout time.Duration) *LocalCifsUserCreateParams {
	return &LocalCifsUserCreateParams{
		timeout: timeout,
	}
}

// NewLocalCifsUserCreateParamsWithContext creates a new LocalCifsUserCreateParams object
// with the ability to set a context for a request.
func NewLocalCifsUserCreateParamsWithContext(ctx context.Context) *LocalCifsUserCreateParams {
	return &LocalCifsUserCreateParams{
		Context: ctx,
	}
}

// NewLocalCifsUserCreateParamsWithHTTPClient creates a new LocalCifsUserCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalCifsUserCreateParamsWithHTTPClient(client *http.Client) *LocalCifsUserCreateParams {
	return &LocalCifsUserCreateParams{
		HTTPClient: client,
	}
}

/* LocalCifsUserCreateParams contains all the parameters to send to the API endpoint
   for the local cifs user create operation.

   Typically these are written to a http.Request.
*/
type LocalCifsUserCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.LocalCifsUser

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local cifs user create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsUserCreateParams) WithDefaults() *LocalCifsUserCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local cifs user create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsUserCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := LocalCifsUserCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the local cifs user create params
func (o *LocalCifsUserCreateParams) WithTimeout(timeout time.Duration) *LocalCifsUserCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local cifs user create params
func (o *LocalCifsUserCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local cifs user create params
func (o *LocalCifsUserCreateParams) WithContext(ctx context.Context) *LocalCifsUserCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local cifs user create params
func (o *LocalCifsUserCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local cifs user create params
func (o *LocalCifsUserCreateParams) WithHTTPClient(client *http.Client) *LocalCifsUserCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local cifs user create params
func (o *LocalCifsUserCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the local cifs user create params
func (o *LocalCifsUserCreateParams) WithInfo(info *models.LocalCifsUser) *LocalCifsUserCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the local cifs user create params
func (o *LocalCifsUserCreateParams) SetInfo(info *models.LocalCifsUser) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the local cifs user create params
func (o *LocalCifsUserCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *LocalCifsUserCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the local cifs user create params
func (o *LocalCifsUserCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *LocalCifsUserCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
