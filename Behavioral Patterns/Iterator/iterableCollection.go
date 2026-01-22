package main

type IterableCollection[T any] interface {
	CreateIterator() Iterator[T]
}
