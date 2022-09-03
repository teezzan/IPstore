Answers
=======

What would you do differently if you had more time?
--------
- I would use a data structure with less spacetime complexity. Hashtables trade speed for space. I tested this with 5 billion IP addresses and it made my laptop stutter - a little bit.
- I would make the storage struct more generic(using Interface{} instead of strings) to allow for more flexibility.

What is the runtime complexity of each function?
-------
`RequestHandled`: Amortized `O(1)`

`Clear`: `O(1)`

`Top100`: `O(n)` (worst case) where n is the number of requested IP address. since n=100, the time complexity tends to `O(1)`

How does your code work?
-------
The most important part of the codebase is the "Storage" implementation. The "Storage" consists of two storage varibles or mechanism.
- *`IpAddressTallyMap`*: A hashtable for mapping an item (IP address) to a count(number).
- *`FrequencyLookupTable`*: An array whose index is the frequency of occurence of items and the value, the collection of such item. To ensure an efficient search, insert and delete operation, the collection is modelled as a hashtable with a runtime complexity of an amortized `O(1)`.

The methods of the "Storage" implementation are enumerated as follows.
- *`Truncate`*: This properly reinitializes the storage variables. A dummy hashmap is insterted at index 0 of *`FrequencyLookupTable`* because there should be no values there.
- *`Init`*: This initializes the storage variables using *Truncate*.
- *`Insert`*: During an insert operation, the code searches *`IpAddressTallyMap`* for the existence of the argument `ip_address`. The current "count" value of key `ip_address` (defaults to `0`) in  *`IpAddressTallyMap`* is incremented by one and updated. In the array *`FrequencyLookupTable`*, at the index value of `ip_address` current "count" value, `ip_address` is added to the hashtable found there and removed from its former index.
- *`Fetch`*: A fetch involves starting from the largest index 0f the array (*`FrequencyLookupTable`*) and extracting the keys inside the hashable found there. This continues until the `limit` (number of requested top items) is reached. It is optimized to skip over empty arrays.

The exposed functions of the library are similar to the ones explained above but with validation of IP address inputs.

The most intersting thing about this implementation is its heavy dependence on the runtime complexity of a hashtable being an amortized O(1). I used Go maps since it is a great implementation of an efficient hashtable.


What other approaches did you decide not to pursue?
----
The first data structure that comes to mind after reading the problem is a heap. However, its runtime complexity of `O(nlogn)` made it less attractive. Coupled with the fact that reads(pops) can be a little distructive for a heap if one doesn't make a copy, I had to discard the thought overall. 

How would you test this?
----
//stress test