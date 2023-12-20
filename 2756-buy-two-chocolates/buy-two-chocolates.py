class Solution:
    def buyChoco(self, prices: List[int], money: int) -> int:
        if len(prices) < 2:
            return money
        arr = sorted(prices)
        price = arr[0] + arr[1]

        if price > money:
            return money
        else:
            return money - price