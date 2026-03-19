package main

// PageData is the view model passed to the HTML template.
// It stores the latest form values plus either a render result or an error
// so the page can be re-rendered after every submission.
type PageData struct {
	Input  string
	Banner string
	Result string
	Error  string
	Color  string
}

// asciiArtRequest is shared by both the form handler and the JSON API.
// The HTML route only fills a subset of these fields, while the API can
// provide the full payload including alignment and substring coloring.
type asciiArtRequest struct {
	Text           string `json:"text"`
	Banner         string `json:"banner"`
	Align          string `json:"align"`
	Color          string `json:"color"`
	ColorSubstring string `json:"color_substring"`
}

// asciiArtResponse is the JSON envelope returned by the API route.
// A successful response uses Result, while failures use Error.
type asciiArtResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// appError lets the shared service layer return an HTTP status together with
// the user-facing message that handlers should render or encode as JSON.
type appError struct {
	Status  int
	Message string
}
