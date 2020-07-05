//序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方
//式重构得到原数据。
//
// 请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串
//反序列化为原始的树结构。
//
// 示例:
//
// 你可以将以下二叉树：
//
//    1
//   / \
//  2   3
//     / \
//    4   5
//
//序列化为 "[1,2,3,null,null,4,5]"
//
// 提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这
//个问题。
//
// 说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。
// Related Topics 树 设计

package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	str []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	//先序遍历构建
	//退出条件
	if root == nil {
		return "null"
	}
	//本层的处理
	str := strconv.Itoa(root.Val)
	//下层的处理
	str += "," + this.serialize(root.Left)
	str += "," + this.serialize(root.Right)
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	//解析的路子是遍历元素 遍历到的元素作为根节点 然后继续遍历获取他的左右子树
	nodeArr := strings.Split(data, ",")
	for _, node := range nodeArr {
		if node != "" {
			this.str = append(this.str, node)
		}
	}
	return this.deserializeHelper()
}

func (this *Codec) deserializeHelper() *TreeNode {
	if this.str[0] == "null" {
		this.str = this.str[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.str[0])
	root := &TreeNode{Val: val}
	this.str = this.str[1:]
	root.Left = this.deserializeHelper()
	root.Right = this.deserializeHelper()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
//leetcode submit region end(Prohibit modification and deletion)
