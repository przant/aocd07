package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "sort"
    "strconv"
    "strings"
)

type Hand struct {
    Rank  int
    Cards string
    Bid   int
}

const (
    HightCard = iota
    OnePair
    TwoPair
    ThreeOAK
    FullHouse
    FourOAK
    FiveOAK
)

var (
    labelsMap = map[rune]rune{
        'J': 'A',
        '2': 'B',
        '3': 'C',
        '4': 'D',
        '5': 'E',
        '6': 'F',
        '7': 'G',
        '8': 'H',
        '9': 'I',
        'T': 'K',
        'Q': 'L',
        'K': 'M',
        'A': 'N',
    }

    revertMap = map[rune]rune{
        'A': 'J',
        'B': '2',
        'C': '3',
        'D': '4',
        'E': '5',
        'F': '6',
        'G': '7',
        'H': '8',
        'I': '9',
        'K': 'T',
        'L': 'Q',
        'M': 'K',
        'N': 'A',
    }
)

type ByHandCards []Hand

func main() {
    pf, err := os.Open("input.txt")
    if err != nil {
        log.Fatalf("while opening file %q: %s", pf.Name(), err)
    }
    defer pf.Close()

    scnr := bufio.NewScanner(pf)

    hands := make(ByHandCards, 0)

    for scnr.Scan() {
        fields := strings.Split(scnr.Text(), " ")
        Cards := strings.Map(func(r rune) rune { return labelsMap[r] }, strings.TrimSpace(fields[0]))
        Bid, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
        hand := Hand{
            Cards: Cards,
            Bid:   Bid,
        }
        handKind(&hand)
        hands = append(hands, hand)
    }

    sort.SliceStable(hands, func(i, j int) bool {
        return hands[i].Rank < hands[j].Rank
    })
    sort.SliceStable(hands, func(i, j int) bool {
        return hands[i].Cards < hands[j].Cards && hands[i].Rank == hands[j].Rank
    })

    result := int64(0)
    for i := 0; i < len(hands); i++ {
        result += int64((i + 1) * hands[i].Bid)
    }

    for _, hand := range hands {
        hand.Cards = strings.Map(func(r rune) rune { return revertMap[r] }, hand.Cards)
        fmt.Println(hand)
    }
    fmt.Println()
    fmt.Println(result)
}

func handKind(hand *Hand) {
    max := 0
    labelMap := make(map[byte]int)
    for i := 0; i < len(hand.Cards); i++ {
        labelMap[hand.Cards[i]]++
        if max < labelMap[hand.Cards[i]] {
            max = labelMap[hand.Cards[i]]
        }
    }

    jkrs := labelMap['A']

    labels := len(labelMap)
    switch {
    case labels == 1 && max == 5:
        hand.Rank = FiveOAK
    case labels == 2 && max == 4:
        hand.Rank = FourOAK
    case labels == 2 && max == 3:
        hand.Rank = FullHouse
        if jkrs > 0 {
            hand.Rank = FiveOAK
        }
    case labels == 3 && max == 3:
        hand.Rank = ThreeOAK
        if jkrs > 0 {
            hand.Rank = FourOAK
        }

    case labels == 3 && max == 2:
        hand.Rank = TwoPair
        if jkrs == 1 {
            hand.Rank = FullHouse
        }
        if jkrs == 2 {
            hand.Rank = FourOAK
        }
    case labels == 4 && max == 2:
        hand.Rank = OnePair
        if jkrs == 1 || jkrs == 2 {
            hand.Rank = ThreeOAK
        }
    case labels == 5 && max == 1:
        hand.Rank = HightCard
        if jkrs > 0 {
            hand.Rank = OnePair
        }
    }
}
