// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewExportRuleModifyParams creates a new ExportRuleModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleModifyParams() *ExportRuleModifyParams {
	return &ExportRuleModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleModifyParamsWithTimeout creates a new ExportRuleModifyParams object
// with the ability to set a timeout on a request.
func NewExportRuleModifyParamsWithTimeout(timeout time.Duration) *ExportRuleModifyParams {
	return &ExportRuleModifyParams{
		timeout: timeout,
	}
}

// NewExportRuleModifyParamsWithContext creates a new ExportRuleModifyParams object
// with the ability to set a context for a request.
func NewExportRuleModifyParamsWithContext(ctx context.Context) *ExportRuleModifyParams {
	return &ExportRuleModifyParams{
		Context: ctx,
	}
}

// NewExportRuleModifyParamsWithHTTPClient creates a new ExportRuleModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleModifyParamsWithHTTPClient(client *http.Client) *ExportRuleModifyParams {
	return &ExportRuleModifyParams{
		HTTPClient: client,
	}
}

/* ExportRuleModifyParams contains all the parameters to send to the API endpoint
   for the export rule modify operation.

   Typically these are written to a http.Request.
*/
type ExportRuleModifyParams struct {

	/* Index.

	   Export Rule Index
	*/
	IndexPathParameter int64

	/* Info.

	   Info specification
	*/
	Info *models.ExportRule

	/* NewIndex.

	   New Export Rule Index
	*/
	NewIndexQueryParameter *int64

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyIDPathParameter int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export rule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleModifyParams) WithDefaults() *ExportRuleModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the export rule modify params
func (o *ExportRuleModifyParams) WithTimeout(timeout time.Duration) *ExportRuleModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule modify params
func (o *ExportRuleModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule modify params
func (o *ExportRuleModifyParams) WithContext(ctx context.Context) *ExportRuleModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule modify params
func (o *ExportRuleModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule modify params
func (o *ExportRuleModifyParams) WithHTTPClient(client *http.Client) *ExportRuleModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule modify params
func (o *ExportRuleModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexPathParameter adds the index to the export rule modify params
func (o *ExportRuleModifyParams) WithIndexPathParameter(index int64) *ExportRuleModifyParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the export rule modify params
func (o *ExportRuleModifyParams) SetIndexPathParameter(index int64) {
	o.IndexPathParameter = index
}

// WithInfo adds the info to the export rule modify params
func (o *ExportRuleModifyParams) WithInfo(info *models.ExportRule) *ExportRuleModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the export rule modify params
func (o *ExportRuleModifyParams) SetInfo(info *models.ExportRule) {
	o.Info = info
}

// WithNewIndexQueryParameter adds the newIndex to the export rule modify params
func (o *ExportRuleModifyParams) WithNewIndexQueryParameter(newIndex *int64) *ExportRuleModifyParams {
	o.SetNewIndexQueryParameter(newIndex)
	return o
}

// SetNewIndexQueryParameter adds the newIndex to the export rule modify params
func (o *ExportRuleModifyParams) SetNewIndexQueryParameter(newIndex *int64) {
	o.NewIndexQueryParameter = newIndex
}

// WithPolicyIDPathParameter adds the policyID to the export rule modify params
func (o *ExportRuleModifyParams) WithPolicyIDPathParameter(policyID int64) *ExportRuleModifyParams {
	o.SetPolicyIDPathParameter(policyID)
	return o
}

// SetPolicyIDPathParameter adds the policyId to the export rule modify params
func (o *ExportRuleModifyParams) SetPolicyIDPathParameter(policyID int64) {
	o.PolicyIDPathParameter = policyID
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.IndexPathParameter)); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.NewIndexQueryParameter != nil {

		// query param new_index
		var qrNewIndex int64

		if o.NewIndexQueryParameter != nil {
			qrNewIndex = *o.NewIndexQueryParameter
		}
		qNewIndex := swag.FormatInt64(qrNewIndex)
		if qNewIndex != "" {

			if err := r.SetQueryParam("new_index", qNewIndex); err != nil {
				return err
			}
		}
	}

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyIDPathParameter)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
