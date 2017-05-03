/* 
 * Uber API
 *
 * Move your app forward with the Uber API
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Profile struct {

	// First name of the Uber user.
	FirstName string `json:"first_name,omitempty"`

	// Last name of the Uber user.
	LastName string `json:"last_name,omitempty"`

	// Email address of the Uber user
	Email string `json:"email,omitempty"`

	// Image URL of the Uber user.
	Picture string `json:"picture,omitempty"`

	// Promo code of the Uber user.
	PromoCode string `json:"promo_code,omitempty"`
}