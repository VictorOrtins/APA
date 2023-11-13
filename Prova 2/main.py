import numpy as np

def knapsack(M, N, V, P):
    matriz = np.zeros( (N+1,M+1))
    for i in range(0, N + 1):
        for j in range(0, M + 1):
            if i == 0 or j == 0:
                matriz[i,j] = 0
            elif P[i - 1] <= j:
                matriz[i,j] = matriz[i-1,j]
            else:
                matriz[i,j] = max(matriz[i-1,j-P[i - 1] + V[i-1]], matriz[i-1,j])

    print(matriz)
    return matriz

matriz = knapsack(7, 4, [1,4,5,7], [1,3,4,5])
print(matriz)