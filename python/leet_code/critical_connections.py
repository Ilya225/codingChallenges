from typing import List
import sys
sys.setrecursionlimit(10000)

def criticalConnections(n: int, connections: List[List[int]]) -> List[List[int]]:
    graph = {}
    visited = [False] * n
    low_links = [None] * n
    ids = [0]*n
    bridges = []

    for con in connections:
        if con[0] not in graph:
            graph[con[0]] = []
        graph[con[0]].append(con[1])

        if con[1] not in graph:
            graph[con[1]] = []
        graph[con[1]].append(con[0])

    def dfs(node, parent, id):
        visited[node] = True
        ids[node] = id
        low_links[node] = id
        id += 1
        for link in graph[node]:
            if link == parent: continue
            if not visited[link]:
                dfs(link, node, id)
                low_links[node] = min(low_links[node], low_links[link])
                if ids[node] < low_links[link]:
                    bridges.append([node, link])
            else:
                low_links[node] = min(low_links[node], ids[link])
                # low_links[link] = min(low_links[node], low_links[link])

    dfs(0, -1, 0)

    return bridges



def main():
    g = [[0, 1], [1, 2], [2, 0], [1, 3]]
    print(criticalConnections(4, g))



if __name__ == "__main__":
    main()