package mikanani

type UpdateAnimeFormat struct {
	Name        string `json:"name" binding:"required"`
	RssUrl      string `json:"rss_url" binding:"required"`
	RuleVersion string `json:"rule_version" binding:"required"`
	RuleRegex   string `json:"rule_regex" binding:"required"`
	IsActive    bool   `json:"is_active" binding:"required"`
}
