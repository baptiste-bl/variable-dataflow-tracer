using System;

namespace ExampleGlobal
{
    class Program
    {
        public static string secret = "MYSUPERSECRETKEY";
        public static int not_after = 60; // 1 minute

        static void Main(string[] args)
        {
            string username = "user";
            string password = "pass";
            string generatedToken = Keygen(username, password, true);
            if (!string.IsNullOrEmpty(generatedToken))
            {
                Console.WriteLine("Generated token: " + generatedToken);
            }
            else
            {
                Console.WriteLine("Login failed.");
            }
        }

        public static bool Login(string username, string password)
        {
            // Mock login function
            return !string.IsNullOrEmpty(username) && !string.IsNullOrEmpty(password);
        }

        public static string EncodeToken(string username, int expiry, string secret)
        {
            // Simplified token encoding for demonstration
            secret = secret ??
            return $"encoded({username},{expiry},{secret})";
        }

        public static string Keygen(string username, string password, bool loginRequired = true)
        {
            if (loginRequired)
            {
                if (!Login(username, password))
                {
                    return null;
                }
            }

            int now = 100; // Mock current time
            string token = EncodeToken(username, now + not_after, secret);

            return token;
        }
    }
}
