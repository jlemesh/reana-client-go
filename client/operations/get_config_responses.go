// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetConfigReader is a Reader for the GetConfig structure.
type GetConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetConfigInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigOK creates a GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {
	return &GetConfigOK{}
}

/* GetConfigOK describes a response with status code 200, with default header values.

Configuration information to consume by Reana-UI.
*/
type GetConfigOK struct {
	Payload interface{}
}

func (o *GetConfigOK) Error() string {
	return fmt.Sprintf("[GET /api/config][%d] getConfigOK  %+v", 200, o.Payload)
}
func (o *GetConfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigInternalServerError creates a GetConfigInternalServerError with default headers values
func NewGetConfigInternalServerError() *GetConfigInternalServerError {
	return &GetConfigInternalServerError{}
}

/* GetConfigInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type GetConfigInternalServerError struct {
	Payload *GetConfigInternalServerErrorBody
}

func (o *GetConfigInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/config][%d] getConfigInternalServerError  %+v", 500, o.Payload)
}
func (o *GetConfigInternalServerError) GetPayload() *GetConfigInternalServerErrorBody {
	return o.Payload
}

func (o *GetConfigInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetConfigInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetConfigInternalServerErrorBody get config internal server error body
swagger:model GetConfigInternalServerErrorBody
*/
type GetConfigInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get config internal server error body
func (o *GetConfigInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get config internal server error body based on context it is used
func (o *GetConfigInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetConfigInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
