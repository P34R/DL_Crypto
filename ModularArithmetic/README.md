# Modular arithmetic

## Task statement
Create a program that provides arithmetic operations modulo some number and finding the inverse element in the corresponding group (G).  
The software implementation should have the following capabilities (**_The corresponding implementation functions are indicated in parentheses_**):  
set the module m, according to which the calculations will be carried out. (`SetModule(m uint64)`)  
solve equations of the form a mod m = x. (`LinearEquation(a int64)`)  
solve equations of the form a^b mod m = x. (`PowEquation(a,b int64)` - may return 0 if gcd(a,modulo) != 1 )  
solve equations of the form a*x ≡ b mod m. (`MulEquation(a,b int64)`)  
generate a prime number in the range A to B. (`PrimeRange(a,b int64)`)  

## How to use

The software implementation is written in the Golang language.  

Clone it using - `git clone https://github.com/P34R/DL_Crypto`  
If you want to test it (set your own module, a and b numbers, etc...), you should edit `main.go` (there is all functions, you just need to change integers inside brackets, e.g. line 10: `ar.SetModule(Modulo integer)`)     
To launch it, you should have Golang installed (1.18)  
Use `go run main.go` in **ModularArithmetic** directory after cloning.
