def f():
	return 4

def mul(x, y):
	return x * y

def mulr(x, y):
	if y == 0:
		return 0;
	return x + mulr(x, y - 1)

def islol(x, lol = 707):
	return x == lol

islol(1337)

#f()

#print("The answer is", mul(2, 5) + mulr(4, 8), sep=" ")
