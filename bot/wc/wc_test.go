package wc_test

import (
	"testing"

	"github.com/ravil23/lingualynda/bot/wc"
)

func TestWordCounts(t *testing.T) {
	t.Run("EmptyText", func(t *testing.T) {
		text := ""
		expected := map[string]uint{}

		actual := wc.WordCounts(text)

		requireWordCountsEqual(t, expected, actual)
	})

	t.Run("TextWithUniqWords", func(t *testing.T) {
		text := "a b c"
		expected := map[string]uint{"a": 1, "b": 1, "c": 1}

		actual := wc.WordCounts(text)

		requireWordCountsEqual(t, expected, actual)
	})

	t.Run("TextWithExtraSymbols", func(t *testing.T) {
		text := " a,  1b c?"
		expected := map[string]uint{"a": 1, "b": 1, "c": 1}

		actual := wc.WordCounts(text)

		requireWordCountsEqual(t, expected, actual)
	})

	t.Run("TextWithDuplicatedWords", func(t *testing.T) {
		text := "Ab ab aB AB"
		expected := map[string]uint{"ab": 4}

		actual := wc.WordCounts(text)

		requireWordCountsEqual(t, expected, actual)
	})
}

func TestSortByValuesAndKeys(t *testing.T) {
	t.Run("EmptyCollection", func(t *testing.T) {
		collection := map[string]uint{}
		expected := []wc.WordCountPair{}

		actual := wc.SortByValuesAndKeys(collection, false)

		requireSortedItemsEqual(t, expected, actual)
	})

	t.Run("CollectionAscendingWithUniqValues", func(t *testing.T) {
		collection := map[string]uint{"a": 1, "b": 3, "c": 2}
		expected := []wc.WordCountPair{
			{Word: "a", Count: 1},
			{Word: "c", Count: 2},
			{Word: "b", Count: 3},
		}

		actual := wc.SortByValuesAndKeys(collection, false)

		requireSortedItemsEqual(t, expected, actual)
	})

	t.Run("CollectionDescendingWithUniqValues", func(t *testing.T) {
		collection := map[string]uint{"a": 1, "b": 3, "c": 2}
		expected := []wc.WordCountPair{
			{Word: "b", Count: 3},
			{Word: "c", Count: 2},
			{Word: "a", Count: 1},
		}

		actual := wc.SortByValuesAndKeys(collection, true)

		requireSortedItemsEqual(t, expected, actual)
	})

	t.Run("CollectionAscendingWithEqualValues", func(t *testing.T) {
		collection := map[string]uint{"b": 0, "c": 0, "a": 0}
		expected := []wc.WordCountPair{
			{Word: "a", Count: 0},
			{Word: "b", Count: 0},
			{Word: "c", Count: 0},
		}

		actual := wc.SortByValuesAndKeys(collection, false)

		requireSortedItemsEqual(t, expected, actual)
	})

	t.Run("CollectionDescendingWithEqualValues", func(t *testing.T) {
		collection := map[string]uint{"a": 0, "c": 0, "b": 0}
		expected := []wc.WordCountPair{
			{Word: "c", Count: 0},
			{Word: "b", Count: 0},
			{Word: "a", Count: 0},
		}

		actual := wc.SortByValuesAndKeys(collection, true)

		requireSortedItemsEqual(t, expected, actual)
	})
}

func TestOfIntegration(t *testing.T) {
	t.Run("Example", func(t *testing.T) {
		text := "To be, or not to be?"
		expected := []wc.WordCountPair{
			{Word: "to", Count: 2},
			{Word: "be", Count: 2},
			{Word: "or", Count: 1},
			{Word: "not", Count: 1},
		}

		collection := wc.WordCounts(text)
		actual := wc.SortByValuesAndKeys(collection, true)

		requireSortedItemsEqual(t, expected, actual)
	})
}

func requireWordCountsEqual(
	t *testing.T,
	expected map[string]uint,
	actual map[string]uint,
) {

	if len(expected) != len(actual) {
		t.Errorf("Invalid length: %d != %d", len(expected), len(actual))
	}
	for word, count := range expected {
		if count != actual[word] {
			t.Errorf("Invalid count for word '%s': %d != %d", word, count, actual[word])
		}
	}
}

func requireSortedItemsEqual(
	t *testing.T,
	expected []wc.WordCountPair,
	actual []wc.WordCountPair,
) {

	if len(expected) != len(actual) {
		t.Errorf("Invalid length: %d != %d", len(expected), len(actual))
	}
	for i, pair := range expected {
		if pair.Word != actual[i].Word {
			t.Errorf("Invalid word: '%s' != '%s'", pair.Word, actual[i].Word)
		}
		if pair.Count != actual[i].Count {
			t.Errorf("Invalid count: '%d' != '%d'", pair.Count, actual[i].Count)
		}
	}
}
