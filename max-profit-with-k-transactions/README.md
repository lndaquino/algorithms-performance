## Max Profit with k transactions (maxProfitWithKTransactions)
This algorithm returns maximum profit on buying/selling stocks making until k transactions.
It receives an array of prices representing a single stock (each position in this array is the price in a different day) and k (maximum number of transactions allowed). Each transaction represents a buy/sell movement and you´re allowed to buy if you aren´t holding a stock.
&nbsp;

### Implementations
There is a GO and a Python implementation. In GO we´ve done optimizations to reduce space complexity.
&nbsp;

### Testing
GO code has also tests that can be run using: go test -v