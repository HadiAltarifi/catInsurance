/*
 * Cat Insurance API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ContractRes struct {

	Id string `json:"id"`

	StartDate string `json:"startDate"`

	EndDate string `json:"endDate"`

	Coverage float64 `json:"coverage"`

	CatName string `json:"catName"`

	Breed string `json:"breed"`

	Color string `json:"color"`

	BirthDate string `json:"birthDate"`

	Neutered bool `json:"neutered"`

	Personality string `json:"personality"`

	Environment string `json:"environment"`
	// In Gramm
	Weight float64 `json:"weight"`

	CustomerId string `json:"customerId"`
}
