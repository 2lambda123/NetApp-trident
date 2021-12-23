// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// LunCreateReader is a Reader for the LunCreate structure.
type LunCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewLunCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunCreateCreated creates a LunCreateCreated with default headers values
func NewLunCreateCreated() *LunCreateCreated {
	return &LunCreateCreated{}
}

/* LunCreateCreated describes a response with status code 201, with default header values.

Created
*/
type LunCreateCreated struct {
	Payload *models.LunResponse
}

func (o *LunCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/luns][%d] lunCreateCreated  %+v", 201, o.Payload)
}
func (o *LunCreateCreated) GetPayload() *models.LunResponse {
	return o.Payload
}

func (o *LunCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunCreateDefault creates a LunCreateDefault with default headers values
func NewLunCreateDefault(code int) *LunCreateDefault {
	return &LunCreateDefault{
		_statusCode: code,
	}
}

/* LunCreateDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 917927 | The specified volume was not found. |
| 918236 | The specified `location.volume.uuid` and `location.volume.name` do not refer to the same volume. |
| 2621462 | The specified SVM does not exist. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5242927 | The specified qtree was not found. |
| 5242950 | The specified `location.qtree.id` and `location.qtree.name` do not refer to the same qtree. |
| 5374121 | A LUN name can only contain characters A-Z, a-z, 0-9, "-", ".", "_", "{" and "}". |
| 5374123 | A negative size was provided for the LUN. |
| 5374124 | The specified size is too small for the LUN. |
| 5374125 | The specified size is too large for the LUN. |
| 5374129 | LUNs cannot be created on a load sharing mirror volume. |
| 5374130 | An invalid size value was provided. |
| 5374237 | LUNs cannot be created on an SVM root volume. |
| 5374238 | LUNs cannot be created in Snapshot copies. |
| 5374241 | A size value with invalid units was provided. |
| 5374242 | A LUN or NVMe namespace already exists at the specified path. |
| 5374352 | An invalid name was provided for the LUN. |
| 5374707 | Creating a LUN in the specific volume is not allowed because the volume is reserved for an application. |
| 5374858 | The volume specified by `name` is not the same as that specified by `location.volume`. |
| 5374859 | No volume was specified for the LUN. |
| 5374860 | The qtree specified by `name` is not the same as that specified by `location.qtree`. |
| 5374861 | The LUN base name specified by `name` is not the same as that specified by `location.logical_unit`. |
| 5374862 | No LUN path base name was provided for the LUN. |
| 5374863 | An error occurred after successfully creating the LUN. Some properties were not set. |
| 5374874 | The specified `clone.source.uuid` and `clone.source.name` do not refer to the same LUN. |
| 5374875 | The specified `clone.source` was not found. |
| 5374876 | The specified `clone.source` was not found. |
| 5374883 | The property cannot be specified when creating a LUN clone. The `target` property of the error object identifies the property. |
| 5374884 | A property that is required when creating a new LUN that is not a LUN clone or LUN copy was not supplied. The `target` property of the error object identifies the property. |
| 5374886 | An error occurred after successfully creating the LUN preventing the retrieval of its properties. |
| 5374899 | The `clone.source.uuid` property is not supported when specifying a source LUN from a Snapshot copy. |
| 5374928 | An incomplete attribute name/value pair was supplied. |
| 5374929 | The combined sizes of an attribute name and value are too large. |
| 5374932 | A name for an attribute was duplicated. |
| 13565952 | The LUN clone request failed. |

*/
type LunCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun create default response
func (o *LunCreateDefault) Code() int {
	return o._statusCode
}

func (o *LunCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/luns][%d] lun_create default  %+v", o._statusCode, o.Payload)
}
func (o *LunCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
