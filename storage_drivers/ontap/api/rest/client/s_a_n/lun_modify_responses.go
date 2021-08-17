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

// LunModifyReader is a Reader for the LunModify structure.
type LunModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunModifyOK creates a LunModifyOK with default headers values
func NewLunModifyOK() *LunModifyOK {
	return &LunModifyOK{}
}

/* LunModifyOK describes a response with status code 200, with default header values.

OK
*/
type LunModifyOK struct {
}

func (o *LunModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/luns/{uuid}][%d] lunModifyOK ", 200)
}

func (o *LunModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLunModifyDefault creates a LunModifyDefault with default headers values
func NewLunModifyDefault(code int) *LunModifyDefault {
	return &LunModifyDefault{
		_statusCode: code,
	}
}

/* LunModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 917927 | The specified volume was not found. |
| 918236 | The specified `location.volume.uuid` and `location.volume.name` do not refer to the same volume. |
| 5242927 | The specified qtree was not found. |
| 5242950 | The specified `location.qtree.id` and `location.qtree.name` do not refer to the same qtree. |
| 5374124 | The specified LUN size is too small. |
| 5374125 | The specified LUN size is too large. |
| 5374130 | An invalid size value was provided. |
| 5374241 | A size value with invalid units was provided. |
| 5374480 | Modifying the LUN is not allowed because it is in a foreign LUN import relationship. |
| 5374858 | The volume specified by `name` is not the same as that specified by `location.volume`. |
| 5374860 | The qtree specified by `name` is not the same as that specified by `location.qtree`. |
| 5374861 | The LUN base name specified by `name` is not the same as that specified by `location.logical_unit`. |
| 5374864 | An error occurred after successfully overwriting data for the LUN as a clone. Some properties were not modified. |
| 5374865 | The LUN's aggregate is offline. The aggregate must be online to modify or remove the LUN. |
| 5374866 | The LUN's volume is offline. The volume must be online to modify or remove the LUN. |
| 5374874 | The specified `clone.source.uuid` and `clone.source.name` do not refer to the same LUN. |
| 5374875 | The specified LUN was not found. This can apply to `clone.source` or the target LUN. The `target` property of the error object identifies the property. |
| 5374876 | The specified LUN was not found. This can apply to `clone.source` or the target LUN. The `target` property of the error object identifies the property. |
| 5374885 | An error occurred after successfully modifying some of the properties of the LUN. Some properties were not modified. |
| 5374889 | An invalid value was specified for `movement.progress.state`. Active LUN movement operations can be PATCHed to only _paused_ or _replicating_. |
| 5374892 | An attempt was made to reduce the size of a LUN. |
| 13565952 | The LUN clone request failed. |

*/
type LunModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun modify default response
func (o *LunModifyDefault) Code() int {
	return o._statusCode
}

func (o *LunModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/luns/{uuid}][%d] lun_modify default  %+v", o._statusCode, o.Payload)
}
func (o *LunModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
