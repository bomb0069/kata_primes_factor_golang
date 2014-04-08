package primes

type Primes int

func (primes Primes) GetPrimesFactor() (primesFactor []int) {
  if int(primes) == 1 {
	primesFactor = []int{}	
  }else{
    for index := 2; index <= int(primes); index++ {
      if int(primes) % index == 0 {
          primesFactor = append(primesFactor,index)
          primes = Primes(int(primes) / index)
          index--
      }
    }
  }
  return 
}