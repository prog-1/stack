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

### Example 1

```
Input: []string{"3", "4", "+", "2", "*"}
Output: 14
```

#### Explanation

Starting from backside

```
* : () * ()
2 : () * (2)
+ : (() + ()) * (2)
4 : (() + (4)) * (2)
3 : ((3) + (4)) * (2)
((3) + (4)) * (2) = 14
```

### Example 2

```
Input: []string{"4", "11", "3", "/", "+"}
Output: 7
```

#### Explanation

Starting from backside

```
+ : () + ()
/ : () + (() / ())
3 : () + (() / (3))
11 : () + ((11) / (3))
4 : (4) + ((11) / (3))
(4) + ((11) / (3)) = 7
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
```
