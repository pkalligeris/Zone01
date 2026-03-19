package main

type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
	Color  string
}

type asciiArtRequest struct {
	Text           string `json:"text"`
	Banner         string `json:"banner"`
	Align          string `json:"align"`
	Color          string `json:"color"`
	ColorSubstring string `json:"color_substring"`
}

type asciiArtResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

type appError struct {
	Status  int
	Message string
}
