package googleauth_service

import (
	"context"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	helpers_json "go-api-echo/internal/pkg/helpers/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	googleOauth2 "google.golang.org/api/oauth2/v2"
)

// Replace your OAuth client ID and client secret here
var oauthConfig = &oauth2.Config{
	ClientID:     config.GlobalEnv.Google.ClientID,
	ClientSecret: config.GlobalEnv.Google.ClientSecret,
	Endpoint:     google.Endpoint,
}

func ValidateIdToken(idToken string) (*idtoken.Payload, *helpers_errors.Error) {
	payload, err := idtoken.Validate(context.Background(), idToken, ``)
	if err != nil {
		comErr := *&helpers_errors.UnauthorizedError
		comErr.AddError(`Invalid idtoken`)
		return nil, comErr
	}
	return payload, nil
}

func ValidateAccessToken(accessToken string) (*googleOauth2.Userinfo, *helpers_errors.Error) {
	token := oauth2.Token{
		AccessToken: accessToken,
		TokenType:   `Bearer`,
	}

	client := oauthConfig.Client(context.Background(), &token)

	// Fetch and print user info
	svc, err := googleOauth2.New(client)
	if err != nil {
		comErr := *&helpers_errors.InternalServerError
		comErr.AddError(`Internal Server Error`)
		return nil, comErr
	}

	// Call the 'userinfo.v2.me.get' method to get user information.
	userinfo, err := svc.Userinfo.V2.Me.Get().Do()
	if err != nil {
		helpers_json.Print(err.Error())
		comErr := *&helpers_errors.UnauthorizedError
		comErr.AddError(`Failed to get user info`)
		return nil, comErr
	}

	return userinfo, nil
}
