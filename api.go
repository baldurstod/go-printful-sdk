package printfulsdk

type RequestBodyKey int64

const (
	FileRole RequestBodyKey = iota
	URL
	Filename
	FileVisible
)

func BuildRequestBody(o options, keys ...RequestBodyKey) map[string]interface{} {
	body := map[string]interface{}{}

	for _, key := range keys {
		switch key {
		case FileRole:
			if o.fileRole != "" {
				body["role"] = o.fileRole
			}
		case URL:
			body["url"] = o.url
		case Filename:
			if o.filename != "" {
				body["filename"] = o.filename
			}
		case FileVisible:
			body["visible"] = o.fileVisible
		}
	}
	return body
}
