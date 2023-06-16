package aifiver

// Personality represents a personality that is a combination of facet scores and
// expressed traits.
type Personality struct {
	t         *Traiter // Traiter reference for re-evaluating the expressed traits
	BigModel           // The model which influences which traits are expressed
	Expressed []*Trait // Traits expressed based on the underlying 5 factor model.
	// TODO: Calculate compatibility with other personalities
	// TODO: Proneness to fall victim to cetain fates (accidental death, falling victim to intrigue)
	// TODO: Traits as expression of facets (Cruel, Sadistic, Diplomatic, ...)
	// TODO: Temporary facet modifiers (Griefing, Recent Betrayal, Setback, ...)
	// Modifiers []*Modifier
}

// NewPersonality returns a new, pretty boring personality.
func NewPersonality(t *Traiter) *Personality {
	return &Personality{
		t: t,
	}
}

// NewPersonalityFromPreset returns a new personality based on the given preset.
func NewPersonalityFromPreset(t *Traiter, ps map[Facet]int) *Personality {
	p := NewPersonality(t)
	for key, val := range ps {
		p.BigModel.Facets[key] = val
	}
	p.Rebuild()
	return p
}

// NewPersonalityRandomized returns a new personality with random facet values.
// TODO: Allow custom rand source or seed.
func NewPersonalityRandomized(t *Traiter) *Personality {
	p := NewPersonality(t)
	p.BigModel = *RandomBig()
	p.Rebuild()
	return p
}

// Rebuild re-evaluates expressed traits based on the facet ratings.
func (p *Personality) Rebuild() {
	p.Expressed = nil
	for _, t := range p.t.Traits {
		if t.IsExpressedBy(p) {
			p.addTrait(t)
		}
	}
}

// Log logs the personality.
func (p *Personality) Log() {
	p.BigModel.Log()
	for _, t := range p.Expressed {
		t.Log()
	}
}

// Stats returns the stats for the personality.
func (p *Personality) Stats() *Stats {
	// Rebuild the traits.
	p.Rebuild()

	// Calculate the stats.
	s := NewStats()
	for _, t := range p.Expressed {
		s.add(t.Stats)
	}
	return s
}

// addTrait adds the given traits to the personality.
func (p *Personality) addTrait(tt ...*Trait) {
	p.Expressed = mergeTraitSets(p.Expressed, tt)
}
