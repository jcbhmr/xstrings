# xstrings

🌸 `xstrings` is to `strings` as `xiter` is to `iter`

## Installation

```sh
go get github.com/jcbhmr/xstrings
```

## Usage

[📚 See the documentation for details.](https://pkg.go.dev/github.com/jcbhmr/xstrings)

```go
// https://www.poetryfoundation.org/poems/42916/jabberwocky
fmt.Println(xstrings.Dedent(`
    ’Twas brillig, and the slithy toves
          Did gyre and gimble in the wabe:
    All mimsy were the borogoves,
          And the mome raths outgrabe.

    ...
`))
// Output:
// ’Twas brillig, and the slithy toves
//        Did gyre and gimble in the wabe:
//  All mimsy were the borogoves,
//        And the mome raths outgrabe.
```

## Development

- Make sure each function has at least one example
- Make sure each function has at least one test
