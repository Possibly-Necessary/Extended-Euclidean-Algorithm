""" A Python script that calculates the Extended Euclidean Algorithm """

# Function that computes the Extended Euclidean Algorithm
def extdEuclid(a, b):

    k=2 # Index

    err_str = None

    # Check if conditions are satisfied, then begin the algorithm
    if not (a>=0 and b>0 and a>=b):
        str = "Condition are not satisfied... "
        return 0, 0, 0, err_str # if conditions are not satisfied

    # Initilize the (empty) arrays r, s and t of size k
    r = [None for _ in range(k)]
    s = [None for _ in range(k)]
    t = [None for _ in range(k)]

    # Insert the integers a and b
    r[0], r[1] = a, b
    s[0], s[1] = 1, 0
    t[0], t[1] = 0, 1
    # while b (in rk-1) is not 0
    while r[k-1] != 0:
        q = r[k-2]//r[k-1]
        r.append(int(r[k-2]%r[k-1])) # encapsulate results to int(), otherwise it will append float values
        s.append(int(s[k-2]-(q*s[k-1])))
        t.append(int(t[k-2]-(q*t[k-1])))
        k += 1
    return r[len(r)-2], s[len(s)-2], t[len(t)-2], err_str

def gcd(a,b):
    if b==0:
        return a
    return gcd(b, a%b)

# Function to handle user input
def userInpt():

    print("Enter two integers, \n")

    # Enforce only integer inputs for a and b
    # In case float values are entered, only the integer part will be taken
    while True:
        try:
            a = int(float(input("a: ")))
            break
        except ValueError:
            print("Invalid input. Enter a strictly integer value.")

    while True:
        try:
            b = int(float(input("b: ")))
            break
        except ValueError:
            print("Invalid input. Enter a strictly integer value.")

    return a, b


# --------- 'main' --------------------#


# Get user input
a, b = userInpt()

# Execute the Extended Euclidean function and get the results
r, s, t, strr= extdEuclid(a, b)

# Check if the string is not empty, that means the conditions were not satisfied
if strr != None:
    print(strr)

print(f'r: {r}, s: {s}, t: {t}')

# Calculate the gcd(a,b) for the lefthand side in gcd(a,b) = s*a + t*d
gcdResult = gcd(a,b)
print(f'gcd({a},{b}) = {gcdResult}')

# Check if the equation gcd(a,b) = s * a + t * b holds
if gcdResult == s * a + t * b:
    print(f'gcd({a}, {b}) = ({s}) * {a} + ({t}) * {b}')
    print(f'{gcdResult} = {s*a + t*b}')
    print(f'Equation holds.')
else:
    print(f'gcd({a}, {b}) = ({s}) * {a} + ({t}) * {b}')
    print(f'{gcdResult} = {s*a + t*b}')
    print(f'Equation does not hold.')
