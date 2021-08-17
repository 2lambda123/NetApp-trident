// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// AuditModifyReader is a Reader for the AuditModify structure.
type AuditModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuditModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAuditModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuditModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuditModifyAccepted creates a AuditModifyAccepted with default headers values
func NewAuditModifyAccepted() *AuditModifyAccepted {
	return &AuditModifyAccepted{}
}

/* AuditModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AuditModifyAccepted struct {
}

func (o *AuditModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /protocols/audit/{svm.uuid}][%d] auditModifyAccepted ", 202)
}

func (o *AuditModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAuditModifyDefault creates a AuditModifyDefault with default headers values
func NewAuditModifyDefault(code int) *AuditModifyDefault {
	return &AuditModifyDefault{
		_statusCode: code,
	}
}

/* AuditModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 9699340    | SVM UUID lookup failed                                                        |
| 9699343    | Audit configuration is absent for modification                                |
| 9699358    | Audit configuration is absent for enabling                                    |
| 9699359    | Audit configuration is already enabled                                        |
| 9699360    | Final consolidation is in progress, audit enable failed                       |
| 9699365    | Enabling of audit configuration failed                                        |
| 9699373    | Audit configuration is absent for disabling                                   |
| 9699374    | Audit configuration is already disabled                                       |
| 9699375    | Disabling of audit configuration failed                                       |
| 9699384    | The specified log_path does not exist                                         |
| 9699385    | The log_path must be a directory                                              |
| 9699386    | The log_path must be a canonical path in the SVMs namespace                   |
| 9699387    | The log_path cannot be empty                                                  |
| 9699388    | Rotate size must be greater than or equal to 1024 KB                          |
| 9699389    | The log_path must not contain a symbolic link                                 |
| 9699398    | The log_path exceeds a maximum supported length of characters                 |
| 9699399    | The log_path contains an unsupported read-only (DP/LS) volume                 |
| 9699400    | The specified log_path is not a valid destination for SVM                     |
| 9699402    | The log_path contains an unsupported snaplock volume                          |
| 9699403    | The log_path cannot be accessed for validation                                |
| 9699406    | The log_path validation failed                                                |
| 9699407    | Additional fields are provided                                                |
| 9699409    | Failed to enable multiproto.audit.evtxlog.support support capability          |
| 9699410    | Failed to disable multiproto.audit.evtxlog.support support capability         |
| 9699418    | Audit configuration is absent for rotate                                      |
| 9699419    | Failed to rotate audit log                                                    |
| 9699420    | Cannot rotate audit log, auditing is not enabled for this SVM                 |
| 9699428    | All nodes need to run ONTAP 8.3.0 release to audit CIFS logon-logoff events   |
| 9699429    | Failed to enable multiproto.audit.cifslogonlogoff.support support capability  |
| 9699430    | Failed to disable multiproto.audit.cifslogonlogoff.support support capability |
| 9699431    | All nodes need to run ONTAP 8.3.0 release to audit CAP staging events         |
| 9699432    | Failed to enable multiproto.audit.capstaging.support support capability       |
| 9699433    | Failed to disable multiproto.audit.capstaging.support support capability      |

*/
type AuditModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the audit modify default response
func (o *AuditModifyDefault) Code() int {
	return o._statusCode
}

func (o *AuditModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/audit/{svm.uuid}][%d] audit_modify default  %+v", o._statusCode, o.Payload)
}
func (o *AuditModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AuditModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
