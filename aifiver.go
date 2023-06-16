// Package aifiver implements a simplified representation of the five factor model.
package aifiver

import (
	"log"
	"math/rand"
)

// Factor represents one of the five factors in the five factor model.
type Factor int

// The five factors.
const (
	FactorOpenness Factor = iota
	FactorConscientiousness
	FactorExtraversion
	FactorAgreeableness
	FactorNeuroticism
)

// Facet represents specific facets within the five factors.
type Facet int

// The facets.
const (
	FacetOpenFantasy             Facet = iota // Proneness to imagination, day-dreaming, and creating
	FacetOpenAesthetics                       // Appreciation for beauty in e.g. art, music, poetry, or nature
	FacetOpenFeelings                         // Receptivity to and intensity of experienced emotions
	FacetOpenAdventurousness                  // The tendency to choose novelty over the familiar
	FacetOpenIdeas                            // The degree of interest and curiosity in entertaining new thoughts and ideas
	FacetOpenValues                           // Willingness to re-evaluate norms and values
	FacetConsCompetence                       // Belief in one's own capacity to handle life's many challenges
	FacetConsOrder                            // Degree of neatness and orderliness
	FacetConsDutifulness                      // How strongly ethical principles guide action
	FacetConsAchievementStriving              // Aspiration-level, the willingness to work towards goals
	FacetConsSelfDicipline                    // The ability to follow through on tasks despite boredom
	FacetConsCautiousness                     // The care and thought put into actions
	FacetExtrWarmth                           // The degree of displayed affection and closeness in relationships
	FacetExtrGregariousness                   // The tendency of seeking the company of others
	FacetExtrAssertiveness                    // The degree of dominance in social interaction
	FacetExtrActivity                         // The level of energy and activity in daily life
	FacetExtrExcitementSeeking                // The need for thrills and intense stimulation
	FacetExtrPositiveEmotions                 // The tendency to be happy, excited, and cheerful
	FacetAgreTrust                            // The general level of wariness or suspicion in contact with other people
	FacetAgreStraightforwardness              // Degree of sincerity vs shrewdness
	FacetAgreAltruism                         // Active concern for the well-being of others
	FacetAgreCompliance                       // Inhibiting vs expressing agression towards others in conflict
	FacetAgreModesty                          // Degree of humility vs arrogance
	FacetAgreTenderMindedness                 // Propensity to empathize with others
	FacetNeurAnxiety                          // Proneness to worry and rumination
	FacetNeurAngryHostility                   // The readiness to experience frustration, anger, and bitterness
	FacetNeurDepression                       // The tendency for guilt, sadness, lonliness, and hopelessness
	FacetNeurSelfConciousness                 // Sensitivity in social situations, such as ridicule, rejection, or awkwardness
	FacetNeurImpulsiveness                    // The ability to tolerate frustration and to control urges, cravings, and desires
	FacetNeurVulnerability                    // The ability to cope with stress
)

// Factor returns the factor associated with the given facet.
func (f Facet) Factor() Factor {
	return facetToFactor[f]
}

// facetToFactor maps a facet to its respective factor.
var facetToFactor = map[Facet]Factor{
	FacetOpenFantasy:             FactorOpenness,
	FacetOpenAesthetics:          FactorOpenness,
	FacetOpenFeelings:            FactorOpenness,
	FacetOpenAdventurousness:     FactorOpenness,
	FacetOpenIdeas:               FactorOpenness,
	FacetOpenValues:              FactorOpenness,
	FacetConsCompetence:          FactorConscientiousness,
	FacetConsOrder:               FactorConscientiousness,
	FacetConsDutifulness:         FactorConscientiousness,
	FacetConsAchievementStriving: FactorConscientiousness,
	FacetConsSelfDicipline:       FactorConscientiousness,
	FacetConsCautiousness:        FactorConscientiousness,
	FacetExtrWarmth:              FactorExtraversion,
	FacetExtrGregariousness:      FactorExtraversion,
	FacetExtrAssertiveness:       FactorExtraversion,
	FacetExtrActivity:            FactorExtraversion,
	FacetExtrExcitementSeeking:   FactorExtraversion,
	FacetExtrPositiveEmotions:    FactorExtraversion,
	FacetAgreTrust:               FactorAgreeableness,
	FacetAgreStraightforwardness: FactorAgreeableness,
	FacetAgreAltruism:            FactorAgreeableness,
	FacetAgreCompliance:          FactorAgreeableness,
	FacetAgreModesty:             FactorAgreeableness,
	FacetAgreTenderMindedness:    FactorAgreeableness,
	FacetNeurAnxiety:             FactorNeuroticism,
	FacetNeurAngryHostility:      FactorNeuroticism,
	FacetNeurDepression:          FactorNeuroticism,
	FacetNeurSelfConciousness:    FactorNeuroticism,
	FacetNeurImpulsiveness:       FactorNeuroticism,
	FacetNeurVulnerability:       FactorNeuroticism,
}

// String returns the string representation of the facet.
func (f Facet) String() string {
	return facetToString[f]
}

// facetToString maps a facet to its string representation.
var facetToString = map[Facet]string{
	FacetOpenFantasy:             "O: Fantasy",
	FacetOpenAesthetics:          "O: Aesthetics",
	FacetOpenFeelings:            "O: Feelings",
	FacetOpenAdventurousness:     "O: Adventurousness",
	FacetOpenIdeas:               "O: Ideas",
	FacetOpenValues:              "O: Values",
	FacetConsCompetence:          "C: Competence",
	FacetConsOrder:               "C: Order",
	FacetConsDutifulness:         "C: Dutifulness",
	FacetConsAchievementStriving: "C: Achievement Striving",
	FacetConsSelfDicipline:       "C: Self Dicipline",
	FacetConsCautiousness:        "C: Cautiousness",
	FacetExtrWarmth:              "E: Warmth",
	FacetExtrGregariousness:      "E: Gregariousness",
	FacetExtrAssertiveness:       "E: Assertiveness",
	FacetExtrActivity:            "E: Activity",
	FacetExtrExcitementSeeking:   "E: Excitement Seeking",
	FacetExtrPositiveEmotions:    "E: Positive Emotions",
	FacetAgreTrust:               "A: Trust",
	FacetAgreStraightforwardness: "A: Straightforwardness",
	FacetAgreAltruism:            "A: Altruism",
	FacetAgreCompliance:          "A: Compliance",
	FacetAgreModesty:             "A: Modesty",
	FacetAgreTenderMindedness:    "A: Tender Mindedness",
	FacetNeurAnxiety:             "N: Anxiety",
	FacetNeurAngryHostility:      "N: Angry Hostility",
	FacetNeurDepression:          "N: Depression",
	FacetNeurSelfConciousness:    "N: Self Conciousness",
	FacetNeurImpulsiveness:       "N: Impulsiveness",
	FacetNeurVulnerability:       "N: Vulnerability",
}

// Note to self:
// Here are some interesting approaches..
// - https://github.com/JacobStoneman/NPCProcGen/blob/master/Library/Collab/Download/Assets/Scripts/NPCController.cs

// Model represents a personality model.
type Model interface {
	Get(factor Factor) int
	GetFacet(facet Facet) int
	Log()
}

// SmallModel represents a simplified five factor model.
// Please note that this is not granular enough to simulate real personalities.
type SmallModel [5]int

// Get returns the value of the given factor.
func (p *SmallModel) Get(factor Factor) int {
	return p[factor]
}

// GetFacet returns the value of the given facet.
func (p *SmallModel) GetFacet(facet Facet) int {
	return p.Get(facet.Factor())
}

// Log logs the small model.
func (p *SmallModel) Log() {
	log.Printf("O: %d, C: %d, E: %d, A: %d, N: %d\n", p[FactorOpenness], p[FactorConscientiousness], p[FactorExtraversion], p[FactorAgreeableness], p[FactorNeuroticism])
}

// RandomSmall randomizes a small model by distributing half of the total range (-10 to 10 = 20)
// of each factor randomly.
// We use a point buy system to distribute the points randomly, however there in reality some factors
// are more likely to be linked to each other than others.
func RandomSmall() *SmallModel {
	p := &SmallModel{}
	totalRange := 20
	maxPoints := totalRange * 5 / 2

	// Distribute points randomly.
	for _, i := range rand.Perm(len(p)) {
		// Maximum number of points per factor is 20 (10 to -10).
		forThisFactor := rand.Intn(maxPoints + 1)
		p[i] = forThisFactor - 10
		maxPoints -= forThisFactor
	}

	// If we have points left, distribute them randomly.
	for maxPoints > 0 {
		for _, i := range rand.Perm(len(p)) {
			if p[i] < 10 {
				p[i]++
				maxPoints--
			}
			if maxPoints <= 0 {
				break
			}
		}
	}

	return p
}

// BigModel represents a more granular five factor model.
type BigModel struct {
	Factors [5]int
	Facets  [30]int
}

// Get returns the value of the given factor.
func (p *BigModel) Get(f Factor) int {
	return p.Factors[f]
}

// GetFacet returns the value of the given facet.
func (p *BigModel) GetFacet(f Facet) int {
	return p.Facets[f]
}

// Log logs the big model.
func (p *BigModel) Log() {
	log.Printf("O: %d, C: %d, E: %d, A: %d, N: %d\n", p.Factors[FactorOpenness], p.Factors[FactorConscientiousness], p.Factors[FactorExtraversion], p.Factors[FactorAgreeableness], p.Factors[FactorNeuroticism])
	for i, v := range p.Facets {
		log.Printf("%s: %d\n", Facet(i), v)
	}
}

// RandomBig randomizes a big model by distributing half of the total range (-10 to 10 = 20)
// of each factor and facet randomly.
// We use a point buy system to distribute the points randomly, however there in reality some factors
// are more likely to be linked to each other than others.
// TODO: Find a way to better distribute the points according to the real world.
func RandomBig() *BigModel {
	p := &BigModel{}
	totalRange := 20
	maxPointsFactors := totalRange * 5 / 2

	// Distribute points for factors randomly.
	remainingPoints := maxPointsFactors
	for _, i := range rand.Perm(len(p.Factors)) {
		// Maximum number of points per factor is 20 (10 to -10).
		forThisFactor := rand.Intn(totalRange + 1)

		// If the number of points for this factor is higher than the remaining points, reduce it.
		if remainingPoints < forThisFactor {
			forThisFactor = remainingPoints
		}

		// Normalize the points from 0 to 20 to -10 to 10.
		p.Factors[i] = forThisFactor - 10

		// Reduce the remaining points.
		remainingPoints -= forThisFactor
	}

	// If we have points left, distribute them randomly.
	for remainingPoints > 0 {
		for _, i := range rand.Perm(len(p.Factors)) {
			if p.Factors[i] < 10 {
				p.Factors[i]++
				remainingPoints--
			}
			if remainingPoints <= 0 {
				break
			}
		}
	}

	// Now for each facet calculate the points based on the value we have for the factor.
	factorPool := make(map[Factor]int)
	maxPointsFacets := totalRange * 30 / 2
	for factor, fvalue := range p.Factors {
		// Calculate the pool of points for this factor.
		// Add 10 to the factor value to get a value between 0 and 20.
		facetPool := int(float64(fvalue+10) * float64(maxPointsFacets) / float64(maxPointsFactors))
		factorPool[Factor(factor)] = facetPool
	}

	// Distribute points for facets randomly.
	for _, i := range rand.Perm(len(p.Facets)) {
		// Get the facet pool for the factor this facet belongs to.
		facetPool := factorPool[Facet(i).Factor()]
		if facetPool <= 0 {
			continue
		}
		// Maximum number of points per facet is 20 (10 to -10).
		forThisFacet := rand.Intn(totalRange + 1)
		if facetPool < forThisFacet {
			forThisFacet = facetPool
		}
		// Assign the normalized points to this facet.
		p.Facets[i] = forThisFacet - 10

		// Reduce the facet pool.
		factorPool[Facet(i).Factor()] -= forThisFacet
	}

	// If we have points left, distribute them randomly.
	// NOTE: I know that this should be in a loop until all points are guaranteed to be distributed,
	// but I'm too lazy to do that right now.
	for _, i := range rand.Perm(len(p.Facets)) {
		rem := factorPool[Facet(i).Factor()]
		if rem <= 0 {
			continue
		}
		if p.Facets[i] < 10 {
			p.Facets[i]++
			factorPool[Facet(i).Factor()]--
		}
	}
	return p
}

// Compatibility returns a compatibility value for the given personalities.
// TODO: Provide function to calculate conflict potential of larger groups.
// https://github.com/FZarattini/PirateShip/blob/master/PirateShip/Assets/Scripts/AI/Empathy.cs
func Compatibility(p1, p2 Model) int {
	// Likelyhood of initiating communication is low if two introverts meet.
	ex := (p1.Get(FactorExtraversion) + p2.Get(FactorExtraversion)) / 2
	if ex <= 0 {
		// Well, that's two people that don't want to interact.
		return 0
	}

	// Potential for conflict if the total value for agreeableness is below zero.
	ag := (p1.Get(FactorAgreeableness) + p2.Get(FactorAgreeableness)) / 2
	if ag < 0 {
		// We are in a very disagreeable situation: Conflict!
		return -1
	}
	return (ex + ag) / 2
}
