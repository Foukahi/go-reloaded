# Go Text Modifier

This Go project is designed to perform a series of text modifications on a given input file and output the modified text to a specified output file. The tool adheres to best practices and includes unit tests to ensure functionality.

## Features

- *Hexadecimal to Decimal Conversion*: Converts hexadecimal numbers preceding (hex) to their decimal equivalent.
- *Binary to Decimal Conversion*: Converts binary numbers preceding (bin) to their decimal equivalent.
- *Case Modification*: Changes the case of words preceding (up), (low), or (cap) to uppercase, lowercase, or capitalized, respectively.
- *Case Modification with Count*: Modifies the case of a specified number of words when followed by (up, <number>), (low, <number>), or (cap, <number>).
- *Punctuation Formatting*: Ensures correct spacing and placement of punctuation marks.
- *Quotation Formatting*: Correctly places quotation marks ' around words without additional spaces.
- *Article Adjustment*: Changes a to an before words starting with a vowel or h.

## Getting Started

Compile the Go code and run the resulting binary with the input and output file names as arguments to start using the tool.

## Usage

```bash
go run main/main.go sample.txt result.txt
