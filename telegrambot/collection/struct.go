package collection

import (
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
	collection map[Term][]Translation

	allTerms        []Term
	allTranslations []Translation
}

func NewEmptyVocabulary() *Vocabulary {
	return &Vocabulary{
		collection: make(map[Term][]Translation),
	}
}

func NewVocabulary(collection map[Term][]Translation) *Vocabulary {
	v := &Vocabulary{
		collection:      collection,
		allTerms:        make([]Term, 0, len(collection)),
		allTranslations: make([]Translation, 0, len(collection)),
	}
	for term, translations := range collection {
		v.allTerms = append(v.allTerms, term)
		v.allTranslations = append(v.allTranslations, translations...)
	}
	return v
}

func (v *Vocabulary) GetRandomTerm() Term {
	return v.allTerms[rand.Intn(len(v.allTerms))]
}

func (v *Vocabulary) GetTranslations(term Term) []Translation {
	return v.collection[term]
}

func (v *Vocabulary) GetRandomTranslation() Translation {
	return v.allTranslations[rand.Intn(len(v.allTranslations))]
}

func (v *Vocabulary) Update(other *Vocabulary) *Vocabulary {
	for term, translations := range other.collection {
		v.collection[term] = translations
	}
	v.allTerms = append(v.allTerms, other.allTerms...)
	v.allTranslations = append(v.allTranslations, other.allTranslations...)
	return v
}
