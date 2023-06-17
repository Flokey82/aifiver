package aifiver

// Modifier is a personality modifier which can be applied to a personality.
// TODO:
// - Add persistance to indicate how long this would last.
// - Should this maybe amplify pre-existing traits instead of simply modifying them?
// - Should we derive the factor modifiers from the facet modifiers?
type Modifier struct {
	Name    string  // Name of the modifier
	Factors [5]int  // Modifiers for the five factors
	Facets  [30]int // Modifiers for the 30 facets
}

var (
	ModRecentBetrayal = &Modifier{
		Name: "recent betrayal",
		Factors: [5]int{
			0,  // Openness
			0,  // Conscientiousness
			0,  // Extraversion
			-1, // Agreeableness
			1,  // Neuroticism
		},
		Facets: [30]int{
			0, 0, 0, 0, 0, // Openness
			0, 0, 0, 0, 0, // Conscientiousness
			0, 0, 0, 0, -1, // Extraversion (Positive Emotions)
			-5, -1, 0, 0, 0, // Agreeableness (Trust, Straightforwardness)
			3, 1, 1, 1, 0, // Neuroticism (Anxiety, Angry Hostility, Depression, Self-Consciousness)
		},
	}
	ModGriefing = &Modifier{
		Name: "griefing",
		Factors: [5]int{
			0,  // Openness
			0,  // Conscientiousness
			-1, // Extraversion
			0,  // Agreeableness
			1,  // Neuroticism
		},
		Facets: [30]int{
			0, 0, 0, 0, 0, // Openness
			0, 0, 0, 0, 0, // Conscientiousness
			0, 0, 0, 0, -2, // Extraversion (Positive Emotions)
			0, 0, 0, 0, 0, // Agreeableness
			0, 1, 2, 0, 0, // Neuroticism (Anger, Depression)
		},
	}
	ModSetback = &Modifier{
		Name: "setback",
		Factors: [5]int{
			0, // Openness
			1, // Conscientiousness
			0, // Extraversion
			0, // Agreeableness
			1, // Neuroticism
		},
		Facets: [30]int{
			0, 0, 0, 0, 0, // Openness
			0, 0, 0, 0, 2, // Conscientiousness (Cautiousness)
			0, 0, 0, 0, 0, // Extraversion
			0, 0, 0, 0, 0, // Agreeableness
			0, 2, 1, 2, 0, // Neuroticism (Anger, Depression, Self-Consciousness)
		},
	}
)
