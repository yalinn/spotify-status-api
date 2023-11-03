package functions

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/url"

	"github.com/tantoony/spotify-status-api-golang/config"
)

func FetchSpotifyToken(code string) error {
	jsonBody := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {"http://localhost:3333/auth/spotify"},
		"client_id":     {config.SPOTIFY_CLIENT_ID},
		"client_secret": {config.SPOTIFY_CLIENT_SECRET},
	}.Encode()
	var jsonData = []byte(jsonBody)
	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBuffer(jsonData))
	auth_type_input := "Basic " + client_token()
	request.Header.Set("Authorization-Type", auth_type_input)
	return err
}

func bufferToBase64(buf *bytes.Buffer) string {
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func client_token() string {
	var strBuffer bytes.Buffer
	strBuffer.WriteString(config.SPOTIFY_CLIENT_ID)
	strBuffer.WriteString(":")
	strBuffer.WriteString(config.SPOTIFY_CLIENT_SECRET)
	return bufferToBase64(&strBuffer)
}
