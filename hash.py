# TESTING hashlib   

import hashlib

#print(hashlib.sha256(b"sarthak").hexdigest())

#text = "hello world"
#hash_value = hashlib.sha256(text.encode()).hexdigest()
#print(hash_value)
 
 
print(hashlib.sha256(b'hello').hexdigest())

print(hashlib.sha256(b'Hello').hexdigest())

text = "sarthak"
print(hashlib.sha256(text.encode()).hexdigest())
