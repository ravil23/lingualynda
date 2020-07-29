package entity

const (
	fineCoefficientForCorrect = 0.1
	fineCoefficientForMistake = 10
	diffCoefficientForCorrect = 1
	diffCoefficientForMistake = 2
)

type UserProfile struct {
	userID                    UserID
	correctlyTranslatedTerms  map[Term]int
	mistakenlyTranslatedTerms map[Term]int
}

func NewUserProfile(userID UserID) *UserProfile {
	return &UserProfile{
		userID:                    userID,
		correctlyTranslatedTerms:  make(map[Term]int),
		mistakenlyTranslatedTerms: make(map[Term]int),
	}
}

func (p *UserProfile) AddCorrectlyTranslatedTerm(term Term) {
	p.correctlyTranslatedTerms[term]++
}

func (p *UserProfile) AddMistakenlyTranslatedTerm(term Term) {
	p.mistakenlyTranslatedTerms[term]++
}

func (p *UserProfile) GetMemorizationWeight(term Term) (float64, bool) {
	correctTranslations, foundInCorrect := p.correctlyTranslatedTerms[term]
	mistakeTranslations, foundInMistake := p.mistakenlyTranslatedTerms[term]
	if !foundInCorrect && !foundInMistake {
		return 1, false
	}
	diff := float64(diffCoefficientForCorrect*correctTranslations - diffCoefficientForMistake*mistakeTranslations)
	if diff == 0 {
		return 1, true
	} else if diff < 0 {
		return fineCoefficientForMistake * -diff, true
	} else {
		return fineCoefficientForCorrect * 1 / (1 + diff), true
	}
}

func (p *UserProfile) IsCorrectlyMemorized(term Term) bool {
	correctTranslations := p.correctlyTranslatedTerms[term]
	mistakeTranslations := p.mistakenlyTranslatedTerms[term]
	return correctTranslations > mistakeTranslations
}
