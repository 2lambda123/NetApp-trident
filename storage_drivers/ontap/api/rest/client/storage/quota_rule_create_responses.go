// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// QuotaRuleCreateReader is a Reader for the QuotaRuleCreate structure.
type QuotaRuleCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotaRuleCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewQuotaRuleCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQuotaRuleCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotaRuleCreateAccepted creates a QuotaRuleCreateAccepted with default headers values
func NewQuotaRuleCreateAccepted() *QuotaRuleCreateAccepted {
	return &QuotaRuleCreateAccepted{}
}

/* QuotaRuleCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QuotaRuleCreateAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *QuotaRuleCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /storage/quota/rules][%d] quotaRuleCreateAccepted  %+v", 202, o.Payload)
}
func (o *QuotaRuleCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QuotaRuleCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotaRuleCreateDefault creates a QuotaRuleCreateDefault with default headers values
func NewQuotaRuleCreateDefault(code int) *QuotaRuleCreateDefault {
	return &QuotaRuleCreateDefault{
		_statusCode: code,
	}
}

/* QuotaRuleCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 917927 | The specified volume was not found. |
| 918232 | Either `volume.name` or `volume.uuid` must be provided. |
| 918236 | The specified `volume.uuid` and `volume.name` refer to different volumes. |
| 2621462 | The specified SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5308501 | Mapping from Windows user to UNIX user for user rule was unsuccessful. |
| 5308502 | Mapping from UNIX user to Windows user for user rule was unsuccessful. |
| 5308552 | Failed to get default quota policy name for SVM. |
| 5308561 | Failed to obtain volume quota state or invalid quota state obtained for volume. |
| 5308562 | users is a required input for creating a user rule and group is not allowed. |
| 5308563 | group is a required input for creating a group rule and users is not allowed. |
| 5308564 | qtree.name is a required input for creating a tree rule and users and group are not allowed. |
| 5308565 | Only one of name or id is allowed for each entry in the users array. |
| 5308566 | Only one of name or id is allowed for group. |
| 5308568 | Quota policy rule create operation succeeded, but quota resize failed due to internal error. To activate the rule, disable and enable quotas for this volume. |
| 5308571 | Quota policy rule create operation succeeded, but quota resize is skipped. To activate the rule, disable and enable quotas for this volume. |
| 5308573 | Input value is greater than limit for field. |
| 5308574 | Input value is out of range for field. |
| 5308575 | Input value is incorrectly larger than listed field. |

*/
type QuotaRuleCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the quota rule create default response
func (o *QuotaRuleCreateDefault) Code() int {
	return o._statusCode
}

func (o *QuotaRuleCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/quota/rules][%d] quota_rule_create default  %+v", o._statusCode, o.Payload)
}
func (o *QuotaRuleCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuotaRuleCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
