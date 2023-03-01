# data-algorithms-challenges

## Challenge 1 : Maps

Implement a Map structure in Go without using the built-in map type. The struct has to have the following interface:

```
set(key, value): oldValue → Adds or sets an element with the given key. If an element was already in the map with the same key, the old values is returned. Else nil is returned.

get(key): value → The value with the given key or nil if no element is in the map with the given key

remove(key): oldValue → Removes the element with the given key. Returns the removed element or nil

has(key): bool → Returns true if the key is in the map
```

### Requirements:

Key and value can be any kind of type. For example, you can go with int key, string value for simplification.

Optionally you can use generics for the argument types

The methods (get / set / remove / has) have to be faster on your map on average than storing the keys and values in a simple slice (let’s call it trivial solution). Any kind of algorithm can be used under the hood for this.

Prove it in a test with measurements on some generated dataset

What is the complexity ( O(?) ) of the trivial solution and what is the complexity of your solution for the different methods (set / get / remove / has)