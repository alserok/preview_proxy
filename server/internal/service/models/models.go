package models

type DownloadThumbnailsReq struct {
	VideoURLs []string `json:"video_urls"`
	Async     bool     `json:"async"`
}

type DownloadThumbnailsRes struct {
	Failed uint32  `json:"failed"`
	Total  uint32  `json:"total"`
	Videos []Video `json:"videos"`
}

type Video struct {
	VideoURL     string `json:"video_url"`
	ThumbnailURL string `json:"thumbnail_url"`
}
