S = input()
A=[]

for i in range(9):    
    w=str(i+1)
    A.append("ABC"+"0"+"0"+w)

for i in range(9,99):    
    w=str(i+1)
    A.append("ABC"+"0"+w)

for i in range(99,349):
    w=str(i+1)
    A.append("ABC"+w)

A.remove("ABC316")

if S in A:
  print("Yes")
else:
  print("No")