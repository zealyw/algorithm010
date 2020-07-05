## 深度优先搜索 DFS

### 代码模板
#### 递归写法
```
func dfs(node,visited){
    if visited[node] {
        //already visited
        return
    }
    visited[node] = true
    //process current node here
    ...
    for nextNode := range node.Childrens{
        if !vivited[node]{
            dfs(nextNode,vivisted)
        } 
    }
}
```
#### 非递归写法
通过维护一个栈来实现
````
func dfs(root){
    if root == nil{
        return []
    }
    visited,stack := map[root]bool{},[]node{root}
    for len(stack) != 0 {
        node = stack[len(stack)-1]
        visited[node] = true
        process(node)
        nodes = generate_related_nodes(node)
        stack.append(nodes)
    }
}

````
## 广度优先搜索 BFS
通过一个队列来进行实现
```
func(graph,start,end){
    visited := map[node]bool{}
    queue := []node{}
    queue = queud.append(start)
    for len(queue) !=0 {
        node = queue[0]
        visited[node] = true
        process(node)
        nodes = generate_related_nodes(node) 
        queue = queue.appends(nodes)
    }
    ...
}

```
## 贪心算法 Greedy
在每一步选择中都采取在当前状态下最好或者最优的选择
贪心与动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退，
动态规划则会保存以前的结果，并根据以前的结果对当前进行选择，有回退的功能

比较有意思的是 不一定是从前往后贪心 也可能是切入或者从后往前贪心

## 二分查找
1.目标函数的单调性 （单调递增、递减）
2.存在上下界
3.能够通过索引访问
```
left, right = 0, len(array) - 1 
while left <= right: 
	  mid = (left + right) / 2 
	  if array[mid] == target: 
		    # find the target!! 
		    break or return result 
	  elif array[mid] < target: 
		    left = mid + 1 
	  else: 
		    right = mid - 1
```