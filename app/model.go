package main

type experience struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Timeframe string `json:"timeframe"`
}

type climateData struct {
	Heading string `json:"heading"`
	Content string `json:"content"`
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

func getClimateData() []climateData {
	climateData := []climateData{
		{
			Heading: "Klimaschutz",
			Content: "Für eine ambitionierte Klimapolitik einsetzen, um katastrophale Auswirkungen des Klimawandels zu verhindern.",
		},
		{
			Heading: "Energiewende",
			Content: "Forderung einer raschen Umstellung auf erneuerbare Energien, um fossile Brennstoffe zu reduzieren und eine nachhaltige Zukunft zu gewährleisten.",
		},
		{
			Heading: "Grünes Wachstum",
			Content: "Ich sehe den Klimawandel als Chance für grünes Wachstum und es sollen Investitionen in nachhaltige Technologien und Arbeitsplätze gefördert werden.",
		},
		{
			Heading: "Mobilität",
			Content: "Anstreben einer schrittweise Abschaffung von Verbrennungsmotoren zugunsten von emissionsfreien Alternativen wie Elektroantrieben.",
		},
		{
			Heading: "Heizen",
			Content: "Für die Umstellung von fossilen Brennstoffen auf erneuerbare Energien in der Gebäudeheizung einsetzen, um den CO2-Ausstoß zu reduzieren und die Energiewende voranzutreiben.",
		},
	}

	return climateData
}
