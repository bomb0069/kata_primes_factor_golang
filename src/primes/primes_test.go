package primes_test

import (
  	. "primes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  	"testing"
)

func TestPrimes(t *testing.T) {
  RegisterFailHandler(Fail)
  RunSpecs(t, "Primes Suite")
}

var _ = Describe("Primes", func() {
    It("should return [] when input is 1", func() {
      primes := Primes(1)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{}))
    })
    It("should return [2] when input is 2", func() {
      primes := Primes(2)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{2}))
    })
    It("should return [3] when input is 3", func() {
      primes := Primes(3)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{3}))
    })
    It("should return [2, 2] when input is 4", func() {
      primes := Primes(4)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{2, 2}))
    })
    It("should return [5] when input is 5", func() {
      primes := Primes(5)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{5}))
    })
    It("should return [2, 3] when input is 6", func() {
      primes := Primes(6)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{2, 3}))
    })
    It("should return [2, 2, 2] when input is 8", func() {
      primes := Primes(8)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{2, 2, 2}))
    })
    It("should return [3, 3] when input is 9", func() {
      primes := Primes(9)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{3, 3}))
    })
    It("should return [2, 2, 7] when input is 28", func() {
      primes := Primes(28)
      Expect(primes.GetPrimesFactor()).To(Equal([]int{2, 2, 7}))
    })
})
var _ = Describe("Minimum PrimesFactor", func() {
    It("should return 2 when input is 2", func() {
      primes := Primes(2)
      Expect(primes.GetMinPrimesFactor()).To(Equal(2))
    })
    It("should return 2 when input is 4", func() {
      primes := Primes(4)
      Expect(primes.GetMinPrimesFactor()).To(Equal(2))
    })
    It("should return 3 when input is 15", func() {
      primes := Primes(15)
      Expect(primes.GetMinPrimesFactor()).To(Equal(3))
    })
})