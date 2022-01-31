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
    private static extern int Initialize();

    [DllImport("libspeechdwrapper.so")]
    private static extern int Speak(GoString text, bool interrupt);

    [DllImport("libspeechdwrapper.so")]
    private static extern int Close();

    private static bool initialized = false;

    static void Main(string[] args) {

      // To Initialize
      int res = Initialize();
      if(res==1)
        initialized = true;

      // To Speak
      if(initialized){
        string msg = "Hello from C#!";
        GoString str = new GoString(msg, msg.Length);
        Speak(str, false);

        msg = "I interrupted";
        str.msg = msg;
        str.len = msg.Length;
        Speak(str, true);
      }

      // To close. IMPORTANT!!
      if(initialized){
        Close();
        initialized = false;
      }
    }
   }
}