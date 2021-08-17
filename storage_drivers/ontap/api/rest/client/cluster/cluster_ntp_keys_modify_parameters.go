// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewClusterNtpKeysModifyParams creates a new ClusterNtpKeysModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterNtpKeysModifyParams() *ClusterNtpKeysModifyParams {
	return &ClusterNtpKeysModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterNtpKeysModifyParamsWithTimeout creates a new ClusterNtpKeysModifyParams object
// with the ability to set a timeout on a request.
func NewClusterNtpKeysModifyParamsWithTimeout(timeout time.Duration) *ClusterNtpKeysModifyParams {
	return &ClusterNtpKeysModifyParams{
		timeout: timeout,
	}
}

// NewClusterNtpKeysModifyParamsWithContext creates a new ClusterNtpKeysModifyParams object
// with the ability to set a context for a request.
func NewClusterNtpKeysModifyParamsWithContext(ctx context.Context) *ClusterNtpKeysModifyParams {
	return &ClusterNtpKeysModifyParams{
		Context: ctx,
	}
}

// NewClusterNtpKeysModifyParamsWithHTTPClient creates a new ClusterNtpKeysModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterNtpKeysModifyParamsWithHTTPClient(client *http.Client) *ClusterNtpKeysModifyParams {
	return &ClusterNtpKeysModifyParams{
		HTTPClient: client,
	}
}

/* ClusterNtpKeysModifyParams contains all the parameters to send to the API endpoint
   for the cluster ntp keys modify operation.

   Typically these are written to a http.Request.
*/
type ClusterNtpKeysModifyParams struct {

	/* ID.

	   Key identifier
	*/
	IDPathParameter int64

	/* Info.

	   Information specification
	*/
	Info *models.NtpKey

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster ntp keys modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNtpKeysModifyParams) WithDefaults() *ClusterNtpKeysModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster ntp keys modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNtpKeysModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) WithTimeout(timeout time.Duration) *ClusterNtpKeysModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) WithContext(ctx context.Context) *ClusterNtpKeysModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) WithHTTPClient(client *http.Client) *ClusterNtpKeysModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDPathParameter adds the id to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) WithIDPathParameter(id int64) *ClusterNtpKeysModifyParams {
	o.SetIDPathParameter(id)
	return o
}

// SetIDPathParameter adds the id to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) SetIDPathParameter(id int64) {
	o.IDPathParameter = id
}

// WithInfo adds the info to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) WithInfo(info *models.NtpKey) *ClusterNtpKeysModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster ntp keys modify params
func (o *ClusterNtpKeysModifyParams) SetInfo(info *models.NtpKey) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterNtpKeysModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.IDPathParameter)); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
