// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// IgroupCreateReader is a Reader for the IgroupCreate structure.
type IgroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IgroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIgroupCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIgroupCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIgroupCreateCreated creates a IgroupCreateCreated with default headers values
func NewIgroupCreateCreated() *IgroupCreateCreated {
	return &IgroupCreateCreated{}
}

/* IgroupCreateCreated describes a response with status code 201, with default header values.

Created
*/
type IgroupCreateCreated struct {
	Payload *models.IgroupResponse
}

func (o *IgroupCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/san/igroups][%d] igroupCreateCreated  %+v", 201, o.Payload)
}
func (o *IgroupCreateCreated) GetPayload() *models.IgroupResponse {
	return o.Payload
}

func (o *IgroupCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IgroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIgroupCreateDefault creates a IgroupCreateDefault with default headers values
func NewIgroupCreateDefault(code int) *IgroupCreateDefault {
	return &IgroupCreateDefault{
		_statusCode: code,
	}
}

/* IgroupCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 2621462 | The supplied SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5373958 | An invalid initiator group name was supplied. |
| 5373966 | An initiator group cannot be created in an SVM that is configured for NVMe. |
| 5373969 | A supplied initiator name looks like an iSCSI IQN initiator, but the portions after the prefix are missing. |
| 5373971 | A supplied initiator name looks like an iSCSI IQN initiator, but the date portion is invalid. |
| 5373972 | A supplied initiator name looks like an iSCSI IQN initiator, but the naming authority portion is invalid. |
| 5373977 | A supplied initiator name looks like an iSCSI EUI initiator, but the length is invalid. |
| 5373978 | A supplied initiator name looks like an iSCSI EUI initiator, but the format is invalid. |
| 5373992 | A supplied initiator name was too long to be valid. |
| 5373993 | A supplied initiator name did not match any valid format. |
| 5374023 | An initiator group with the same name already exists. |
| 5374038 | An invalid Fibre Channel WWPN was supplied. |
| 5374039 | An invalid iSCSI initiator name was supplied. |
| 5374732 | An initiator is already in another initiator group with a conflicting operating system type. |

*/
type IgroupCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the igroup create default response
func (o *IgroupCreateDefault) Code() int {
	return o._statusCode
}

func (o *IgroupCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/san/igroups][%d] igroup_create default  %+v", o._statusCode, o.Payload)
}
func (o *IgroupCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *IgroupCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
