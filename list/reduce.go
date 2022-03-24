package list

// Reduce reduces slc down to a single element by using fn to build
// up an accumulator from each value in the list, starting with
// start.
//
// For example, with a list of [1, 2, 3] a start value of "0",
// and a function that converts each integer in the list to its
// string representation and appends it to the end of the
// accumulated string, the execution of this function would look
// like:
//
// 	accumulator := "0", index = 0, element = 1
//		new accumulator = "0" + "1" = "01"
// 	accumulator = "01", index = 1, element = 2
//		new accumulator = "01" + "2" = "012"
// 	accumulator = "012", index = 2, element = 3
//		new accumulator = "012" + "3" = "0123"
//	final result = "0123"
func Reduce[ValT, AccumT any](
	slc []ValT,
	start AccumT,
	fn func(AccumT, ValT) AccumT,
) AccumT {
	ret := start
	for _, t := range slc {
		ret = fn(ret, t)
	}
	return ret
}
