import unittest
import code

class MyTestCase(unittest.TestCase):
    def test_something(self):
        sol = code.Solution()
        self.assertEqual("0246813579", sol.convert("0123456789", 2))
        self.assertEqual("0481357926", sol.convert("0123456789", 3))
        self.assertEqual("0615724839", sol.convert("0123456789", 4))
        self.assertEqual("0817926354", sol.convert("0123456789", 5))
        self.assertEqual("06157b248a39", sol.convert("0123456789ab", 4))
        self.assertEqual("A", sol.convert("A", 1))
        self.assertEqual("AB", sol.convert("AB", 1))
        self.assertEqual("ACB", sol.convert("ABC", 2))
        self.assertEqual("ABDC", sol.convert("ABCD", 3))

if __name__ == '__main__':
    unittest.main()
