package drinks

type CreateDrinkRequest struct {
	Name       string `json:"name"`
	Alcohol    bool   `json:"alcohol"`
	LocationId string `json:"location_id"`
}

type DrinksByLocationRequest struct {
	LocationId string `json:"location_id"`
}
