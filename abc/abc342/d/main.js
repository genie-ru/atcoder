let lines = require("fs").readFileSync("/dev/stdin","utf8").trim().split("\n");
let N = Number(lines[0]);
let A = lines[1].split(" ").map(Number);
let d = new Array(200001).fill(0);
let co = 0,zc = 0;
for(let i=0;i<N;i++){
  if(A[i]!=0){
    let num = A[i],tmp = Math.trunc(Math.sqrt(num));
    for(;;){
      let n = Math.pow(tmp,2)
      if(A[i]%n==0||tmp==0){
        d[A[i]/n] += 1;
        co += (d[A[i]/n]-1);
        break;
      }
      tmp--;
    }
  }else{
    co += ((N-1)-zc);
    zc++;
  } 
}
console.log(co);