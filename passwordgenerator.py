import random
import string

length = int(input("Enter the length of the password: "))

if length < 8:
    print("Password length must be at least 8 characters.")
    exit()

special_characters = "!@#$&*"
all_characters = string.ascii_lowercase + string.ascii_uppercase + string.digits + special_characters

password = [random.choice(string.ascii_lowercase),random.choice(string.ascii_uppercase),
random.choice(string.digits),random.choice(special_characters)]

for _ in range(length - 4):
    password.append(random.choice(all_characters))
random.shuffle(password)

password = "".join(password)
print(password)

