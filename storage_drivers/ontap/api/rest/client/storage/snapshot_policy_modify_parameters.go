// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NewSnapshotPolicyModifyParams creates a new SnapshotPolicyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotPolicyModifyParams() *SnapshotPolicyModifyParams {
	return &SnapshotPolicyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotPolicyModifyParamsWithTimeout creates a new SnapshotPolicyModifyParams object
// with the ability to set a timeout on a request.
func NewSnapshotPolicyModifyParamsWithTimeout(timeout time.Duration) *SnapshotPolicyModifyParams {
	return &SnapshotPolicyModifyParams{
		timeout: timeout,
	}
}

// NewSnapshotPolicyModifyParamsWithContext creates a new SnapshotPolicyModifyParams object
// with the ability to set a context for a request.
func NewSnapshotPolicyModifyParamsWithContext(ctx context.Context) *SnapshotPolicyModifyParams {
	return &SnapshotPolicyModifyParams{
		Context: ctx,
	}
}

// NewSnapshotPolicyModifyParamsWithHTTPClient creates a new SnapshotPolicyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotPolicyModifyParamsWithHTTPClient(client *http.Client) *SnapshotPolicyModifyParams {
	return &SnapshotPolicyModifyParams{
		HTTPClient: client,
	}
}

/* SnapshotPolicyModifyParams contains all the parameters to send to the API endpoint
   for the snapshot policy modify operation.

   Typically these are written to a http.Request.
*/
type SnapshotPolicyModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.SnapshotPolicy

	/* UUID.

	   Snapshot copy policy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyModifyParams) WithDefaults() *SnapshotPolicyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotPolicyModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) WithTimeout(timeout time.Duration) *SnapshotPolicyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) WithContext(ctx context.Context) *SnapshotPolicyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) WithHTTPClient(client *http.Client) *SnapshotPolicyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) WithInfo(info *models.SnapshotPolicy) *SnapshotPolicyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) SetInfo(info *models.SnapshotPolicy) {
	o.Info = info
}

// WithUUIDPathParameter adds the uuid to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) WithUUIDPathParameter(uuid string) *SnapshotPolicyModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the snapshot policy modify params
func (o *SnapshotPolicyModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotPolicyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
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
