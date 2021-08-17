// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/v21/storage_drivers/ontap/api/rest/models"
)

// ConfigurationBackupFileDeleteReader is a Reader for the ConfigurationBackupFileDelete structure.
type ConfigurationBackupFileDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigurationBackupFileDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigurationBackupFileDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConfigurationBackupFileDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConfigurationBackupFileDeleteOK creates a ConfigurationBackupFileDeleteOK with default headers values
func NewConfigurationBackupFileDeleteOK() *ConfigurationBackupFileDeleteOK {
	return &ConfigurationBackupFileDeleteOK{}
}

/* ConfigurationBackupFileDeleteOK describes a response with status code 200, with default header values.

OK
*/
type ConfigurationBackupFileDeleteOK struct {
}

func (o *ConfigurationBackupFileDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /support/configuration-backup/backups/{node.uuid}/{name}][%d] configurationBackupFileDeleteOK ", 200)
}

func (o *ConfigurationBackupFileDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewConfigurationBackupFileDeleteDefault creates a ConfigurationBackupFileDeleteDefault with default headers values
func NewConfigurationBackupFileDeleteDefault(code int) *ConfigurationBackupFileDeleteDefault {
	return &ConfigurationBackupFileDeleteDefault{
		_statusCode: code,
	}
}

/* ConfigurationBackupFileDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 5963826 | Failed to delete backup file. |

*/
type ConfigurationBackupFileDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the configuration backup file delete default response
func (o *ConfigurationBackupFileDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ConfigurationBackupFileDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /support/configuration-backup/backups/{node.uuid}/{name}][%d] configuration_backup_file_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ConfigurationBackupFileDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigurationBackupFileDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
