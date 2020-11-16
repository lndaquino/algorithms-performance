# run example bellow in powershell
# & fibonacci.ps1 200000

param([Int32]$pos=1)

node fibonacci.js $pos
python fibonacci.py $pos
.\fibonacci.exe $pos