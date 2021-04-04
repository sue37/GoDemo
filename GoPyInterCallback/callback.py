import ctypes

def register_callback():
    ctypes.CFUNCTYPE(ctypes.c_int, ctypes.c_int)
