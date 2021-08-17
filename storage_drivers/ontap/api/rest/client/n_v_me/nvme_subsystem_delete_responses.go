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

// NvmeSubsystemDeleteReader is a Reader for the NvmeSubsystemDelete structure.
type NvmeSubsystemDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeSubsystemDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemDeleteOK creates a NvmeSubsystemDeleteOK with default headers values
func NewNvmeSubsystemDeleteOK() *NvmeSubsystemDeleteOK {
	return &NvmeSubsystemDeleteOK{}
}

/* NvmeSubsystemDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NvmeSubsystemDeleteOK struct {
}

func (o *NvmeSubsystemDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystems/{uuid}][%d] nvmeSubsystemDeleteOK ", 200)
}

func (o *NvmeSubsystemDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeSubsystemDeleteDefault creates a NvmeSubsystemDeleteDefault with default headers values
func NewNvmeSubsystemDeleteDefault(code int) *NvmeSubsystemDeleteDefault {
	return &NvmeSubsystemDeleteDefault{
		_statusCode: code,
	}
}

/* NvmeSubsystemDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 72090023 | The NVMe subsystem contains one or more mapped namespaces. Use the `allow_delete_while_mapped` query parameter to delete an NVMe subsystem with mapped NVMe namespaces. |
| 72090024 | The NVMe subsystem contains one or more NVMe hosts. Use the `allow_delete_with_hosts` query parameter to delete an NVMe subsystem with NVMe hosts. |

*/
type NvmeSubsystemDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme subsystem delete default response
func (o *NvmeSubsystemDeleteDefault) Code() int {
	return o._statusCode
}

func (o *NvmeSubsystemDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/nvme/subsystems/{uuid}][%d] nvme_subsystem_delete default  %+v", o._statusCode, o.Payload)
}
func (o *NvmeSubsystemDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
