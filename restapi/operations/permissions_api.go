package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	httpkit "github.com/go-swagger/go-swagger/httpkit"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"
	spec "github.com/go-swagger/go-swagger/spec"
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"permissions/restapi/operations/permissions"
	"permissions/restapi/operations/resource_types"
	"permissions/restapi/operations/resources"
	"permissions/restapi/operations/status"
	"permissions/restapi/operations/subjects"
)

// NewPermissionsAPI creates a new Permissions instance
func NewPermissionsAPI(spec *spec.Document) *PermissionsAPI {
	o := &PermissionsAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*PermissionsAPI Manages Permissions for the CyVerse Discovery Environment and related applications. */
type PermissionsAPI struct {
	spec            *spec.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer httpkit.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer httpkit.Producer

	// ResourceTypesDeleteResourceTypesIDHandler sets the operation handler for the delete resource types ID operation
	ResourceTypesDeleteResourceTypesIDHandler resource_types.DeleteResourceTypesIDHandler
	// StatusGetHandler sets the operation handler for the get operation
	StatusGetHandler status.GetHandler
	// ResourceTypesGetResourceTypesHandler sets the operation handler for the get resource types operation
	ResourceTypesGetResourceTypesHandler resource_types.GetResourceTypesHandler
	// ResourceTypesPostResourceTypesHandler sets the operation handler for the post resource types operation
	ResourceTypesPostResourceTypesHandler resource_types.PostResourceTypesHandler
	// ResourceTypesPutResourceTypesIDHandler sets the operation handler for the put resource types ID operation
	ResourceTypesPutResourceTypesIDHandler resource_types.PutResourceTypesIDHandler
	// ResourcesAddResourceHandler sets the operation handler for the add resource operation
	ResourcesAddResourceHandler resources.AddResourceHandler
	// SubjectsAddSubjectHandler sets the operation handler for the add subject operation
	SubjectsAddSubjectHandler subjects.AddSubjectHandler
	// ResourcesDeleteResourceHandler sets the operation handler for the delete resource operation
	ResourcesDeleteResourceHandler resources.DeleteResourceHandler
	// SubjectsDeleteSubjectHandler sets the operation handler for the delete subject operation
	SubjectsDeleteSubjectHandler subjects.DeleteSubjectHandler
	// PermissionsGrantPermissionHandler sets the operation handler for the grant permission operation
	PermissionsGrantPermissionHandler permissions.GrantPermissionHandler
	// PermissionsListPermissionsHandler sets the operation handler for the list permissions operation
	PermissionsListPermissionsHandler permissions.ListPermissionsHandler
	// ResourcesListResourcesHandler sets the operation handler for the list resources operation
	ResourcesListResourcesHandler resources.ListResourcesHandler
	// SubjectsListSubjectsHandler sets the operation handler for the list subjects operation
	SubjectsListSubjectsHandler subjects.ListSubjectsHandler
	// ResourcesUpdateResourceHandler sets the operation handler for the update resource operation
	ResourcesUpdateResourceHandler resources.UpdateResourceHandler
	// SubjectsUpdateSubjectHandler sets the operation handler for the update subject operation
	SubjectsUpdateSubjectHandler subjects.UpdateSubjectHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup
}

// SetDefaultProduces sets the default produces media type
func (o *PermissionsAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *PermissionsAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *PermissionsAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *PermissionsAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *PermissionsAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *PermissionsAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the PermissionsAPI
func (o *PermissionsAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ResourceTypesDeleteResourceTypesIDHandler == nil {
		unregistered = append(unregistered, "resource_types.DeleteResourceTypesIDHandler")
	}

	if o.StatusGetHandler == nil {
		unregistered = append(unregistered, "status.GetHandler")
	}

	if o.ResourceTypesGetResourceTypesHandler == nil {
		unregistered = append(unregistered, "resource_types.GetResourceTypesHandler")
	}

	if o.ResourceTypesPostResourceTypesHandler == nil {
		unregistered = append(unregistered, "resource_types.PostResourceTypesHandler")
	}

	if o.ResourceTypesPutResourceTypesIDHandler == nil {
		unregistered = append(unregistered, "resource_types.PutResourceTypesIDHandler")
	}

	if o.ResourcesAddResourceHandler == nil {
		unregistered = append(unregistered, "resources.AddResourceHandler")
	}

	if o.SubjectsAddSubjectHandler == nil {
		unregistered = append(unregistered, "subjects.AddSubjectHandler")
	}

	if o.ResourcesDeleteResourceHandler == nil {
		unregistered = append(unregistered, "resources.DeleteResourceHandler")
	}

	if o.SubjectsDeleteSubjectHandler == nil {
		unregistered = append(unregistered, "subjects.DeleteSubjectHandler")
	}

	if o.PermissionsGrantPermissionHandler == nil {
		unregistered = append(unregistered, "permissions.GrantPermissionHandler")
	}

	if o.PermissionsListPermissionsHandler == nil {
		unregistered = append(unregistered, "permissions.ListPermissionsHandler")
	}

	if o.ResourcesListResourcesHandler == nil {
		unregistered = append(unregistered, "resources.ListResourcesHandler")
	}

	if o.SubjectsListSubjectsHandler == nil {
		unregistered = append(unregistered, "subjects.ListSubjectsHandler")
	}

	if o.ResourcesUpdateResourceHandler == nil {
		unregistered = append(unregistered, "resources.UpdateResourceHandler")
	}

	if o.SubjectsUpdateSubjectHandler == nil {
		unregistered = append(unregistered, "subjects.UpdateSubjectHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *PermissionsAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *PermissionsAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]httpkit.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *PermissionsAPI) ConsumersFor(mediaTypes []string) map[string]httpkit.Consumer {

	result := make(map[string]httpkit.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *PermissionsAPI) ProducersFor(mediaTypes []string) map[string]httpkit.Producer {

	result := make(map[string]httpkit.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *PermissionsAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *PermissionsAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/resource_types/{id}"] = resource_types.NewDeleteResourceTypesID(o.context, o.ResourceTypesDeleteResourceTypesIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/"] = status.NewGet(o.context, o.StatusGetHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resource_types"] = resource_types.NewGetResourceTypes(o.context, o.ResourceTypesGetResourceTypesHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resource_types"] = resource_types.NewPostResourceTypes(o.context, o.ResourceTypesPostResourceTypesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/resource_types/{id}"] = resource_types.NewPutResourceTypesID(o.context, o.ResourceTypesPutResourceTypesIDHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/resources"] = resources.NewAddResource(o.context, o.ResourcesAddResourceHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/subjects"] = subjects.NewAddSubject(o.context, o.SubjectsAddSubjectHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/resources/{id}"] = resources.NewDeleteResource(o.context, o.ResourcesDeleteResourceHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/subjects/{id}"] = subjects.NewDeleteSubject(o.context, o.SubjectsDeleteSubjectHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/permissions"] = permissions.NewGrantPermission(o.context, o.PermissionsGrantPermissionHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/permissions"] = permissions.NewListPermissions(o.context, o.PermissionsListPermissionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/resources"] = resources.NewListResources(o.context, o.ResourcesListResourcesHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/subjects"] = subjects.NewListSubjects(o.context, o.SubjectsListSubjectsHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/resources/{id}"] = resources.NewUpdateResource(o.context, o.ResourcesUpdateResourceHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/subjects/{id}"] = subjects.NewUpdateSubject(o.context, o.SubjectsUpdateSubjectHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *PermissionsAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
