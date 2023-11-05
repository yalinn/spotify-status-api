package functions

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/tantoony/spotify-status-api-golang/config"
)

type AuthorizationResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	Error       string `json:"error"`
}

type UserMetaResponse struct {
	ID    string `json:"id"`
	Name  string `json:"display_name"`
	Error string `json:"error"`
}

type UserResponse struct {
	User UserMetaResponse `json:"user"`
}

func AuthorizeSpotify(code string) string {
	reqBody := url.Values{}
	reqBody.Set("grant_type", "authorization_code")
	reqBody.Set("code", code)
	reqBody.Add("redirect_uri", "http://127.0.0.1:3000/api/spotify")
	reqBody.Add("client_id", config.SPOTIFY_CLIENT_ID)
	reqBody.Add("client_secret", config.SPOTIFY_CLIENT_SECRET)
	encodedData := reqBody.Encode()
	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(encodedData))

	if err != nil {
		fmt.Print(err)
		return ""
	}
	auth_type_input := "Basic " + client_token()
	request.Header.Set("Authorization", auth_type_input)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response := responseString(request)
	return response
}

func FetchSpotifyUser(access_token string) string {
	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/me", nil)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	auth_type_input := "Bearer " + access_token
	request.Header.Set("Authorization", auth_type_input)
	response := responseString(request)
	return response
}

func client_token() string {
	return base64.StdEncoding.EncodeToString([]byte(config.SPOTIFY_CLIENT_ID + ":" + config.SPOTIFY_CLIENT_SECRET))
}

func responseString(request *http.Request) string {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return ""
}

func UserPlaying(access_token string) string {
	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	auth_type_input := "Bearer " + access_token
	request.Header.Set("Authorization", auth_type_input)
	response := responseString(request)
	return response
}

func UserQueue(access_token string) string {
	request, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/queue", nil)
	if err != nil {
		fmt.Print(err)
		return ""
	}
	auth_type_input := "Bearer " + access_token
	request.Header.Set("Authorization", auth_type_input)
	response := responseString(request)
	return response
}
