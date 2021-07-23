// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetCwfNetworkIDHaPairsReader is a Reader for the GetCwfNetworkIDHaPairs structure.
type GetCwfNetworkIDHaPairsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDHaPairsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDHaPairsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDHaPairsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDHaPairsOK creates a GetCwfNetworkIDHaPairsOK with default headers values
func NewGetCwfNetworkIDHaPairsOK() *GetCwfNetworkIDHaPairsOK {
	return &GetCwfNetworkIDHaPairsOK{}
}

/*GetCwfNetworkIDHaPairsOK handles this case with default header values.

All high availability gateway pairs in Carrier Wifi network
*/
type GetCwfNetworkIDHaPairsOK struct {
	Payload map[string]models.CwfHaPair
}

func (o *GetCwfNetworkIDHaPairsOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/ha_pairs][%d] getCwfNetworkIdHaPairsOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDHaPairsOK) GetPayload() map[string]models.CwfHaPair {
	return o.Payload
}

func (o *GetCwfNetworkIDHaPairsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDHaPairsDefault creates a GetCwfNetworkIDHaPairsDefault with default headers values
func NewGetCwfNetworkIDHaPairsDefault(code int) *GetCwfNetworkIDHaPairsDefault {
	return &GetCwfNetworkIDHaPairsDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDHaPairsDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDHaPairsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID ha pairs default response
func (o *GetCwfNetworkIDHaPairsDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDHaPairsDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/ha_pairs][%d] GetCwfNetworkIDHaPairs default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDHaPairsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDHaPairsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}