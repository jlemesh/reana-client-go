// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RequestTokenReader is a Reader for the RequestToken structure.
type RequestTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRequestTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRequestTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRequestTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /api/token] request_token", response, response.Code())
	}
}

// NewRequestTokenOK creates a RequestTokenOK with default headers values
func NewRequestTokenOK() *RequestTokenOK {
	return &RequestTokenOK{}
}

/*
RequestTokenOK describes a response with status code 200, with default header values.

User information correspoding to the session cookie sent in the request.
*/
type RequestTokenOK struct {
	Payload *RequestTokenOKBody
}

// IsSuccess returns true when this request token o k response has a 2xx status code
func (o *RequestTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this request token o k response has a 3xx status code
func (o *RequestTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request token o k response has a 4xx status code
func (o *RequestTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this request token o k response has a 5xx status code
func (o *RequestTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this request token o k response a status code equal to that given
func (o *RequestTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the request token o k response
func (o *RequestTokenOK) Code() int {
	return 200
}

func (o *RequestTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenOK %s", 200, payload)
}

func (o *RequestTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenOK %s", 200, payload)
}

func (o *RequestTokenOK) GetPayload() *RequestTokenOKBody {
	return o.Payload
}

func (o *RequestTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RequestTokenOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestTokenUnauthorized creates a RequestTokenUnauthorized with default headers values
func NewRequestTokenUnauthorized() *RequestTokenUnauthorized {
	return &RequestTokenUnauthorized{}
}

/*
RequestTokenUnauthorized describes a response with status code 401, with default header values.

Error message indicating that the uses is not authenticated.
*/
type RequestTokenUnauthorized struct {
	Payload *RequestTokenUnauthorizedBody
}

// IsSuccess returns true when this request token unauthorized response has a 2xx status code
func (o *RequestTokenUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this request token unauthorized response has a 3xx status code
func (o *RequestTokenUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request token unauthorized response has a 4xx status code
func (o *RequestTokenUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this request token unauthorized response has a 5xx status code
func (o *RequestTokenUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this request token unauthorized response a status code equal to that given
func (o *RequestTokenUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the request token unauthorized response
func (o *RequestTokenUnauthorized) Code() int {
	return 401
}

func (o *RequestTokenUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenUnauthorized %s", 401, payload)
}

func (o *RequestTokenUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenUnauthorized %s", 401, payload)
}

func (o *RequestTokenUnauthorized) GetPayload() *RequestTokenUnauthorizedBody {
	return o.Payload
}

func (o *RequestTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RequestTokenUnauthorizedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestTokenForbidden creates a RequestTokenForbidden with default headers values
func NewRequestTokenForbidden() *RequestTokenForbidden {
	return &RequestTokenForbidden{}
}

/*
RequestTokenForbidden describes a response with status code 403, with default header values.

Request failed. User token not valid.
*/
type RequestTokenForbidden struct {
	Payload *RequestTokenForbiddenBody
}

// IsSuccess returns true when this request token forbidden response has a 2xx status code
func (o *RequestTokenForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this request token forbidden response has a 3xx status code
func (o *RequestTokenForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request token forbidden response has a 4xx status code
func (o *RequestTokenForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this request token forbidden response has a 5xx status code
func (o *RequestTokenForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this request token forbidden response a status code equal to that given
func (o *RequestTokenForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the request token forbidden response
func (o *RequestTokenForbidden) Code() int {
	return 403
}

func (o *RequestTokenForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenForbidden %s", 403, payload)
}

func (o *RequestTokenForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenForbidden %s", 403, payload)
}

func (o *RequestTokenForbidden) GetPayload() *RequestTokenForbiddenBody {
	return o.Payload
}

func (o *RequestTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RequestTokenForbiddenBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestTokenInternalServerError creates a RequestTokenInternalServerError with default headers values
func NewRequestTokenInternalServerError() *RequestTokenInternalServerError {
	return &RequestTokenInternalServerError{}
}

/*
RequestTokenInternalServerError describes a response with status code 500, with default header values.

Request failed. Internal server error.
*/
type RequestTokenInternalServerError struct {
	Payload *RequestTokenInternalServerErrorBody
}

// IsSuccess returns true when this request token internal server error response has a 2xx status code
func (o *RequestTokenInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this request token internal server error response has a 3xx status code
func (o *RequestTokenInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this request token internal server error response has a 4xx status code
func (o *RequestTokenInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this request token internal server error response has a 5xx status code
func (o *RequestTokenInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this request token internal server error response a status code equal to that given
func (o *RequestTokenInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the request token internal server error response
func (o *RequestTokenInternalServerError) Code() int {
	return 500
}

func (o *RequestTokenInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenInternalServerError %s", 500, payload)
}

func (o *RequestTokenInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/token][%d] requestTokenInternalServerError %s", 500, payload)
}

func (o *RequestTokenInternalServerError) GetPayload() *RequestTokenInternalServerErrorBody {
	return o.Payload
}

func (o *RequestTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RequestTokenInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
RequestTokenForbiddenBody request token forbidden body
swagger:model RequestTokenForbiddenBody
*/
type RequestTokenForbiddenBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this request token forbidden body
func (o *RequestTokenForbiddenBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this request token forbidden body based on context it is used
func (o *RequestTokenForbiddenBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RequestTokenForbiddenBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RequestTokenForbiddenBody) UnmarshalBinary(b []byte) error {
	var res RequestTokenForbiddenBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RequestTokenInternalServerErrorBody request token internal server error body
swagger:model RequestTokenInternalServerErrorBody
*/
type RequestTokenInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this request token internal server error body
func (o *RequestTokenInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this request token internal server error body based on context it is used
func (o *RequestTokenInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RequestTokenInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RequestTokenInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res RequestTokenInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RequestTokenOKBody request token o k body
swagger:model RequestTokenOKBody
*/
type RequestTokenOKBody struct {

	// reana token
	ReanaToken *RequestTokenOKBodyReanaToken `json:"reana_token,omitempty"`
}

// Validate validates this request token o k body
func (o *RequestTokenOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReanaToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RequestTokenOKBody) validateReanaToken(formats strfmt.Registry) error {
	if swag.IsZero(o.ReanaToken) { // not required
		return nil
	}

	if o.ReanaToken != nil {
		if err := o.ReanaToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestTokenOK" + "." + "reana_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestTokenOK" + "." + "reana_token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this request token o k body based on the context it is used
func (o *RequestTokenOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateReanaToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RequestTokenOKBody) contextValidateReanaToken(ctx context.Context, formats strfmt.Registry) error {

	if o.ReanaToken != nil {

		if swag.IsZero(o.ReanaToken) { // not required
			return nil
		}

		if err := o.ReanaToken.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requestTokenOK" + "." + "reana_token")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("requestTokenOK" + "." + "reana_token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *RequestTokenOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RequestTokenOKBody) UnmarshalBinary(b []byte) error {
	var res RequestTokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RequestTokenOKBodyReanaToken request token o k body reana token
swagger:model RequestTokenOKBodyReanaToken
*/
type RequestTokenOKBodyReanaToken struct {

	// requested at
	RequestedAt string `json:"requested_at,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this request token o k body reana token
func (o *RequestTokenOKBodyReanaToken) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this request token o k body reana token based on context it is used
func (o *RequestTokenOKBodyReanaToken) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RequestTokenOKBodyReanaToken) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RequestTokenOKBodyReanaToken) UnmarshalBinary(b []byte) error {
	var res RequestTokenOKBodyReanaToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
RequestTokenUnauthorizedBody request token unauthorized body
swagger:model RequestTokenUnauthorizedBody
*/
type RequestTokenUnauthorizedBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this request token unauthorized body
func (o *RequestTokenUnauthorizedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this request token unauthorized body based on context it is used
func (o *RequestTokenUnauthorizedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RequestTokenUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RequestTokenUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res RequestTokenUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
