package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/http/response"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/keyauth/apps/service"
)

var (
	h = &handler{}
)

type handler struct {
	service service.MetaServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(service.AppName)
	h.service = app.GetGrpcApp(service.AppName).(service.MetaServiceServer)
	return nil
}

func (h *handler) Name() string {
	return service.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"services"}

	ws.Route(ws.POST("/").To(h.CreateService).
		Doc("create a service").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(service.CreateServiceRequest{}).
		Writes(response.NewData(service.Service{})))

	ws.Route(ws.GET("/").To(h.QueryService).
		Doc("get all service").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Metadata("action", "list").
		Reads(service.QueryServiceRequest{}).
		Writes(response.NewData(service.ServiceSet{})).
		Returns(200, "OK", service.ServiceSet{}))

	ws.Route(ws.GET("/{id}").To(h.DescribeService).
		Doc("get a service").
		Param(ws.PathParameter("id", "identifier of the service").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(response.NewData(service.Service{})).
		Returns(200, "OK", response.NewData(service.Service{})).
		Returns(404, "Not Found", nil))

	ws.Route(ws.PUT("/{id}").To(h.UpdateService).
		Doc("update a service").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(service.CreateServiceRequest{}))

	ws.Route(ws.PATCH("/{id}").To(h.PatchService).
		Doc("patch a service").
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(service.CreateServiceRequest{}))

	ws.Route(ws.DELETE("/{id}").To(h.DeleteService).
		Doc("delete a service").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Param(ws.PathParameter("id", "identifier of the service").DataType("string")))
}

func init() {
	app.RegistryRESTfulApp(h)
}
