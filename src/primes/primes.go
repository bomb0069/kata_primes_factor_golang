package primes

type Primes int

func (primes Primes) GetPrimesFactor() (primesFactor []int) {
  if int(primes) == 1 {
	primesFactor = []int{}	
  } else {
    minimumPrimesFactor := primes.GetMinPrimesFactor()
    primesFactor = []int{minimumPrimesFactor}

    subPrimes := Primes(int(primes) / minimumPrimesFactor)
    primesFactor = append (primesFactor, subPrimes.GetPrimesFactor()...)
  }
  return 
}

func (primes Primes) GetMinPrimesFactor() (minimumPrimeFactor int) {
    for testFactor := 2; testFactor <= int(primes); testFactor++ {
      if int(primes) % testFactor == 0 {
        return testFactor
      }
    }
  	return
}