import ctypes

lib = ctypes.cdll.LoadLibrary("./pygo.so")
lib2 = ctypes.CDLL("./pygo.so")

result = lib.Sum(100, 200)
print(result)

lib2.PrintDll()
