package entity

import (
	"math"
	"math/rand"
)

type Term string

func (t Term) String() string {
	return string(t)
}

type Translation string

func (t Translation) String() string {
	return string(t)
}

type Vocabulary struct {
	translations map[Term][]Translation

	allTerms        []Term
	allTranslations []Translation
}

func NewEmptyVocabulary() *Vocabulary {
	return &Vocabulary{
		translations: make(map[Term][]Translation),
	}
}

func NewVocabulary(wordList map[Term][]Translation) *Vocabulary {
	v := &Vocabulary{
		translations:    wordList,
		allTerms:        make([]Term, 0, len(wordList)),
		allTranslations: make([]Translation, 0, len(wordList)),
	}
	for term, translations := range wordList {
		v.allTerms = append(v.allTerms, term)
		v.allTranslations = append(v.allTranslations, translations...)
	}
	return v
}

func (v *Vocabulary) GetRandomTerm() Term {
	return v.allTerms[rand.Intn(len(v.allTerms))]
}

func (v *Vocabulary) GetTermByUserProfile(userProfile *UserProfile) (Term, float64, bool) {
	weights := make(map[Term]float64, len(v.allTerms))
	weightsSum := 0.
	weightsMax := 0.
	for _, term := range v.allTerms {
		weight := userProfile.GetMemorizationWeight(term)
		weights[term] = weight
		weightsSum += weight
		weightsMax = math.Max(weightsMax, weight)
	}
	randomPoint := rand.Float64() * weightsSum
	allTermsMemorized := weightsMax < 1
	var left, right float64
	for term, weight := range weights {
		right = left + weight
		if left <= randomPoint && randomPoint < right {
			return term, weight, allTermsMemorized
		}
		left = right
	}
	return "", 0, false
}

func (v *Vocabulary) GetCorrectMemorizedTermsCount(userProfile *UserProfile) int {
	correctMemorizedTermsCount := 0
	for _, term := range v.allTerms {
		if userProfile.IsCorrectMemorized(term) {
			correctMemorizedTermsCount++
		}
	}
	return correctMemorizedTermsCount
}

func (v *Vocabulary) GetTermsCount() int {
	return len(v.allTerms)
}

func (v *Vocabulary) GetTranslations(term Term) []Translation {
	return v.translations[term]
}

func (v *Vocabulary) GetRandomTranslation() Translation {
	return v.allTranslations[rand.Intn(len(v.allTranslations))]
}

func (v *Vocabulary) Update(other *Vocabulary) *Vocabulary {
	for term, translations := range other.translations {
		v.translations[term] = translations
	}
	v.allTerms = append(v.allTerms, other.allTerms...)
	v.allTranslations = append(v.allTranslations, other.allTranslations...)
	return v
}

func (v *Vocabulary) MakeInvertedVocabulary() *Vocabulary {
	invertedWordList := make(map[Term][]Translation)
	for term, translations := range v.translations {
		for _, translation := range translations {
			if invertedTranslations, found := invertedWordList[Term(translation)]; found {
				alreadyAdded := false
				for _, invertedTranslation := range invertedTranslations {
					if invertedTranslation == Translation(term) {
						alreadyAdded = true
						break
					}
				}
				if !alreadyAdded {
					invertedTranslations = append(invertedTranslations, Translation(term))
				}
			} else {
				invertedWordList[Term(translation)] = []Translation{Translation(term)}
			}
		}
	}
	return NewVocabulary(invertedWordList)
}
