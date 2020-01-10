import itertools

N = int(input())
P = "".join(input().split())
Q = "".join(input().split())
L = []
for p in itertools.permutations(map(str, range(1,N+1))):
    L.append("".join(p))
L.sort()
p = -1
q = -1
i = 0
for elem in L:
    if P == elem:
        p = i
    if Q == elem:
        q = i
    i+=1
ans = abs(q-p)
print(ans)
