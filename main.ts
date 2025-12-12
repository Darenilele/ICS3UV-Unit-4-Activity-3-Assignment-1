/**
 * @author Daren Ileleji
 * @version 1.0.0
 * @date 2025-12-10
 * @fileoverview This program will ask how many items and what items you want to purchase, then calculate the price
 */

// Variables
let itemString: string = "";
let itemNum: number = 0;

let total: number = 0;
let subtotal: number = 0;
let tax: number = 0;
let discount: number = 0;

// Arrays
const priceHold: number[] = new Array();
const itemsNameHold: string[] = [];

// Input
itemString = prompt("How many items would you like to buy?") || "Invalid amount"
itemNum = parseInt(itemString);

// Processing
for (let itemC = 0; itemC < itemNum; itemC++ ) {
  const itemNameString = prompt(`What is the name of item ${itemC + 1} ?`) || "Invalid item"
  const itemCostString = prompt(`What is the cost of said item ${itemC + 1} ?`) || "Invalid amount"

  const itemCostNumber = parseFloat(itemCostString);

  itemsNameHold[itemC] = itemNameString;
  priceHold[itemC] = itemCostNumber;

}

for (let itemC = 0; itemC < itemsNameHold.length; itemC++) {
  subtotal += priceHold[itemC]

}

if (subtotal > 350) {
  discount = 0.10 * subtotal;
}

tax = (subtotal - discount) * 0.13;
total = subtotal - discount + tax;

// Output
console.log(`Items in shopping cart: ${itemsNameHold}`);
console.log(`You have purchased: ${itemNum} items`);
console.log(`Your subtotal is ${subtotal.toFixed(2)}`);
console.log(`You are able to get a discount of ${discount.toFixed(2)}`);
console.log(`the HST is ${tax.toFixed(2)}`);
console.log(`Your total is ${total.toFixed(2)}`);

console.log("\nDone.")