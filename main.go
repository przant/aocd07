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
    BiD   int
}

const (
    handLen = 5
)

var (
    labelValues = map[rune]int{
        '2': 0,
        '3': 1,
        '4': 2,
        '5': 3,
        '6': 4,
        '7': 5,
        '8': 6,
        '9': 7,
        'T': 8,
        'Q': 9,
        'K': 10,
        'A': 11,
    }
)

type ByHandCards []Hand

func (hs ByHandCards) Len() int      { return len(hs) }
func (hs ByHandCards) Swap(i, j int) { hs[i], hs[j] = hs[j], hs[i] }
func (hs ByHandCards) Less(i, j int) bool {
    for c := 0; c < handLen; c++ {
        if labelValues[rune(hs[i].Cards[c])] < labelValues[rune(hs[j].Cards[c])] &&
            hs[i].Rank <= hs[j].Rank {
            return true
        }
    }
    return false
}

func main() {
    pf, err := os.Open("example.txt")
    if err != nil {
        log.Fatalf("while opening file %q: %s", pf.Name(), err)
    }
    defer pf.Close()

    scnr := bufio.NewScanner(pf)

    hands := make(ByHandCards, 0)

    for scnr.Scan() {
        fields := strings.Split(scnr.Text(), " ")
        Bid, _ := strconv.Atoi(strings.TrimSpace(fields[1]))
        hand := Hand{
            Cards: strings.TrimSpace(fields[0]),
            BiD:   Bid,
        }
        switch {
        case hc(hand.Cards):
            hand.Rank = 1
        case pair(hand.Cards):
            hand.Rank = 2
        case pair2(hand.Cards):
            hand.Rank = 3
        case kao3(hand.Cards):
            hand.Rank = 4
        case fh(hand.Cards):
            hand.Rank = 5
        case kao4(hand.Cards):
            hand.Rank = 6
        case kao5(hand.Cards):
            hand.Rank = 7

        }
        hands = append(hands, hand)
    }

    sort.Sort(hands)

    sort.Slice(hands, func(i, j int) bool {
        return hands[i].Rank < hands[j].Rank
    })

    for i := 0; i < len(hands); i++ {
        hands[i].Rank = i + 1
    }

    result := int64(0)
    for i := 0; i < len(hands); i++ {
        result += int64(hands[i].Rank * hands[i].BiD)
    }

    for _, hand := range hands {
        fmt.Println(hand)
    }
    fmt.Println()
    fmt.Println(result)
}

func kao5(hand string) bool {
    label := hand[0]
    for _, card := range hand {
        if card != rune(label) {
            return false
        }
    }

    return true
}

func kao4(hand string) bool {
    swaps := 0
    label := hand[0]
    for _, card := range hand {
        if card != rune(label) {
            swaps++
        }
    }

    if swaps > 2 {
        return false
    }

    return true
}

func fh(hand string) bool {
    swaps := 0
    label := hand[0]
    for _, card := range hand {
        if card != rune(label) {
            swaps++
        }
    }

    if swaps > 3 {
        return false
    }

    return true
}

func kao3(hand string) bool {
    cards := make(map[rune]int)
    for _, card := range hand {
        cards[card]++
    }

    for _, value := range cards {
        if value == 3 {
            return true
        }
    }

    return false
}

func pair2(hand string) bool {
    pairs := 0
    cards := make(map[rune]int)
    for _, card := range hand {
        cards[card]++
    }

    for _, value := range cards {
        if value == 2 {
            pairs++
        }
    }

    return pairs == 2
}

func pair(hand string) bool {
    pairs := 0
    cards := make(map[rune]int)
    for _, card := range hand {
        cards[card]++
    }

    for _, value := range cards {
        if value == 2 {
            pairs++
        }
    }

    return pairs == 1
}

func hc(hand string) bool {
    if hand[0] != hand[1] && hand[1] != hand[2] &&
        hand[2] != hand[3] && hand[3] != hand[4] {
        return true
    }
    return false
}
