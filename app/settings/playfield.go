package settings

var Playfield = initPlayfield()

func initPlayfield() *playfield {
	return &playfield{
		DrawObjects: true,
		DrawCursors: true,
		Scale:       1,
		OsuShift:    false,
		LeadInTime:  5,
		LeadInHold:  2,
		FadeOutTime: 5,
		Seizure: &seizure{
			Enabled:  true,
			Duration: 5,
		},
		Background: &background{
			LoadStoryboards: true,
			FlashToTheBeat:  true,
			Dim: &dim{
				Intro:  0,
				Normal: 0.95,
				Breaks: 0.5,
			},
			Parallax: &parallax{
				Amount: 0.1,
				Speed:  0.5,
			},
			Blur: &blur{
				Enabled: true,
				Values: &dim{
					Intro:  0,
					Normal: 0.6,
					Breaks: 0.3,
				},
			},
		},
		Logo: &logo{
			DrawSpectrum: false,
			Dim: &dim{
				Intro:  0,
				Normal: 1,
				Breaks: 1,
			},
		},
		Bloom: &bloom{
			Enabled:           true,
			BloomToTheBeat:    true,
			BloomBeatAddition: 0.3,
			Threshold:         0.0,
			Blur:              0.6,
			Power:             0.7,
		},
	}
}

type playfield struct {
	DrawObjects bool
	DrawCursors bool
	Scale       float64 //1, scale the playfield (1 means that 384 will be rescaled to 900 on FullHD monitor)
	OsuShift    bool    //false, offset the playfield like in osu!
	LeadInTime  float64 //5
	LeadInHold  float64 //2
	FadeOutTime float64 //5
	Seizure     *seizure
	Background  *background
	Logo        *logo
	Bloom       *bloom
}

type seizure struct {
	// Whether seizure warning should be displayed before intro
	Enabled bool

	Duration float64
}

// Background controls
type background struct {
	// Whether storyboards should be loaded
	LoadStoryboards bool

	FlashToTheBeat bool

	// Dim controls
	Dim *dim

	Parallax *parallax

	// Blur controls
	Blur *blur
}

type parallax struct {
	// Amount of parallax, also scales bg by (1+Amount), set to 0 to disable it
	Amount float64

	// Speed of parallax
	Speed float64
}

type blur struct {
	Enabled bool

	Values *dim
}

type logo struct {
	DrawSpectrum bool
	Dim          *dim
}

type dim struct {
	// Value before drain time start
	Intro float64

	// Value during drain time
	Normal float64

	// Value during breaks
	Breaks float64
}

type bloom struct {
	Enabled           bool
	BloomToTheBeat    bool
	BloomBeatAddition float64
	Threshold         float64
	Blur              float64
	Power             float64
}
