package primes

type Primes int

func (primes Primes) GetPrimesFactor() (primesFactor []int) {
  if int(primes) == 1 {
	primesFactor = []int{}	
  } else {
    
    for testFactor := 2; testFactor <= int(primes); testFactor++ {
      canDivideByTestFactor := int(primes) % testFactor == 0
      if canDivideByTestFactor {
        newPrimes := Primes(int(primes) / testFactor)
        primesFactor = append (primesFactor, testFactor)
        primesFactor = append (primesFactor, newPrimes.GetPrimesFactor()...)
        break;
      }
    }
  }
  return 
}