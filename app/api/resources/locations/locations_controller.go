package locations

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
)

type LocationsController interface {
	CreateLocation(w http.ResponseWriter, r *http.Request)
	GetLocationsByDrinkId(w http.ResponseWriter, r *http.Request)
}

type locationsController struct {
	service LocationsService
}

func NewLocationsController(service LocationsService) LocationsController {
	return &locationsController{
		service: service,
	}
}

func (controller *locationsController) GetLocationsByDrinkId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := &LocationsByDrinkRequest{}
	err := util.FromJSON(request, r.Body)

	if cancel := shared.HandleParseError(err, w); cancel {
		return
	}

	locations, err := controller.service.GetLocationsByDrinkId(request)

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
