package statuslog

type StatuslogBioData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
		Bio     string `json:"bio"`
		CSS     string `json:"css"`
		Head    string `json:"head"`
	} `json:"response"`
}


type StatuslogListData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message  string `json:"message"`
		Statuses []struct {
			ID           string `json:"id"`
			Address      string `json:"address"`
			Created      string `json:"created"`
			RelativeTime string `json:"relative_time"`
			Emoji        string `json:"emoji"`
			Content      string `json:"content"`
			ExternalURL  string `json:"external_url"`
		} `json:"statuses"`
	} `json:"response"`
}

type StatuslogRetrieveData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
		Status  struct {
			ID           string `json:"id"`
			Address      string `json:"address"`
			Created      string `json:"created"`
			RelativeTime string `json:"relative_time"`
			Emoji        string `json:"emoji"`
			Content      string `json:"content"`
			ExternalURL  string `json:"external_url"`
		} `json:"status"`
	} `json:"response"`
}

type NewStatusData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message     string `json:"message"`
		ID          string `json:"id"`
		Status      string `json:"status"`
		URL         string `json:"url"`
		ExternalURL string `json:"external_url"`
	} `json:"response"`
}
