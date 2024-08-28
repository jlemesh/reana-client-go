// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PingReader is a Reader for the Ping structure.
type PingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/ping] ping", response, response.Code())
	}
}

// NewPingOK creates a PingOK with default headers values
func NewPingOK() *PingOK {
	return &PingOK{}
}

/*
PingOK describes a response with status code 200, with default header values.

Ping succeeded. Service is running and accessible.
*/
type PingOK struct {
	Payload *PingOKBody
}

// IsSuccess returns true when this ping o k response has a 2xx status code
func (o *PingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ping o k response has a 3xx status code
func (o *PingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping o k response has a 4xx status code
func (o *PingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping o k response has a 5xx status code
func (o *PingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ping o k response a status code equal to that given
func (o *PingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ping o k response
func (o *PingOK) Code() int {
	return 200
}

func (o *PingOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/ping][%d] pingOK %s", 200, payload)
}

func (o *PingOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/ping][%d] pingOK %s", 200, payload)
}

func (o *PingOK) GetPayload() *PingOKBody {
	return o.Payload
}

func (o *PingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PingOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PingOKBody ping o k body
swagger:model PingOKBody
*/
type PingOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this ping o k body
func (o *PingOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ping o k body based on context it is used
func (o *PingOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PingOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PingOKBody) UnmarshalBinary(b []byte) error {
	var res PingOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
