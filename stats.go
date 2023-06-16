package aifiver

import "log"

// Opinion by type
// - Same trait
// - Opposite trait
// - Attraction
type Opinion int

const (
	TOpinionSame       Opinion = iota // Same trait
	TOpinionOpposite                  // Opposite trait
	TOpinionAttraction                // Attraction
)

// String implements the stringer function for an opinion.
func (o Opinion) String() string {
	switch o {
	case TOpinionSame:
		return "same trait"
	case TOpinionOpposite:
		return "opposite trait"
	case TOpinionAttraction:
		return "attraction"
	}
	return "unknown"
}

// Fate that can befall an individual
// - Early death
// - Capture in battle
// - Death in battle
// - Accidental injury
// - Bad investment
// - Scheme fallacy
// - Scheme discovery
type Fate int

const (
	TFateEarlyDeath      Fate = iota // Early death
	TFateCapture                     // Capture in battle
	TFateDeath                       // Death in battle
	TFateAccident                    // Accidental injury
	TFateInvestment                  // Bad investment
	TFateSchemeFallacy               // Scheme fallacy
	TFateSchemeDiscovery             // Scheme discovery
)

// String implements the stringer function for a fate.
func (f Fate) String() string {
	switch f {
	case TFateEarlyDeath:
		return "early death"
	case TFateCapture:
		return "capture"
	case TFateDeath:
		return "death"
	case TFateAccident:
		return "accident"
	case TFateInvestment:
		return "investment"
	case TFateSchemeFallacy:
		return "scheme fallacy"
	case TFateSchemeDiscovery:
		return "scheme discovery"
	}
	return "unknown"
}

// AIMod influences a specific aspect of decision making
// - AI Boldness
// - AI Compassion
// - AI Energy
// - AI Greed
// - AI Honor
// - AI Rationality
// - AI Sociability
// - AI Vengefulness
type AIMod int

const (
	TAIModBoldness     AIMod = iota // AI Boldness
	TAIModCompassion                // AI Compassion
	TAIModEnergy                    // AI Energy
	TAIModGreed                     // AI Greed
	TAIModHonor                     // AI Honor
	TAIModRationality               // AI Rationality
	TAIModSociability               // AI Sociability
	TAIModVengefulness              // AI Vengefulness
)

// String implements the stringer function for an ai modifier.
func (a AIMod) String() string {
	switch a {
	case TAIModBoldness:
		return "boldness"
	case TAIModCompassion:
		return "compassion"
	case TAIModEnergy:
		return "energy"
	case TAIModGreed:
		return "greed"
	case TAIModHonor:
		return "honor"
	case TAIModRationality:
		return "rationality"
	case TAIModSociability:
		return "sociability"
	case TAIModVengefulness:
		return "vengefulness"
	}
	return "unknown"
}

type Stats struct {
	Opinion map[Opinion]int // Opinion modifiers
	Fate    map[Fate]int    // Chance modifier for fates
	Skill   map[Skill]int   // Skill modifiers
	AI      map[AIMod]int   // AI modifiers
}

func NewStats() *Stats {
	return &Stats{
		Opinion: make(map[Opinion]int),
		Fate:    make(map[Fate]int),
		Skill:   make(map[Skill]int),
		AI:      make(map[AIMod]int),
	}
}

func (s *Stats) add(st *Stats) {
	for key, val := range st.Opinion {
		s.Opinion[key] += val
	}
	for key, val := range st.Fate {
		s.Fate[key] += val
	}
	for key, val := range st.Skill {
		s.Skill[key] += val
	}
	for key, val := range st.AI {
		s.AI[key] += val
	}
}

// Log logs the stats.
func (s *Stats) Log() {
	log.Println("Stats:")
	for key, val := range s.Opinion {
		log.Printf("  Opinion %s: %v\n", key, val)
	}
	for key, val := range s.Fate {
		log.Printf("  Fate %s: %v\n", key, val)
	}
	for key, val := range s.Skill {
		log.Printf("  Skill %s: %v\n", key, val)
	}
	for key, val := range s.AI {
		log.Printf("  AI %s: %v\n", key, val)
	}
}
