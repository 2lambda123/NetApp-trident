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
)

// NewLocalCifsGroupCollectionGetParams creates a new LocalCifsGroupCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLocalCifsGroupCollectionGetParams() *LocalCifsGroupCollectionGetParams {
	return &LocalCifsGroupCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLocalCifsGroupCollectionGetParamsWithTimeout creates a new LocalCifsGroupCollectionGetParams object
// with the ability to set a timeout on a request.
func NewLocalCifsGroupCollectionGetParamsWithTimeout(timeout time.Duration) *LocalCifsGroupCollectionGetParams {
	return &LocalCifsGroupCollectionGetParams{
		timeout: timeout,
	}
}

// NewLocalCifsGroupCollectionGetParamsWithContext creates a new LocalCifsGroupCollectionGetParams object
// with the ability to set a context for a request.
func NewLocalCifsGroupCollectionGetParamsWithContext(ctx context.Context) *LocalCifsGroupCollectionGetParams {
	return &LocalCifsGroupCollectionGetParams{
		Context: ctx,
	}
}

// NewLocalCifsGroupCollectionGetParamsWithHTTPClient creates a new LocalCifsGroupCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLocalCifsGroupCollectionGetParamsWithHTTPClient(client *http.Client) *LocalCifsGroupCollectionGetParams {
	return &LocalCifsGroupCollectionGetParams{
		HTTPClient: client,
	}
}

/* LocalCifsGroupCollectionGetParams contains all the parameters to send to the API endpoint
   for the local cifs group collection get operation.

   Typically these are written to a http.Request.
*/
type LocalCifsGroupCollectionGetParams struct {

	/* Description.

	   Filter by description
	*/
	DescriptionQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* MembersName.

	   Filter by members.name
	*/
	MembersNameQueryParameter *string

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* Sid.

	   Filter by sid
	*/
	SIDQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the local cifs group collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupCollectionGetParams) WithDefaults() *LocalCifsGroupCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the local cifs group collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LocalCifsGroupCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := LocalCifsGroupCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithTimeout(timeout time.Duration) *LocalCifsGroupCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithContext(ctx context.Context) *LocalCifsGroupCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithHTTPClient(client *http.Client) *LocalCifsGroupCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescriptionQueryParameter adds the description to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithDescriptionQueryParameter(description *string) *LocalCifsGroupCollectionGetParams {
	o.SetDescriptionQueryParameter(description)
	return o
}

// SetDescriptionQueryParameter adds the description to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetDescriptionQueryParameter(description *string) {
	o.DescriptionQueryParameter = description
}

// WithFieldsQueryParameter adds the fields to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithFieldsQueryParameter(fields []string) *LocalCifsGroupCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *LocalCifsGroupCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithMembersNameQueryParameter adds the membersName to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithMembersNameQueryParameter(membersName *string) *LocalCifsGroupCollectionGetParams {
	o.SetMembersNameQueryParameter(membersName)
	return o
}

// SetMembersNameQueryParameter adds the membersName to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetMembersNameQueryParameter(membersName *string) {
	o.MembersNameQueryParameter = membersName
}

// WithNameQueryParameter adds the name to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithNameQueryParameter(name *string) *LocalCifsGroupCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *LocalCifsGroupCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *LocalCifsGroupCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *LocalCifsGroupCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSIDQueryParameter adds the sid to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithSIDQueryParameter(sid *string) *LocalCifsGroupCollectionGetParams {
	o.SetSIDQueryParameter(sid)
	return o
}

// SetSIDQueryParameter adds the sid to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetSIDQueryParameter(sid *string) {
	o.SIDQueryParameter = sid
}

// WithSVMNameQueryParameter adds the svmName to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *LocalCifsGroupCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *LocalCifsGroupCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the local cifs group collection get params
func (o *LocalCifsGroupCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LocalCifsGroupCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DescriptionQueryParameter != nil {

		// query param description
		var qrDescription string

		if o.DescriptionQueryParameter != nil {
			qrDescription = *o.DescriptionQueryParameter
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.MembersNameQueryParameter != nil {

		// query param members.name
		var qrMembersName string

		if o.MembersNameQueryParameter != nil {
			qrMembersName = *o.MembersNameQueryParameter
		}
		qMembersName := qrMembersName
		if qMembersName != "" {

			if err := r.SetQueryParam("members.name", qMembersName); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
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

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SIDQueryParameter != nil {

		// query param sid
		var qrSid string

		if o.SIDQueryParameter != nil {
			qrSid = *o.SIDQueryParameter
		}
		qSid := qrSid
		if qSid != "" {

			if err := r.SetQueryParam("sid", qSid); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLocalCifsGroupCollectionGet binds the parameter fields
func (o *LocalCifsGroupCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamLocalCifsGroupCollectionGet binds the parameter order_by
func (o *LocalCifsGroupCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
