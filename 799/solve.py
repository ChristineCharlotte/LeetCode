class Solution:
    def champagneTower(self, poured: int, query_row: int, query_glass: int) -> float:
        ca = [[0 for i in range(query_row + 2)] for j in range(query_row + 2)]
        ca[0][0] = poured
        for r in range(query_row + 1):
            for l in range(r + 1):
                d = (ca[r][l] - 1) / 2
                if d > 0:
                    ca[r][l] = 1
                    ca[r+1][l] += d
                    ca[r+1][l+1] += d
        re = ca[query_row][query_glass]
        if re <= 1:
            return re
        else:
            return 1