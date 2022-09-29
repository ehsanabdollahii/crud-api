package api

const (
	ErrorCode400InvalidJson    = "invalid_json"
	ErrorMessage400InvalidJson = "The submitted Json is invalid"

	ErrorCode400InvalidData    = "invalid_data"
	ErrorMessage400InvalidData = "The submitted form is invalid"

	// city
	ErrorCode400CityExists    = "city_exist"
	ErrorMessage400CityExists = "There is already a city with the given name"

	ErrorCode400CityNotExists    = "city_not_exist"
	ErrorMessage400CityNotExists = "There is no city with the given name"

	ErrorCode400CityNameEnglishOnly    = "city_name_english_only"
	ErrorMessage400CityNameEnglishOnly = "The city name must be English    "
)
