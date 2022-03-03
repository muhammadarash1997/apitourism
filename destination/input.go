package destination

type DestinationInput struct {
	Name    string `json:"name" binding:"required"`
	Type    string `json:"type" binding:"required"`
	GeoHash string `json:"geo_hash" binding:"required"`
}

// Perlu menggunakan tag json agar saat pengambilan payload dengan fungsi ShouldBindJSON tidak error