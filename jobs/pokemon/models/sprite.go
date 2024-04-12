package models

type DreamWorld struct {
	FrontDefault string      `json:"front_default"`
	FrontFemale  interface{} `json:"front_female"`
}

type Home struct {
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
	FrontShiny   string `json:"front_shiny"`
}

type Showdown struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
}

type SpriteOther struct {
	DreamWorld      DreamWorld      `json:"dream_world"`
	Home            Home            `json:"home"`
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
	Showdown        Showdown        `json:"showdown"`
}

type Sprites struct {
	BackDefault      string      `json:"back_default"`
	BackFemale       interface{} `json:"back_female"`
	BackShiny        string      `json:"back_shiny"`
	BackShinyFemale  interface{} `json:"back_shiny_female"`
	FrontDefault     string      `json:"front_default"`
	FrontFemale      interface{} `json:"front_female"`
	FrontShiny       string      `json:"front_shiny"`
	FrontShinyFemale interface{} `json:"front_shiny_female"`
	Other            SpriteOther `json:"other"`
	Versions         struct {
		GenerationI struct {
			RedBlue struct {
				BackDefault      string `json:"back_default"`
				BackGray         string `json:"back_gray"`
				BackTransparent  string `json:"back_transparent"`
				FrontDefault     string `json:"front_default"`
				FrontGray        string `json:"front_gray"`
				FrontTransparent string `json:"front_transparent"`
			} `json:"red-blue"`
			Yellow struct {
				BackDefault      string `json:"back_default"`
				BackGray         string `json:"back_gray"`
				BackTransparent  string `json:"back_transparent"`
				FrontDefault     string `json:"front_default"`
				FrontGray        string `json:"front_gray"`
				FrontTransparent string `json:"front_transparent"`
			} `json:"yellow"`
		} `json:"generation-i"`
		GenerationIi struct {
			Crystal struct {
				BackDefault           string `json:"back_default"`
				BackShiny             string `json:"back_shiny"`
				BackShinyTransparent  string `json:"back_shiny_transparent"`
				BackTransparent       string `json:"back_transparent"`
				FrontDefault          string `json:"front_default"`
				FrontShiny            string `json:"front_shiny"`
				FrontShinyTransparent string `json:"front_shiny_transparent"`
				FrontTransparent      string `json:"front_transparent"`
			} `json:"crystal"`
			Gold struct {
				BackDefault      string `json:"back_default"`
				BackShiny        string `json:"back_shiny"`
				FrontDefault     string `json:"front_default"`
				FrontShiny       string `json:"front_shiny"`
				FrontTransparent string `json:"front_transparent"`
			} `json:"gold"`
			Silver struct {
				BackDefault      string `json:"back_default"`
				BackShiny        string `json:"back_shiny"`
				FrontDefault     string `json:"front_default"`
				FrontShiny       string `json:"front_shiny"`
				FrontTransparent string `json:"front_transparent"`
			} `json:"silver"`
		} `json:"generation-ii"`
		GenerationIii struct {
			Emerald struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"emerald"`
			FireredLeafgreen struct {
				BackDefault  string `json:"back_default"`
				BackShiny    string `json:"back_shiny"`
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"firered-leafgreen"`
			RubySapphire struct {
				BackDefault  string `json:"back_default"`
				BackShiny    string `json:"back_shiny"`
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"ruby-sapphire"`
		} `json:"generation-iii"`
		GenerationIv struct {
			DiamondPearl struct {
				BackDefault      string      `json:"back_default"`
				BackFemale       interface{} `json:"back_female"`
				BackShiny        string      `json:"back_shiny"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"diamond-pearl"`
			HeartgoldSoulsilver struct {
				BackDefault      string      `json:"back_default"`
				BackFemale       interface{} `json:"back_female"`
				BackShiny        string      `json:"back_shiny"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"heartgold-soulsilver"`
			Platinum struct {
				BackDefault      string      `json:"back_default"`
				BackFemale       interface{} `json:"back_female"`
				BackShiny        string      `json:"back_shiny"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"platinum"`
		} `json:"generation-iv"`
		GenerationV struct {
			BlackWhite struct {
				Animated struct {
					BackDefault      string      `json:"back_default"`
					BackFemale       interface{} `json:"back_female"`
					BackShiny        string      `json:"back_shiny"`
					BackShinyFemale  interface{} `json:"back_shiny_female"`
					FrontDefault     string      `json:"front_default"`
					FrontFemale      interface{} `json:"front_female"`
					FrontShiny       string      `json:"front_shiny"`
					FrontShinyFemale interface{} `json:"front_shiny_female"`
				} `json:"animated"`
				BackDefault      string      `json:"back_default"`
				BackFemale       interface{} `json:"back_female"`
				BackShiny        string      `json:"back_shiny"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"black-white"`
		} `json:"generation-v"`
		GenerationVi struct {
			OmegarubyAlphasapphire struct {
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"omegaruby-alphasapphire"`
			XY struct {
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"x-y"`
		} `json:"generation-vi"`
		GenerationVii struct {
			Icons struct {
				FrontDefault string      `json:"front_default"`
				FrontFemale  interface{} `json:"front_female"`
			} `json:"icons"`
			UltraSunUltraMoon struct {
				FrontDefault     string      `json:"front_default"`
				FrontFemale      interface{} `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale interface{} `json:"front_shiny_female"`
			} `json:"ultra-sun-ultra-moon"`
		} `json:"generation-vii"`
		GenerationViii struct {
			Icons struct {
				FrontDefault string      `json:"front_default"`
				FrontFemale  interface{} `json:"front_female"`
			} `json:"icons"`
		} `json:"generation-viii"`
	} `json:"versions"`
}
