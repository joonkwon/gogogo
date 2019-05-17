# Implement SumOfMultiples function
```
func SumOfMultiples(limit int, divisors []int) int{
    ...
}
```
**SumOfMultiples** will accept limit (integer) and divisors (slice of integers) and return the sum of multiples of divisors (integer) that are less than the limit. If two multiples are the same, only one will be calculated into the sum.

## Examples
SumOfMultiples(12, [3,5]) => 33
multiples of 3 below 12: 3, 6, 9
multiples of 5 below 12: 5, 10

SumOfMultiples(15, [3,4,5]) => 57
multiples of 3 below 15: 3, 6, 9,12
multiples of 4 below 15: 4, 8, 12
multiples of 5 below 15: 5, 10
sum: 3 + 4 + 5 + 6 + 8 + 9 + 10 + 12 = 57

SumOfMultiples(0, [3,4]) => 0

SumOfMultiples(5, []) => 0

## Computing Union of Multiples
In summary, when there is common multiples for divisors, only one of those repeated multiples are included in the sum - only union of multiples of each divisors are summed. 
For this, we can loop through and remove the repitition.

However, we can use some data structure like Set here since a set has only unique elements. It might be not so efficient. But think of creating a Set-like data structure to remove repitition and make code looks cleaner. 

