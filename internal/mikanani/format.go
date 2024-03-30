package mikanani

type UpdateAnimeFormat struct {
	Name        string `json:"name" binding:"required"`
	RssUrl      string `json:"rss_url" binding:"required"`
	RuleVersion string `json:"rule_version" binding:"required"`
	RuleRegex   string `json:"rule_regex" binding:"required"`
	IsActive    int32  `json:"is_active" binding:"required"`
}

// ----- v2
type AnimeDocFormat struct {
	Uid    int64  `json:"uid"`
	RssUrl string `json:"rss_url,omitempty"`
	Rule   string `json:"rule,omitempty"`
	Regex  string `json:"regex,omitempty"`
}

type AnimeMetaFormat struct {
	Uid            int64    `json:"uid"`
	Name           string   `json:"name,omitempty"`
	DownloadBitmap int64    `json:"download_bitmap,omitempty"`
	IsActive       int32    `json:"is_active,omitempty"`
	Tags           []string `json:"tags,omitempty"`
}

type ListAnimeMetaFormat struct {
	StartIdx     int64 `form:"start" binding:"required"`
	EndIdx       int64 `form:"end" binding:"required"`
	StatusFilter int32 `form:"status_filter" binding:"required"`
}

type GetAnimeDocFormat struct {
	Uid int64 `form:"uid" binding:"required"`
}

type UpdateAnimeDocFormat struct {
	AnimeDocFormat
	Uid string `json:"uid" binding:"required"`
}

type UpdateAnimeMetaFormat struct {
	AnimeMetaFormat
	Uid string `json:"uid" binding:"required"`
}

type InsertAnimeItemFormat struct {
	Name     string   `json:"name" binding:"required"`
	Tags     []string `json:"tags,omitempty"`
	RssUrl   string   `json:"rss_url" binding:"required"`
	Rule     string   `json:"rule" binding:"required"`
	Regex    string   `json:"regex" binding:"required"`
	IsActive int32    `json:"is_active" binding:"required"`
}

type DeleteAnimeItemFormat struct {
	Uid int64 `form:"uid" binding:"required"`
}
