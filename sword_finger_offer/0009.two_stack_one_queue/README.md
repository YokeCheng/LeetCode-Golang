#

 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

示例 1：

    输入：
    ["CQueue","appendTail","deleteHead","deleteHead"]
    [[],[3],[],[]]
    输出：[null,null,3,-1] 

示例 1：

    输入： 
    ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"] 
    [[],[],[5],[2],[],[]] 
    输出：[null,-1,null,null,5,2] 

## 提示：

1 <= values <= 10000
最多会对 appendTail、deleteHead 进行 10000 次调用

## 解题思路：

A,B两个栈
appendTail时，向A中添加
deleteHead时，
 - 当B中有值，直接pop一个出来
 - 当A存在，B中不存在，从A中转移到B中(相当于倒序出来了)，然后pop一个
 - 当AB都不存在，返回空
