// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// NewSnaplockLegalHoldCreateParams creates a new SnaplockLegalHoldCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldCreateParams() *SnaplockLegalHoldCreateParams {
	return &SnaplockLegalHoldCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldCreateParamsWithTimeout creates a new SnaplockLegalHoldCreateParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldCreateParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldCreateParams {
	return &SnaplockLegalHoldCreateParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldCreateParamsWithContext creates a new SnaplockLegalHoldCreateParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldCreateParamsWithContext(ctx context.Context) *SnaplockLegalHoldCreateParams {
	return &SnaplockLegalHoldCreateParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldCreateParamsWithHTTPClient creates a new SnaplockLegalHoldCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldCreateParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldCreateParams {
	return &SnaplockLegalHoldCreateParams{
		HTTPClient: client,
	}
}

/* SnaplockLegalHoldCreateParams contains all the parameters to send to the API endpoint
   for the snaplock legal hold create operation.

   Typically these are written to a http.Request.
*/
type SnaplockLegalHoldCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SnaplockLegalHoldOperation

	/* LitigationID.

	   Litigation ID
	*/
	LitigationIDPathParameter string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldCreateParams) WithDefaults() *SnaplockLegalHoldCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := SnaplockLegalHoldCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) WithContext(ctx context.Context) *SnaplockLegalHoldCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) WithInfo(info *models.SnaplockLegalHoldOperation) *SnaplockLegalHoldCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) SetInfo(info *models.SnaplockLegalHoldOperation) {
	o.Info = info
}

// WithLitigationIDPathParameter adds the litigationID to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) WithLitigationIDPathParameter(litigationID string) *SnaplockLegalHoldCreateParams {
	o.SetLitigationIDPathParameter(litigationID)
	return o
}

// SetLitigationIDPathParameter adds the litigationId to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) SetLitigationIDPathParameter(litigationID string) {
	o.LitigationIDPathParameter = litigationID
}

// WithReturnRecords adds the returnRecords to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) WithReturnRecords(returnRecords *bool) *SnaplockLegalHoldCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snaplock legal hold create params
func (o *SnaplockLegalHoldCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param litigation.id
	if err := r.SetPathParam("litigation.id", o.LitigationIDPathParameter); err != nil {
		return err
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
