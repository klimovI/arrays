package arrays

/*
Find returns a pointer to the first element in the provided slice for which
the specified predicate function returns true, or nil if no such element is found.

Parameters:
  - predicate: A function that takes an index and a value, and returns true if the
    value satisfies the desired condition.
  - slice: The slice to search.

Returns:
  - A pointer to the first element in the slice for which the predicate function
    returns true, or nil if no such element is found.
*/
func Find[V any](predicate func(index int, value V) bool, slice []V) *V {
	for i := 0; i < len(slice); i++ {
		if predicate(i, slice[i]) {
			return &slice[i]
		}
	}

	return nil
}

/*
FindIndex returns the index of the first element in the provided slice for which
the specified predicate function returns true, or -1 if no such element is found.

Parameters:
  - predicate: A function that takes an index and a value, and returns true if the
    value satisfies the desired condition.
  - slice: The slice to search.

Returns:
  - The index of the first element in the slice for which the predicate function
    returns true, or -1 if no such element is found.
*/
func FindIndex[V any](predicate func(index int, value V) bool, slice []V) int {
	for i := 0; i < len(slice); i++ {
		if predicate(i, slice[i]) {
			return i
		}
	}

	return -1
}

/*
Filter returns a new slice containing only the elements from the provided slice
for which the specified predicate function returns true.

Parameters:
  - predicate: A function that takes an index and a value, and returns true if the
    value should be included in the result slice.
  - slice: The slice to filter.

Returns:
  - A new slice containing only the elements from the provided slice for which the
    predicate function returns true.
*/
func Filter[V any](predicate func(index int, value V) bool, slice []V) []V {
	result := make([]V, 0, len(slice))

	for i, v := range slice {
		if predicate(i, v) {
			result = append(result, v)
		}
	}

	return result
}

/*
ForEach applies the specified action function to each element of the provided slice.

Parameters:
- action: A function that takes an index and a value, and performs some action on
  the value.
- slice: The slice to iterate over.
*/
// p.s. just use range :)
func ForEach[V any](action func(index int, value V), slice []V) {
	for i, v := range slice {
		action(i, v)
	}
}

/*
Map applies the specified transform function to each element of the provided slice,
and returns a new slice containing the transformed values.

Parameters:
  - transform: A function that takes an index and a value, and returns the transformed
    value.
  - slice: The slice to transform.

Returns:
- A new slice containing the transformed values.
*/
func Map[V, R any](transform func(index int, value V) R, slice []V) []R {
	result := make([]R, len(slice))

	for i, v := range slice {
		result[i] = transform(i, v)
	}

	return result
}

/*
Reduce applies the specified reducer function to the elements of the provided slice,
and returns a single result value.

Parameters:
  - reducer: A function that takes an accumulator value, an index, and a value, and
    returns a new accumulator value.
  - slice: The slice to reduce.
  - initialAccumulator: The initial value for the accumulator.

Returns:
- The final accumulator value.
*/
func Reduce[V, A any](
	reducer func(accumulator A, index int, value V) A,
	slice []V,
	initialAccumulator A,
) A {
	acc := initialAccumulator

	for i, v := range slice {
		acc = reducer(acc, i, v)
	}

	return acc
}
