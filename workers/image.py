def swap(arr):
    for i in range(0,len(arr)-1,2):
        print(i, end=" ") # current
        print(i+1) # next
        temp = arr[i]
        arr[i] = arr[i+1]
        arr[i+1] = temp

def swapAlternative(arr, n):
    for i in range(n):
        if (i+1 < n):
            swap(arr)

def printArray(arr, n):
    for i in range(n):
        print(arr[i])


arr = [1,3,5,4,6,7]
# brr = [4,5,7,8,5]

swapAlternative(arr, 6)
# reverse(arr, 5)

printArray(arr, 6)