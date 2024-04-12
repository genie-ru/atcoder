import * as fs from "fs";

const inputsStr = fs.readFileSync("/dev/stdin", "utf8");
const inputs = inputsStr.split('\n');

var a = false
var b = false
var c = false
 
for (let i = 0; i < parseInt(inputs[0]); i++) {
  if (inputs[1].charAt(i) === 'A') a = true
  if (inputs[1].charAt(i) === 'B') b = true
  if (inputs[1].charAt(i) === 'C') c = true
  if (a && b && c) {
    console.log(i+1)
    break
  }
}