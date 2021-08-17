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

// AggregateModifyReader is a Reader for the AggregateModify structure.
type AggregateModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewAggregateModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregateModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregateModifyAccepted creates a AggregateModifyAccepted with default headers values
func NewAggregateModifyAccepted() *AggregateModifyAccepted {
	return &AggregateModifyAccepted{}
}

/* AggregateModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type AggregateModifyAccepted struct {
	Payload *models.JobLinkResponse
}

func (o *AggregateModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /storage/aggregates/{uuid}][%d] aggregateModifyAccepted  %+v", 202, o.Payload)
}
func (o *AggregateModifyAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *AggregateModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateModifyDefault creates a AggregateModifyDefault with default headers values
func NewAggregateModifyDefault(code int) *AggregateModifyDefault {
	return &AggregateModifyDefault{
		_statusCode: code,
	}
}

/* AggregateModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 262247 | The value is invalid for the field. |
| 460777 | Failed to get information on the job. |
| 786434 | Cannot connect to node where the aggregate resides. |
| 786435 | Internal Error. Failed to create a communication handle. |
| 786439 | An aggregate already uses the specified name. |
| 786447 | Failed to modify aggregate. |
| 786456 | Failed to add disks to aggregate. |
| 786458 | Failed to rename aggregate. |
| 786468 | VLDB is offline. |
| 786472 | Node that hosts the aggregate is offline. |
| 786479 | Cannot find node ID for the node. |
| 786730 | Internal Error |
| 786771 | Aggregate does not exist. |
| 786787 | Aggregate is not online. |
| 786808 | Aggregate mirror failed. |
| 786867 | Specified aggregate resides on the remote cluster. |
| 786911 | Not every node in the cluster has the Data ONTAP version required for the feature. |
| 786923 | This operation is disallowed during pre-commit phase of 7-mode to clustered Data ONTAP transition. |
| 786924 | Internal Error for an aggregate that is in pre-commit phase of a 7-mode to clustered Data ONTAP transition. |
| 786955 | Modifying raidtype to raid_tec requires a minimum of six disks in the RAID Group. |
| 786956 | Modifying raidtype to raid_dp requires a minimum of four disks in the RAID Group. |
| 786965 | Spare Selection in userspace failed. |
| 787046 | Mirroring of a FabricPool is not allowed. |
| 787092 | The target field cannot be specified for this operation. |
| 787144 | Aggregate is not a FabricPool. |
| 787156 | Modifying the attributes of mirror object store is not allowed. |
| 787169 | Only one field can be modified per operation. |
| 787170 | Failed to patch the \"block_storage.primary.disk_count\" because the disk count specified is smaller than existing disk count. |
| 787178 | Unmirroring an aggregate with a PATCH operation is not supported. |
| 2425736 | No matching node found for the UUID provided. |
| 13108106 | Cannot run aggregate relocation because volume expand is in progress. |
| 26542083 | Destination node is at higher Data ONTAP version than source node. |
| 26542084 | Source node is at higher Data ONTAP version than destination node. |
| 26542097 | Unable to get D-blade ID of destination. |
| 26542101 | Unable to contact source node. |
| 26542102 | Unable to contact destination node. |
| 26542120 | A Vserver migrate operation is in progress. When the migrate operation completes, try the operation again. |
| 26542121 | A MetroCluster disaster recovery operation is in progress. When the recovery operation completes, try the operation again. |

*/
type AggregateModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the aggregate modify default response
func (o *AggregateModifyDefault) Code() int {
	return o._statusCode
}

func (o *AggregateModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/aggregates/{uuid}][%d] aggregate_modify default  %+v", o._statusCode, o.Payload)
}
func (o *AggregateModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AggregateModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
