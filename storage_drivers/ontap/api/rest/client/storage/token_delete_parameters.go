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
)

// NewTokenDeleteParams creates a new TokenDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTokenDeleteParams() *TokenDeleteParams {
	return &TokenDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTokenDeleteParamsWithTimeout creates a new TokenDeleteParams object
// with the ability to set a timeout on a request.
func NewTokenDeleteParamsWithTimeout(timeout time.Duration) *TokenDeleteParams {
	return &TokenDeleteParams{
		timeout: timeout,
	}
}

// NewTokenDeleteParamsWithContext creates a new TokenDeleteParams object
// with the ability to set a context for a request.
func NewTokenDeleteParamsWithContext(ctx context.Context) *TokenDeleteParams {
	return &TokenDeleteParams{
		Context: ctx,
	}
}

// NewTokenDeleteParamsWithHTTPClient creates a new TokenDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTokenDeleteParamsWithHTTPClient(client *http.Client) *TokenDeleteParams {
	return &TokenDeleteParams{
		HTTPClient: client,
	}
}

/* TokenDeleteParams contains all the parameters to send to the API endpoint
   for the token delete operation.

   Typically these are written to a http.Request.
*/
type TokenDeleteParams struct {

	/* NodeUUID.

	   Node UUID
	*/
	NodeUUIDPathParameter string

	/* UUID.

	   Token UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the token delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenDeleteParams) WithDefaults() *TokenDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the token delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TokenDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the token delete params
func (o *TokenDeleteParams) WithTimeout(timeout time.Duration) *TokenDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the token delete params
func (o *TokenDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the token delete params
func (o *TokenDeleteParams) WithContext(ctx context.Context) *TokenDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the token delete params
func (o *TokenDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the token delete params
func (o *TokenDeleteParams) WithHTTPClient(client *http.Client) *TokenDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the token delete params
func (o *TokenDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeUUIDPathParameter adds the nodeUUID to the token delete params
func (o *TokenDeleteParams) WithNodeUUIDPathParameter(nodeUUID string) *TokenDeleteParams {
	o.SetNodeUUIDPathParameter(nodeUUID)
	return o
}

// SetNodeUUIDPathParameter adds the nodeUuid to the token delete params
func (o *TokenDeleteParams) SetNodeUUIDPathParameter(nodeUUID string) {
	o.NodeUUIDPathParameter = nodeUUID
}

// WithUUIDPathParameter adds the uuid to the token delete params
func (o *TokenDeleteParams) WithUUIDPathParameter(uuid string) *TokenDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the token delete params
func (o *TokenDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *TokenDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param node.uuid
	if err := r.SetPathParam("node.uuid", o.NodeUUIDPathParameter); err != nil {
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
