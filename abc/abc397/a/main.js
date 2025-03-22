const readline = require('readline');

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('', (input) => {
  const X = parseFloat(input);

  if (X >= 38.0) {
    console.log(1);
  } else if (X >= 37.5) {
    console.log(2);
  } else {
    console.log(3);
  }

  rl.close();
});