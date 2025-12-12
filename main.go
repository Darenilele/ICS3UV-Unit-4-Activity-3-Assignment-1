/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-09
 * @fileoverview This program will ask how many items and what items you want to purchase, then calculate the price
 */

package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {

// Variables
  var itemString string = ""
  var itemNum int = 0

  var subtotal float64 = 0
  var tax float64 = 0
  var discount float64 = 0
  var total float64 = 0

// Input number of items
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("How many items would you like to buy? ")
  itemString, _ = reader.ReadString('\n')
  itemString = strings.TrimSpace(itemString)
  itemNum, _ = strconv.Atoi(itemString)

// Arrays with fixed size
  var priceHold [100]float64
  var itemsNameHold [100]string

// Processing
    for itemC := 0; itemC < itemNum; itemC++ {
    fmt.Print("What is the name of item ", itemC+1, " ? ")
    itemNameString, _ := reader.ReadString('\n')
    itemNameString = strings.TrimSpace(itemNameString)

      fmt.Print("What is the cost of said item ", itemC+1, " ? ")
      itemCostString, _ := reader.ReadString('\n')
      itemCostString = strings.TrimSpace(itemCostString)
      itemCostNumber, _ := strconv.ParseFloat(itemCostString, 64)

        itemsNameHold[itemC] = itemNameString
        priceHold[itemC] = itemCostNumber
    }

    for itemC := 0; itemC < itemNum; itemC++ {
        subtotal = subtotal + priceHold[itemC]
    }

    if subtotal > 350 {
        discount = subtotal * 0.10
    }

    tax = (subtotal - discount) * 0.13
    total = subtotal - discount + tax

// Output
    fmt.Print("Items in shopping cart: ")
    for i := 0; i < itemNum; i++ {
        fmt.Print(itemsNameHold[i])
        if i < itemNum-1 {
            fmt.Print(", ")
        }
    }
    fmt.Println()
    fmt.Println("You have purchased:", itemNum, "items")
    fmt.Println("Your subtotal is", subtotal)
    fmt.Println("You are able to get a discount of", discount)
    fmt.Println("The HST is", tax)
    fmt.Println("Your total is", total)

    fmt.Println("\nDone.")
}
