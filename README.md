# LRU CACHE (LEAST RECENTLY USED CACHE)

This means that the eviction policy moves the least recently used cached data to the end of the list. When the capacity is full, the data at the end of the list is dropped, and new data is added to the head of the list.

- CACHE CAPACITY IS 5

- A <-> B <-> C <-> D <-> E

- NEW DATA IS 'F' 

- CAPACITY IS FULL

- REMOVE E FROM TAIL, ADD 'F' TO HEAD

- F <-> A <-> B <-> C <-> D