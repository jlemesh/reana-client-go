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

// CreateWorkflowReader is a Reader for the CreateWorkflow structure.
type CreateWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateWorkflowCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateWorkflowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateWorkflowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewCreateWorkflowNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateWorkflowCreated creates a CreateWorkflowCreated with default headers values
func NewCreateWorkflowCreated() *CreateWorkflowCreated {
	return &CreateWorkflowCreated{}
}

/* CreateWorkflowCreated describes a response with status code 201, with default header values.

Request succeeded. The workflow has been created.
*/
type CreateWorkflowCreated struct {
	Payload *CreateWorkflowCreatedBody
}

func (o *CreateWorkflowCreated) Error() string {
	return fmt.Sprintf("[POST /api/workflows][%d] createWorkflowCreated  %+v", 201, o.Payload)
}
func (o *CreateWorkflowCreated) GetPayload() *CreateWorkflowCreatedBody {
	return o.Payload
}

func (o *CreateWorkflowCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateWorkflowCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkflowBadRequest creates a CreateWorkflowBadRequest with default headers values
func NewCreateWorkflowBadRequest() *CreateWorkflowBadRequest {
	return &CreateWorkflowBadRequest{}
}

/* CreateWorkflowBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed
*/
type CreateWorkflowBadRequest struct {
	Payload *CreateWorkflowBadRequestBody
}

func (o *CreateWorkflowBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/workflows][%d] createWorkflowBadRequest  %+v", 400, o.Payload)
}
func (o *CreateWorkflowBadRequest) GetPayload() *CreateWorkflowBadRequestBody {
	return o.Payload
}

func (o *CreateWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateWorkflowBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkflowForbidden creates a CreateWorkflowForbidden with default headers values
func NewCreateWorkflowForbidden() *CreateWorkflowForbidden {
	return &CreateWorkflowForbidden{}
}

/* CreateWorkflowForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type CreateWorkflowForbidden struct {
	Payload *CreateWorkflowForbiddenBody
}

func (o *CreateWorkflowForbidden) Error() string {
	return fmt.Sprintf("[POST /api/workflows][%d] createWorkflowForbidden  %+v", 403, o.Payload)
}
func (o *CreateWorkflowForbidden) GetPayload() *CreateWorkflowForbiddenBody {
	return o.Payload
}

func (o *CreateWorkflowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateWorkflowForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkflowNotFound creates a CreateWorkflowNotFound with default headers values
func NewCreateWorkflowNotFound() *CreateWorkflowNotFound {
	return &CreateWorkflowNotFound{}
}

/* CreateWorkflowNotFound describes a response with status code 404, with default header values.

Request failed. User does not exist.
*/
type CreateWorkflowNotFound struct {
	Payload *CreateWorkflowNotFoundBody
}

func (o *CreateWorkflowNotFound) Error() string {
	return fmt.Sprintf("[POST /api/workflows][%d] createWorkflowNotFound  %+v", 404, o.Payload)
}
func (o *CreateWorkflowNotFound) GetPayload() *CreateWorkflowNotFoundBody {
	return o.Payload
}

func (o *CreateWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateWorkflowNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkflowInternalServerError creates a CreateWorkflowInternalServerError with default headers values
func NewCreateWorkflowInternalServerError() *CreateWorkflowInternalServerError {
	return &CreateWorkflowInternalServerError{}
}

/* CreateWorkflowInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type CreateWorkflowInternalServerError struct {
	Payload *CreateWorkflowInternalServerErrorBody
}

func (o *CreateWorkflowInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/workflows][%d] createWorkflowInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateWorkflowInternalServerError) GetPayload() *CreateWorkflowInternalServerErrorBody {
	return o.Payload
}

func (o *CreateWorkflowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateWorkflowInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateWorkflowNotImplemented creates a CreateWorkflowNotImplemented with default headers values
func NewCreateWorkflowNotImplemented() *CreateWorkflowNotImplemented {
	return &CreateWorkflowNotImplemented{}
}

/* CreateWorkflowNotImplemented describes a response with status code 501, with default header values.

Request failed. Not implemented.
*/
type CreateWorkflowNotImplemented struct {
}

func (o *CreateWorkflowNotImplemented) Error() string {
	return fmt.Sprintf("[POST /api/workflows][%d] createWorkflowNotImplemented ", 501)
}

func (o *CreateWorkflowNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CreateWorkflowBadRequestBody create workflow bad request body
swagger:model CreateWorkflowBadRequestBody
*/
type CreateWorkflowBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this create workflow bad request body
func (o *CreateWorkflowBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create workflow bad request body based on context it is used
func (o *CreateWorkflowBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateWorkflowBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateWorkflowBadRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateWorkflowBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateWorkflowCreatedBody create workflow created body
swagger:model CreateWorkflowCreatedBody
*/
type CreateWorkflowCreatedBody struct {

	// message
	Message string `json:"message,omitempty"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`

	// workflow name
	WorkflowName string `json:"workflow_name,omitempty"`
}

// Validate validates this create workflow created body
func (o *CreateWorkflowCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create workflow created body based on context it is used
func (o *CreateWorkflowCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateWorkflowCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateWorkflowCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateWorkflowCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateWorkflowForbiddenBody create workflow forbidden body
swagger:model CreateWorkflowForbiddenBody
*/
type CreateWorkflowForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this create workflow forbidden body
func (o *CreateWorkflowForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create workflow forbidden body based on context it is used
func (o *CreateWorkflowForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateWorkflowForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateWorkflowForbiddenBody) UnmarshalBinary(b []byte) error {
	var res CreateWorkflowForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateWorkflowInternalServerErrorBody create workflow internal server error body
swagger:model CreateWorkflowInternalServerErrorBody
*/
type CreateWorkflowInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this create workflow internal server error body
func (o *CreateWorkflowInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create workflow internal server error body based on context it is used
func (o *CreateWorkflowInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateWorkflowInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateWorkflowInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res CreateWorkflowInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateWorkflowNotFoundBody create workflow not found body
swagger:model CreateWorkflowNotFoundBody
*/
type CreateWorkflowNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this create workflow not found body
func (o *CreateWorkflowNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create workflow not found body based on context it is used
func (o *CreateWorkflowNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateWorkflowNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateWorkflowNotFoundBody) UnmarshalBinary(b []byte) error {
	var res CreateWorkflowNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
