package locations

type CreateLocationRequest struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type LocationsByDrinkRequest struct {
	DrinkId string `json:"drink_id"`
}
