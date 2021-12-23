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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewLocalCifsGroupMembersBulkDeleteParams creates a new LocalCifsGroupMembersBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalCifsGroupMembersBulkDeleteParams() *LocalCifsGroupMembersBulkDeleteParams {
	return &LocalCifsGroupMembersBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalCifsGroupMembersBulkDeleteParamsWithTimeout creates a new LocalCifsGroupMembersBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewLocalCifsGroupMembersBulkDeleteParamsWithTimeout(timeout time.Duration) *LocalCifsGroupMembersBulkDeleteParams {
	return &LocalCifsGroupMembersBulkDeleteParams{
		timeout: timeout,
	}
}

// NewLocalCifsGroupMembersBulkDeleteParamsWithContext creates a new LocalCifsGroupMembersBulkDeleteParams object
// with the ability to set a context for a request.
func NewLocalCifsGroupMembersBulkDeleteParamsWithContext(ctx context.Context) *LocalCifsGroupMembersBulkDeleteParams {
	return &LocalCifsGroupMembersBulkDeleteParams{
		Context: ctx,
	}
}

// NewLocalCifsGroupMembersBulkDeleteParamsWithHTTPClient creates a new LocalCifsGroupMembersBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalCifsGroupMembersBulkDeleteParamsWithHTTPClient(client *http.Client) *LocalCifsGroupMembersBulkDeleteParams {
	return &LocalCifsGroupMembersBulkDeleteParams{
		HTTPClient: client,
	}
}

/* LocalCifsGroupMembersBulkDeleteParams contains all the parameters to send to the API endpoint
   for the local cifs group members bulk delete operation.

   Typically these are written to a http.Request.
*/
type LocalCifsGroupMembersBulkDeleteParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.LocalCifsGroupMembers

	/* LocalCifsGroupSid.

	   Local group SID
	*/
	LocalCifsGroupSIDPathParameter string

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local cifs group members bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupMembersBulkDeleteParams) WithDefaults() *LocalCifsGroupMembersBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local cifs group members bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupMembersBulkDeleteParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := LocalCifsGroupMembersBulkDeleteParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) WithTimeout(timeout time.Duration) *LocalCifsGroupMembersBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) WithContext(ctx context.Context) *LocalCifsGroupMembersBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) WithHTTPClient(client *http.Client) *LocalCifsGroupMembersBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) WithInfo(info *models.LocalCifsGroupMembers) *LocalCifsGroupMembersBulkDeleteParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) SetInfo(info *models.LocalCifsGroupMembers) {
	o.Info = info
}

// WithLocalCifsGroupSIDPathParameter adds the localCifsGroupSid to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) WithLocalCifsGroupSIDPathParameter(localCifsGroupSid string) *LocalCifsGroupMembersBulkDeleteParams {
	o.SetLocalCifsGroupSIDPathParameter(localCifsGroupSid)
	return o
}

// SetLocalCifsGroupSIDPathParameter adds the localCifsGroupSid to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) SetLocalCifsGroupSIDPathParameter(localCifsGroupSid string) {
	o.LocalCifsGroupSIDPathParameter = localCifsGroupSid
}

// WithReturnRecordsQueryParameter adds the returnRecords to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) WithReturnRecordsQueryParameter(returnRecords *bool) *LocalCifsGroupMembersBulkDeleteParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithSVMUUIDPathParameter adds the svmUUID to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *LocalCifsGroupMembersBulkDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the local cifs group members bulk delete params
func (o *LocalCifsGroupMembersBulkDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LocalCifsGroupMembersBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param local_cifs_group.sid
	if err := r.SetPathParam("local_cifs_group.sid", o.LocalCifsGroupSIDPathParameter); err != nil {
		return err
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
