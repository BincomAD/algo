# Массивы
# Задача "Найти максимальный элемент":

arr = [1, 2, 4, 5, 6, 10]

print(max(arr))

# Задача "Обратить список":

arr = [1, 2, 3, 4, 5]

print(arr[::-1])

# Задача "Проверка наличия элемента":

arr = [1, 2, 3, 4, 5, 6]

print(6 in arr)
print(9 in arr)

# Задача "Слияние двух отсортированных списков":

# Вариант 1. В тупую
arr2 = [1, 2, 5, 7, 9, 10]
arr = [2, 4, 6, 8, 11]
print(sorted(arr+arr2))

# Вариант 2. По умном O(n)

i, j = 0, 0
res = []
while i < len(arr) or j < len(arr2):
    if i < len(arr) and (j == len(arr2) or (arr[i] < arr2[j])):
        res.append(arr[i])
        i += 1
    elif j < len(arr2):
        res.append(arr2[j])
        j += 1

print(res)

# Односвязные списки

class ListNode:
    def __init__(self, value = 0, next = None):
        self.value = value
        self.next = next

# Добавление элемента в конец списка:

listTail = ListNode(1)
listHead = ListNode(2)

listTail.next = listHead

list = ListNode(3)
listHead.next = list
listHead = list

next = listTail

while next != None:
    print(next.value)
    next = next.next

# Поиск элемента в списке:

arr= [1, 2, 3, 4, 5, 6, 7, 8]

listTail = ListNode(arr[0])
listHead = ListNode(arr[1])

listTail.next = listHead

for i in range(2,len(arr)):
    list = ListNode(arr[i])
    listHead.next = list
    listHead = list

target = 10

node = listTail
while node != None:
    if node.value == target:
        print(True)
    node = node.next


# Двусвязный список
class DoublyNodeList:
    def __init__(self, value=0, next=None, prev=None):
        self.value = value
        self.next = next
        self.prev = prev

arr = [1, 2, 3, 4, 5, 6, 7]
#Добавление элемента в связный список:
tailNode = DoublyNodeList(arr[0])
headNode = DoublyNodeList(arr[1])
tailNode.next = headNode
headNode.prev = tailNode

for i in range(2,len(arr)):
    node = DoublyNodeList(arr[i])
    headNode.next = node
    node.prev = headNode
    headNode = node

node = tailNode
while node != None:
    print(node.value)
    node = node.next

#Поиск элемента в связном списке:
target = 5

node = tailNode
while node != None:
    if node.value == target:
        print(True)

    node = node.next

#Удаление элемента из двусвязного списка
node = tailNode
while node != None:
    if node.value == target:
        if node.prev:
            node.prev.next = node.next

        if node.next:
            node.next.prev = node.prev

    node = node.next

#Обратный обход двусвязного списка:

node = headNode
while node != None:
    print(node.value)
    node = node.prev