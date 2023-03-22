# data-algorithms-challenges

## 1. Task: Map - Simple solution

Implement a Map structure in Go without using the built-in map type. The struct has to have the following interface:

```
set(key, value): oldValue → Adds or sets an element with the given key. If an element was already in the map with the same key, the old values is returned. Else nil is returned.

get(key): value → The value with the given key or nil if no element is in the map with the given key

remove(key): oldValue → Removes the element with the given key. Returns the removed element or nil

has(key): bool → Returns true if the key is in the map
```

Requirements:

Key and value can be any kind of type. For example, you can go with int key, string value for simplification.

Optionally you can use generics for the argument types

Use a simple Go slice for your "Simple" solution

 

## 2. Task: Map - Hashing solution

Implement a Hash Map structure in Go without using the built-in map type. The struct has to have the same interface as your Simple solution

Requirements:

The methods (get / set / remove / has) have to be faster on your hash map on average than storing the keys and values in a simple solution from Task 1.

Prove it in a test with measurements on some generated dataset

Use hashing and Hash Buckets for this solution 

What is the complexity ( O(?) ) of the trivial solution and what is the complexity of your solution for the different methods (set / get / remove / has)

 

## 3. Task: Map - Tree solution

Implement a Tree Map structure in Go without using the built-in map type. The struct has to have the same interface as your Simple and Hash Map solution from previous tasks.

Requirements:

The Search complexity has to be O(log(N))

Look for a Tree variant that gives you this complexity

What is the complexity of all other methods (add, remove) on your map?

 

## 4. Task: Making your solutions Thread Safe

Implement a thread-safe layer on the top of all three solutions that makes a Map thread-safe.

Create a runnable go application (main function) that shows that your solutions have issues using them in a multi-threaded environment

Create a common interface that all your solutions implement (make them compatible if they are not right now)

Create a thread-safe layer that accepts (wraps) any of your map solutions and makes it thread-safe

Amend the code in your main with a test to show that it using your wrapper indeed thread-safe

Does the thread safe solution degrades the performance in single-threaded environment? Amend your main function with a test that shows if it does.
