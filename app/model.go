package main

type experience struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Timeframe string `json:"timeframe"`
}

func getData() []experience {
	experiences := []experience{
		{
			Role:      "Solution Architect",
			Company:   "Red Bull",
			City:      "Salzburg",
			Timeframe: "April 2020 - now",
		},
		{
			Role:      "Senior Software Engineer",
			Company:   "FLYERALARM",
			City:      "Würzburg",
			Timeframe: "June 2018 - April 2020",
		},
		{
			Role:      "Lead Developer",
			Company:   "FLYERALARM",
			City:      "Würzburg",
			Timeframe: "October 2017 - June 2018",
		},
		{
			Role:      "Software Engineer",
			Company:   "FLYERALARM",
			City:      "Würzburg",
			Timeframe: "August 2016 - October 2017",
		},
	}

	return experiences
}
