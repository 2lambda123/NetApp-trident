// Code generated by go-swagger; DO NOT EDIT.

package support

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

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NewAutosupportModifyParams creates a new AutosupportModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAutosupportModifyParams() *AutosupportModifyParams {
	return &AutosupportModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAutosupportModifyParamsWithTimeout creates a new AutosupportModifyParams object
// with the ability to set a timeout on a request.
func NewAutosupportModifyParamsWithTimeout(timeout time.Duration) *AutosupportModifyParams {
	return &AutosupportModifyParams{
		timeout: timeout,
	}
}

// NewAutosupportModifyParamsWithContext creates a new AutosupportModifyParams object
// with the ability to set a context for a request.
func NewAutosupportModifyParamsWithContext(ctx context.Context) *AutosupportModifyParams {
	return &AutosupportModifyParams{
		Context: ctx,
	}
}

// NewAutosupportModifyParamsWithHTTPClient creates a new AutosupportModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAutosupportModifyParamsWithHTTPClient(client *http.Client) *AutosupportModifyParams {
	return &AutosupportModifyParams{
		HTTPClient: client,
	}
}

/* AutosupportModifyParams contains all the parameters to send to the API endpoint
   for the autosupport modify operation.

   Typically these are written to a http.Request.
*/
type AutosupportModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Autosupport

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the autosupport modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutosupportModifyParams) WithDefaults() *AutosupportModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the autosupport modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutosupportModifyParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := AutosupportModifyParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the autosupport modify params
func (o *AutosupportModifyParams) WithTimeout(timeout time.Duration) *AutosupportModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the autosupport modify params
func (o *AutosupportModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the autosupport modify params
func (o *AutosupportModifyParams) WithContext(ctx context.Context) *AutosupportModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the autosupport modify params
func (o *AutosupportModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the autosupport modify params
func (o *AutosupportModifyParams) WithHTTPClient(client *http.Client) *AutosupportModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the autosupport modify params
func (o *AutosupportModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the autosupport modify params
func (o *AutosupportModifyParams) WithInfo(info *models.Autosupport) *AutosupportModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the autosupport modify params
func (o *AutosupportModifyParams) SetInfo(info *models.Autosupport) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the autosupport modify params
func (o *AutosupportModifyParams) WithReturnRecords(returnRecords *bool) *AutosupportModifyParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the autosupport modify params
func (o *AutosupportModifyParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *AutosupportModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
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
