type (
	youtubeVideo struct {
		Entry ytEntry `json:"entry"`
	}

	ytEntry struct {
		Title      ytTitle      `json:"title"`
		MediaGroup ytMediaGroup `json:"media$group"`
	}

	ytMediaGroup struct {
		MediaThumbnail []ytMediaThumbnail `json:"media$thumbnail"`
	}

	ytMediaThumbnail struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"with"`
		Name   string `json:"yt$name"`
	}

	ytTitle struct {
		T string `json:"$t"`
	}
)