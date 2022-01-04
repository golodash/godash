# What Is Godash?

Godash is inspired by [lodash](https://github.com/lodash/lodash) nodejs utility library delivering modularity.\
Basically we are going to add their functionality, in our functionality.

If you love to contribute, please do, we appreciate it.

And ofcourse don't forget to read [CONTRIBUTING](/CONTRIBUTING.md) documentation to know about how to contribute in this project.

## Our Goals

This is a growing list, if we achieve our goal on consistency, this list will be extended.

### Array:

1. - [X] Chunk (array, [size=1]) [source](https://lodash.com/docs/latest#chunk)</details>
2. - [ ] Compact (array) [source](https://lodash.com/docs/latest#compact)</details>
3. - [ ] Concat (array, [values]) [source](https://lodash.com/docs/latest#concat)</details>
4. - [ ] Difference (array, [values]) [source](https://lodash.com/docs/latest#difference)</details>
5. - [ ] DifferenceBy (array, [values], [iteratee=_.identity]) [source](https://lodash.com/docs/latest#differenceBy)</details>
6. - [ ] DifferenceWith (array, [values], [comparator]) [source](https://lodash.com/docs/latest#differenceWith)</details>
7. - [ ] Drop (array, [n=1]) [source](https://lodash.com/docs/latest#drop)</details>
8. - [ ] DropRight (array, [n=1]) [source](https://lodash.com/docs/latest#dropRight)</details>
9. - [ ] DropWhile (array, [predicate=_.identity]) [source](https://lodash.com/docs/latest#dropWhile)</details>
10. - [ ] Fill (array, value, [start=0], [end=array.length]) [source](https://lodash.com/docs/latest#fill)</details>

## For developers

### Go1.18beta

How to install and work with it?

Just take a look [here](https://go.dev/dl/#go1.18beta1).
Until go1.18 release(February 2022 - [draft release notes](https://tip.golang.org/doc/go1.18)) the project uses go1.18beta version.

### How to run benchmarks and test cases?

#### Test Cases

```bash
go1.18beta1 test github.com/gotorn/godash/<your desired package>
```

#### Benchmarks

```bash
go1.18beta1 test -bench=. github.com/gotorn/godash/<your desired package>
```