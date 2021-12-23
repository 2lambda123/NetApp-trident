// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewLunMapReportingNodeDeleteParams creates a new LunMapReportingNodeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunMapReportingNodeDeleteParams() *LunMapReportingNodeDeleteParams {
	return &LunMapReportingNodeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunMapReportingNodeDeleteParamsWithTimeout creates a new LunMapReportingNodeDeleteParams object
// with the ability to set a timeout on a request.
func NewLunMapReportingNodeDeleteParamsWithTimeout(timeout time.Duration) *LunMapReportingNodeDeleteParams {
	return &LunMapReportingNodeDeleteParams{
		timeout: timeout,
	}
}

// NewLunMapReportingNodeDeleteParamsWithContext creates a new LunMapReportingNodeDeleteParams object
// with the ability to set a context for a request.
func NewLunMapReportingNodeDeleteParamsWithContext(ctx context.Context) *LunMapReportingNodeDeleteParams {
	return &LunMapReportingNodeDeleteParams{
		Context: ctx,
	}
}

// NewLunMapReportingNodeDeleteParamsWithHTTPClient creates a new LunMapReportingNodeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunMapReportingNodeDeleteParamsWithHTTPClient(client *http.Client) *LunMapReportingNodeDeleteParams {
	return &LunMapReportingNodeDeleteParams{
		HTTPClient: client,
	}
}

/* LunMapReportingNodeDeleteParams contains all the parameters to send to the API endpoint
   for the lun map reporting node delete operation.

   Typically these are written to a http.Request.
*/
type LunMapReportingNodeDeleteParams struct {

	/* IgroupUUID.

	   The unique identifier of the igroup.

	*/
	IgroupUUIDPathParameter string

	/* LunUUID.

	   The unique identifier of the LUN.

	*/
	LunUUIDPathParameter string

	/* UUID.

	   The unique identifier of the node.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun map reporting node delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunMapReportingNodeDeleteParams) WithDefaults() *LunMapReportingNodeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun map reporting node delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunMapReportingNodeDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) WithTimeout(timeout time.Duration) *LunMapReportingNodeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) WithContext(ctx context.Context) *LunMapReportingNodeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) WithHTTPClient(client *http.Client) *LunMapReportingNodeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIgroupUUIDPathParameter adds the igroupUUID to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) WithIgroupUUIDPathParameter(igroupUUID string) *LunMapReportingNodeDeleteParams {
	o.SetIgroupUUIDPathParameter(igroupUUID)
	return o
}

// SetIgroupUUIDPathParameter adds the igroupUuid to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) SetIgroupUUIDPathParameter(igroupUUID string) {
	o.IgroupUUIDPathParameter = igroupUUID
}

// WithLunUUIDPathParameter adds the lunUUID to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) WithLunUUIDPathParameter(lunUUID string) *LunMapReportingNodeDeleteParams {
	o.SetLunUUIDPathParameter(lunUUID)
	return o
}

// SetLunUUIDPathParameter adds the lunUuid to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) SetLunUUIDPathParameter(lunUUID string) {
	o.LunUUIDPathParameter = lunUUID
}

// WithUUIDPathParameter adds the uuid to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) WithUUIDPathParameter(uuid string) *LunMapReportingNodeDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the lun map reporting node delete params
func (o *LunMapReportingNodeDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *LunMapReportingNodeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param igroup.uuid
	if err := r.SetPathParam("igroup.uuid", o.IgroupUUIDPathParameter); err != nil {
		return err
	}

	// path param lun.uuid
	if err := r.SetPathParam("lun.uuid", o.LunUUIDPathParameter); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
