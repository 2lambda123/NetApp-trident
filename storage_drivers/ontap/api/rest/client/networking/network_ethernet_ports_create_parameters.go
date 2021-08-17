// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewNetworkEthernetPortsCreateParams creates a new NetworkEthernetPortsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkEthernetPortsCreateParams() *NetworkEthernetPortsCreateParams {
	return &NetworkEthernetPortsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkEthernetPortsCreateParamsWithTimeout creates a new NetworkEthernetPortsCreateParams object
// with the ability to set a timeout on a request.
func NewNetworkEthernetPortsCreateParamsWithTimeout(timeout time.Duration) *NetworkEthernetPortsCreateParams {
	return &NetworkEthernetPortsCreateParams{
		timeout: timeout,
	}
}

// NewNetworkEthernetPortsCreateParamsWithContext creates a new NetworkEthernetPortsCreateParams object
// with the ability to set a context for a request.
func NewNetworkEthernetPortsCreateParamsWithContext(ctx context.Context) *NetworkEthernetPortsCreateParams {
	return &NetworkEthernetPortsCreateParams{
		Context: ctx,
	}
}

// NewNetworkEthernetPortsCreateParamsWithHTTPClient creates a new NetworkEthernetPortsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkEthernetPortsCreateParamsWithHTTPClient(client *http.Client) *NetworkEthernetPortsCreateParams {
	return &NetworkEthernetPortsCreateParams{
		HTTPClient: client,
	}
}

/* NetworkEthernetPortsCreateParams contains all the parameters to send to the API endpoint
   for the network ethernet ports create operation.

   Typically these are written to a http.Request.
*/
type NetworkEthernetPortsCreateParams struct {

	/* Info.

	   Port parameters
	*/
	Info *models.Port

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network ethernet ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetPortsCreateParams) WithDefaults() *NetworkEthernetPortsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network ethernet ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkEthernetPortsCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NetworkEthernetPortsCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) WithTimeout(timeout time.Duration) *NetworkEthernetPortsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) WithContext(ctx context.Context) *NetworkEthernetPortsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) WithHTTPClient(client *http.Client) *NetworkEthernetPortsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) WithInfo(info *models.Port) *NetworkEthernetPortsCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) SetInfo(info *models.Port) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) WithReturnRecords(returnRecords *bool) *NetworkEthernetPortsCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the network ethernet ports create params
func (o *NetworkEthernetPortsCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkEthernetPortsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
