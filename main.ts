/**
 * @author Rabira Motuma
 * @version 1.0.0
 * @date 2025-11-25
 * @fileoverview This program rounds off numbers to a specified number of decimal places.
 */

// variables
let user: string;
let number1AsString: string;
let number2AsString: string;
let number3AsString: string;
let number4AsString: string;
let number1AsNumber: number;
let number2AsNumber: number;
let number3AsNumber: number;
let number4AsNumber: number;

// inputs
user = prompt("What is your name?") || "No name entered!";
number1AsString = prompt("What is your first number?") || "0";
number2AsString = prompt("What is your second number?") || "0";
number3AsString = prompt("What is your third number?") || "0";
number4AsString = prompt("What is your fourth number?") || "0";

// process
number1AsNumber = parseFloat(number1AsString);
number2AsNumber = parseFloat(number2AsString);
number3AsNumber = parseFloat(number3AsString);
number4AsNumber = parseFloat(number4AsString);

// output
console.log("\n");
console.log(`Hello, ${user}.`);
console.log(`${number1AsNumber.toFixed(3).padStart(5)}`);
console.log(`${number2AsNumber.toFixed(5).padStart(1)}`);
console.log(`${number3AsNumber.toFixed(1).padStart(2)}`);
console.log(`${number4AsNumber.toFixed(1).padStart(0)}`);

console.log("\nDone.");
