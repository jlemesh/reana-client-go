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

// GetWorkflowDiffReader is a Reader for the GetWorkflowDiff structure.
type GetWorkflowDiffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowDiffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowDiffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkflowDiffBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkflowDiffForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowDiffNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowDiffInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}] get_workflow_diff", response, response.Code())
	}
}

// NewGetWorkflowDiffOK creates a GetWorkflowDiffOK with default headers values
func NewGetWorkflowDiffOK() *GetWorkflowDiffOK {
	return &GetWorkflowDiffOK{}
}

/*
GetWorkflowDiffOK describes a response with status code 200, with default header values.

Request succeeded. Info about a workflow, including the status is returned.
*/
type GetWorkflowDiffOK struct {
	Payload *GetWorkflowDiffOKBody
}

// IsSuccess returns true when this get workflow diff o k response has a 2xx status code
func (o *GetWorkflowDiffOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow diff o k response has a 3xx status code
func (o *GetWorkflowDiffOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow diff o k response has a 4xx status code
func (o *GetWorkflowDiffOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow diff o k response has a 5xx status code
func (o *GetWorkflowDiffOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow diff o k response a status code equal to that given
func (o *GetWorkflowDiffOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workflow diff o k response
func (o *GetWorkflowDiffOK) Code() int {
	return 200
}

func (o *GetWorkflowDiffOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffOK %s", 200, payload)
}

func (o *GetWorkflowDiffOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffOK %s", 200, payload)
}

func (o *GetWorkflowDiffOK) GetPayload() *GetWorkflowDiffOKBody {
	return o.Payload
}

func (o *GetWorkflowDiffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiffOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiffBadRequest creates a GetWorkflowDiffBadRequest with default headers values
func NewGetWorkflowDiffBadRequest() *GetWorkflowDiffBadRequest {
	return &GetWorkflowDiffBadRequest{}
}

/*
GetWorkflowDiffBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming payload seems malformed.
*/
type GetWorkflowDiffBadRequest struct {
	Payload *GetWorkflowDiffBadRequestBody
}

// IsSuccess returns true when this get workflow diff bad request response has a 2xx status code
func (o *GetWorkflowDiffBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow diff bad request response has a 3xx status code
func (o *GetWorkflowDiffBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow diff bad request response has a 4xx status code
func (o *GetWorkflowDiffBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow diff bad request response has a 5xx status code
func (o *GetWorkflowDiffBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow diff bad request response a status code equal to that given
func (o *GetWorkflowDiffBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get workflow diff bad request response
func (o *GetWorkflowDiffBadRequest) Code() int {
	return 400
}

func (o *GetWorkflowDiffBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffBadRequest %s", 400, payload)
}

func (o *GetWorkflowDiffBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffBadRequest %s", 400, payload)
}

func (o *GetWorkflowDiffBadRequest) GetPayload() *GetWorkflowDiffBadRequestBody {
	return o.Payload
}

func (o *GetWorkflowDiffBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiffBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiffForbidden creates a GetWorkflowDiffForbidden with default headers values
func NewGetWorkflowDiffForbidden() *GetWorkflowDiffForbidden {
	return &GetWorkflowDiffForbidden{}
}

/*
GetWorkflowDiffForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type GetWorkflowDiffForbidden struct {
	Payload *GetWorkflowDiffForbiddenBody
}

// IsSuccess returns true when this get workflow diff forbidden response has a 2xx status code
func (o *GetWorkflowDiffForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow diff forbidden response has a 3xx status code
func (o *GetWorkflowDiffForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow diff forbidden response has a 4xx status code
func (o *GetWorkflowDiffForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow diff forbidden response has a 5xx status code
func (o *GetWorkflowDiffForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow diff forbidden response a status code equal to that given
func (o *GetWorkflowDiffForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get workflow diff forbidden response
func (o *GetWorkflowDiffForbidden) Code() int {
	return 403
}

func (o *GetWorkflowDiffForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffForbidden %s", 403, payload)
}

func (o *GetWorkflowDiffForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffForbidden %s", 403, payload)
}

func (o *GetWorkflowDiffForbidden) GetPayload() *GetWorkflowDiffForbiddenBody {
	return o.Payload
}

func (o *GetWorkflowDiffForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiffForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiffNotFound creates a GetWorkflowDiffNotFound with default headers values
func NewGetWorkflowDiffNotFound() *GetWorkflowDiffNotFound {
	return &GetWorkflowDiffNotFound{}
}

/*
GetWorkflowDiffNotFound describes a response with status code 404, with default header values.

Request failed. Either user or workflow does not exist.
*/
type GetWorkflowDiffNotFound struct {
	Payload *GetWorkflowDiffNotFoundBody
}

// IsSuccess returns true when this get workflow diff not found response has a 2xx status code
func (o *GetWorkflowDiffNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow diff not found response has a 3xx status code
func (o *GetWorkflowDiffNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow diff not found response has a 4xx status code
func (o *GetWorkflowDiffNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow diff not found response has a 5xx status code
func (o *GetWorkflowDiffNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow diff not found response a status code equal to that given
func (o *GetWorkflowDiffNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workflow diff not found response
func (o *GetWorkflowDiffNotFound) Code() int {
	return 404
}

func (o *GetWorkflowDiffNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffNotFound %s", 404, payload)
}

func (o *GetWorkflowDiffNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffNotFound %s", 404, payload)
}

func (o *GetWorkflowDiffNotFound) GetPayload() *GetWorkflowDiffNotFoundBody {
	return o.Payload
}

func (o *GetWorkflowDiffNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiffNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiffInternalServerError creates a GetWorkflowDiffInternalServerError with default headers values
func NewGetWorkflowDiffInternalServerError() *GetWorkflowDiffInternalServerError {
	return &GetWorkflowDiffInternalServerError{}
}

/*
GetWorkflowDiffInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type GetWorkflowDiffInternalServerError struct {
	Payload *GetWorkflowDiffInternalServerErrorBody
}

// IsSuccess returns true when this get workflow diff internal server error response has a 2xx status code
func (o *GetWorkflowDiffInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow diff internal server error response has a 3xx status code
func (o *GetWorkflowDiffInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow diff internal server error response has a 4xx status code
func (o *GetWorkflowDiffInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow diff internal server error response has a 5xx status code
func (o *GetWorkflowDiffInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workflow diff internal server error response a status code equal to that given
func (o *GetWorkflowDiffInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get workflow diff internal server error response
func (o *GetWorkflowDiffInternalServerError) Code() int {
	return 500
}

func (o *GetWorkflowDiffInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffInternalServerError %s", 500, payload)
}

func (o *GetWorkflowDiffInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name_a}/diff/{workflow_id_or_name_b}][%d] getWorkflowDiffInternalServerError %s", 500, payload)
}

func (o *GetWorkflowDiffInternalServerError) GetPayload() *GetWorkflowDiffInternalServerErrorBody {
	return o.Payload
}

func (o *GetWorkflowDiffInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiffInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetWorkflowDiffBadRequestBody get workflow diff bad request body
swagger:model GetWorkflowDiffBadRequestBody
*/
type GetWorkflowDiffBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow diff bad request body
func (o *GetWorkflowDiffBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow diff bad request body based on context it is used
func (o *GetWorkflowDiffBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiffBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiffBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiffBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiffForbiddenBody get workflow diff forbidden body
swagger:model GetWorkflowDiffForbiddenBody
*/
type GetWorkflowDiffForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow diff forbidden body
func (o *GetWorkflowDiffForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow diff forbidden body based on context it is used
func (o *GetWorkflowDiffForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiffForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiffForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiffForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiffInternalServerErrorBody get workflow diff internal server error body
swagger:model GetWorkflowDiffInternalServerErrorBody
*/
type GetWorkflowDiffInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow diff internal server error body
func (o *GetWorkflowDiffInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow diff internal server error body based on context it is used
func (o *GetWorkflowDiffInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiffInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiffInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiffInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiffNotFoundBody get workflow diff not found body
swagger:model GetWorkflowDiffNotFoundBody
*/
type GetWorkflowDiffNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow diff not found body
func (o *GetWorkflowDiffNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow diff not found body based on context it is used
func (o *GetWorkflowDiffNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiffNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiffNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiffNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiffOKBody get workflow diff o k body
swagger:model GetWorkflowDiffOKBody
*/
type GetWorkflowDiffOKBody struct {

	// reana specification
	ReanaSpecification string `json:"reana_specification,omitempty"`

	// workspace listing
	WorkspaceListing string `json:"workspace_listing,omitempty"`
}

// Validate validates this get workflow diff o k body
func (o *GetWorkflowDiffOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow diff o k body based on context it is used
func (o *GetWorkflowDiffOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiffOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiffOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiffOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
