# split vs index

Why does the split generate more trash than the index function in part1 folder?

Split function uses index function under hood. And the split function creates a slice and if you don't need all elements of the slice or you can move by the string without the slice, then maybe you will better use the index function.
