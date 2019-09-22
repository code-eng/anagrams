# Anagrams

This is a simple server that has two endpoints:

```
POST /load - takes array of words
GET /get?word=<word> - finds all anagrams for a given word among loaded words
```

## Example of usage

```
curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'

curl 'localhost:8080/get?word=foobar' => ["foobar","boofar"]
curl 'localhost:8080/get?word=raboof' => ["foobar","boofar"]
curl 'localhost:8080/get?word=abba' => ["aabb","baba"]
curl 'localhost:8080/get?word=test' => ["test"]
curl 'localhost:8080/get?word=qwerty' => null
```

## Runing server

```
make build
make run
```

## Runing tests

```
make test
```

# Implementation details

Here described implementation, how words are stored to fast anagram retrieval and some research results.

## How words are stored

To store words, modification of [Anatree](https://en.wikipedia.org/wiki/Anatree) is used. In this modification
there are multiple roots of a tree stored as a `map`. Check `Anatree` type in `anatree/tree.go` for more details.
As for roots and branches sorted order of words is used, because sorting small words is really fast. The algorithm
of construction is:

- take a word. For example: `"=)Sasha!"`
- throw away non-letter characters and convert letters to lower case. `"=)Sasha!" => "sasha"`
- sort letters. `"sasha" => "aahss"`
- make a frequency list (not a map, because with a map order would be lost). `"aahss" => [(a, 2), (h, 1), (s, 2)]`
- add a word `"=)Sasha!"` to Anatree by traversing it using frequency list `[(a, 2), (h, 1), (s, 2)]`
- now we will have a tree like this: `<(a, 2)> => <(h, 1)> => <(s, 2) words: ["=)Sasha!"]>`
- take next word and repeat the process

Eventually there would be a map of maps, that is fast to traverse

## How anagrams are found

The process is similar to adding words to `Anatree`:
- take a word
- convert it to frequency list (same as for adding word algorithm)
- traverse Anatree using frequency list
- get anagrams:)

## Speed discussion

Retrieval of anagrams is so fast, that handler that just returns 200 OK have same latency on my PC. Thus I
decided to not use caching.

I tried server on 270_000 words. Anatree is build in 800ms on average and retrieval happens on 3ms.

In order to make things faster, I tried alternative algorithms of sorting in linear time (Radix, Bucket).
But on words golang's quicksort is faster.

Also I tried to speed up `Anatree` construction with concurrency, but I only achieved more complicated
slower version of tree :(. To build Anatree concurrently it requires whole redesign.

## Problems with this implementation

This implementation strongly relies on loaded words dictionary. For example it is possible to just load a huge dict of
same words. Or it is possible to load same words with arbitrary garbage like `"@123(#%#@  car"`
