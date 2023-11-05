package functions

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

type SpotifyResponse struct {
	IsActive     bool   `json:"is_active"`
	Type         string `json:"type"`
	ShuffleState bool   `json:"shuffle_state"`
	RepeatState  string `json:"repeat_state"`
	IsPlaying    bool   `json:"is_playing"`
	TimeStamp    int64  `json:"timestamp"`
	Song         string `json:"song"`
	Progress     struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"progress"`
	Artists    []Artist     `json:"artists"`
	ProgressMs int          `json:"progress_ms"`
	DurationMs int          `json:"duration_ms"`
	Image      Image        `json:"image"`
	Url        string       `json:"url"`
	ReqTime    string       `json:"reqTime"`
	Queue      []QueuedSong `json:"queue"`
}

type QueuedSong struct {
	Name    string `json:"name"`
	Artists string `json:"artists"`
	Image   Image  `json:"image"`
}

type Image struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Artist struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type UserPlayingResponse struct {
	Device struct {
		ID               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	} `json:"device"`
	ShuffleState bool   `json:"shuffle_state"`
	RepeatState  string `json:"repeat_state"`
	Timestamp    int    `json:"timestamp"`
	Context      struct {
		ExternalURLs struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"context"`
	ProgressMs int `json:"progress_ms"`
	Item       struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalURLs struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			AvailableMarkets []string `json:"available_markets"`
			ExternalURLs     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalURLs struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIDs      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalURLs struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		Disallows struct {
			Resuming bool `json:"resuming"`
		} `json:"disallows"`
	} `json:"actions"`
	IsPlaying bool `json:"is_playing"`
}

type UserPlayerQueueResponse struct {
	Currently_Playing struct {
		Album struct {
			Album_Type string `json:"album_type"`
			Artists    []struct {
				External_Urls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			Available_Markets []string `json:"available_markets"`
			External_Urls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                   string `json:"name"`
			Release_Date           string `json:"release_date"`
			Release_Date_Precision string `json:"release_date_precision"`
			Total_Tracks           int    `json:"total_tracks"`
			Type                   string `json:"type"`
			URI                    string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			External_Urls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		Available_Markets []string `json:"available_markets"`
		Disc_Number       int      `json:"disc_number"`
		Duration_Ms       int      `json:"duration_ms"`
		Explicit          bool     `json:"explicit"`
		External_Ids      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		External_Urls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href         string `json:"href"`
		ID           string `json:"id"`
		Is_Local     bool   `json:"is_local"`
		Name         string `json:"name"`
		Popularity   int    `json:"popularity"`
		Preview_Url  string `json:"preview_url"`
		Track_Number int    `json:"track_number"`
		Type         string `json:"type"`
		URI          string `json:"uri"`
	} `json:"currently_playing"`
	Queue []struct {
		Album struct {
			Album_Type string `json:"album_type"`

			Artists []struct {
				External_Urls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			Available_Markets []string `json:"available_markets"`
			External_Urls     struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                   string `json:"name"`
			Release_Date           string `json:"release_date"`
			Release_Date_Precision string `json:"release_date_precision"`
			Total_Tracks           int    `json:"total_tracks"`
			Type                   string `json:"type"`
			URI                    string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			External_Urls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		Available_Markets []string `json:"available_markets"`
		Disc_Number       int      `json:"disc_number"`
		Duration_Ms       int      `json:"duration_ms"`
		Explicit          bool     `json:"explicit"`
		External_Ids      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		External_Urls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href         string `json:"href"`
		ID           string `json:"id"`
		Is_Local     bool   `json:"is_local"`
		Name         string `json:"name"`
		Popularity   int    `json:"popularity"`
		Preview_Url  string `json:"preview_url"`
		Track_Number int    `json:"track_number"`
		Type         string `json:"type"`
		URI          string `json:"uri"`
	} `json:"queue"`
}
