# Дана строка, содержащая символы '(', ')', '{', '}', '[' и ']'.
#  Нужно определить, является ли вводная строка валидной скобочной последовательностью.

s = "({)}"
t = {"(":")", "{":"}", "[":"]"}

stack = []
for i in s:
    if i in t:
        stack.append(i)
    elif not stack or i in t.values():
        if t[stack[-1]] != i:
            break
        stack.pop()

print(not stack)

# Дан стек чисел, нужно отсортировать его в порядке возрастания,
# используя только дополнительный стек.

stack = [2, 1, 5, 3, 1, -1, 0]
dop_stack = []

while stack:
    i = stack.pop()
    while dop_stack and i < dop_stack[-1]:
        stack.append(dop_stack.pop())
    dop_stack.append(i)

print(dop_stack)