## LeetCode 142 环形链表 II [中等]

### 题目描述

- https://leetcode-cn.com/problems/linked-list-cycle-ii/

给定一个链表，返回链表开始入环的第一个节点。  如果链表无环，则返回  null。

如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。

不允许修改 链表。

示例 1：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist.png)

```
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
```

示例  2：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test2.png)

```
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
```

示例 3：
![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/07/circularlinkedlist_test3.png)

```
输入：head = [1], pos = -1
输出：返回 null
解释：链表中没有环。
```

提示：

- 链表中节点的数目范围在范围 [0, 10^4] 内
- -10^5 <= Node.val <= 10^5
- pos 的值为 -1 或者链表中的一个有效索引

进阶：你是否可以使用 O(1) 空间解决此题？

---

### 题解 1 (哈希表)

#### 基本流程

```
1. 创建哈希表
2. 遍历链表
3. 将链表中的节点依次放入哈希表
4. 当将放入节点已存在哈希表中则该节点为入环第一个节点
```

### 题解 2 (快慢指针)

#### 基本流程

```
1. 快慢指针同时从头节点出发
2. 快指针走两步
3. 慢指针走一步
4. 当快慢指针相遇则表示有环
5. 头节点和慢指针同时走一步
6. 相遇节点为入环第一个节点
```
