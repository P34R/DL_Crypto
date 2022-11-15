# ECPoint
## Implement ECPoint class to work with algebra on Elliptic Curves
All declared functions have been implemented. You can change curve parameters inside `ECPoint.go` file, but it's hardcoded (they are global variables also). Also, generator point is static `G=(15,13)`. I chose `y^2=x^3+7 % 17` curve.  
Added PrintlnECPoint function (Print but with '\n' at the end).

## How to run
The software implementation is written in the Golang language.

Clone it using - `git clone https://github.com/P34R/DL_Crypto` command.  
There is some "native" tests inside main.go.     
To launch it, you should have Golang installed (1.18)  
Use `go run main.go` in **ECPoint** directory after cloning.
