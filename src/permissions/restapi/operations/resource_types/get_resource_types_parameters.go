package resource_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetResourceTypesParams creates a new GetResourceTypesParams object
// with the default values initialized.
func NewGetResourceTypesParams() GetResourceTypesParams {
	var ()
	return GetResourceTypesParams{}
}

// GetResourceTypesParams contains all the bound params for the get resource types operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetResourceTypes
type GetResourceTypesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The resource type name to search for.
	  In: query
	*/
	ResourceTypeName *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetResourceTypesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := httpkit.Values(r.URL.Query())

	qResourceTypeName, qhkResourceTypeName, _ := qs.GetOK("resource_type_name")
	if err := o.bindResourceTypeName(qResourceTypeName, qhkResourceTypeName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetResourceTypesParams) bindResourceTypeName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ResourceTypeName = &raw

	return nil
}
