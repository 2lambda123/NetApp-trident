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

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NewAccountModifyParams creates a new AccountModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountModifyParams() *AccountModifyParams {
	return &AccountModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountModifyParamsWithTimeout creates a new AccountModifyParams object
// with the ability to set a timeout on a request.
func NewAccountModifyParamsWithTimeout(timeout time.Duration) *AccountModifyParams {
	return &AccountModifyParams{
		timeout: timeout,
	}
}

// NewAccountModifyParamsWithContext creates a new AccountModifyParams object
// with the ability to set a context for a request.
func NewAccountModifyParamsWithContext(ctx context.Context) *AccountModifyParams {
	return &AccountModifyParams{
		Context: ctx,
	}
}

// NewAccountModifyParamsWithHTTPClient creates a new AccountModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountModifyParamsWithHTTPClient(client *http.Client) *AccountModifyParams {
	return &AccountModifyParams{
		HTTPClient: client,
	}
}

/* AccountModifyParams contains all the parameters to send to the API endpoint
   for the account modify operation.

   Typically these are written to a http.Request.
*/
type AccountModifyParams struct {

	/* Info.

	   User account details
	*/
	Info *models.Account

	/* Name.

	   User account name
	*/
	NamePathParameter string

	/* OwnerUUID.

	   Account owner UUID
	*/
	OwnerUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountModifyParams) WithDefaults() *AccountModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the account modify params
func (o *AccountModifyParams) WithTimeout(timeout time.Duration) *AccountModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account modify params
func (o *AccountModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account modify params
func (o *AccountModifyParams) WithContext(ctx context.Context) *AccountModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account modify params
func (o *AccountModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account modify params
func (o *AccountModifyParams) WithHTTPClient(client *http.Client) *AccountModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account modify params
func (o *AccountModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the account modify params
func (o *AccountModifyParams) WithInfo(info *models.Account) *AccountModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the account modify params
func (o *AccountModifyParams) SetInfo(info *models.Account) {
	o.Info = info
}

// WithNamePathParameter adds the name to the account modify params
func (o *AccountModifyParams) WithNamePathParameter(name string) *AccountModifyParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the account modify params
func (o *AccountModifyParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithOwnerUUIDPathParameter adds the ownerUUID to the account modify params
func (o *AccountModifyParams) WithOwnerUUIDPathParameter(ownerUUID string) *AccountModifyParams {
	o.SetOwnerUUIDPathParameter(ownerUUID)
	return o
}

// SetOwnerUUIDPathParameter adds the ownerUuid to the account modify params
func (o *AccountModifyParams) SetOwnerUUIDPathParameter(ownerUUID string) {
	o.OwnerUUIDPathParameter = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AccountModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
