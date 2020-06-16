For educational purposes, I recently took a stab at implementing a doubly linked-list in Go. 
Unknown to me, there is a container/list package provided by the Go standard library that implements
a generic doubly-linked list. I decided to dive into the code to compare its implementation with my 
own. The standard library implementation was much better than my own and I learned a few things from
it. Here I will do a deep dive on the container/list package focusing on its design, nifty tricks it
has employed, and how it was tested.

First, a linked-list is a data structure that provides a collection of data elements, whose order is
not determined by placement in memory (like arrays). Instead, each element of a linked-list points to the next. Each element simply contains data and a pointer to the next element. In the case of a 
doubly linked-list, each element also contains a pointer to the previous element. This allows the
list to be iterated in both forward and reverse directions.

### Defined Types

Go's implementation of a doubly-linked list can be found [here](https://github.com/golang/go/blob/master/src/container/list/list.go). There are two fundamental types exported by the package, `Element`
and `List`. 

The `Element` struct below represents the element type that can be stored in a list. It contains
pointers to the `next` and `prev` elements in the list. Also, it contains a pointer to the `list`
it belongs to and of course the `Value` the element is storing. Note that `Value` is exported
so users can easily get access to the data contained within an `Element`. The other fields
are not exported, but more on this later.

```
type Element struct {
	next, prev *Element
	list *List // The list to which this element belongs.
	Value interface{} // The value stored with this element.
}
```

The `List` struct below represents the doubly-linked list. It contains two fields, a `root` element
and `len`, the length of the list. 

The `root` element is referred in the comments as the "sentinel"
element. For those unaware, a sentinel is defined as a person or
thing that watches. In this case the comments are trying to indicate that the sentinel element is 
used to keep track of the first and last element. The `root` field is a struct, so each time a `List` is created, a zero-valued `root` will also be created. This is important, every `List` will always
contain a root element.

The reason there is a field for the length of the list is so the list does not need to be iterated
every time a user wants to know the length. This makes returning the length O(1) instead of O(n).

```
type List struct {
	root Element // sentinel list element, only &root, root.prev, and root.next are used
	len  int // current list length excluding (this) sentinel element
}
```

### List Data Type

Great, so now we understand the simple data structures used to implement the linked-list. But what
type of data can a `List` hold? Is it strongly typed like an array or slice? The answer is no. 
If we look at the `Element` definition, we can see that the `Value` field is of type `interface{}`
This means that `Value` can be a type that implements the methods contained within `interface{}`.
Since there are no methods contained within `{}` every type can be used. To access data stored 
within an `Element`, we must do a type assertion on `Value`.

```
element := Element{Value: "foo"}
foo := element.Value.(string) 
```

Go does not have a language feature called "Generics" so this is the easiest way to make `List`
capable of containing elements of any data type.

### Initialization

So, now we understand the data structures that are provided and how any type of data can be stored
in the list. How do we create a new list and what happens under the hood when we do so?

A list should be created using the provided builder function `New()`.

`func New() *List { return new(List).Init() }`

This functions returns a pointer to a `List` and does so by executing two steps. Lets unpack what is
going on in this function. The first thing the function does is `new(List)`. The `new` keyword in 
Go returns a zero-value initialized pointer to the provided type. So in this case it returns a pointer
to a `List` where the `root` element has nil pointers for the `next` and `prev` fields, a nil
pointer to the `list` it belongs to and an unspecified `Value`. The second step is calling `Init()` on
the zero-valued `List` pointer.

```
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}
```

Currently, the `List` is empty so the `Init` function sets `next` and `prev` elements to the `root`
itself, and sets the length of the list to 0. The length is already 0, but it is explicitly set here
in case someone wants to use `Init` to reset a list that was not empty.

### Lazy Initialization

What happens if someone creates a list like l := List{}? Well, `Init()` won't be called
and the list will not be ready for use. However, the authors of this package were smart and
implemented lazy initialization, such that this is not a problem. 

```
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}
```

In any function that modifies the list, this function is the first thing that is called. It checks
if the next element of the root is nil and if so perform initialization. This works because
root is only nil if Init() was not called.

### Insertion

Now how do we add items into the list. The package provides 4 methods for adding single elements to
the list; `PushFront`, `PushBack`, `InsertBefore`, and `InsertAfter`. Under the hood the functions
rely on a single helper function called `insert`. Let's see how this works.

```
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}
```

This function will insert `e` after `at`, returning `e`. What it needs to do is break existing links
between `at` and the next element `at.next` and reconstruct links so that `e` is between `at` and 
`at.next`. I could describe this process but I think the graphic below does a better job.

![alt text](../static/asset/linked_list.png)










- Insert helper
	- Insert
	- AddAfter
	- AddBefore
- Iterating
	- Next and Prev functions return nil when sentinel element is hit
- Test helper
	- CheckLen
	- CheckPointers

