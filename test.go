package godor

type Request struct {
	Header  string `json:"header" godor:"in=header"`
	Path    string `json:"path" godor:"in=path,required"`
	Query   string `json:"query" godor:"in=query"`
	Body    int64  `json:"body" godor:"in=body,required"`
	NoJson  string `godor:"in=query"`
	Invalid string `godor:"no_in_or_out"`

	Nohodor    int
	unexported bool
}

type request struct {
	Header  string `json:"header" godor:"in=header"`
	Path    string `json:"path" godor:"in=path,required"`
	Query   string `json:"query" godor:"in=query"`
	Body    int64  `json:"body" godor:"in=body,required"`
	NoJson  string `godor:"in=query"`
	Invalid string `godor:"no_in_or_out"`

	Nohodor    int
	unexported bool
}

type Response struct {
	Header  string `json:"header" godor:"out=header"`
	Path    string `json:"path" godor:"out=path,required"`
	Query   string `json:"query" godor:"out=query"`
	Body    int64  `json:"body" godor:"out=body,required"`
	NoJson  string `godor:"out=query"`
	Invalid string `godor:"no_in_or_out"`

	Nohodor    int
	unexported bool
}

type ReqRes struct {
	Header  string `json:"header" godor:"out=header"`
	Path    string `json:"path" godor:"out=path,required"`
	Query   string `json:"query" godor:"out=query"`
	Body    int64  `json:"body" godor:"out=body,required"`
	NoJson  string `godor:"out=query"`
	Invalid string `godor:"no_in_or_out"`

	Mix        string `godor:"in=path,out=path"`
	Nohodor    int
	unexported bool
}
