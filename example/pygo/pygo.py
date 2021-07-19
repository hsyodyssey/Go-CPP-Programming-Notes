import ctypes

lib = ctypes.cdll.LoadLibrary("./pygo.so")

result = lib.Sum(100, 200)
print(result)

# lib.PrintDll()
