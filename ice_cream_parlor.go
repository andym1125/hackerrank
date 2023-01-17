package main

/* Challenge: www.hackerrank.com/challenges/ctci-ice-cream-parlor */

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'whatFlavors' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY cost
 *  2. INTEGER money
 */

func whatFlavors(cost []int32, money int32) {
    // Write your code here
    costs := make(map[int32]int32) //cost->repetitions
    locs := make(map[int32][]int32) //cost->list of index+1
    for i, c := range cost {
        costs[c]++
        locs[c] = append(locs[c], int32(i+1))
    }
    
    for currCost, num := range costs {
        if currCost >= money {
            continue
        }
        
        _, exists := costs[money-currCost]
        if exists {
            // Same cost
            if money-currCost == currCost && num >= 2 {
                if locs[currCost][0] < locs[currCost][1] {
                    fmt.Printf("%d %d\n", locs[currCost][0], locs[currCost][1])
                } else {
                    fmt.Printf("%d %d\n", locs[currCost][0], locs[currCost][1])
                }
                return                
            //Different costs
            } else if money-currCost != currCost {
                if locs[currCost][0] < locs[money-currCost][0] {
                    fmt.Printf("%d %d\n", locs[currCost][0], locs[money-currCost][0])            
                } else {
                    fmt.Printf("%d %d\n", locs[money-currCost][0], locs[currCost][0])            
                }
                return
            }
            //currCost is half of money
            continue
        }
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        moneyTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        money := int32(moneyTemp)

        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        costTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var cost []int32

        for i := 0; i < int(n); i++ {
            costItemTemp, err := strconv.ParseInt(costTemp[i], 10, 64)
            checkError(err)
            costItem := int32(costItemTemp)
            cost = append(cost, costItem)
        }

        whatFlavors(cost, money)
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
