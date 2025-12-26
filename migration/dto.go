package migration

type newsData struct {
	Data data `json:"data"`
}

type data struct {
	NewsList []news `json:"news_list"`
}

type news struct {
	NewsObject newsObject `json:"news_obj"`
}

type newsObject struct {
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	SourceURL       string   `json:"source_url"`
	CreatedAt       int64    `json:"created_at"`
	SourceName      string   `json:"source_name"`
	CategoryNames   []string `json:"category_names"`
	ImpressiveScore float64  `json:"impressive_score"`
	// Latitude        float64   `json:"latitude"`
	// Longitude       float64   `json:"longitude"`
}
