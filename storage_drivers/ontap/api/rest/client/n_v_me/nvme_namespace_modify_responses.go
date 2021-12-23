// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NvmeNamespaceModifyReader is a Reader for the NvmeNamespaceModify structure.
type NvmeNamespaceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeNamespaceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeNamespaceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeNamespaceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeNamespaceModifyOK creates a NvmeNamespaceModifyOK with default headers values
func NewNvmeNamespaceModifyOK() *NvmeNamespaceModifyOK {
	return &NvmeNamespaceModifyOK{}
}

/* NvmeNamespaceModifyOK describes a response with status code 200, with default header values.

OK
*/
type NvmeNamespaceModifyOK struct {
}

func (o *NvmeNamespaceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvmeNamespaceModifyOK ", 200)
}

func (o *NvmeNamespaceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNvmeNamespaceModifyDefault creates a NvmeNamespaceModifyDefault with default headers values
func NewNvmeNamespaceModifyDefault(code int) *NvmeNamespaceModifyDefault {
	return &NvmeNamespaceModifyDefault{
		_statusCode: code,
	}
}

/* NvmeNamespaceModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 13565952 | The namespace clone request failed. |
| 72089724 | The specified namespace size is too large. |
| 72089730 | The specified namespace cannot be updated as it resides in a Snapshot copy. |
| 72090005 | The specified `clone.source.uuid` and `clone.source.name` do not refer to the same LUN. |
| 72090006 | The specified namespace was not found. This can apply to `clone.source` or the target namespace. The `target` property of the error object identifies the property. |
| 72090007 | The specified namespace was not found. This can apply to `clone.source` or the target namespace. The `target` property of the error object identifies the property. |
| 72090010 | An error occurred after successfully overwriting data for the namespace as a clone. Some properties were not modified. |
| 72090011 | An error occurred after successfully modifying some of the properties of the namespace. Some properties were not modified. |
| 72090016 | The namespace's aggregate is offline. The aggregate must be online to modify or remove the namespace. |
| 72090017 | The namespace's volume is offline. The volume must be online to modify or remove the namespace. |
| 72090038 | An attempt was made to reduce the size of the specified namespace. |

*/
type NvmeNamespaceModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme namespace modify default response
func (o *NvmeNamespaceModifyDefault) Code() int {
	return o._statusCode
}

func (o *NvmeNamespaceModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/namespaces/{uuid}][%d] nvme_namespace_modify default  %+v", o._statusCode, o.Payload)
}
func (o *NvmeNamespaceModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeNamespaceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
