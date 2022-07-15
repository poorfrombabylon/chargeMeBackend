// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package schema

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// Список станций по адресу
type AddressStations struct {
	Access    int32     `json:"access"`
	Address   string    `json:"address"`
	Icon      string    `json:"icon"`
	IconType  string    `json:"icon_type"`
	Id        int32     `json:"id"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	Name      string    `json:"name"`
	Score     float32   `json:"score"`
	Stations  []Station `json:"stations"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// Сущность разъема.
type Outlet struct {
	Connector int32 `json:"connector"`
	Id        int32 `json:"id"`
	Kilowatts int32 `json:"kilowatts"`
	Power     int32 `json:"power"`
}

// Сущность станции.
type Station struct {
	Id      int32    `json:"id"`
	Outlets []Outlet `json:"outlets"`
}

// GetChargingStationsParams defines parameters for GetChargingStations.
type GetChargingStationsParams struct {
	LatitudeMin  *float32 `form:"latitudeMin,omitempty" json:"latitudeMin,omitempty"`
	LongitudeMin *float32 `form:"longitudeMin,omitempty" json:"longitudeMin,omitempty"`
	LatitudeMax  *float32 `form:"latitudeMax,omitempty" json:"latitudeMax,omitempty"`
	LongitudeMax *float32 `form:"longitudeMax,omitempty" json:"longitudeMax,omitempty"`
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Проверка сервиса
	// (GET /healthz)
	HealthCheck(ctx echo.Context) error
	// Получение списка зарядных станций в пределах координат
	// (GET /v1/stations)
	GetChargingStations(ctx echo.Context, params GetChargingStationsParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// HealthCheck converts echo context to params.
func (w *ServerInterfaceWrapper) HealthCheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.HealthCheck(ctx)
	return err
}

// GetChargingStations converts echo context to params.
func (w *ServerInterfaceWrapper) GetChargingStations(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetChargingStationsParams
	// ------------- Optional query parameter "latitudeMin" -------------

	err = runtime.BindQueryParameter("form", true, false, "latitudeMin", ctx.QueryParams(), &params.LatitudeMin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter latitudeMin: %s", err))
	}

	// ------------- Optional query parameter "longitudeMin" -------------

	err = runtime.BindQueryParameter("form", true, false, "longitudeMin", ctx.QueryParams(), &params.LongitudeMin)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter longitudeMin: %s", err))
	}

	// ------------- Optional query parameter "latitudeMax" -------------

	err = runtime.BindQueryParameter("form", true, false, "latitudeMax", ctx.QueryParams(), &params.LatitudeMax)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter latitudeMax: %s", err))
	}

	// ------------- Optional query parameter "longitudeMax" -------------

	err = runtime.BindQueryParameter("form", true, false, "longitudeMax", ctx.QueryParams(), &params.LongitudeMax)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter longitudeMax: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetChargingStations(ctx, params)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/healthz", wrapper.HealthCheck)
	router.GET(baseURL+"/v1/stations", wrapper.GetChargingStations)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xWz2/bNhT+VwRuR8N2k6UodBuKYduh2KHHIRhY+sVmK5EqSafNAgNOgmzYD2CX3Yfd",
	"dlTTePXSWfkXHv+j4VFSrMRKIgxbbxRJvfd933v8yEMmdJppBcpZFh8yKyaQ8jD8dDQyYO1Tx53UKkyN",
	"wAojM/pmMcPf8RKX/ggLvIj8kT/GHFf+O1ziXxFeYhFhjud+jgt/5E9Yj2VGZ2CchBCLCwE2jPa0Sblj",
	"MZPKbW+xHnMHGZSfMAbDZj3GSyy0u1q0zkg1pjUpCM0tC9+Us22ro465E+6km46aUdQ0fVYtajW+fVXx",
	"tD25Fdq0/2IbcksHaRh8bGCPxeyjwbpYg6pSg6o+9G8VjBvDD9hs1mMGXk6lgRGLv671XmtZKdfUKajS",
	"YNzkV7GpsTeQ7l5l1s+eg3AE5TNjtCHo14sudKlUB9lTsJaP2+S7QSzEXO9vQ/PV1CXg2jrYn/gfcIUF",
	"ta//OfJzzPGd/xEX+Dfm/Y2mFVopEK6k1oFE5yZ7IRP9ijvX9UBk+hV0A7GhVs2gKvY6cx21TcG6yzpI",
	"2DSC5aaEnSXRoWjdD0JV5PvOQSBdx96kStul2tObTB9PuBnDE+o1J10C16f2wdhy34P+sD8MBDJQPJMs",
	"Ztv9YX+blOBuEngMJsATN/mWxuPWzvzNz7HAM1z4OV5gHvmjMDwLhptHuKI56tY3WPhjUp8sN5jxm3U1",
	"WEBhQu2+HLGYfRHyPp6AeMFIF5tpZcvCbA2HrR5/lbaZL8eFPw7a2mmacnNwP2bSjY8t1cCBmCgpeMJ2",
	"KcRg/8Gg6Xu3KIIFvvcn/ntc4AqXuIgCZ4pNyfAd5n7uf8FzXPmf/OnGjXQW4SXdRniOC3yPuT+N8AIL",
	"LPwcz3FJivrjDcE+BxeqLNX46iakQhqeggNDdA6ZJIAvp2AOapuMr0z0iVTBNKlJW1j9gUtSjaCW+HI8",
	"wwLf0vVJLP/EFS7Cpz/Bt4SbjhS85mlGLbjzsP/J9lbv5j0y690Cqrbze1D9SlpT1g+FqxaLv/63YoW2",
	"86f/k1x347pbrk7IHu60I9ttP6VCKwcqHBOeZYkUoTUHz23p0WukdznmzRdeML/r1J5Oy2fDrMd2/sPM",
	"5fOgJR/5v1E8iSyYfTARVBuv+8wHcoLarkRlANH6yROglxhLC5iahMVs4lxm48Eg0YInE21d/Gj4aMhm",
	"u7N/AgAA///RnGF1aQsAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
