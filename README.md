## Chapter 1

### Make map

https://blog.golang.org/maps

```
m := make(map[string]int)
```

### short hand declaration

```
m := 0.0
``

### Underscore (Blank identifier)

```
_, err := test()

```

### Channel operator <- arrow

```
ch chan<-string
```

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
``

### fmt printF arg

fmt (Args) | Describe
------------ | -------------
%d | decimal integer 
%x, | %o, %b integer in hexadecimal, octal, binary 
%f, | %g, %e floating-point number: 3.141593 3.141592653589793 3.141593e+00 
%t | boolean: true or false 
%c | rune (Unicode code point) 
%s | string 
%q | quoted string "abc" or rune 'c' 
%v | any value in a natural format 
%T | type of any value 
%% | literal percent sign (no operand)

