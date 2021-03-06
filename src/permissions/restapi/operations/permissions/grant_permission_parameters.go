package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"permissions/models"
)

// NewGrantPermissionParams creates a new GrantPermissionParams object
// with the default values initialized.
func NewGrantPermissionParams() GrantPermissionParams {
	var ()
	return GrantPermissionParams{}
}

// GrantPermissionParams contains all the bound params for the grant permission operation
// typically these are obtained from a http.Request
//
// swagger:parameters grantPermission
type GrantPermissionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Information about the permission to add.
	  Required: true
	  In: body
	*/
	PermissionGrantRequest *models.PermissionGrantRequest
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GrantPermissionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if httpkit.HasBody(r) {
		defer r.Body.Close()
		var body models.PermissionGrantRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("permissionGrantRequest", "body"))
			} else {
				res = append(res, errors.NewParseError("permissionGrantRequest", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.PermissionGrantRequest = &body
			}
		}

	} else {
		res = append(res, errors.Required("permissionGrantRequest", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
