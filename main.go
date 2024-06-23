package go_cache

type Node struct {
	key        string
	value      []byte
	prev, next *Node // Doubly Linked List
}

type LRUCache struct {
	capacity   int              // Cache capacity
	cacheMap   map[string]*Node // KV pair
	head, tail *Node
}

func New(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cacheMap: make(map[string]*Node),
		head:     head,
		tail:     tail,
	}
}

func (cache *LRUCache) Get(key string) []byte {
	if node, exists := cache.cacheMap[key]; exists {
		cache.moveToFront(node)
		return node.value
	}
	return nil
}

func (cache *LRUCache) Put(key string, value []byte) {
	if node, exist := cache.cacheMap[key]; exist {
		node.value = value
		cache.moveToFront(node)
	} else {
		if len(cache.cacheMap) >= cache.capacity {
			delete(cache.cacheMap, cache.tail.prev.key) // HEAD <-> D <-> A <-> B <-> C <-> E <-> TAIL ---- DELETE E KEY
			cache.removeNode(cache.tail.prev)           // HEAD <-> D <-> A <-> B <-> C <-> TAIL ----- REMOVED E
		}
		newNode := &Node{
			key:   key,
			value: value,
		}
		cache.cacheMap[key] = newNode
		cache.addNode(newNode)
	}
}

func (cache *LRUCache) addNode(node *Node) {
	node.next = cache.head.next // D -> A
	node.prev = cache.head      // HEAD <- D -> A
	cache.head.next.prev = node // HEAD <- D <-> A
	cache.head.next = node      // HEAD <-> D <-> A

}

func (cache *LRUCache) removeNode(node *Node) {
	// A-B-C-D-E
	// NODE IS D
	prev := node.prev // previous node is C
	next := node.next // next node is E
	prev.next = next  // C CONNECTED TO AS NEXT C->E
	next.prev = prev  // C CONNECTED TO AS PREVIOUS C<-E
	// A-B-C-E
}

func (cache *LRUCache) moveToFront(node *Node) {
	cache.removeNode(node)
	cache.addNode(node)
}
