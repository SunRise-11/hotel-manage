import unittest

import requests

def url(path):
    return "http://localhost:3000%s" % path

class IntegrationTest(unittest.TestCase):
    def setUp(self):
        pass

    def testHealthCheck(self):
        r = requests.get(url('/health'))
        self.assertEqual(200, r.status_code)

    def test404(self):
        r = requests.get(url('/foo'))
        self.assertEqual(404, r.status_code)
        r = requests.get(url('/bar'))
        self.assertEqual(404, r.status_code)

if __name__ == '__main__':
    unittest.main()
