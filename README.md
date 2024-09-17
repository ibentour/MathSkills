# Statistics Calculator

This Go program calculates basic statistical measures (average, median, variance, and standard deviation) from a set of numbers provided in a text file.

## Features

- Reads numbers from a text file
- Calculates and displays:
  - Average
  - Median
  - Variance
  - Standard Deviation

## Requirements

- Go programming language (https://golang.org/)

## Usage

1. Prepare a text file named `data.txt` in the same directory as the program. Each number should be on a separate line.

2. Run the program using the following command:

   ```bash
   go run main.go data.txt
   ```

3. The program will output the calculated statistics.

## Input File Format

The input file (`data.txt`) should contain one integer per line. Empty lines are ignored.

Example `data.txt`:

```bash
10
15
20
25
30
```

## Output

The program will print the calculated statistics to the console:

```
Average: [calculated average]
Median: [calculated median]
Variance: [calculated variance]
Standard Deviation: [calculated standard deviation]
```

## Error Handling

- If the input file is not provided or is not named `data.txt`, the program will display a usage message.
- If the file cannot be read, an error message will be displayed.
- If any line in the file cannot be converted to an integer, an error message will be shown.

## Functions

- `Average`: Calculates the mean of the input numbers.
- `Median`: Calculates the median of the input numbers.
- `Variance`: Calculates the variance of the input numbers.
- `Standard`: Calculates the standard deviation from the variance.

## Note

All calculations are rounded to the nearest integer for simplicity.
