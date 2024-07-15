package drinks

import (
	"net/http"
	"noTirT/alcotracker/app/api/shared"
	"noTirT/alcotracker/util"
)

type DrinksController interface {
	CreateDrink(w http.ResponseWriter, r *http.Request)
	GetDrinksByLocationId(w http.ResponseWriter, r *http.Request)
}

type drinksController struct {
	service DrinksService
}

func NewDrinksController(service DrinksService) DrinksController {
	return &drinksController{
		service: service,
	}
}

// GetDrinksByLocationId godoc
//
//	@Id				GetDrinksByLocationID
//	@Summary		Get Drinks By Location ID
//	@Description	Get Drinks by Location ID
//	@Tags			Drinks
//	@Produce		json
//	@Param			locationId	path		string	true "Id of location"
//	@Success		200			{object}	GenericResponse
//	@Failure		500			{object}	GenericResponse	"Server or Database internal error"
//	@Router			/drinks/{locationId} [GET]
//	@Security		Bearer
func (controller *drinksController) GetDrinksByLocationId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	locationId := r.PathValue("locationId")

	drinks, err := controller.service.GetDrinksByLocationId(locationId)

	if cancel := shared.HandleGetError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Retrieved drinks by location",
		Data:    drinks,
	}, w)
}

// CreateDrink godoc
//
//	@Id				CreateDrink
//	@Summary		Create Drink
//	@Description	Create new Drink
//	@Tags			Drinks
//	@Produce		json
//	@Param			CreateDrinkRequest	body		CreateDrinkRequest	true "Drink info"
//	@Success		200					{object}	GenericResponse
//	@Failure		500					{object}	GenericResponse	"Server or Database internal error"
//	@Router			/drinks [POST]
//	@Security		Bearer
func (controller *drinksController) CreateDrink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := &CreateDrinkRequest{}
	err := util.FromJSON(request, r.Body)

	if cancel := shared.HandleParseError(err, w); cancel {
		return
	}

	err = controller.service.CreateDrink(request)

	if cancel := shared.HandlePostError(err, w); cancel {
		return
	}

	w.WriteHeader(http.StatusOK)
	util.ToJSON(&shared.GenericResponse{
		Status:  true,
		Message: "Created drink successfully",
	}, w)
}
