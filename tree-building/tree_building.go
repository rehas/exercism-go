package tree

import(
    "fmt"
    "errors"
    "sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

func (r Record) isRootRecord() bool{
    return r.ID == 0 && r.Parent == 0
}
func (r Record) isValid() bool{
    return r.ID > r.Parent 
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func (n *Node) isSameId(r Record) bool{
    return r.ID == n.ID
}

func (n *Node) isParent(r Record) bool{
    return r.Parent == n.ID
}

// size finds the total size of the tree
func (n *Node) size() int{
    if len(n.Children)== 0{
        return 1
    }
    var count int
    for _, c := range n.Children{
        count += c.size() 
    }
	return count +1 // count root
}

func (n *Node) addChild(r Record) error{
    for _, child := range n.Children{
        if child.isSameId(r){
            return errors.New("duplicate children")
        }
    }
	n.Children = append(n.Children, &Node{ID:r.ID})
    // children need to be sorted
    sort.SliceStable(n.Children, func(i, j int) bool {
        return n.Children[i].ID < n.Children[j].ID
    })
    return nil
}

func (n *Node) locateRecord(rec Record) error{
    if n.isSameId(rec){
        return errors.New("duplicate IDs")
    }else if n.isParent(rec){
        return n.addChild(rec)
    }
	// This is where the possiblities left are
    // This record is a grand child
    // Or this record is disconnected to the root
    // Disconnected roots are handled after tree is constructed
    for _, cn := range n.Children{
        err := cn.locateRecord(rec)
        if err != nil{
            return err
        }
    }
    return nil
}

func isValidRecordSet(records []Record) bool{
    size := len(records)
    for _, rec := range records{
        if rec.ID >= size{
            return false
        }
    	if rec.ID > 0 && rec.Parent > 0 && !rec.isValid(){
            return false
        }
    }
return true
}

func Build(records []Record) (*Node, error) {
	// if there are no records return
    // for every record
    // check if the ID matched current NodeId (step1)
    // if match, return error
    // if not, check the parentId against the current NodeId
    // if match, check if this node exists in children 
    // if exists return error
    // if not, append to children
    // if parent Id does not match current NodeId
    // repeat this process against all the children of current Node starting from step 1

    if !isValidRecordSet(records){
        return nil, errors.New("invalid record set")
    }
    
    recordsSize := len(records)
    if recordsSize < 1 {
        return nil, nil
    }

    var root *Node
    // find and remove root record from list
    for i, record := range records {
        if record.isRootRecord(){
            root = &Node{ID:record.ID}
            fmt.Println(records, i)
            records = append(records[:i], records[i+1:]...)
            break
        }
    }

    if root == nil{
        return nil, errors.New("invalid root")
    }

    // sort based on parentIDs to prevent revisiting same record twice
    sort.SliceStable(records, func(i, j int) bool {
        return records[i].Parent < records[j].Parent
    })

    for _, record := range records{
        err := root.locateRecord(record)
		fmt.Println(records, record, root.Children, err)    
        if err!=nil{
            return nil, err
        }
    }
    
	if root.size() != recordsSize{
        return nil, errors.New("disconnected node")
    }
    return root, nil
}


