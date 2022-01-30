using System;
using System.Runtime.InteropServices;

namespace HelloWorldApplication {

    public struct GoString
    {
        public string msg;
        public long len;
        public GoString(string msg, long len)
        {
            this.msg = msg;
            this.len = len;
        }
    }
   class HelloWorld {

    [DllImport("libspeechdwrapper.so")]
    private static extern void Initialize();

    [DllImport("libspeechdwrapper.so")]
    private static extern void Speak(GoString text, bool interrupt);

    [DllImport("libspeechdwrapper.so")]
    private static extern void Close();
      static void Main(string[] args) {
        Initialize();

        string msg = "Hello from C#!";
        GoString str = new GoString(msg, msg.Length);
        Speak(str, false);

        msg = "I interrupted";
        str.msg = msg;
        str.len = msg.Length;
        Speak(str, true);

        Close();

      }
   }
}