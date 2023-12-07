# Подсчет частоты элементов в списке: Используйте словарь для подсчета, сколько
# раз каждый элемент встречается в списке.

dict = ['a', 'a', 'a', 'a', 'a', 'b', 'c', 'e', 'c', 'c', 'e']

my_dict = {}

for i in dict:
    my_dict[i] = my_dict.get(i, 0) + 1

print(my_dict)

# Две суммы: Дан массив чисел и целевое число, определите, есть ли два числа в
# массиве, которые в сумме дают целевое число. Используйте словарь для хранения
# разностей и индексов.

arr = [8,9,1,2,-3,0,1,2]
target = 10

res = {}
for index, i in enumerate(arr):
    comp = target - i
    if comp in res:
        print(arr[res[comp]], arr[index])
        break
    res[i] = index

# Группировка анаграмм: Дан список строк, сгруппируйте анаграммы вместе.
# Словарь может использоваться для хранения отсортированных строк как ключей.

arr = ["mama", "amam", "banan", "nanba"]

for i in range(len(arr)):
    arr[i] = ''.join(sorted(arr[i]))

res = {}
for i in arr:
    res[i] = res.get(i, 0) + 1

print(res)

# Проверка баланса скобок в выражении: Используйте словарь для сопоставления
# открывающих и закрывающих скобок, чтобы проверить, правильно ли они
# сбалансированы в выражении.

stack = []
arr = "()"
dict = {"(": ")", "{": "}", "[": "]"}

stack = []
for i in arr:
    if i in dict:
        stack.append(i)
    elif i in dict.values():
        if not stack or dict[stack.pop()] != i:
            print(False)
            break
else:
    print(not stack)

