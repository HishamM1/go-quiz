# GO Quiz

A simple timed terminal quiz game written in Go. with the benifit of go routines and channels.

## Installation

1. Clone the repository
2. Run `go build -o go-quiz` in the repository directory
3. Run `./go-quiz` to start the game

## Flags

- `-csv`: Path to the CSV file containing the quiz questions
- `-limit`: Time limit for the quiz in seconds

## Example

`./go-quiz -csv=problems.csv -limit=30`

## CSV Format

The CSV file should have two columns: question and answer. The first row should be the header row.
like this:

```csv
2+2,4
3+3,6
```
