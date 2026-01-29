# DisoScript
This is the script which i will base a future project on where you can graph mathematical functions or expressions and solve for them
It will also be useful to automate tools like desmos. The Script has a very particular way of writing. The next Section describes the
features of it and how to write them.

## Syntax
Here are some keywords used in the file

### Keywords
- `( LeftP`
- `+ Plus`
- `% Modulo`
- `!= Inequal`
- `?= IsEqual`
- `) RightP`
- `- Minus`
- `@ Sqrt`
- `== Equal`
- `; Semi-colon`
- `, Comma`
- `* Times`
- `/ Divide`
- `= Declare`

This section will go through each part of the script...

### Variable Declerations
Variables are declares using the var keyword and a identifier
> `var c = 6`
> or
> `var x = y * @x`
 
### Function Declerations
Functions can be declared using using parentheses after a name
> `f(x) = @x`
they can also have multiple parameters using the comma
> `f(x, y) = y * @x`

### Binary and Unary Operations
Some operations like +, -, *, /, or % require 2 operands
> `x = 2 + 4` <br>
> `y = 5 - 6` <br>
> `f(x) = x * 6` <br>
> `x / 4 ?= 24` <br>

Others use one
> `-4` <br> 
> `g(x) = @x` <br>
