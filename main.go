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

var (
    labelsMap = map[rune]rune{
        '2': 'A',
        '3': 'B',
        '4': 'C',
        '5': 'D',
        '6': 'E',
        '7': 'F',
        '8': 'G',
        '9': 'H',
        'T': 'I',
        'J': 'J',
        'Q': 'K',
        'K': 'L',
        'A': 'M',
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
        fmt.Println(hand)
    }
    fmt.Println()
    fmt.Println(result)
}

func kao5(hand string) bool {
    max := 0
    labels := make(map[byte]int)
    for i := 0; i < len(hand); i++ {
        labels[hand[i]]++
        if max < labels[hand[i]] {
            max = labels[hand[i]]
        }
    }

    return len(labels) == 1 && max == 5
}

func kao4(hand string) bool {
    max := 0
    labels := make(map[rune]int)
    for _, label := range hand {
        labels[label]++
        if max < labels[label] {
            max = labels[label]
        }
    }

    return len(labels) == 2 && max == 4
}

func fh(hand string) bool {
    max := 0
    labels := make(map[rune]int)
    for _, label := range hand {
        labels[label]++
        if max < labels[label] {
            max = labels[label]
        }
    }

    return len(labels) == 2 && max == 3
}

func kao3(hand string) bool {
    max := 0
    labels := make(map[rune]int)
    for _, label := range hand {
        labels[label]++
        if max < labels[label] {
            max = labels[label]
        }
    }

    return len(labels) == 3 && max == 3
}

func pair2(hand string) bool {
    max := 0
    labels := make(map[rune]int)
    for _, label := range hand {
        labels[label]++
        if max < labels[label] {
            max = labels[label]
        }
    }

    return len(labels) == 3 && max == 2
}

func pair(hand string) bool {
    max := 0
    labels := make(map[rune]int)
    for _, label := range hand {
        labels[label]++
        if max < labels[label] {
            max = labels[label]
        }
    }

    return len(labels) == 4 && max == 2
}

func hc(hand string) bool {
    labels := make(map[rune]int)
    for _, label := range hand {
        labels[label]++
    }

    return len(labels) == 5
}
