package entity

const fineCoefficientForMistake = 10

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
	if mistakeTranslations > correctTranslations {
		return fineCoefficientForMistake * float64(mistakeTranslations-correctTranslations)
	} else {
		return 1 / float64(1+(correctTranslations-mistakeTranslations))
	}
}
