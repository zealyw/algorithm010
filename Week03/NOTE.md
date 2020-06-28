## 递归以及树的遍历
### 树的遍历
#### 前序
`
func preorderTraversal(root *TreeNode) []int {
 	//根左右
 	if root == nil {
 		return []int{}
 	}
 	path := append([]int{root.Val}, preorderTraversal(root.Left)...)
 	path = append(path, preorderTraversal(root.Right)...)
 	return path
 }
 `
#### 中序
`func inorderTraversal(root *TreeNode) []int {
 	//左根右
 	if root == nil {
 		return []int{}
 	}
 	path := append(inorderTraversal(root.Left), root.Val)
 	path = append(path, inorderTraversal(root.Right)...)
 	return path
 }`
#### 后序
`func postorderTraversal(root *TreeNode) []int {
 	//左右根
 	if root == nil {
 		return []int{}
 	}
 	path := append(postorderTraversal(root.Left), postorderTraversal(root.Right)...)
 	path = append(path, root.Val)
 	return path
 }`
#### 前序中序构建二叉树
重点在于通过前序确认根节点的位置，然后通过根节点在中序中的位置将中序拆分出左右子树包含的部分
`func buildTree(inorder []int, postorder []int) *TreeNode {
 	//inorder [[左子树][根][右子树]]
 	//poster[[左子树][右子树][根]]
 	if len(inorder) == 0 {
 		return nil
 	}
 	root := &TreeNode{postorder[len(postorder)-1],nil,nil}
 	i:= 0
 	for ;i<len(inorder);i++ {
 		if inorder[i] == postorder[len(postorder)-1] {
 			break
 		}
 	}
 	root.Left = buildTree(inorder[:i],postorder[:i])
 	root.Right = buildTree(inorder[i+1:],postorder[i:len(postorder)-1])
 	return root
 }`

#### 中序后序构建二叉树
思路的话同前中序构建树，不同的点在于这回通过后续来找到根节点的位置
`func buildTree(inorder []int, postorder []int) *TreeNode {
 	//inorder [[左子树][根][右子树]]
 	//poster[[左子树][右子树][根]]
 	if len(inorder) == 0 {
 		return nil
 	}
 	root := &TreeNode{postorder[len(postorder)-1],nil,nil}
 	i:= 0
 	for ;i<len(inorder);i++ {
 		if inorder[i] == postorder[len(postorder)-1] {
 			break
 		}
 	}
 	root.Left = buildTree(inorder[:i],postorder[:i])
 	root.Right = buildTree(inorder[i+1:],postorder[i:len(postorder)-1])
 	return root
 }`
 
 ### 递归
 找重复性的问题来进行递归操作, 忌人肉递归，找到最近最简问题（重复子问题）
 #### 递归的代码模板
 `
 func recursion(level,params...){
    # recursion terminator
    if level > MAX_LEVLE{
        process_result
        return
    }
    #process logic in current level
    process(level,data...)
    #drill down
    this.recursion(level+1,params...)
    # reverse the current level status if needed 
 }
 `
 #### 分治
 同常规递归的不同点在于需要组合各个子问题的处理结果
 `
 func divide_conquer(problem, params...){
    #recurison terminator
    if problem is None{
        return 
    }
    #prepare data
    data = prepare_data(problem)
    subProblems = split_problem(problem,data)
    #conquer subProblems
    subresult1 = this.divide_conquer(subProblems[0],parmas)
    subresult2 = this.divide_conquer(subProblems[1],parmas)
    subresult3 = this.divide_conquer(subProblems[2],parmas)
    ...
    #process and generate the final result 
    result = process_result(subresult1,subresult2,subresult3...)
    #revert the current level states
 }
 `
 
 #### 回溯
 回溯---不断的去试错。一种暴力破解的方式，可以通过一些判断来尽可能早的判错来减少操作次数