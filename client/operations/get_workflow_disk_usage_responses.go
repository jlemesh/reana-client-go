// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetWorkflowDiskUsageReader is a Reader for the GetWorkflowDiskUsage structure.
type GetWorkflowDiskUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowDiskUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkflowDiskUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkflowDiskUsageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkflowDiskUsageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkflowDiskUsageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkflowDiskUsageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/workflows/{workflow_id_or_name}/disk_usage] get_workflow_disk_usage", response, response.Code())
	}
}

// NewGetWorkflowDiskUsageOK creates a GetWorkflowDiskUsageOK with default headers values
func NewGetWorkflowDiskUsageOK() *GetWorkflowDiskUsageOK {
	return &GetWorkflowDiskUsageOK{}
}

/*
GetWorkflowDiskUsageOK describes a response with status code 200, with default header values.

Request succeeded. Info about the disk usage is returned.
*/
type GetWorkflowDiskUsageOK struct {
	Payload *GetWorkflowDiskUsageOKBody
}

// IsSuccess returns true when this get workflow disk usage o k response has a 2xx status code
func (o *GetWorkflowDiskUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workflow disk usage o k response has a 3xx status code
func (o *GetWorkflowDiskUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow disk usage o k response has a 4xx status code
func (o *GetWorkflowDiskUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow disk usage o k response has a 5xx status code
func (o *GetWorkflowDiskUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow disk usage o k response a status code equal to that given
func (o *GetWorkflowDiskUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get workflow disk usage o k response
func (o *GetWorkflowDiskUsageOK) Code() int {
	return 200
}

func (o *GetWorkflowDiskUsageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageOK %s", 200, payload)
}

func (o *GetWorkflowDiskUsageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageOK %s", 200, payload)
}

func (o *GetWorkflowDiskUsageOK) GetPayload() *GetWorkflowDiskUsageOKBody {
	return o.Payload
}

func (o *GetWorkflowDiskUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiskUsageOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiskUsageBadRequest creates a GetWorkflowDiskUsageBadRequest with default headers values
func NewGetWorkflowDiskUsageBadRequest() *GetWorkflowDiskUsageBadRequest {
	return &GetWorkflowDiskUsageBadRequest{}
}

/*
GetWorkflowDiskUsageBadRequest describes a response with status code 400, with default header values.

Request failed. The incoming data specification seems malformed.
*/
type GetWorkflowDiskUsageBadRequest struct {
	Payload *GetWorkflowDiskUsageBadRequestBody
}

// IsSuccess returns true when this get workflow disk usage bad request response has a 2xx status code
func (o *GetWorkflowDiskUsageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow disk usage bad request response has a 3xx status code
func (o *GetWorkflowDiskUsageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow disk usage bad request response has a 4xx status code
func (o *GetWorkflowDiskUsageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow disk usage bad request response has a 5xx status code
func (o *GetWorkflowDiskUsageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow disk usage bad request response a status code equal to that given
func (o *GetWorkflowDiskUsageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get workflow disk usage bad request response
func (o *GetWorkflowDiskUsageBadRequest) Code() int {
	return 400
}

func (o *GetWorkflowDiskUsageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageBadRequest %s", 400, payload)
}

func (o *GetWorkflowDiskUsageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageBadRequest %s", 400, payload)
}

func (o *GetWorkflowDiskUsageBadRequest) GetPayload() *GetWorkflowDiskUsageBadRequestBody {
	return o.Payload
}

func (o *GetWorkflowDiskUsageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiskUsageBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiskUsageForbidden creates a GetWorkflowDiskUsageForbidden with default headers values
func NewGetWorkflowDiskUsageForbidden() *GetWorkflowDiskUsageForbidden {
	return &GetWorkflowDiskUsageForbidden{}
}

/*
GetWorkflowDiskUsageForbidden describes a response with status code 403, with default header values.

Request failed. User is not allowed to access workflow.
*/
type GetWorkflowDiskUsageForbidden struct {
	Payload *GetWorkflowDiskUsageForbiddenBody
}

// IsSuccess returns true when this get workflow disk usage forbidden response has a 2xx status code
func (o *GetWorkflowDiskUsageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow disk usage forbidden response has a 3xx status code
func (o *GetWorkflowDiskUsageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow disk usage forbidden response has a 4xx status code
func (o *GetWorkflowDiskUsageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow disk usage forbidden response has a 5xx status code
func (o *GetWorkflowDiskUsageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow disk usage forbidden response a status code equal to that given
func (o *GetWorkflowDiskUsageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get workflow disk usage forbidden response
func (o *GetWorkflowDiskUsageForbidden) Code() int {
	return 403
}

func (o *GetWorkflowDiskUsageForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageForbidden %s", 403, payload)
}

func (o *GetWorkflowDiskUsageForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageForbidden %s", 403, payload)
}

func (o *GetWorkflowDiskUsageForbidden) GetPayload() *GetWorkflowDiskUsageForbiddenBody {
	return o.Payload
}

func (o *GetWorkflowDiskUsageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiskUsageForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiskUsageNotFound creates a GetWorkflowDiskUsageNotFound with default headers values
func NewGetWorkflowDiskUsageNotFound() *GetWorkflowDiskUsageNotFound {
	return &GetWorkflowDiskUsageNotFound{}
}

/*
GetWorkflowDiskUsageNotFound describes a response with status code 404, with default header values.

Request failed. User does not exist.
*/
type GetWorkflowDiskUsageNotFound struct {
	Payload *GetWorkflowDiskUsageNotFoundBody
}

// IsSuccess returns true when this get workflow disk usage not found response has a 2xx status code
func (o *GetWorkflowDiskUsageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow disk usage not found response has a 3xx status code
func (o *GetWorkflowDiskUsageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow disk usage not found response has a 4xx status code
func (o *GetWorkflowDiskUsageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workflow disk usage not found response has a 5xx status code
func (o *GetWorkflowDiskUsageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workflow disk usage not found response a status code equal to that given
func (o *GetWorkflowDiskUsageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get workflow disk usage not found response
func (o *GetWorkflowDiskUsageNotFound) Code() int {
	return 404
}

func (o *GetWorkflowDiskUsageNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageNotFound %s", 404, payload)
}

func (o *GetWorkflowDiskUsageNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageNotFound %s", 404, payload)
}

func (o *GetWorkflowDiskUsageNotFound) GetPayload() *GetWorkflowDiskUsageNotFoundBody {
	return o.Payload
}

func (o *GetWorkflowDiskUsageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiskUsageNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkflowDiskUsageInternalServerError creates a GetWorkflowDiskUsageInternalServerError with default headers values
func NewGetWorkflowDiskUsageInternalServerError() *GetWorkflowDiskUsageInternalServerError {
	return &GetWorkflowDiskUsageInternalServerError{}
}

/*
GetWorkflowDiskUsageInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal controller error.
*/
type GetWorkflowDiskUsageInternalServerError struct {
	Payload *GetWorkflowDiskUsageInternalServerErrorBody
}

// IsSuccess returns true when this get workflow disk usage internal server error response has a 2xx status code
func (o *GetWorkflowDiskUsageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workflow disk usage internal server error response has a 3xx status code
func (o *GetWorkflowDiskUsageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workflow disk usage internal server error response has a 4xx status code
func (o *GetWorkflowDiskUsageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workflow disk usage internal server error response has a 5xx status code
func (o *GetWorkflowDiskUsageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workflow disk usage internal server error response a status code equal to that given
func (o *GetWorkflowDiskUsageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get workflow disk usage internal server error response
func (o *GetWorkflowDiskUsageInternalServerError) Code() int {
	return 500
}

func (o *GetWorkflowDiskUsageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageInternalServerError %s", 500, payload)
}

func (o *GetWorkflowDiskUsageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/workflows/{workflow_id_or_name}/disk_usage][%d] getWorkflowDiskUsageInternalServerError %s", 500, payload)
}

func (o *GetWorkflowDiskUsageInternalServerError) GetPayload() *GetWorkflowDiskUsageInternalServerErrorBody {
	return o.Payload
}

func (o *GetWorkflowDiskUsageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWorkflowDiskUsageInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetWorkflowDiskUsageBadRequestBody get workflow disk usage bad request body
swagger:model GetWorkflowDiskUsageBadRequestBody
*/
type GetWorkflowDiskUsageBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow disk usage bad request body
func (o *GetWorkflowDiskUsageBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow disk usage bad request body based on context it is used
func (o *GetWorkflowDiskUsageBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiskUsageBody get workflow disk usage body
swagger:model GetWorkflowDiskUsageBody
*/
type GetWorkflowDiskUsageBody struct {

	// search
	Search string `json:"search,omitempty"`

	// summarize
	Summarize bool `json:"summarize,omitempty"`
}

// Validate validates this get workflow disk usage body
func (o *GetWorkflowDiskUsageBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow disk usage body based on context it is used
func (o *GetWorkflowDiskUsageBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiskUsageForbiddenBody get workflow disk usage forbidden body
swagger:model GetWorkflowDiskUsageForbiddenBody
*/
type GetWorkflowDiskUsageForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow disk usage forbidden body
func (o *GetWorkflowDiskUsageForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow disk usage forbidden body based on context it is used
func (o *GetWorkflowDiskUsageForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageForbiddenBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiskUsageInternalServerErrorBody get workflow disk usage internal server error body
swagger:model GetWorkflowDiskUsageInternalServerErrorBody
*/
type GetWorkflowDiskUsageInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow disk usage internal server error body
func (o *GetWorkflowDiskUsageInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow disk usage internal server error body based on context it is used
func (o *GetWorkflowDiskUsageInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiskUsageNotFoundBody get workflow disk usage not found body
swagger:model GetWorkflowDiskUsageNotFoundBody
*/
type GetWorkflowDiskUsageNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get workflow disk usage not found body
func (o *GetWorkflowDiskUsageNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow disk usage not found body based on context it is used
func (o *GetWorkflowDiskUsageNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiskUsageOKBody get workflow disk usage o k body
swagger:model GetWorkflowDiskUsageOKBody
*/
type GetWorkflowDiskUsageOKBody struct {

	// disk usage info
	DiskUsageInfo []*GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0 `json:"disk_usage_info"`

	// user
	User string `json:"user,omitempty"`

	// workflow id
	WorkflowID string `json:"workflow_id,omitempty"`

	// workflow name
	WorkflowName string `json:"workflow_name,omitempty"`
}

// Validate validates this get workflow disk usage o k body
func (o *GetWorkflowDiskUsageOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDiskUsageInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowDiskUsageOKBody) validateDiskUsageInfo(formats strfmt.Registry) error {
	if swag.IsZero(o.DiskUsageInfo) { // not required
		return nil
	}

	for i := 0; i < len(o.DiskUsageInfo); i++ {
		if swag.IsZero(o.DiskUsageInfo[i]) { // not required
			continue
		}

		if o.DiskUsageInfo[i] != nil {
			if err := o.DiskUsageInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowDiskUsageOK" + "." + "disk_usage_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkflowDiskUsageOK" + "." + "disk_usage_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get workflow disk usage o k body based on the context it is used
func (o *GetWorkflowDiskUsageOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDiskUsageInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowDiskUsageOKBody) contextValidateDiskUsageInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DiskUsageInfo); i++ {

		if o.DiskUsageInfo[i] != nil {

			if swag.IsZero(o.DiskUsageInfo[i]) { // not required
				return nil
			}

			if err := o.DiskUsageInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getWorkflowDiskUsageOK" + "." + "disk_usage_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getWorkflowDiskUsageOK" + "." + "disk_usage_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageOKBody) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0 get workflow disk usage o k body disk usage info items0
swagger:model GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0
*/
type GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0 struct {

	// name
	Name string `json:"name,omitempty"`

	// size
	Size *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size `json:"size,omitempty"`
}

// Validate validates this get workflow disk usage o k body disk usage info items0
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0) validateSize(formats strfmt.Registry) error {
	if swag.IsZero(o.Size) { // not required
		return nil
	}

	if o.Size != nil {
		if err := o.Size.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get workflow disk usage o k body disk usage info items0 based on the context it is used
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if o.Size != nil {

		if swag.IsZero(o.Size) { // not required
			return nil
		}

		if err := o.Size.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("size")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("size")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size get workflow disk usage o k body disk usage info items0 size
swagger:model GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size
*/
type GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size struct {

	// human readable
	HumanReadable string `json:"human_readable,omitempty"`

	// raw
	Raw int64 `json:"raw,omitempty"`
}

// Validate validates this get workflow disk usage o k body disk usage info items0 size
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get workflow disk usage o k body disk usage info items0 size based on context it is used
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size) UnmarshalBinary(b []byte) error {
	var res GetWorkflowDiskUsageOKBodyDiskUsageInfoItems0Size
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
