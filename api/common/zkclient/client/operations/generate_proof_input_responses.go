// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/0glabs/0g-serving-broker/common/zkclient/models"
)

// GenerateProofInputReader is a Reader for the GenerateProofInput structure.
type GenerateProofInputReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateProofInputReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenerateProofInputOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateProofInputDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateProofInputOK creates a GenerateProofInputOK with default headers values
func NewGenerateProofInputOK() *GenerateProofInputOK {
	return &GenerateProofInputOK{}
}

/*
GenerateProofInputOK describes a response with status code 200, with default header values.

OK
*/
type GenerateProofInputOK struct {
	Payload models.AdditionalProperties
}

// IsSuccess returns true when this generate proof input o k response has a 2xx status code
func (o *GenerateProofInputOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate proof input o k response has a 3xx status code
func (o *GenerateProofInputOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate proof input o k response has a 4xx status code
func (o *GenerateProofInputOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate proof input o k response has a 5xx status code
func (o *GenerateProofInputOK) IsServerError() bool {
	return false
}

// IsCode returns true when this generate proof input o k response a status code equal to that given
func (o *GenerateProofInputOK) IsCode(code int) bool {
	return code == 200
}

func (o *GenerateProofInputOK) Error() string {
	return fmt.Sprintf("[POST /proof-input][%d] generateProofInputOK  %+v", 200, o.Payload)
}

func (o *GenerateProofInputOK) String() string {
	return fmt.Sprintf("[POST /proof-input][%d] generateProofInputOK  %+v", 200, o.Payload)
}

func (o *GenerateProofInputOK) GetPayload() models.AdditionalProperties {
	return o.Payload
}

func (o *GenerateProofInputOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateProofInputDefault creates a GenerateProofInputDefault with default headers values
func NewGenerateProofInputDefault(code int) *GenerateProofInputDefault {
	return &GenerateProofInputDefault{
		_statusCode: code,
	}
}

/*
GenerateProofInputDefault describes a response with status code -1, with default header values.

Error
*/
type GenerateProofInputDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the generate proof input default response
func (o *GenerateProofInputDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this generate proof input default response has a 2xx status code
func (o *GenerateProofInputDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate proof input default response has a 3xx status code
func (o *GenerateProofInputDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate proof input default response has a 4xx status code
func (o *GenerateProofInputDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate proof input default response has a 5xx status code
func (o *GenerateProofInputDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate proof input default response a status code equal to that given
func (o *GenerateProofInputDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GenerateProofInputDefault) Error() string {
	return fmt.Sprintf("[POST /proof-input][%d] generateProofInput default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateProofInputDefault) String() string {
	return fmt.Sprintf("[POST /proof-input][%d] generateProofInput default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateProofInputDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GenerateProofInputDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GenerateProofInputBody generate proof input body
swagger:model GenerateProofInputBody
*/
type GenerateProofInputBody struct {

	// l
	L int64 `json:"l,omitempty"`

	// pubkey
	Pubkey models.PublicKey `json:"pubkey"`

	// requests
	Requests []*models.Request `json:"requests"`

	// signatures
	Signatures models.Signatures `json:"signatures"`
}

// Validate validates this generate proof input body
func (o *GenerateProofInputBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubkey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateProofInputBody) validatePubkey(formats strfmt.Registry) error {
	if swag.IsZero(o.Pubkey) { // not required
		return nil
	}

	if err := o.Pubkey.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "pubkey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "pubkey")
		}
		return err
	}

	return nil
}

func (o *GenerateProofInputBody) validateRequests(formats strfmt.Registry) error {
	if swag.IsZero(o.Requests) { // not required
		return nil
	}

	for i := 0; i < len(o.Requests); i++ {
		if swag.IsZero(o.Requests[i]) { // not required
			continue
		}

		if o.Requests[i] != nil {
			if err := o.Requests[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GenerateProofInputBody) validateSignatures(formats strfmt.Registry) error {
	if swag.IsZero(o.Signatures) { // not required
		return nil
	}

	if err := o.Signatures.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signatures")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signatures")
		}
		return err
	}

	return nil
}

// ContextValidate validate this generate proof input body based on the context it is used
func (o *GenerateProofInputBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubkey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRequests(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateProofInputBody) contextValidatePubkey(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Pubkey.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "pubkey")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "pubkey")
		}
		return err
	}

	return nil
}

func (o *GenerateProofInputBody) contextValidateRequests(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Requests); i++ {

		if o.Requests[i] != nil {
			if err := o.Requests[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "requests" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GenerateProofInputBody) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	if err := o.Signatures.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("body" + "." + "signatures")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("body" + "." + "signatures")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GenerateProofInputBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenerateProofInputBody) UnmarshalBinary(b []byte) error {
	var res GenerateProofInputBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
