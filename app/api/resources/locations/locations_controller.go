package locations

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
)

type LocationsController interface {
	CreateLocation(w http.ResponseWriter, r *http.Request)
	GetLocationsByDrinkId(w http.ResponseWriter, r *http.Request)
	GetAllLocations(w http.ResponseWriter, r *http.Request)
}

type locationsController struct {
	service LocationsService
}

func NewLocationsController(service LocationsService) LocationsController {
	return &locationsController{
		service: service,
	}
}

// GetAllLocations godoc
//
//	@Id				GetAllLocations
//	@Summary		Get All Locations
//	@Description	Get all Locations
//	@Tags			Locations
//	@Produce		json
//	@Success		200	{object}	GenericResponse
//	@Failure		500	{object}	GenericResponse	"Server or Database internal error"
//	@Router			/locations [GET]
//	@Security		Bearer
func (controller *locationsController) GetAllLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	locations, err := controller.service.GetAllLocations()

	if shared.HandleGetError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Locations retrieved successfully",
		Data:    locations,
	}, w)
}

// GetLocationsByDrinkId godoc
//
//	@Id				GetLocationsByDrinkId
//	@Summary		Get Locations by Drink ID
//	@Description	Get locations by drink id
//	@Tags			Locations
//	@Produce		json
//	@Param			drinkId	path		string	true "ID of drink"
//	@Success		200		{object}	GenericResponse
//	@Failure		500		{object}	GenericResponse	"Server or Database internal error"
//	@Router			/locations/{drinkId} [GET]
//	@Security		Bearer
func (controller *locationsController) GetLocationsByDrinkId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	drinkId := r.PathValue("drinkId")

	locations, err := controller.service.GetLocationsByDrinkId(drinkId)

	if cancel := shared.HandleGetError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Retrieved locations successfully",
		Data:    locations,
	}, w)
}

// CreateLocation godoc
//
//	@Id				CreateLocation
//	@Summary		Create location
//	@Description	Create new location
//	@Tags			Locations
//	@Produce		json
//	@Param			LocationInfo	body		CreateLocationRequest	true "Location info"
//	@Success		200				{object}	GenericResponse
//	@Failure		500				{object}	GenericResponse	"Server or Database internal error"
//	@Router			/locations [POST]
//	@Security		Bearer
func (controller *locationsController) CreateLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := &CreateLocationRequest{}
	err := util.FromJSON(request, r.Body)

	if cancel := shared.HandleParseError(err, w); cancel {
		return
	}

	err = controller.service.CreateLocation(request)

	if cancel := shared.HandlePostError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Location created successfully",
	}, w)
}
