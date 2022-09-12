package main

import "time"

type UserPayload struct {
	User struct {
		ID           int    `json:"id"`
		Email        string `json:"email"`
		IsAdmin      int    `json:"is_admin"`
		HasLogged    int    `json:"has_logged"`
		Status       int    `json:"status"`
		UserSettings struct {
			UserID                    int    `json:"user_id"`
			Name                      string `json:"name"`
			Phone                     string `json:"phone"`
			Status                    int    `json:"status"`
			ShortDateFormat           string `json:"short_date_format"`
			ShortTimeFormat           string `json:"short_time_format"`
			DecimalSeparators         string `json:"decimal_separators"`
			ThousandsSeparators       string `json:"thousands_separators"`
			DistanceUnit              string `json:"distance_unit"`
			Language                  string `json:"language"`
			Country                   string `json:"country"`
			Timezone                  string `json:"timezone"`
			VolumetricMeasurementUnit int    `json:"volumetric_measurement_unit"`
			Currency                  string `json:"currency"`
		} `json:"user_settings"`
		OrganizationSettings struct {
			Currency       string      `json:"currency"`
			Country        string      `json:"country"`
			ConsultantUser interface{} `json:"consultant_user"`
		} `json:"organization_settings"`
		NeedsToAnswerSuccessMeasurement int    `json:"needs_to_answer_success_measurement"`
		OrganizationID                  int    `json:"organization_id"`
		OrganizationStatus              int    `json:"organization_status"`
		RoleID                          int    `json:"role_id"`
		CompanyName                     string `json:"company_name"`
		Vehicles                        []int  `json:"vehicles"`
		Groups                          []int  `json:"groups"`
		Clients                         []int  `json:"clients"`
		Token                           string `json:"token"`
	} `json:"user"`
}

type GenerateResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type PrivateClaims struct {
	Email    string `json:"email"`
	UserInfo struct {
		User struct {
			ID           int    `json:"id"`
			Email        string `json:"email"`
			IsAdmin      int    `json:"is_admin"`
			HasLogged    int    `json:"has_logged"`
			Status       int    `json:"status"`
			UserSettings struct {
				UserID                    int       `json:"user_id"`
				Name                      string    `json:"name"`
				Phone                     string    `json:"phone"`
				Status                    int       `json:"status"`
				ShortDateFormat           string    `json:"short_date_format"`
				ShortTimeFormat           string    `json:"short_time_format"`
				DecimalSeparators         string    `json:"decimal_separators"`
				ThousandsSeparators       string    `json:"thousands_separators"`
				DistanceUnit              string    `json:"distance_unit"`
				Language                  string    `json:"language"`
				Country                   string    `json:"country"`
				Timezone                  string    `json:"timezone"`
				VolumetricMeasurementUnit int       `json:"volumetric_measurement_unit"`
				Created                   time.Time `json:"created"`
				Modified                  time.Time `json:"modified"`
				Currency                  string    `json:"currency"`
			} `json:"user_settings"`
			OrganizationSettings struct {
				Currency       string `json:"currency"`
				Country        string `json:"country"`
				ConsultantUser string `json:"consultant_user"`
			} `json:"organization_settings"`
			OrganizationID     int           `json:"organization_id"`
			OrganizationStatus int           `json:"organization_status"`
			RoleID             int           `json:"role_id"`
			CompanyName        string        `json:"company_name"`
			Vehicles           []int         `json:"vehicles"`
			Groups             []interface{} `json:"groups"`
			Clients            []interface{} `json:"clients"`
		} `json:"user"`
	} `json:"userInfo"`
}

type ClaimSet struct {
	Iss   string `json:"iss"`             // email address of the client_id of the application making the access token request
	Scope string `json:"scope,omitempty"` // space-delimited list of the permissions the application requests
	Aud   string `json:"aud"`             // descriptor of the intended target of the assertion (Optional).
	Exp   int64  `json:"exp"`             // the expiration time of the assertion (seconds since Unix epoch)
	Iat   int64  `json:"iat"`             // the time the assertion was issued (seconds since Unix epoch)
	Typ   string `json:"typ,omitempty"`   // token type (Optional).

	// Email for which the application is requesting delegated access (Optional).
	Sub string `json:"sub,omitempty"`

	// The old name of Sub. Client keeps setting Prn to be
	// complaint with legacy OAuth 2.0 providers. (Optional)
	Prn string `json:"prn,omitempty"`

	// See http://tools.ietf.org/html/draft-jones-json-web-token-10#section-4.3
	// This array is marshalled using custom code (see (c *ClaimSet) encode()).
	PrivateClaims map[string]interface{} `json:"-"`
}

type TokenDecoded struct {
	Iss      string `json:"iss"`
	Aud      string `json:"aud"`
	Exp      int    `json:"exp"`
	Iat      int    `json:"iat"`
	Sub      string `json:"sub"`
	Email    string `json:"email"`
	UserInfo struct {
		User struct {
			ID           int    `json:"id"`
			Email        string `json:"email"`
			IsAdmin      int    `json:"is_admin"`
			HasLogged    int    `json:"has_logged"`
			Status       int    `json:"status"`
			UserSettings struct {
				UserID                    int    `json:"user_id"`
				Name                      string `json:"name"`
				Phone                     string `json:"phone"`
				Status                    int    `json:"status"`
				ShortDateFormat           string `json:"short_date_format"`
				ShortTimeFormat           string `json:"short_time_format"`
				DecimalSeparators         string `json:"decimal_separators"`
				ThousandsSeparators       string `json:"thousands_separators"`
				DistanceUnit              string `json:"distance_unit"`
				Language                  string `json:"language"`
				Country                   string `json:"country"`
				Timezone                  string `json:"timezone"`
				VolumetricMeasurementUnit int    `json:"volumetric_measurement_unit"`
				Currency                  string `json:"currency"`
			} `json:"user_settings"`
			OrganizationSettings struct {
				Currency       string      `json:"currency"`
				Country        string      `json:"country"`
				ConsultantUser interface{} `json:"consultant_user"`
			} `json:"organization_settings"`
			NeedsToAnswerSuccessMeasurement int    `json:"needs_to_answer_success_measurement"`
			OrganizationID                  int    `json:"organization_id"`
			OrganizationStatus              int    `json:"organization_status"`
			RoleID                          int    `json:"role_id"`
			CompanyName                     string `json:"company_name"`
			Vehicles                        []int  `json:"vehicles"`
			Groups                          []int  `json:"groups"`
			Clients                         []int  `json:"clients"`
			Token                           string `json:"token"`
		} `json:"user"`
	} `json:"userInfo"`
}
