// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// NvmeSubsystemMapDeleteReader is a Reader for the NvmeSubsystemMapDelete structure.
type NvmeSubsystemMapDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemMapDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemMapDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemMapDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemMapDeleteOK creates a NvmeSubsystemMapDeleteOK with default headers values
func NewNvmeSubsystemMapDeleteOK() *NvmeSubsystemMapDeleteOK {
	return &NvmeSubsystemMapDeleteOK{}
}

/* NvmeSubsystemMapDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemMapDeleteOK struct {
	Payload *models.NvmeSubsystemMapResponse
}

func (o *NvmeSubsystemMapDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystem-maps/{subsystem.uuid}/{namespace.uuid}][%d] nvmeSubsystemMapDeleteOK  %+v", 200, o.Payload)
}
func (o *NvmeSubsystemMapDeleteOK) GetPayload() *models.NvmeSubsystemMapResponse {
	return o.Payload
}

func (o *NvmeSubsystemMapDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeSubsystemMapResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemMapDeleteDefault creates a NvmeSubsystemMapDeleteDefault with default headers values
func NewNvmeSubsystemMapDeleteDefault(code int) *NvmeSubsystemMapDeleteDefault {
	return &NvmeSubsystemMapDeleteDefault{
		_statusCode: code,
	}
}

/* NvmeSubsystemMapDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 72090019 | The specified NVMe namespace is not mapped to the specified NVMe subsystem. |

*/
type NvmeSubsystemMapDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme subsystem map delete default response
func (o *NvmeSubsystemMapDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemMapDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystem-maps/{subsystem.uuid}/{namespace.uuid}][%d] nvme_subsystem_map_delete default  %+v", o._statusCode, o.Payload)
}
func (o *NvmeSubsystemMapDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemMapDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
