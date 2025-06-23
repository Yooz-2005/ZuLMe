package request

type GetPresignedUrlRequest struct {
	Bucket      string `json:"bucket" form:"bucket"`
	ObjectName  string `json:"object_name"  form:"object_name"`
	Expires     int64  `json:"expires"  form:"expires"`
	ContentType string `json:"content_type"  form:"content_type"`
}
