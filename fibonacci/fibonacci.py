import datetime, sys

n = int(sys.argv[1])

def fib_pd(pos):
  mem = [0]*(pos+1)
  for i in range(pos+1):
    f=0
    if i<=2:
      f=1
    else:
      f=mem[i-1]+mem[i-2]
    mem[i]=f
  return mem[pos]

beginTime = datetime.datetime.now()

# print(fib_pd(n))
fib_pd(n)

totalTime = datetime.datetime.now() - beginTime
print("\nIn Python - Fibonnaci series in position %d took: %sms\n" % (n, str(totalTime)))