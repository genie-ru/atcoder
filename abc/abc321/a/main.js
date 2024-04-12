const input = require('fs').readFileSync('/dev/stdin', 'utf8')
let lines = input.split('\n')
let N = lines[0].trim()

let ans = 'Yes'

for (let i = 1; i < N.length; i++) {
  if (N[i - 1] <= N[i]) ans = 'No'
}

console.log(ans)
