package functions

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func fetchToken(code string) error {
	jsonBody := url.Values{
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {"http://localhost:3333/auth/spotify"},
		"client_id":     {"c0d0f0b0c0d0f0b0c0d0f0b0c0d0f0b0"},
		"client_secret": {"c0d0f0b0c0d0f0b0c0d0f0b0c0d0f0b0"},
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
	strBuffer.WriteString(os.Getenv("SPOTIFY_CLIENT_ID"))
	strBuffer.WriteString(":")
	strBuffer.WriteString(os.Getenv("SPOTIFY_CLIENT_SECRET"))
	return bufferToBase64(&strBuffer)
}

func getOs() error {
	osy := os.Getenv("CLIENT_ID")
	fmt.Println(osy)
	return nil
}

var asd = getOs()
