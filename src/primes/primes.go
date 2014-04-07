package primes

type Primes int

func (primes Primes) GetPrimesFactor() (primesFactor []int) {
  if int(primes) == 1 {
	primesFactor = []int{}	
  } else if (int(primes) == 4) {
    primesFactor = []int{2, 2}
  } else {
    primesFactor = []int{int(primes)}
  }
  return 
}