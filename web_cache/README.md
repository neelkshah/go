### Things I learned ###

* Variable passing is always by value in Golang; this can make pointers confusing if you're already used to the C 
version.
* You can't reuse http response bodies. At least not directly. Which is why the cached values were always printing as 
blank.
* GoRoutines for background tasks
* Structs as an abstraction in place of classes
* The refreshCache conundrum: The way I visualized the cache was that each layer of the cache would have a configurable 
time-to-live (TTL) parameter at the end of which duration the cache layer would expire and all its contents would be 
shifted to the next farther layer. The contents of the farthest layer would be lost and the closest layer would be 
empty. While at first glance, it seems foolish since there isn't really a need for this, it threw up an interesting 
question. If the cache had just two layers, L1 (hot) and L2 (cold) with a TTL of 10s and 20s respectively, a rather 
confounding race condition occurs at t = 20s on account of the asynchronous nature of execution of the refreshCache 
function. If the cold cache dumps its content first, it would yet contain what used to earlier be the contents of the 
hot cache. But if the hot cache dumped its contents first, the cold cache in turn, would dump them too along with its 
own contents. And there's no clean way of resolving this issue. One could prevent the hot cache from voiding its 
contents before the cold cache, but the question emerges - for how long? If the cold cache were to dump it's contents 
later, say by 10 seconds, the cached content in question would only be available for those 10 seconds. The answer then, 
would be to choose the TTL parameters such that their multiples occur as further apart from each other as possible, and 
accept the unexpected voiding of (even the hottest) cache layers as an unfortunate coincidence. Essentially, the point 
is that the problem is not just about concurrency.
