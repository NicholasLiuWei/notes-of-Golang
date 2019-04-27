package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	LTag  bool
	Left  *Hero
	RTag  bool
	Right *Hero
}

//前序遍历【先输出root节点，然后在输出左子树，然后在输出右子书】
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf(" %d ", node.No)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

//中序遍历【先输出root的左子树，在输出root节点，在输出root的右子树】
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf(" %d ", node.No )
		InfixOrder(node.Right)
	}
}

//后续遍历
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf(" %d ", node.No)
	}
}

var i = 0

func ConstructTree(root *Hero, nodes []Hero) {
	fmt.Println(i)
	node := nodes[i]
	fmt.Printf("construct:: %d %s \r\n", node.No, node.Name)
	i++
	if node.LTag == false && node.RTag == false {
		*root = node
		root.Left = nil
		root.Right = nil
	} else if !node.LTag{
		*root = node
		root.Left = nil
		root.Right = &Hero{}
		ConstructTree(root.Right, nodes)
	}else if !node.RTag{
		*root = node
		root.Right = nil
		//root.Left = &Hero{}
		ConstructTree(root.Left, nodes)
	}else{
		*root = node
		root.Left = &Hero{}
		root.Right = &Hero{}
		ConstructTree(root.Left, nodes)
		ConstructTree(root.Right, nodes)
	}
}

func main() {
	nodes := []Hero{}

	newNodes(1, "1", true,true, &nodes)

	newNodes(2, "左2", true,true, &nodes)

	newNodes(11, "左11", false,false, &nodes)

	newNodes(12, "右12", true,true, &nodes)

	newNodes(13, "左13",false,false, &nodes)

	newNodes(14, "右14",false,false, &nodes)

	newNodes(3, "右3",false,true, &nodes)

	newNodes(2, "右2",true,true, &nodes)

	newNodes(14, "左14",false,false, &nodes)

	newNodes(15, "右15",false,false, &nodes)

	fmt.Printf("len(nodes) = %d\n", len(nodes))
	for key, value := range nodes {
		fmt.Printf("构建前==》 %d %d %s\n", key, value.No, value.Name)
	}
	root := &Hero{}
	ConstructTree(root, nodes)
	fmt.Println("root:", root.Name, root.No)
	//PreOrder(root)
	//PostOrder(root)
	InfixOrder(root)
}

func newNodes(no int, name string, ltag bool, rtag bool, nodes *[]Hero) {
	node := Hero{
		No:   no,
		Name: name,
		LTag: ltag,
		RTag: rtag,
	}
	*nodes = append(*nodes, node)
}

func test(node *Hero) {
	test(node.Left)
	test(node.Right)

}
