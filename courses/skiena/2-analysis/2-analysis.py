from typing import List, Union

def polynomial(A: List[Union[int]], x: Union[int]) -> Union[int]:
    n = len(A)
    if n == 0:
        raise ValueError("Coefficient list cannot be empty")

    p = A[-1]
    for i in range(n - 2, -1, -1):
        p = A[i] + x * p
    return p



def fib(n):
    if n == 0:
        return 0
    if n == 1:
        return 1
    if n == 2:
        return 2
    return fib(n-1) + fib(n-2)
