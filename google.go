package validate

type GoogleConfig struct {
	Token string `url:"id_token"`
}

type GoogleAuthenticateSuccessResponse struct {
	// Verify that it is issued by accounts.google.com or
	// https://accounts.google.com.
	IssuedBy string `json:"iss"`

	// An identifier for the user, unique among all Google accounts and never reused.
	// A Google account can have multiple emails at different points in time,
	// but the sub value is never changed.
	// Use sub within your application as the unique-identifier key for the user.
	UserID string `json:"sub"`

	// Client ID of Android
	AndroidAudience string `json:"azp"`

	// Client ID of web
	Audience string `json:"aud"`

	// Convert this to int64 at runtime.
	// Unix Timestamp with double quotes returned.
	// Eg: iat: "1433981153"
	IssuedAt string `json:"iat"`

	// Conver this to int64 at runtime.
	// Unix Timestamp with double quotes returned.
	// Eg: exp: "1433981953"
	ExpiresAt string `json:"exp"`

	// These seven fields are only included when the user has granted the "profile" and
	// "email" OAuth scopes to the application.
	Email string `json:"email"`

	//Convert this to bool at runtime.
	// email_verified :"true" or email_verified: "false"
	IsEmailVerified string `json:"email_verified"`

	// Name as displayed on Google profiles
	DisplayName string `json:"name"`

	// URL to the current display pic
	DisplayPicture string `json:"picture"`

	// Given name / First name of the user
	FirstName string `json:"given_name"`

	// Family name / Last name  of the user
	LastName string `json:"family_name"`

	// Eg: locale: "en"
	Locale string `json:"locale"`
}

func Google(config *GoogleConfig) (GoogleAuthenticateSuccessResponse, error) {
	// TODO
}
