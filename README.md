# Approximate Square Root Using Newton's Method in Go

My solution to the ["A Tour of Go"](https://tour.golang.org/flowcontrol/8) problem for finding the square root of a positive integer using Newton's Method.

The function will run through a maximum of 10 iterations of the algorithm to find a good approximation of the square root of the integer passed into the function.  

The initial seed value is based on the integer passed in to the function in order to help in creating a more accurate starting point for the algorithm.

The function can also be terminated prematurely if the successive iterations of the for loop have not changed the square root approximation by more than 0.0000001 (an arbitrary value chosen by me).

You can find more information about Newton's Method at https://en.wikipedia.org/wiki/Newton%27s_method#Square_root_of_a_number
