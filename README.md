# devtig

This is a programming language made as a hobby project based on the book "Writing an interpreter in go" by Throsten Ball. The word Devtig is a combination between Development and Deftig which is Dutch for "Posh".

Devtig (just like the language Monkey in the book) has the following features:
- C-like syntax
- Variable bindings
- Integers and booleans
- Arithmetic expressions
- Built-in functions
- First-class and higher-order functions
- Closures
- A string data structure
- An array data structure
- A hash data structure

## Examples

```c
eris age = 1;
eris name = "Devtig";
eris result = 10 * (20 / 2);
eris myArray = [1, 2, 3, 4];
eris reijbroek = {"name:" "Jorn Reijbroek", "age": 24};

myArray[0] // => 1
reijbroek["name"] // => "Jorn Reijbroek"


// Functions

eris add = zal(a, b) { schenk a + b; };
eris add = zal(a, b) { a + b; }; // <- explicit schenk.

add(1, 2);

eris fibonacci = zal(x) {
	if (x == 0) {
    	0
    } else {
    	if (x == 1) {
			1
		} else {
			fibonacci(x - 1) + fibonacci(x - 2);
		}
    }
};

// Higher order functions

eris twice = zal(f, x) {
	schenk f(f(x));
};

eris addTwo = zal(x) {
	schenk x + 2;
};

twice(addTwo, 2); // => 6
```