package aifiver

// DefaultTraits adds the default traits to the traiter.
// These are just examples. If you have suggestions for better traits, please
// open an issue or a pull request!
func DefaultTraits(tt *Traiter) {
	traitParanoid := NewTrait("Paranoid", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetAgreTrust) < -6
	})
	traitParanoid.Stats.Skill[TSkillDiplomacy] = -5
	traitParanoid.Stats.Skill[TSkillIntrigue] = 5
	tt.AddTrait(traitParanoid)

	traitGullable := NewTrait("Gullable", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetAgreTrust)+p.GetFacet(FacetAgreCompliance) > 15
	})
	traitGullable.Stats.Skill[TSkillDiplomacy] = -5
	traitGullable.Stats.Skill[TSkillIntrigue] = -5
	tt.AddTrait(traitGullable)

	MarkOppositeTraits(traitParanoid, traitGullable)

	traitChaste := NewTrait("Chaste", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetNeurImpulsiveness) < -5
	})
	traitChaste.Stats.Skill[TSkillLearning] = 2
	tt.AddTrait(traitChaste)

	traitLustful := NewTrait("Lustful", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetNeurImpulsiveness) > 5
	})
	traitLustful.Stats.Skill[TSkillIntrigue] = 2
	tt.AddTrait(traitLustful)

	MarkOppositeTraits(traitChaste, traitLustful)

	traitHonest := NewTrait("Honest", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetNeurImpulsiveness) < -6 &&
			p.GetFacet(FacetConsDutifulness) > 6 &&
			p.GetFacet(FacetAgreAltruism) > 2 &&
			p.GetFacet(FacetAgreStraightforwardness) > 2
	})
	traitHonest.Stats.Skill[TSkillDiplomacy] = 2
	traitHonest.Stats.Skill[TSkillIntrigue] = -4
	tt.AddTrait(traitHonest)

	traitDeceitful := NewTrait("Deceitful", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetAgreAltruism) < 0 &&
			p.GetFacet(FacetAgreStraightforwardness) < 2 &&
			p.GetFacet(FacetAgreModesty) < -5
	})
	traitDeceitful.Stats.Skill[TSkillDiplomacy] = -2
	traitDeceitful.Stats.Skill[TSkillIntrigue] = 4
	tt.AddTrait(traitDeceitful)

	MarkOppositeTraits(traitHonest, traitDeceitful)
}
