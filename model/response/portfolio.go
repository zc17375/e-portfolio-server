package response

type PortfolioListResponse struct {
	Name        string `json:"name" bson:"name"`
	HeadImgPath string `json:"head_img_path" bson:"head_img_path"`
}
