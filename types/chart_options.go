package types

type ChartOptions struct {
	Plugins struct {
		Legend struct {
			Display  bool   `json:"display"`
			Position string `json:"position"`
		} `json:"legend"`
		Title struct {
			Display bool   `json:"display"`
			Text    string `json:"text"`
			Align   string `json:"align"`
			Weight  string `json:"weight"`
			Padding struct {
				Top    int `json:"top"`
				Bottom int `json:"bottom"`
			} `json:"padding"`
		} `json:"title"`
		Tooltip struct {
			Mode      string `json:"mode"`
			Intersect bool   `json:"intersect"`
		} `json:"tooltip"`
	} `json:"plugins"`
	Scales struct {
		X struct {
			Stacked bool `json:"stacked"`
			Ticks   struct {
				Major struct {
					Enabled   bool   `json:"enabled"`
					FontStyle string `json:"fontStyle"`
				} `json:"major"`
				Source          string `json:"source"`
				AutoSkip        bool   `json:"autoSkip"`
				MaxRotation     int    `json:"maxRotation"`
				SampleSize      int    `json:"sampleSize"`
				BackdropPadding int    `json:"backdropPadding"`
			} `json:"ticks"`
		} `json:"x"`
		Y struct {
			Stacked bool `json:"stacked"`
		} `json:"y"`
	} `json:"scales"`
	Layout struct {
		Padding int `json:"padding"`
	} `json:"layout"`
}
