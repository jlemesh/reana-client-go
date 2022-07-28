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

// GetWorkflowParametersReader is a Reader for the GetWorkflowParameters structure.
type GetWorkflowParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkflowParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkflowParametersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowParametersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowParametersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkflowParametersOK creates a GetWorkflowParametersOK with default headers values
func NewGetWorkflowParametersOK() *GetWorkflowParametersOK {
	return &GetWorkflowParametersOK{}
}

/* GetWorkflowParametersOK describes a response with status code 200, with default header values.

Request succeeded. Workflow input parameters, including the status are returned.
*/
type GetWorkflowParametersOK struct {
	Payload *GetWorkflowParametersOKBody
}

func (o *GetWorkflowParametersOK) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/parameters][%d] getWorkflowParametersOK  %+v", 200, o.Payload)
}
func (o *GetWorkflowParametersOK) GetPayload() *GetWorkflowParametersOKBody {
	return o.Payload
}

func (o *GetWorkflowParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowParametersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowParametersBadRequest creates a GetWorkflowParametersBadRequest with default headers values
func NewGetWorkflowParametersBadRequest() *GetWorkflowParametersBadRequest {
	return &GetWorkflowParametersBadRequest{}
}

/* GetWorkflowParametersBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type GetWorkflowParametersBadRequest struct {
	Payload *GetWorkflowParametersBadRequestBody
}

func (o *GetWorkflowParametersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/parameters][%d] getWorkflowParametersBadRequest  %+v", 400, o.Payload)
}
func (o *GetWorkflowParametersBadRequest) GetPayload() *GetWorkflowParametersBadRequestBody {
	return o.Payload
}

func (o *GetWorkflowParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowParametersBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowParametersForbidden creates a GetWorkflowParametersForbidden with default headers values
func NewGetWorkflowParametersForbidden() *GetWorkflowParametersForbidden {
	return &GetWorkflowParametersForbidden{}
}

/* GetWorkflowParametersForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type GetWorkflowParametersForbidden struct {
	Payload *GetWorkflowParametersForbiddenBody
}

func (o *GetWorkflowParametersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/parameters][%d] getWorkflowParametersForbidden  %+v", 403, o.Payload)
}
func (o *GetWorkflowParametersForbidden) GetPayload() *GetWorkflowParametersForbiddenBody {
	return o.Payload
}

func (o *GetWorkflowParametersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowParametersForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowParametersNotFound creates a GetWorkflowParametersNotFound with default headers values
func NewGetWorkflowParametersNotFound() *GetWorkflowParametersNotFound {
	return &GetWorkflowParametersNotFound{}
}

/* GetWorkflowParametersNotFound describes a response with status code 404, with default header values.

Request failed. Either User or Analysis does not exist.
*/
type GetWorkflowParametersNotFound struct {
	Payload *GetWorkflowParametersNotFoundBody
}

func (o *GetWorkflowParametersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/parameters][%d] getWorkflowParametersNotFound  %+v", 404, o.Payload)
}
func (o *GetWorkflowParametersNotFound) GetPayload() *GetWorkflowParametersNotFoundBody {
	return o.Payload
}

func (o *GetWorkflowParametersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowParametersNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowParametersInternalServerError creates a GetWorkflowParametersInternalServerError with default headers values
func NewGetWorkflowParametersInternalServerError() *GetWorkflowParametersInternalServerError {
	return &GetWorkflowParametersInternalServerError{}
}

/* GetWorkflowParametersInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type GetWorkflowParametersInternalServerError struct {
	Payload *GetWorkflowParametersInternalServerErrorBody
}

func (o *GetWorkflowParametersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/parameters][%d] getWorkflowParametersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetWorkflowParametersInternalServerError) GetPayload() *GetWorkflowParametersInternalServerErrorBody {
	return o.Payload
}

func (o *GetWorkflowParametersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowParametersInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWorkflowParametersBadRequestBody get workflow parameters bad request body
swagger:model GetWorkflowParametersBadRequestBody
*/
type GetWorkflowParametersBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow parameters bad request body
func (o *GetWorkflowParametersBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow parameters bad request body based on context it is used
func (o *GetWorkflowParametersBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowParametersBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowParametersBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowParametersBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowParametersForbiddenBody get workflow parameters forbidden body
swagger:model GetWorkflowParametersForbiddenBody
*/
type GetWorkflowParametersForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow parameters forbidden body
func (o *GetWorkflowParametersForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow parameters forbidden body based on context it is used
func (o *GetWorkflowParametersForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowParametersForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowParametersForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowParametersForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowParametersInternalServerErrorBody get workflow parameters internal server error body
swagger:model GetWorkflowParametersInternalServerErrorBody
*/
type GetWorkflowParametersInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow parameters internal server error body
func (o *GetWorkflowParametersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow parameters internal server error body based on context it is used
func (o *GetWorkflowParametersInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowParametersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowParametersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowParametersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowParametersNotFoundBody get workflow parameters not found body
swagger:model GetWorkflowParametersNotFoundBody
*/
type GetWorkflowParametersNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow parameters not found body
func (o *GetWorkflowParametersNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow parameters not found body based on context it is used
func (o *GetWorkflowParametersNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowParametersNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowParametersNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowParametersNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetWorkflowParametersOKBody get workflow parameters o k body
swagger:model GetWorkflowParametersOKBody
*/
type GetWorkflowParametersOKBody struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameters
	Parameters interface{} `json:"parameters,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this get workflow parameters o k body
func (o *GetWorkflowParametersOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow parameters o k body based on context it is used
func (o *GetWorkflowParametersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowParametersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowParametersOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowParametersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
