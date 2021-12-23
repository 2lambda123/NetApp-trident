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
	"github.com/go-openapi/swag"
)

// NewSnapshotCollectionGetParams creates a new SnapshotCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotCollectionGetParams() *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotCollectionGetParamsWithTimeout creates a new SnapshotCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnapshotCollectionGetParamsWithTimeout(timeout time.Duration) *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnapshotCollectionGetParamsWithContext creates a new SnapshotCollectionGetParams object
// with the ability to set a context for a request.
func NewSnapshotCollectionGetParamsWithContext(ctx context.Context) *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		Context: ctx,
	}
}

// NewSnapshotCollectionGetParamsWithHTTPClient creates a new SnapshotCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotCollectionGetParamsWithHTTPClient(client *http.Client) *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		HTTPClient: client,
	}
}

/* SnapshotCollectionGetParams contains all the parameters to send to the API endpoint
   for the snapshot collection get operation.

   Typically these are written to a http.Request.
*/
type SnapshotCollectionGetParams struct {

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTimeQueryParameter *string

	/* ExpiryTime.

	   Filter by expiry_time
	*/
	ExpiryTimeQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* Owners.

	   Filter by owners
	*/
	OwnersQueryParameter *string

	/* ReclaimableSpace.

	   Filter by reclaimable_space
	*/
	ReclaimableSpaceQueryParameter *int64

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

	/* Size.

	   Filter by size
	*/
	SizeQueryParameter *int64

	/* SnaplockExpiryTime.

	   Filter by snaplock_expiry_time
	*/
	SnaplockExpiryTimeQueryParameter *string

	/* SnapmirrorLabel.

	   Filter by snapmirror_label
	*/
	SnapmirrorLabelQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeNameQueryParameter *string

	/* VolumeUUID.

	   Volume
	*/
	VolumeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotCollectionGetParams) WithDefaults() *SnapshotCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SnapshotCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithTimeout(timeout time.Duration) *SnapshotCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithContext(ctx context.Context) *SnapshotCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithHTTPClient(client *http.Client) *SnapshotCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentQueryParameter adds the comment to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithCommentQueryParameter(comment *string) *SnapshotCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithCreateTimeQueryParameter adds the createTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithCreateTimeQueryParameter(createTime *string) *SnapshotCollectionGetParams {
	o.SetCreateTimeQueryParameter(createTime)
	return o
}

// SetCreateTimeQueryParameter adds the createTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetCreateTimeQueryParameter(createTime *string) {
	o.CreateTimeQueryParameter = createTime
}

// WithExpiryTimeQueryParameter adds the expiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithExpiryTimeQueryParameter(expiryTime *string) *SnapshotCollectionGetParams {
	o.SetExpiryTimeQueryParameter(expiryTime)
	return o
}

// SetExpiryTimeQueryParameter adds the expiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetExpiryTimeQueryParameter(expiryTime *string) {
	o.ExpiryTimeQueryParameter = expiryTime
}

// WithFieldsQueryParameter adds the fields to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithFieldsQueryParameter(fields []string) *SnapshotCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SnapshotCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithNameQueryParameter(name *string) *SnapshotCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SnapshotCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithOwnersQueryParameter adds the owners to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithOwnersQueryParameter(owners *string) *SnapshotCollectionGetParams {
	o.SetOwnersQueryParameter(owners)
	return o
}

// SetOwnersQueryParameter adds the owners to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetOwnersQueryParameter(owners *string) {
	o.OwnersQueryParameter = owners
}

// WithReclaimableSpaceQueryParameter adds the reclaimableSpace to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithReclaimableSpaceQueryParameter(reclaimableSpace *int64) *SnapshotCollectionGetParams {
	o.SetReclaimableSpaceQueryParameter(reclaimableSpace)
	return o
}

// SetReclaimableSpaceQueryParameter adds the reclaimableSpace to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetReclaimableSpaceQueryParameter(reclaimableSpace *int64) {
	o.ReclaimableSpaceQueryParameter = reclaimableSpace
}

// WithReturnRecordsQueryParameter adds the returnRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SnapshotCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SnapshotCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSizeQueryParameter adds the size to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSizeQueryParameter(size *int64) *SnapshotCollectionGetParams {
	o.SetSizeQueryParameter(size)
	return o
}

// SetSizeQueryParameter adds the size to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSizeQueryParameter(size *int64) {
	o.SizeQueryParameter = size
}

// WithSnaplockExpiryTimeQueryParameter adds the snaplockExpiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSnaplockExpiryTimeQueryParameter(snaplockExpiryTime *string) *SnapshotCollectionGetParams {
	o.SetSnaplockExpiryTimeQueryParameter(snaplockExpiryTime)
	return o
}

// SetSnaplockExpiryTimeQueryParameter adds the snaplockExpiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSnaplockExpiryTimeQueryParameter(snaplockExpiryTime *string) {
	o.SnaplockExpiryTimeQueryParameter = snaplockExpiryTime
}

// WithSnapmirrorLabelQueryParameter adds the snapmirrorLabel to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSnapmirrorLabelQueryParameter(snapmirrorLabel *string) *SnapshotCollectionGetParams {
	o.SetSnapmirrorLabelQueryParameter(snapmirrorLabel)
	return o
}

// SetSnapmirrorLabelQueryParameter adds the snapmirrorLabel to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSnapmirrorLabelQueryParameter(snapmirrorLabel *string) {
	o.SnapmirrorLabelQueryParameter = snapmirrorLabel
}

// WithStateQueryParameter adds the state to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithStateQueryParameter(state *string) *SnapshotCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithSVMNameQueryParameter adds the svmName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *SnapshotCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SnapshotCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithUUIDQueryParameter(uuid *string) *SnapshotCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WithVolumeNameQueryParameter adds the volumeName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithVolumeNameQueryParameter(volumeName *string) *SnapshotCollectionGetParams {
	o.SetVolumeNameQueryParameter(volumeName)
	return o
}

// SetVolumeNameQueryParameter adds the volumeName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetVolumeNameQueryParameter(volumeName *string) {
	o.VolumeNameQueryParameter = volumeName
}

// WithVolumeUUIDPathParameter adds the volumeUUID to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithVolumeUUIDPathParameter(volumeUUID string) *SnapshotCollectionGetParams {
	o.SetVolumeUUIDPathParameter(volumeUUID)
	return o
}

// SetVolumeUUIDPathParameter adds the volumeUuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetVolumeUUIDPathParameter(volumeUUID string) {
	o.VolumeUUIDPathParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.CreateTimeQueryParameter != nil {

		// query param create_time
		var qrCreateTime string

		if o.CreateTimeQueryParameter != nil {
			qrCreateTime = *o.CreateTimeQueryParameter
		}
		qCreateTime := qrCreateTime
		if qCreateTime != "" {

			if err := r.SetQueryParam("create_time", qCreateTime); err != nil {
				return err
			}
		}
	}

	if o.ExpiryTimeQueryParameter != nil {

		// query param expiry_time
		var qrExpiryTime string

		if o.ExpiryTimeQueryParameter != nil {
			qrExpiryTime = *o.ExpiryTimeQueryParameter
		}
		qExpiryTime := qrExpiryTime
		if qExpiryTime != "" {

			if err := r.SetQueryParam("expiry_time", qExpiryTime); err != nil {
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

	if o.OwnersQueryParameter != nil {

		// query param owners
		var qrOwners string

		if o.OwnersQueryParameter != nil {
			qrOwners = *o.OwnersQueryParameter
		}
		qOwners := qrOwners
		if qOwners != "" {

			if err := r.SetQueryParam("owners", qOwners); err != nil {
				return err
			}
		}
	}

	if o.ReclaimableSpaceQueryParameter != nil {

		// query param reclaimable_space
		var qrReclaimableSpace int64

		if o.ReclaimableSpaceQueryParameter != nil {
			qrReclaimableSpace = *o.ReclaimableSpaceQueryParameter
		}
		qReclaimableSpace := swag.FormatInt64(qrReclaimableSpace)
		if qReclaimableSpace != "" {

			if err := r.SetQueryParam("reclaimable_space", qReclaimableSpace); err != nil {
				return err
			}
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

	if o.SizeQueryParameter != nil {

		// query param size
		var qrSize int64

		if o.SizeQueryParameter != nil {
			qrSize = *o.SizeQueryParameter
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.SnaplockExpiryTimeQueryParameter != nil {

		// query param snaplock_expiry_time
		var qrSnaplockExpiryTime string

		if o.SnaplockExpiryTimeQueryParameter != nil {
			qrSnaplockExpiryTime = *o.SnaplockExpiryTimeQueryParameter
		}
		qSnaplockExpiryTime := qrSnaplockExpiryTime
		if qSnaplockExpiryTime != "" {

			if err := r.SetQueryParam("snaplock_expiry_time", qSnaplockExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.SnapmirrorLabelQueryParameter != nil {

		// query param snapmirror_label
		var qrSnapmirrorLabel string

		if o.SnapmirrorLabelQueryParameter != nil {
			qrSnapmirrorLabel = *o.SnapmirrorLabelQueryParameter
		}
		qSnapmirrorLabel := qrSnapmirrorLabel
		if qSnapmirrorLabel != "" {

			if err := r.SetQueryParam("snapmirror_label", qSnapmirrorLabel); err != nil {
				return err
			}
		}
	}

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
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

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if o.VolumeNameQueryParameter != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeNameQueryParameter != nil {
			qrVolumeName = *o.VolumeNameQueryParameter
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnapshotCollectionGet binds the parameter fields
func (o *SnapshotCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnapshotCollectionGet binds the parameter order_by
func (o *SnapshotCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
