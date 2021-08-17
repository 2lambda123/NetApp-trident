// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewNvmeSubsystemHostCreateParams creates a new NvmeSubsystemHostCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemHostCreateParams() *NvmeSubsystemHostCreateParams {
	return &NvmeSubsystemHostCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemHostCreateParamsWithTimeout creates a new NvmeSubsystemHostCreateParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemHostCreateParamsWithTimeout(timeout time.Duration) *NvmeSubsystemHostCreateParams {
	return &NvmeSubsystemHostCreateParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemHostCreateParamsWithContext creates a new NvmeSubsystemHostCreateParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemHostCreateParamsWithContext(ctx context.Context) *NvmeSubsystemHostCreateParams {
	return &NvmeSubsystemHostCreateParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemHostCreateParamsWithHTTPClient creates a new NvmeSubsystemHostCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemHostCreateParamsWithHTTPClient(client *http.Client) *NvmeSubsystemHostCreateParams {
	return &NvmeSubsystemHostCreateParams{
		HTTPClient: client,
	}
}

/* NvmeSubsystemHostCreateParams contains all the parameters to send to the API endpoint
   for the nvme subsystem host create operation.

   Typically these are written to a http.Request.
*/
type NvmeSubsystemHostCreateParams struct {

	/* Info.

	   The property values for the NVMe subsystem host to add to the NVMe subsystem.

	*/
	Info *models.NvmeSubsystemHost

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* SubsystemUUID.

	   The unique identifier of the NVMe subsystem.

	*/
	SubsystemUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem host create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostCreateParams) WithDefaults() *NvmeSubsystemHostCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem host create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := NvmeSubsystemHostCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) WithTimeout(timeout time.Duration) *NvmeSubsystemHostCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) WithContext(ctx context.Context) *NvmeSubsystemHostCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) WithHTTPClient(client *http.Client) *NvmeSubsystemHostCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) WithInfo(info *models.NvmeSubsystemHost) *NvmeSubsystemHostCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) SetInfo(info *models.NvmeSubsystemHost) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) WithReturnRecords(returnRecords *bool) *NvmeSubsystemHostCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithSubsystemUUIDPathParameter adds the subsystemUUID to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) WithSubsystemUUIDPathParameter(subsystemUUID string) *NvmeSubsystemHostCreateParams {
	o.SetSubsystemUUIDPathParameter(subsystemUUID)
	return o
}

// SetSubsystemUUIDPathParameter adds the subsystemUuid to the nvme subsystem host create params
func (o *NvmeSubsystemHostCreateParams) SetSubsystemUUIDPathParameter(subsystemUUID string) {
	o.SubsystemUUIDPathParameter = subsystemUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemHostCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subsystem.uuid
	if err := r.SetPathParam("subsystem.uuid", o.SubsystemUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
