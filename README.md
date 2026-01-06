# gf2vs
gf2vs is a go package implementing the vector space of the Galois Field of order 2. It is sometimes called bit array (also known as bit map, bit set, bit string, or bit vector).
The vectors are defined as special type.
In addition to math/bits it implements functions of the vector space of a given size.
Each vector is constraint to the vector space given at creation time. The unit vectors
are considered the base of the vector space. There are functions for verifying is a vector
a base vector. The boolean operations and the vector operations are implemented.
The count of ones is considered the norm of the vectors. It is the l_1 norm, or hamming weight
of the vector. Some times this function is named popcount. It is the result of the scalar product.
