function fib(pos) {
  let mem = new Array(pos)

  for(let i=0;i<=pos;i++) {
    let f=0n
    if (i<=2) {
      f=1n
    } else {
      f=mem[i-1] + mem[i-2]
    }
    mem[i]=f
  }

  return mem[pos]
}

n = Number(process.argv[2])
const start = new Date().getTime()
// console.log(fib(n))
fib(n)
const end = new Date().getTime()
const totalTime = end - start
console.log("\n", "In Node - Fibonnaci series in position " + n + " took: " + totalTime + "ms", "\n" )