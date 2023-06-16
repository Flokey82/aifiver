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
	traitChaste.Stats.Opinion[TOpinionSame] = 10
	traitChaste.Stats.Opinion[TOpinionOpposite] = -10
	tt.AddTrait(traitChaste)

	traitLustful := NewTrait("Lustful", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetNeurImpulsiveness) > 5
	})
	traitLustful.Stats.Skill[TSkillIntrigue] = 2
	traitLustful.Stats.Opinion[TOpinionSame] = 10
	traitLustful.Stats.Opinion[TOpinionOpposite] = -10
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
	traitHonest.Stats.Opinion[TOpinionSame] = 10
	traitHonest.Stats.Opinion[TOpinionOpposite] = -10
	tt.AddTrait(traitHonest)

	traitDeceitful := NewTrait("Deceitful", TTypePersonality, func(p *Personality) bool {
		return p.GetFacet(FacetAgreAltruism) < 0 &&
			p.GetFacet(FacetAgreStraightforwardness) < 2 &&
			p.GetFacet(FacetAgreModesty) < -5
	})
	traitDeceitful.Stats.Skill[TSkillDiplomacy] = -2
	traitDeceitful.Stats.Skill[TSkillIntrigue] = 4
	traitDeceitful.Stats.Opinion[TOpinionOpposite] = -10
	tt.AddTrait(traitDeceitful)

	MarkOppositeTraits(traitHonest, traitDeceitful)

	/*
		traitGenerous := NewTrait("Generous", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetAgreAltruism) > 6
		})
		traitGenerous.Stats.Skill[TSkillDiplomacy] = 2
		traitGenerous.Stats.Skill[TSkillIntrigue] = -2
		traitGenerous.Stats.Opinion[TOpinionOpposite] = -15
		tt.AddTrait(traitGenerous)

		traitGreedy := NewTrait("Greedy", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetAgreAltruism) < -6
		})
		traitGreedy.Stats.Skill[TSkillDiplomacy] = -2
		traitGreedy.Stats.Skill[TSkillIntrigue] = 2
		tt.AddTrait(traitGreedy)

		MarkOppositeTraits(traitGenerous, traitGreedy)

		traitJust := NewTrait("Just", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetAgreAltruism) > 6 &&
				p.GetFacet(FacetAgreStraightforwardness) > 6
		})
		traitJust.Stats.Skill[TSkillDiplomacy] = 2
		traitJust.Stats.Skill[TSkillIntrigue] = -2
		traitJust.Stats.Skill[TSkillParenting] = 2
		tt.AddTrait(traitJust)

		traitArbitrary := NewTrait("Arbitrary", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetAgreAltruism) < -6 &&
				p.GetFacet(FacetAgreStraightforwardness) < -6
		})
		traitArbitrary.Stats.Skill[TSkillDiplomacy] = -2
		traitArbitrary.Stats.Skill[TSkillIntrigue] = 2
		traitArbitrary.Stats.Skill[TSkillParenting] = -2
		tt.AddTrait(traitArbitrary)

		MarkOppositeTraits(traitJust, traitArbitrary)

		traitBrave := NewTrait("Brave", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetNeurAnxiety) < -6 &&
				p.GetFacet(FacetNeurAngryHostility) < -6
		})
		traitBrave.Stats.Skill[TSkillMartial] = 2
		traitBrave.Stats.Skill[TSkillProwess] = 3
		traitBrave.Stats.Skill[TSkillIntrigue] = -2
		tt.AddTrait(traitBrave)

		traitCowardly := NewTrait("Cowardly", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetNeurAnxiety) > 6 &&
				p.GetFacet(FacetNeurAngryHostility) > 6
		})
		traitCowardly.Stats.Skill[TSkillMartial] = -2
		traitCowardly.Stats.Skill[TSkillProwess] = -3
		traitCowardly.Stats.Skill[TSkillIntrigue] = 2
		tt.AddTrait(traitCowardly)

		MarkOppositeTraits(traitBrave, traitCowardly)

		traitPatient := NewTrait("Patient", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetNeurAnxiety) < -6 &&
				p.GetFacet(FacetNeurAngryHostility) > 6 &&
				p.GetFacet(FacetNeurImpulsiveness) < -4
		})
		traitPatient.Stats.Skill[TSkillMartial] = 2
		traitPatient.Stats.Skill[TSkillIntrigue] = -2
		traitPatient.Stats.Skill[TSkillParenting] = 2
		tt.AddTrait(traitPatient)

		traitImpatient := NewTrait("Impatient", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetNeurAnxiety) > 6 &&
				p.GetFacet(FacetNeurAngryHostility) < -6 &&
				p.GetFacet(FacetNeurImpulsiveness) > 4
		})
		traitImpatient.Stats.Skill[TSkillMartial] = -2
		traitImpatient.Stats.Skill[TSkillIntrigue] = 2
		traitImpatient.Stats.Skill[TSkillParenting] = -2
		tt.AddTrait(traitImpatient)

		MarkOppositeTraits(traitPatient, traitImpatient)

		traitHumble := NewTrait("Humble", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetAgreModesty) > 6
		})
		traitHumble.Stats.Skill[TSkillDiplomacy] = 2
		traitHumble.Stats.Skill[TSkillIntrigue] = -2
		tt.AddTrait(traitHumble)

		traitProud := NewTrait("Proud", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetAgreModesty) < -6
		})
		traitProud.Stats.Skill[TSkillDiplomacy] = -2
		traitProud.Stats.Skill[TSkillIntrigue] = 2
		tt.AddTrait(traitProud)

		MarkOppositeTraits(traitHumble, traitProud)

		traitGregarious := NewTrait("Gregarious", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetExtrGregariousness) > 6
		})
		traitGregarious.Stats.Skill[TSkillDiplomacy] = 2
		traitGregarious.Stats.Opinion[TOpinionSame] = 10
		tt.AddTrait(traitGregarious)

		traitSolitary := NewTrait("Solitary", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetExtrGregariousness) < -6
		})
		traitSolitary.Stats.Skill[TSkillLearning] = 2
		traitSolitary.Stats.Opinion[TOpinionSame] = 10
		tt.AddTrait(traitSolitary)

		MarkOppositeTraits(traitGregarious, traitSolitary)

		traitCalm := NewTrait("Calm", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetNeurAnxiety) < -6 &&
				p.GetFacet(FacetNeurAngryHostility) < -6
		})
		traitCalm.Stats.Skill[TSkillDiplomacy] = 1
		traitCalm.Stats.Skill[TSkillIntrigue] = 1
		traitCalm.Stats.Opinion[TOpinionSame] = 10
		traitCalm.Stats.Opinion[TOpinionOpposite] = -10
		tt.AddTrait(traitCalm)

		traitVengeful := NewTrait("Vengeful", TTypePersonality, func(p *Personality) bool {
			return p.GetFacet(FacetNeurAnxiety) > 6 &&
				p.GetFacet(FacetNeurAngryHostility) > 6
		})
		traitVengeful.Stats.Skill[TSkillDiplomacy] = -1
		traitVengeful.Stats.Skill[TSkillMartial] = 4
		tt.AddTrait(traitVengeful)

		MarkOppositeTraits(traitCalm, traitVengeful)
	*/
}
