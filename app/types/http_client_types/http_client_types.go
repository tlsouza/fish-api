package http_client

type ImgSrcSet map[string]string

type Meta struct {
	ScientificClassification struct {
		Domain    string `json:"domain"`
		Kingdom   string `json:"kingdom"`
		Phylum    string `json:"phylum"`
		Class     string `json:"class"`
		Order     string `json:"order"`
		Family    string `json:"family"`
		Subfamily string `json:"subfamily"`
	} `json:"scientific_classification"`
}

type FishExternalApi struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	ImgSrcSet ImgSrcSet `json:"img_src_set"`
	Meta      Meta      `json:"meta"`
}
