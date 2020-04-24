### Things I learned ###

* Variable passing is always by value in Golang; this can make pointers confusing if you're already used to the C 
version.
* You can't reuse http response bodies. At least not directly. Which is why the cached values were always printing as 
blank.
* GoRoutines for background tasks
* Structs as an abstraction in place of classes
