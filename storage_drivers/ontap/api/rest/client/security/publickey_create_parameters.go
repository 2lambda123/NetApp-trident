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

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NewPublickeyCreateParams creates a new PublickeyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublickeyCreateParams() *PublickeyCreateParams {
	return &PublickeyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublickeyCreateParamsWithTimeout creates a new PublickeyCreateParams object
// with the ability to set a timeout on a request.
func NewPublickeyCreateParamsWithTimeout(timeout time.Duration) *PublickeyCreateParams {
	return &PublickeyCreateParams{
		timeout: timeout,
	}
}

// NewPublickeyCreateParamsWithContext creates a new PublickeyCreateParams object
// with the ability to set a context for a request.
func NewPublickeyCreateParamsWithContext(ctx context.Context) *PublickeyCreateParams {
	return &PublickeyCreateParams{
		Context: ctx,
	}
}

// NewPublickeyCreateParamsWithHTTPClient creates a new PublickeyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPublickeyCreateParamsWithHTTPClient(client *http.Client) *PublickeyCreateParams {
	return &PublickeyCreateParams{
		HTTPClient: client,
	}
}

/* PublickeyCreateParams contains all the parameters to send to the API endpoint
   for the publickey create operation.

   Typically these are written to a http.Request.
*/
type PublickeyCreateParams struct {

	/* Info.

	   The public key details for the user account.
	*/
	Info *models.Publickey

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the publickey create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublickeyCreateParams) WithDefaults() *PublickeyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the publickey create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublickeyCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := PublickeyCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the publickey create params
func (o *PublickeyCreateParams) WithTimeout(timeout time.Duration) *PublickeyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the publickey create params
func (o *PublickeyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the publickey create params
func (o *PublickeyCreateParams) WithContext(ctx context.Context) *PublickeyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the publickey create params
func (o *PublickeyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the publickey create params
func (o *PublickeyCreateParams) WithHTTPClient(client *http.Client) *PublickeyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the publickey create params
func (o *PublickeyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the publickey create params
func (o *PublickeyCreateParams) WithInfo(info *models.Publickey) *PublickeyCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the publickey create params
func (o *PublickeyCreateParams) SetInfo(info *models.Publickey) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the publickey create params
func (o *PublickeyCreateParams) WithReturnRecords(returnRecords *bool) *PublickeyCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the publickey create params
func (o *PublickeyCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *PublickeyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
