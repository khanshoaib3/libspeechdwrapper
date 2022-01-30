import ctypes


class GoString(ctypes.Structure):
    _fields_ = [("p", ctypes.c_char_p), ("n", ctypes.c_longlong)]

lib = ctypes.CDLL("libspeechdwrapper.so")

lib.Initialize()

lib.Speak.argtypes = [GoString]

msg = GoString(b"This is a very very long message.", 5)
lib.Speak(msg, False)

msg = GoString(b"I interrupted", 13)
lib.Speak(msg, True)

lib.Close()