# Stack / Recursion

## Exercise 1

[Reverse Polish
Notation](https://en.wikipedia.org/wiki/Reverse_Polish_notation) is a method
for representing expressions in which the operator symbol is placed after the
arguments being operated on.

For example, the following RPN expression will produce the sum of 2 and 3,
namely 5: 2 3 +.

Write a function `func rpn(expr []string) int` that evaluates the value of an
expression in RPN defined by `[]string` type. Each string may be an integer
or an operator. Valid operators are +, -, *, /.

> NOTE: Assume the input is always correct.

When implementing this assignment, follow the following rules:

* You cannot use recursion.
* Input slice elements must be used in the same order. Treat the input slice
  as a console or tape input.
* Your solution must use software stack for storing input numbers. The stack
  must be implemented using `[]int`. You can use slice tricks for the standard
  stack operations like `pop` and `push`.
* It is allowed to access only the top element of the stack directly.

### Examples

```
Input: []string{"3", "4", "+", "2", "*"}
Output: 14

In : Stack
3  : [3]
4  : [3 4]
+  : [7]
2  : [7 2]
*  : [14]
```

```
Input: []string{"4", "11", "3", "/", "+"}
Output: 7

In : Stack
4  : [4]
11 : [4 11]
3  : [4 11 3]
/  : [4 3]       # Integer division 11 / 3 = 3.
+  : [7]
```

## Exercise 2

Write a program that simplifies an absolute path for a file (Unix-style) given
as `string`.

In Unix-style file system:

* A period `.` refers to the current directory.
* A double period `..` refers to the directory up a level.
* Any multiple consecutive slashes `//` are treated as a single slash `/`.

In Simplified Absolute path:

* The path starts with a single slash `/`.
* Any two directories are separated by a single slash `/`.
* The path doesn't end with trailing slashes `/`.
* The path only contains the directories on the path from the root directory to
  the target file or directory (i.e., no period `.` or double period `..`)

### Examples

```
Input: "/home/"
Output: "/home"
```

```
Input: "/../"
Output: "/"
```

```
Input: "/foo/./bar/../../baz/"
Output: "/baz"

/foo : /foo
.    : /foo
/bar : /foo/bar
..   : /foo
..   : /
/baz : /baz
/    : /baz
```

## Exercise 3

https://en.wikipedia.org/wiki/Polish_notation

Polish notation (PN), also known as normal Polish notation (NPN), Łukasiewicz
notation, Warsaw notation, Polish prefix notation or simply prefix notation, is
a mathematical notation in which operators precede their operands, in contrast
to the more common infix notation, in which operators are placed between
operands, as well as reverse Polish notation (RPN). It does not need any
parentheses as long as each operator has a fixed number of operands.

Lisp and related programming languages define their entire syntax in prefix
notation (and others use postfix notation).

For example, the following NPN expression will produce the sum of 2 and 3,
namely 5: `+ 2 3`.

Write a function `func npn(expr []string) int` that evaluates the value of an
expression in NPN defined by `[]string` type. Each string may be an integer
or an operator. Valid operators are +, -, *, /.

> NOTE: Assume the input is always correct.

When implementing this assignment, follow the following rules:

* Your solution must NOT use any software stacks.
* Input slice elements must be used in the same order. Treat the input slice
  as a console or tape input.

> HINT: Go through the examples below and notice that recursive evaluator could
> handle two cases:
> 
> 1) Numbers are returned as is.
> 2) Operators recursively invoke the evaluator to obtain both arguments (operands)
>    and return the result.
### Examples

```
Input: []string{"123"}
Output: 123
```

```
Input: []string{"*", "+", "3", "4", "2"}
Output: 14

In : Operands (_ means a placeholder)
*  : (_) * (_)
+  : ((_) + (_)) * (_)
3  : ((3) + (_)) * (_)
4  : ((3) + (4)) * (_)
2  : ((3) + (4)) * (2)
```

```
Input: []string{"*", "2", "+", "3", "4"}
Output: 14

In : Operands (_ means a placeholder)
*  : (_) * (_)
2  : (2) * (_)
+  : (2) * ((_) + (_))
3  : (2) * ((3) + (_))
4  : (2) * ((3) + (4))
```

```
Input: []string{"+", "4", "/", "11", "3"}
Output: 7

In : Operands (_ means a placeholder)
+  : (_) + (_)
4  : (4) + (_)
/  : (4) + ((_) / (_))
11 : (4) + ((11) / (_))
3  : (4) + ((11) / (3))
```
