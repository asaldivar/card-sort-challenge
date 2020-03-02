package card_sort

type CardSorter interface {
	SortCards(cards []BoardingCard) []BoardingCard
}

type sorter struct {
}

func (s sorter) SortCards(cards []BoardingCard) []BoardingCard {

}
