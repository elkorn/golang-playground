
type (
	youtubeVideo struct {
		Entry      ytEntry      `json:"entry"`
		MediaGroup ytMediaGroup `json:"media$group"`
	}

	ytEntry struct {
		Title ytTitle `json:"title"`
	}

	ytMediaGroup struct {
		MediaThumbnail []ytMediaThumbnail `json:"media$thumbnail"`
	}

	ytMediaThumbnail struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
		Name   string `json:"yt$name"`
	}

	ytTitle struct {
		T string `json:"$t"`
	}
)