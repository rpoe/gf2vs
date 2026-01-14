# Field $\mathbb{F}_2$

The finite field of order 2 has 2 elements $\mathbb{F}_2 = \{0, 1\}$ and
the operations addition  $+$ and multiplication $\cdot$. For the
definition see equation
(<a href="#def:fieldoperations" data-reference-type="ref"
data-reference="def:fieldoperations">[def:fieldoperations]</a>).
$$\begin{array}{r l l l l}
+ : & 0 + 0 = 0, & 0 + 1 = 1, & 1 + 0 = 1, & 1 + 1 = 0, \\
\cdot : & 0 \cdot 0 = 0, & 0 \cdot 1 = 0, & 1 \cdot 0 = 0, & 1 \cdot 1 = 1. \\
\end{array} \label{def:fieldoperations}$$ We may use the notation $a b$
instead of $a \cdot b$ omitting the multiplication sign if there is no
ambiguity.

Each of the 2 operations of the field $\mathbb{F}_2$ satisfy the group
axioms (Wikipedia contributors 2026b) for the groups
$G_+ : (\mathbb{F}_2, +)$ and $G_{\cdot} : (\mathbb{F}_2, \cdot)$. in
addition both operations are commutative. For reference the group axioms
are repeated here. We use the symbol $\circ$ to denote the binary
operations $+, \cdot$.

Associativity  
   
$\forall a, b, c \in G: (a \circ b) \circ c = a \circ (b \circ c)$.

Identity element $e$  
   
$\exists e \in G, \forall a \in G:  e \circ a = a \text{ and }  a \circ e = a$,
$e$ is unique.

Inverse element $a^{-1}$  
   
$\forall a \in G \enspace \exists b \in G : a \circ b = e \text{ and } b \circ a = e, e$
identity element, $b$ is unique $\forall a$, notation $b = a^{-1}$.

Commutivity  
   
$a \circ b= b \circ a$.

We can look at the field from an algebraic point of view or from a logic
view. In logic the field can be seen as the boolean variables $F = 0$
and $T = 1$. The boolean operations are disjunction $\vee$ (Wikipedia
contributors 2025d), contravalence $\oplus$  (Wikipedia contributors
2025a) and conjunction $\wedge$ (Wikipedia contributors 2025c). The
definition is repeated in equation
(<a href="#def:booleanoperation" data-reference-type="ref"
data-reference="def:booleanoperation">[def:booleanoperation]</a>).
$$\begin{array}{r l l l l}
\vee : & 0 \vee 0 = 0, & 0 \vee 1 = 1, & 1 \vee 0 = 1, & 1 \vee 1 = 1, \\
\oplus : & 0 \oplus 0 = 0, & 0 \oplus 1 = 1, & 1 \oplus 0 = 1, & 1 \oplus 1 = 0, \\
\wedge : & 0 \wedge 0 = 0, & 0 \wedge 1 = 0, & 1 \wedge 0 = 0, & 1 \wedge 1 = 1. \\
\end{array} \label{def:booleanoperation}$$ Please note the operations
$\oplus$ and $\wedge$ are identically defined as $+$ and $\cdot$ and
hence satisfy the group axioms. But the operation $\vee$ does not
satisfy the group axioms, there is no inverse element. In the remaining
chapters we will use the notation $+, \cdot, \vee$ for the operations
only.

# Vector Space $\mathbb{F}_2^n$

We define the vector space $\mathbb{F}_2^n$ over the field
$\mathbb{F}_2$ as set $V$ of vectors $v$ of $n$ elements of the field
together with the binary operation addition and the binary function
scalar multiplication
(<a href="#def:vectoroperations" data-reference-type="ref"
data-reference="def:vectoroperations">[def:vectoroperations]</a>).
$$u + v = w, \enspace u, v, w \in V,  \quad
a \cdot v = w, \enspace a \in \mathbb{F}_2, v, w \in V. 
\label{def:vectoroperations}$$ We apply the addition element-wise and we
multiply the scalar with each element of the vector. This definition is
similar to the definition in (Wikipedia contributors 2025e).

We use the notation $(v_i) := v$ for the vector $v$ with the components
$v_i$ .

In addition we define 2 constants of $\mathbb{F}_2^n$:

Zeros  
   
$\vmathbb{0}$ Zero vector were all components are $0$.

Ones  
   
$\vmathbb{1}$ Vector were all components are $1$.

The axioms of a vector space are satisfied (Wikipedia contributors
2025e):

Associativity of vector addition  
   
$u + (v + w) = (u + v) + w, \enspace \forall u, v, w \in \mathbb{F}_2^n$.

Commutativity of vector addition  
   
$u + v = v + u, \enspace \forall u, v \in \mathbb{F}_2^n$.

Identity element of vector addition  
   
$\exists \vmathbb{0} \in \mathbb{F}_2^n : v + \vmathbb{0} = v, \enspace \forall v \in \mathbb{F}_1^n$.

Inverse elements of vector addition  
   
$\forall v \in \mathbb{F}_2^n \enspace \exists -v \in \mathbb{F}_2^n : v + (-v) = \vmathbb{0}, -v = v$,
each vector is its own additive inverse.

Compatibility of scalar multiplication with field multiplication  
   
$a (b v) = (a b ) v, \enspace a, b \in \mathbb{F}_2, v \in \mathbb{F}_2^n$.

Identity element of scalar multiplication  
   
$1 v = v, \enspace 1 \in \mathbb{F}_2, v \in \mathbb{F}_2^n$, $1$ is the
multiplicative identity of $\mathbb{F}_2$.

Distributivity of scalar multiplication with respect to vector addition  
   
$a (u + v) = a u + a v, \enspace a \in \mathbb{F}_2, u, v \in \mathbb{F}_2^n$.

Distributivity of scalar multiplication with respect to field addition  
   
$(a + b) v = a v + b v, \enspace a, b \in \mathbb{F}_2, v \in \mathbb{F}_2^n$.

In this vector space we are not limited to the operations vector
addition and scalar multiplication. We can use the boolean operations
too.

Complement, Not  
   
$\overline v = \vmathbb{1} - v = \vmathbb{1} + v$, swap all bits.

Disjunction, Or  
   
$u \vee v = (u_i) \vee (v_i) = (u_i \vee v_i)$, element wise Or.

Contravalence, Xor  
   
$u \oplus v = u + v = (u_i) + (v_i) = (u_i + v_i)$, element wise xor,
duplicate of vector addition.

Conjunction, And  
   
$u \wedge v = (u_i) \cdot (v_i) = (u_i \cdot v_i)$, element wise And.

As we apply the operations element wise, we satisfy the laws of
associativity and commutativity.

We use some more definitions to cover the further properties of a vector
space:

Unit vector  
   
We define the unit vectors $e_i, i = 1, \dots, n$ of the vector space as
the vectors where the $i$th element is $x_i = 1$ and all other elements
are $0$.

& e\_i = (x\_k), &  
& x\_k = {

ll 1, & k = i,  
0, & k i,  

. x\_k \_2, e\_i \_2^n. &  

Generating system  
   
We define the subspace $\mathbb{E}= \{e_i\}$,
$\mathbb{E}\subset \mathbb{F}_2^n$ of vectors $e_i$. The subspace
$\mathbb{E}$ forms a generating system. Obviously each vector $v$ of
$\mathbb{F}_2^n$ is a linear combination of the scalars $a_i,$ and the
$e_i$.

& v = \_i=1^n a\_i e\_i, a\_i \_2, e\_i , v \_2^n. &

So the subset $\mathbb{E}$ is a span of $\mathbb{F}_2^n$. In this vector
space it is the only span. And the decomposition of a vector $v$ in a
linear combination of unit vectors $e_i$ is unique.

Basis  
   
The subspace $\mathbb{E}$ is the one and only basis of the vector space
$\mathbb{F}_2^n$.

Index  
   
We name $i = 1, \dots n$ of $e_i$ the index of a unit vector in the
basis.

Norm  
   
We define the Norm $|v|$ of a vector $v \in \mathbb{F}_2^n$ to be its
Hamming weight (Wikipedia contributors 2025b). In this case the count of
ones of the vector. The value of the norm is an element of the set
$\{0, 1, \dots n\} \ne \mathbb{F}_2, n > 2$, In contrast to usual vector
spaces for example on $\mathbb{Z}$, where the norm of a vector is an
element of $\mathbb{Z}$.

Inner product  
We define the inner product of 2 vectors, to be the norm of the product
of 2 vectors:  
$< u, v > = | u \cdot v |$.

Orthogonality  
   
$< u, v > = 0$,  
we say 2 vectors are orthogonal if the inner product is $0$. Please note
the inner product of any vector with $\vmathbb{0}$ is $0$.

Poeppel, Ralf. 2026. “Package Documentation Gf2vs.”
<https://pkg.go.dev/github.com/rpoe/gf2vs>.

Wikipedia contributors. 2025a. “Exclusive or — Wikipedia, the Free
Encyclopedia.”
<https://en.wikipedia.org/w/index.php?title=Exclusive_or&oldid=1316886803>.

———. 2025b. “Hamming Weight — Wikipedia, the Free Encyclopedia.”
<https://en.wikipedia.org/w/index.php?title=Hamming_weight&oldid=1306107874>.

———. 2025c. “Logical Conjunction — Wikipedia, the Free Encyclopedia.”
<https://en.wikipedia.org/w/index.php?title=Logical_conjunction&oldid=1324909528>.

———. 2025d. “Logical Disjunction — Wikipedia, the Free Encyclopedia.”
<https://en.wikipedia.org/w/index.php?title=Logical_disjunction&oldid=1317551960>.

———. 2025e. “Vector Space — Wikipedia, the Free Encyclopedia.”
<https://en.wikipedia.org/w/index.php?title=Vector_space&oldid=1326882436>.

———. 2026a. “Finite Field — Wikipedia, the Free Encyclopedia.”
<https://en.wikipedia.org/w/index.php?title=Finite_field&oldid=1330855394>.

———. 2026b. “Group (Mathematics) — Wikipedia, the Free Encyclopedia.”
<https://en.wikipedia.org/w/index.php?title=Group_(mathematics)&oldid=1330839314>.
