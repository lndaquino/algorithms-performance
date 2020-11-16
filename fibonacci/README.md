### Fibonacci series
Run script in powershell: & fibonacci.ps1 200000 to see time performance using diferent techonologies

In my development enviroment Fibonacci series in position 200000 took:
- Node (14.3.0): ~1.5s
- Python (3.7.3): ~4.5s
- GO (1.15.3): ~0.6s

The algorithm uses only one thread in each language.