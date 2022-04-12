package main

import "fmt"

//cmdNode结点
type CmdNode struct {
	pre     *CmdNode //前一个结点
	next    *CmdNode //后一个结点
	cmdName string   //cmd名称
	cmdDesc string   //cmd描述
	cmdFunc callback //定义一个callback接口，该接口必须实现execCMD()方法
}

//callback接口
type callback interface {
	exec()
}

//定义为循环链表，如同STL中cmdList
type CmdList struct {
	head *CmdNode //虚节点
	len  int
}

//cmdList获取函数
func getCmdList() *CmdList {
	head := &CmdNode{nil, nil, "", "", nil}
	head.pre = head
	head.next = head
	return &CmdList{head, 0}
}

//寻找节点
func (this *CmdList) find(cmdName string) *CmdNode {
	t := this.head.next
	for t != this.head && t.cmdName != cmdName {
		t = t.next
	}
	if t == this.head {
		return nil
	} else {
		return t
	}
}

//插入结点
func (this *CmdList) insert(cmdName string, cmdDesc string, cmdFunc callback, isReplace bool) bool {
	//先检查是否有相同名称的命令，如果有则根据isReplace的值进行替换
	t := this.find(cmdName)
	if t == nil {
		t = this.head
		//尾插，默认新添加的命令使用次数会少
		newNode := &CmdNode{t.pre, t, cmdName, cmdDesc, cmdFunc}
		t.pre.next = newNode
		t.pre = newNode
		this.len++
		return true
	}
	//如果允许替换则直接替换
	if isReplace {
		t.cmdDesc = cmdDesc
		t.cmdFunc = cmdFunc
		return true
	}
	return false
}

//删除节点
func (this *CmdList) remove(cmdName string) bool {
	t := this.find(cmdName)
	if t == nil {
		return false
	}
	t.pre.next = t.next
	t.next.pre = t.pre
	this.len--
	return true
}

//执行命令
func (this *CmdList) exec(cmdName string) bool {
	t := this.find(cmdName)
	if t == nil {
		fmt.Println("Cannot find this command \"" + cmdName + "\"")
		return false
	}
	t.cmdFunc.exec()
	return true
}

//打印命令
func (this *CmdList) printCmd() {
	t := this.head.next
	for t != this.head {
		fmt.Println(t.cmdName + "\t" + t.cmdDesc)
		t = t.next
	}
}
