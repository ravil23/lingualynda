package entity

const (
	fineCoefficientForCorrect = 0.1
	fineCoefficientForMistake = 10
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

func (p *UserProfile) GetMemorizationWeight(term Term) float64 {
	correctTranslations := p.correctlyTranslatedTerms[term]
	mistakeTranslations := p.mistakenlyTranslatedTerms[term]
	diff := float64(correctTranslations - mistakeTranslations)
	if diff == 0 {
		return 1
	} else if diff < 0 {
		return fineCoefficientForMistake * -diff
	} else {
		return fineCoefficientForCorrect * 1 / (1 + diff)
	}
}

func (p *UserProfile) IsCorrectMemorized(term Term) bool {
	correctTranslations := p.correctlyTranslatedTerms[term]
	mistakeTranslations := p.mistakenlyTranslatedTerms[term]
	return correctTranslations > mistakeTranslations
}
