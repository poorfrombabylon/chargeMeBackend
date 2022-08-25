// Package schema provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package schema

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// Список станций по адресу
type AddressStationsPreliminary struct {
	Access    int                  `json:"access"`
	Address   string               `json:"address"`
	Icon      *string              `json:"icon,omitempty"`
	IconType  *string              `json:"icon_type,omitempty"`
	Id        string               `json:"id"`
	Latitude  float32              `json:"latitude"`
	Longitude float32              `json:"longitude"`
	Name      string               `json:"name"`
	Score     *float32             `json:"score,omitempty"`
	Stations  []StationPreliminary `json:"stations"`
}

// Удобсва.
type Amenity struct {
	Form       *int   `json:"form,omitempty"`
	Id         string `json:"id"`
	LocationId string `json:"location_id"`
}

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// LocationWithFullStation defines model for LocationWithFullStation.
type LocationWithFullStation struct {
	// Удобства
	Amenities []Amenity `json:"amenities"`

	// Список станций по адресу
	Location StationsFull `json:"location"`

	// Отзывы о локации
	Reviews []Review `json:"reviews"`
}

// Сущность разъема.
type OutletPreliminary struct {
	Connector int      `json:"connector"`
	Id        string   `json:"id"`
	Kilowatts *float32 `json:"kilowatts"`
	Power     int      `json:"power"`
}

// ResponseLocations defines model for ResponseLocations.
type ResponseLocations struct {
	// Результат запроса.
	Locations []AddressStationsPreliminary `json:"locations"`
}

// Отзыв о локации.
type Review struct {
	Comment       *string   `json:"comment,omitempty"`
	ConnectorType *int      `json:"connector_type,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	Id            string    `json:"id"`
	OutletId      string    `json:"outlet_id"`
	Rating        *int      `json:"rating,omitempty"`
	StationId     string    `json:"station_id"`
	UserName      *string   `json:"user_name,omitempty"`
	VehicleName   *string   `json:"vehicle_name,omitempty"`
	VehicleType   *string   `json:"vehicle_type,omitempty"`
}

// Полная информация о станции.
type StationFull struct {
	Available       *int                `json:"available,omitempty"`
	Cost            *int                `json:"cost,omitempty"`
	CostDescription *string             `json:"cost_description,omitempty"`
	Hours           *string             `json:"hours,omitempty"`
	Id              string              `json:"id"`
	Kilowatts       *float32            `json:"kilowatts,omitempty"`
	Manufacturer    *string             `json:"manufacturer,omitempty"`
	Name            *string             `json:"name,omitempty"`
	Outlets         []OutletPreliminary `json:"outlets"`
}

// Сущность станции.
type StationPreliminary struct {
	Id      string              `json:"id"`
	Outlets []OutletPreliminary `json:"outlets"`
}

// Список станций по адресу
type StationsFull struct {
	Access                       *int          `json:"access,omitempty"`
	AccessRestriction            *string       `json:"access_restriction,omitempty"`
	AccessRestrictionDescription *string       `json:"access_restriction_description,omitempty"`
	Address                      string        `json:"address"`
	Amenities                    *[]Amenity    `json:"amenities,omitempty"`
	ComingSoon                   *bool         `json:"coming_soon,omitempty"`
	Cost                         *bool         `json:"cost,omitempty"`
	CostDescription              *string       `json:"cost_description,omitempty"`
	Description                  *string       `json:"description,omitempty"`
	Hours                        *string       `json:"hours,omitempty"`
	IconType                     *string       `json:"icon_type,omitempty"`
	Id                           string        `json:"id"`
	Latitude                     float32       `json:"latitude"`
	Longitude                    float32       `json:"longitude"`
	Name                         string        `json:"name"`
	Open247                      *bool         `json:"open247,omitempty"`
	PhoneNumber                  *string       `json:"phone_number,omitempty"`
	Reviews                      *[]Review     `json:"reviews,omitempty"`
	Score                        *float32      `json:"score,omitempty"`
	Stations                     []StationFull `json:"stations"`
}

// GetLocationsParams defines parameters for GetLocations.
type GetLocationsParams struct {
	LatitudeMin  *float32 `form:"latitudeMin,omitempty" json:"latitudeMin,omitempty"`
	LongitudeMin *float32 `form:"longitudeMin,omitempty" json:"longitudeMin,omitempty"`
	LatitudeMax  *float32 `form:"latitudeMax,omitempty" json:"latitudeMax,omitempty"`
	LongitudeMax *float32 `form:"longitudeMax,omitempty" json:"longitudeMax,omitempty"`
}

// CreateFullLocationJSONBody defines parameters for CreateFullLocation.
type CreateFullLocationJSONBody = StationsFull

// GetChargingStationsByLocationIDParams defines parameters for GetChargingStationsByLocationID.
type GetChargingStationsByLocationIDParams struct {
	LocationId string `form:"locationId" json:"locationId"`
}

// CreateFullLocationJSONRequestBody defines body for CreateFullLocation for application/json ContentType.
type CreateFullLocationJSONRequestBody = CreateFullLocationJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Проверка сервиса
	// (GET /healthz)
	HealthCheck(w http.ResponseWriter, r *http.Request)
	// Получение списка локаций с зарядками в пределах координат.
	// (GET /v1/locations)
	GetLocations(w http.ResponseWriter, r *http.Request, params GetLocationsParams)
	// Создания локации со станциями.
	// (POST /v1/locations)
	CreateFullLocation(w http.ResponseWriter, r *http.Request)
	// Получение списка зарядных станций и удобств на локации.
	// (GET /v1/locations/stations)
	GetChargingStationsByLocationID(w http.ResponseWriter, r *http.Request, params GetChargingStationsByLocationIDParams)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// HealthCheck operation middleware
func (siw *ServerInterfaceWrapper) HealthCheck(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.HealthCheck(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetLocations operation middleware
func (siw *ServerInterfaceWrapper) GetLocations(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLocationsParams

	// ------------- Optional query parameter "latitudeMin" -------------
	if paramValue := r.URL.Query().Get("latitudeMin"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "latitudeMin", r.URL.Query(), &params.LatitudeMin)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "latitudeMin", Err: err})
		return
	}

	// ------------- Optional query parameter "longitudeMin" -------------
	if paramValue := r.URL.Query().Get("longitudeMin"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "longitudeMin", r.URL.Query(), &params.LongitudeMin)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "longitudeMin", Err: err})
		return
	}

	// ------------- Optional query parameter "latitudeMax" -------------
	if paramValue := r.URL.Query().Get("latitudeMax"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "latitudeMax", r.URL.Query(), &params.LatitudeMax)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "latitudeMax", Err: err})
		return
	}

	// ------------- Optional query parameter "longitudeMax" -------------
	if paramValue := r.URL.Query().Get("longitudeMax"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "longitudeMax", r.URL.Query(), &params.LongitudeMax)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "longitudeMax", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetLocations(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// CreateFullLocation operation middleware
func (siw *ServerInterfaceWrapper) CreateFullLocation(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.CreateFullLocation(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetChargingStationsByLocationID operation middleware
func (siw *ServerInterfaceWrapper) GetChargingStationsByLocationID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetChargingStationsByLocationIDParams

	// ------------- Required query parameter "locationId" -------------
	if paramValue := r.URL.Query().Get("locationId"); paramValue != "" {

	} else {
		siw.ErrorHandlerFunc(w, r, &RequiredParamError{ParamName: "locationId"})
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "locationId", r.URL.Query(), &params.LocationId)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "locationId", Err: err})
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetChargingStationsByLocationID(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/healthz", wrapper.HealthCheck)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/locations", wrapper.GetLocations)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/v1/locations", wrapper.CreateFullLocation)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/v1/locations/stations", wrapper.GetChargingStationsByLocationID)
	})

	return r
}
