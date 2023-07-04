package model

type Address struct {
	Street  string       `json:"street"`
	Suite   string       `json:"suite"`
	City    string       `json:"city"`
	Zipcode string       `json:"zipcode"`
	Geo     *GeoLocation `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type GeoLocation struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type User struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Address  *Address `json:"address"`
	Phone    string   `json:"phone"`
	Website  string   `json:"website"`
	Company  *Company `json:"company"`
}
