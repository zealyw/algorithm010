## 字典树
是一种树形结构，典型应用是用于统计和排序大量的字符串，经常被用在搜索引擎系统用于文字词频统计
优点：最大限度减少字符串的比较，查询效率比哈希表高。
基本性质：
1. 节点本身不存完整单词
2. 从根节点到某一个节点，路径经过的字符连接起来，就是该节点对应的字符串
3. 每个节点的所有子节点路径代表的字符都不相同
```go
type Trie struct {
	isEnd bool
	next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	for _,char := range word{
		if this.next[char - 'a'] == nil{
			this.next[char - 'a'] = new(Trie)
		}
		this = this.next[char -'a']
	}
	this.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _,char := range word{
		if this.next[char - 'a'] == nil {
			return false
		}
		this = this.next[char - 'a']
	}
	if this.isEnd == false{
		return false
	}
	return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _,char := range prefix{
		if this.next[char - 'a'] == nil{
			return false
		}
		this = this.next[char - 'a']
	}
	return true
}
```
## 并查集
```go
func find(x int, p []int) {
    r := x
    for p[r] != r {
        r = p[r]
    }
    i := x
    var j int
    for i != r { // 路径压缩，将树高变为2
        j = p[i]
        p[i] = r
        i = j
    }

    return r
}
// 将节点x  y所在集合合并
func union(x, y int, p []int) {
    x_root := find(x, p)
    y_root := find(y, p)
    if x_root != y_root {
        p[x_root] = y_root
    } 
}
```

## AVL树
平衡因子：它的左子树的高度减去它的右子树的高度，值在{-1, 0, 1}。
通过旋转操作来进行平衡
* 左旋
* 右旋
* 左右旋
* 右左旋
## 红黑树
* 红黑树是一种近似平衡的二叉搜索树（Binary Search Tree），它能够确保任何一个结点的左右子树的高度差小于两倍。具体来说，红黑树是满足如下条件的二叉搜索树：
* 每个节点要么是红色，要么是黑色
* 根节点是黑色
* 每个叶节点（NIL结点，空结点）是黑色的
* 不能有相邻接的两个红色结点
* 从任一结点到其每一个叶子的所有路径都包含相同数目的黑色结点。
* 关键性质：从根到叶子的最长的可能路径不多于最短的可能路径的两倍长。
