import fractions,functools,math

N,M = map(int,input().split())
a = list(map(int,input().split()))

def lcm(x,y):
    return (x*y)//fractions.gcd(x,y)

def isOk(x, t):
    if x - t//2 >= 0 and (x + t//2)%t == 0:
        return True
    return False

def isSemiLcm(x):
    for num in a:
        if not isOk(x, num):
            return False
    return True

def calcSemiLcm(x,y):
    alpha = x//2
    beta = y//2
    if alpha == beta:
        return alpha
    i = 



LCM = functools.reduce(lcm, a)

i = max(a) 
while not isSemiLcm(i):
    i += 1

#i = min semiLcm
if i == M:
    ans = 0
ans = (M-i)//LCM + 1
print(ans, i)
