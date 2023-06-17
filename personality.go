package aifiver

// Personality represents a personality that is a combination of facet scores and
// expressed traits.
type Personality struct {
	t         *Traiter // Traiter reference for re-evaluating the expressed traits
	BigModel           // The model which influences which traits are expressed.
	Expressed []*Trait // Traits expressed based on the underlying 5 factor model.
	// TODO: Proneness to fall victim to cetain fates (accidental death, falling victim to intrigue)
	Modifiers []*Modifier // Temporary facet modifiers (Griefing, Recent Betrayal, Setback, ...)
	// TODO: Add a factor of how much the modifiers affect the personality...
	// Over time this should deminish and eventually disappear.
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

// AddModifier adds a modifier to the personality (if it is not already present).
func (p *Personality) AddModifier(m *Modifier) {
	for _, mod := range p.Modifiers {
		if mod == m {
			return
		}
	}
	p.Modifiers = append(p.Modifiers, m)
	p.Rebuild()
}

// RemoveModifier removes a modifier from the personality (if it is present).
func (p *Personality) RemoveModifier(m *Modifier) {
	for i, mod := range p.Modifiers {
		if mod == m {
			p.Modifiers = append(p.Modifiers[:i], p.Modifiers[i+1:]...)
			p.Rebuild()
			return
		}
	}
}

// Get implements the Model interface and applies any modifiers to the base value.
func (p *Personality) Get(f Factor) int {
	base := p.BigModel.Get(f)
	if len(p.Modifiers) == 0 {
		return base
	}

	// Apply the modifiers.
	for _, m := range p.Modifiers {
		base += m.Factors[f]
	}
	return clampTo10(base)
}

// GetFacet implements the Model interface and applies any modifiers to the base value.
func (p *Personality) GetFacet(f Facet) int {
	base := p.BigModel.GetFacet(f)
	if len(p.Modifiers) == 0 {
		return base
	}

	// Apply the modifiers.
	for _, m := range p.Modifiers {
		base += m.Facets[f]
	}
	return clampTo10(base)
}

func clampTo10(val int) int {
	if val > 10 {
		return 10
	} else if val < -10 {
		return -10
	}
	return val
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

// Interaction returns a value for how well the personalities interact face to face.
func (p *Personality) Interaction(b *Personality) int {
	return Compatibility(p, b)
}

// Opinion returns the opinion of the other personality based on their reputation.
// TODO: What about a misperception of the other person?
func (p *Personality) Opinion(b *Personality) int {
	var base int

	// Now check if any of our traits modify the compatibility.
	for _, t := range p.Expressed {
		if t.Stats.Opinion[TOpinionSame] != 0 && b.hasTrait(t) {
			base += t.Stats.Opinion[TOpinionSame]
		} else if t.Stats.Opinion[TOpinionOpposite] != 0 && b.hasTrait(t.Opposite) {
			base += t.Stats.Opinion[TOpinionOpposite]
		}
	}

	// TODO: Add attraction opinion modifiers?
	return base
}

func (p *Personality) hasTrait(t *Trait) bool {
	for _, tt := range p.Expressed {
		if tt == t {
			return true
		}
	}
	return false
}
