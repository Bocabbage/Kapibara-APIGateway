package mikanani

type UpdateAnimeFormat struct {
	Name        string `json:"name" binding:"required"`
	RssUrl      string `json:"rss_url" binding:"required"`
	RuleVersion string `json:"rule_version" binding:"required"`
	RuleRegex   string `json:"rule_regex" binding:"required"`
	IsActive    bool   `json:"is_active" binding:"required"`
}

// ----- v2
type AnimeDocFormat struct {
	Uid    int64  `json:"uid"`
	RssUrl string `json:"rss_url"`
	Rule   string `json:"rule"`
	Regex  string `json:"regex"`
}

type AnimeMetaFormat struct {
	Uid            int64    `json:"uid"`
	Name           string   `json:"name"`
	DownloadBitmap int64    `json:"download_bitmap"`
	IsActive       bool     `json:"is_active"`
	Tags           []string `json:"tags"`
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
	Uid int64 `json:"uid" binding:"required"`
}

type UpdateAnimeMetaFormat struct {
	AnimeMetaFormat
	Uid int64 `json:"uid" binding:"required"`
}

type InsertAnimeItemFormat struct {
	Name   string   `json:"name" binding:"required"`
	Tags   []string `json:"tags"`
	RssUrl string   `json:"rss_url" binding:"required"`
	Rule   string   `json:"rule" binding:"required"`
	Regex  string   `json:"regex" binding:"required"`
}

type DeleteAnimeItemFormat struct {
	Uid int64 `form:"uid" binding:"required"`
}
