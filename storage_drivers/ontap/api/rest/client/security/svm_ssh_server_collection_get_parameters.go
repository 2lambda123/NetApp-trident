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
	"github.com/go-openapi/swag"
)

// NewSvmSSHServerCollectionGetParams creates a new SvmSSHServerCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSvmSSHServerCollectionGetParams() *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSvmSSHServerCollectionGetParamsWithTimeout creates a new SvmSSHServerCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSvmSSHServerCollectionGetParamsWithTimeout(timeout time.Duration) *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		timeout: timeout,
	}
}

// NewSvmSSHServerCollectionGetParamsWithContext creates a new SvmSSHServerCollectionGetParams object
// with the ability to set a context for a request.
func NewSvmSSHServerCollectionGetParamsWithContext(ctx context.Context) *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		Context: ctx,
	}
}

// NewSvmSSHServerCollectionGetParamsWithHTTPClient creates a new SvmSSHServerCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSvmSSHServerCollectionGetParamsWithHTTPClient(client *http.Client) *SvmSSHServerCollectionGetParams {
	return &SvmSSHServerCollectionGetParams{
		HTTPClient: client,
	}
}

/* SvmSSHServerCollectionGetParams contains all the parameters to send to the API endpoint
   for the svm ssh server collection get operation.

   Typically these are written to a http.Request.
*/
type SvmSSHServerCollectionGetParams struct {

	/* Ciphers.

	   Filter by ciphers
	*/
	CiphersQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* KeyExchangeAlgorithms.

	   Filter by key_exchange_algorithms
	*/
	KeyExchangeAlgorithmsQueryParameter *string

	/* MacAlgorithms.

	   Filter by mac_algorithms
	*/
	MacAlgorithmsQueryParameter *string

	/* MaxAuthenticationRetryCount.

	   Filter by max_authentication_retry_count
	*/
	MaxAuthenticationRetryCountQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

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

// WithDefaults hydrates default values in the svm ssh server collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerCollectionGetParams) WithDefaults() *SvmSSHServerCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the svm ssh server collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SvmSSHServerCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SvmSSHServerCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithTimeout(timeout time.Duration) *SvmSSHServerCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithContext(ctx context.Context) *SvmSSHServerCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithHTTPClient(client *http.Client) *SvmSSHServerCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCiphersQueryParameter adds the ciphers to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithCiphersQueryParameter(ciphers *string) *SvmSSHServerCollectionGetParams {
	o.SetCiphersQueryParameter(ciphers)
	return o
}

// SetCiphersQueryParameter adds the ciphers to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetCiphersQueryParameter(ciphers *string) {
	o.CiphersQueryParameter = ciphers
}

// WithFieldsQueryParameter adds the fields to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithFieldsQueryParameter(fields []string) *SvmSSHServerCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithKeyExchangeAlgorithmsQueryParameter adds the keyExchangeAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithKeyExchangeAlgorithmsQueryParameter(keyExchangeAlgorithms *string) *SvmSSHServerCollectionGetParams {
	o.SetKeyExchangeAlgorithmsQueryParameter(keyExchangeAlgorithms)
	return o
}

// SetKeyExchangeAlgorithmsQueryParameter adds the keyExchangeAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetKeyExchangeAlgorithmsQueryParameter(keyExchangeAlgorithms *string) {
	o.KeyExchangeAlgorithmsQueryParameter = keyExchangeAlgorithms
}

// WithMacAlgorithmsQueryParameter adds the macAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithMacAlgorithmsQueryParameter(macAlgorithms *string) *SvmSSHServerCollectionGetParams {
	o.SetMacAlgorithmsQueryParameter(macAlgorithms)
	return o
}

// SetMacAlgorithmsQueryParameter adds the macAlgorithms to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetMacAlgorithmsQueryParameter(macAlgorithms *string) {
	o.MacAlgorithmsQueryParameter = macAlgorithms
}

// WithMaxAuthenticationRetryCountQueryParameter adds the maxAuthenticationRetryCount to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithMaxAuthenticationRetryCountQueryParameter(maxAuthenticationRetryCount *int64) *SvmSSHServerCollectionGetParams {
	o.SetMaxAuthenticationRetryCountQueryParameter(maxAuthenticationRetryCount)
	return o
}

// SetMaxAuthenticationRetryCountQueryParameter adds the maxAuthenticationRetryCount to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetMaxAuthenticationRetryCountQueryParameter(maxAuthenticationRetryCount *int64) {
	o.MaxAuthenticationRetryCountQueryParameter = maxAuthenticationRetryCount
}

// WithMaxRecordsQueryParameter adds the maxRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SvmSSHServerCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SvmSSHServerCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SvmSSHServerCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SvmSSHServerCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *SvmSSHServerCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SvmSSHServerCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the svm ssh server collection get params
func (o *SvmSSHServerCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SvmSSHServerCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CiphersQueryParameter != nil {

		// query param ciphers
		var qrCiphers string

		if o.CiphersQueryParameter != nil {
			qrCiphers = *o.CiphersQueryParameter
		}
		qCiphers := qrCiphers
		if qCiphers != "" {

			if err := r.SetQueryParam("ciphers", qCiphers); err != nil {
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

	if o.KeyExchangeAlgorithmsQueryParameter != nil {

		// query param key_exchange_algorithms
		var qrKeyExchangeAlgorithms string

		if o.KeyExchangeAlgorithmsQueryParameter != nil {
			qrKeyExchangeAlgorithms = *o.KeyExchangeAlgorithmsQueryParameter
		}
		qKeyExchangeAlgorithms := qrKeyExchangeAlgorithms
		if qKeyExchangeAlgorithms != "" {

			if err := r.SetQueryParam("key_exchange_algorithms", qKeyExchangeAlgorithms); err != nil {
				return err
			}
		}
	}

	if o.MacAlgorithmsQueryParameter != nil {

		// query param mac_algorithms
		var qrMacAlgorithms string

		if o.MacAlgorithmsQueryParameter != nil {
			qrMacAlgorithms = *o.MacAlgorithmsQueryParameter
		}
		qMacAlgorithms := qrMacAlgorithms
		if qMacAlgorithms != "" {

			if err := r.SetQueryParam("mac_algorithms", qMacAlgorithms); err != nil {
				return err
			}
		}
	}

	if o.MaxAuthenticationRetryCountQueryParameter != nil {

		// query param max_authentication_retry_count
		var qrMaxAuthenticationRetryCount int64

		if o.MaxAuthenticationRetryCountQueryParameter != nil {
			qrMaxAuthenticationRetryCount = *o.MaxAuthenticationRetryCountQueryParameter
		}
		qMaxAuthenticationRetryCount := swag.FormatInt64(qrMaxAuthenticationRetryCount)
		if qMaxAuthenticationRetryCount != "" {

			if err := r.SetQueryParam("max_authentication_retry_count", qMaxAuthenticationRetryCount); err != nil {
				return err
			}
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

// bindParamSvmSSHServerCollectionGet binds the parameter fields
func (o *SvmSSHServerCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSvmSSHServerCollectionGet binds the parameter order_by
func (o *SvmSSHServerCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
