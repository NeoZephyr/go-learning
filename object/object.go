// 每个目录一个包
// 包可以与目录不同名
package main

import "fmt"

type TreeNode struct {
	value int
	left, right *TreeNode
}

func (node TreeNode) print() {
	fmt.Println(node.value)
}

// 传值，不改变原变量的值
// 有复制的开销
//func (node TreeNode) setValue(value int) {
//	node.value = value
//}

// 传指针
func (node *TreeNode) setValue(value int) {
	node.value = value
}

func (node *TreeNode) traverse() {
	if node == nil {
		return
	}

	node.left.traverse()
	fmt.Print(node.value, " ")
	node.right.traverse()
}

func createTreeNode(value int) *TreeNode {
	// 返回局部变量的地址
	return &TreeNode{value: value}
}

// 通过包装实现扩展
type NewTreeNode struct {
	node *TreeNode
}

func (treeNode *NewTreeNode) postOrder()  {
	if treeNode == nil || treeNode.node == nil {
		return
	}

	left := NewTreeNode{treeNode.node.left}
	right := NewTreeNode{treeNode.node.right}

	left.postOrder()
	right.postOrder()

	treeNode.node.print()
}

// 通过别名实现扩展
type Queue []int

func (queue *Queue) push(v int) {
	*queue = append(*queue, v)
}

func (queue *Queue) pop() int  {
	head := (*queue)[0]
	*queue = (*queue)[1:]
	return head
}

func (queue *Queue) isEmpty() bool {
	return len(*queue) == 0
}

type Customer struct {
	name string
	gender string
	email string
}

func (customer *Customer) printInfo() {
	fmt.Printf("customer: name = %s, gender = %s, email = %s\n", customer.name, customer.gender, customer.email)
}

func (customer *Customer) printAll() {
	fmt.Printf("category: customer\n")
	customer.printInfo()
}

type Member struct {
	Customer
	name string
}

func (member *Member) printInfo()  {
	fmt.Printf("Member: memberName = %s, name = %s, gender = %s, email = %s\n", member.name, member.Customer.name, member.gender, member.email)
}

func main() {
	//testObject1()
	testObject2()
}

func testObject1() {
	var root TreeNode

	root = TreeNode{value: 10}
	root.left = &TreeNode{}
	root.right = &TreeNode{20, nil, nil}
	root.right.left = new(TreeNode)
	root.right.right = createTreeNode(40)

	root.print()
	root.setValue(-1)
	root.print()

	root.traverse()
	fmt.Println()

	nodes := []TreeNode{
		{value: 300},
		{100, nil, nil},
		{200, nil, nil},
	}

	fmt.Println(nodes)
}

func testObject2() {
	customer := Customer{"jack", "male", "a@qq.com"}
	fmt.Printf("customer: %v\n", customer)

	var member Member
	member.name = "sarah"
	member.Customer.name = "sally"
	member.gender = "female"
	member.email = "s@qq.com"

	fmt.Printf("member: %+v\n", member)

	customer.printInfo()
	member.Customer.printInfo()
	member.printInfo()

	customer.printAll()
	// 不会调用重写的 printInfo 方法
	member.printAll()

	// 方法值
	func1 := customer.printInfo
	func1()

	// 方法表达式
	func2 := (*Customer).printInfo
	func2(&customer)
}