# Table of Contents

- [Table of Contents](#table-of-contents)
- [Go Basic Data Types](#go-basic-data-types)
  - [Boolean Type](#boolean-type)
  - [String Type](#string-type)
  - [Signed Integers](#signed-integers)
  - [Unsigned Integers](#unsigned-integers)
  - [Aliases](#aliases)
  - [Floating-Point Numbers](#floating-point-numbers)
  - [Complex Numbers](#complex-numbers)





# Go Basic Data Types

This document provides an overview of the basic data types in Go, including their sizes, ranges, and usage.

---

## Boolean Type
- **bool**: Represents `true` or `false`.  
  - Size: 1 byte (implementation may vary)  
  - Used for logical values.

---

## String Type
- **string**: Sequence of Unicode characters.  
  - Immutable.  
  - Size depends on the content (stores a pointer + length internally).

---

## Signed Integers
| Type    | Size / Range                                                                                                    |
| ------- | --------------------------------------------------------------------------------------------------------------- |
| `int`   | Platform-dependent integer (32 bits on 32-bit systems, 64 bits on 64-bit systems). Can be negative or positive. |
| `int8`  | 8-bit signed integer. Range: -128 to 127                                                                        |
| `int16` | 16-bit signed integer. Range: -32,768 to 32,767                                                                 |
| `int32` | 32-bit signed integer. Range: -2,147,483,648 to 2,147,483,647                                                   |
| `int64` | 64-bit signed integer. Range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807                           |

---

## Unsigned Integers
| Type      | Size / Range                                                                                                  |
| --------- | ------------------------------------------------------------------------------------------------------------- |
| `uint`    | Platform-dependent unsigned integer (32 bits on 32-bit, 64 bits on 64-bit). Range: 0 to max value of int type |
| `uint8`   | 8-bit unsigned integer. Range: 0 to 255                                                                       |
| `uint16`  | 16-bit unsigned integer. Range: 0 to 65,535                                                                   |
| `uint32`  | 32-bit unsigned integer. Range: 0 to 4,294,967,295                                                            |
| `uint64`  | 64-bit unsigned integer. Range: 0 to 18,446,744,073,709,551,615                                               |
| `uintptr` | Unsigned integer large enough to store the uninterpreted bits of a pointer. Used for low-level memory access  |

---

## Aliases
- **byte**: Alias for `uint8`. Commonly used to represent raw bytes of data.  
- **rune**: Alias for `int32`. Represents a Unicode code point, allowing work with characters beyond ASCII.

---

## Floating-Point Numbers
| Type      | Description                                                                                           |
| --------- | ----------------------------------------------------------------------------------------------------- |
| `float32` | 32-bit IEEE 754 floating-point number. Precision: ~7 decimal digits                                   |
| `float64` | 64-bit IEEE 754 floating-point number. Precision: ~15 decimal digits. Most commonly used for decimals |

---

## Complex Numbers
| Type         | Description                                                                |
| ------------ | -------------------------------------------------------------------------- |
| `complex64`  | 64-bit complex number: 32 bits for real part + 32 bits for imaginary part  |
| `complex128` | 128-bit complex number: 64 bits for real part + 64 bits for imaginary part |

---

This file provides a quick reference for Go's fundamental data types.
