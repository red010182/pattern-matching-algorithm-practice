
## run
```
go run main.go
```

## benchmark
Env: MacbookPro OSX, i5 2.3GHz, RAM 16GB
```
Text: bible.txt (4467663 characters)
Randomly picked 100 patterns. Each pattern has min length 7

============= Aho Corasick =============

=> Aho Corasick use 32.251163ms
=> Found 2 patterns, with total 4467663 compare times

============= Wu Manber =============

=> Wu Manber use 32.091946ms
=> Found 2 patterns, with total 797649 compare times

============= Brute Force =============

=> Brute Force use 1.139785543s
=> Found 2 patterns, with total 6339880 compare times

============= KMP =============

=> KMP use 855.909529ms
=> Found 2 patterns, with total 3193027 compare times

============= Boyer Moore =============

=> Boyer Moore use 1.576866672s
=> Found 2 patterns, with total 446914 compare times

============= Robin Karp =============

=> Robin Karp use 1.315759813s
=> Found 2 patterns, with total 3139711 compare times
```
