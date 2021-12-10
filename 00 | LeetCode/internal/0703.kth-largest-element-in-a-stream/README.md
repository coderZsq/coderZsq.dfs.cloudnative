## LeetCode 703 数据流中的第 K 大元素 [简单]

### 题目描述

- https://leetcode-cn.com/problems/kth-largest-element-in-a-stream/

设计一个找到数据流中第 k 大元素的类（class）。注意是排序后的第 k 大元素，不是第 k 个不同的元素。

请实现 KthLargest  类：

KthLargest(int k, int[] nums) 使用整数 k 和整数流 nums 初始化对象。
int add(int val) 将 val 插入数据流 nums 后，返回当前数据流中第 k 大的元素。

示例：

```
输入：
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
输出：
[null, 4, 5, 5, 8, 8]

解释：
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
```

提示：

- 1 <= k <= 104
- 0 <= nums.length <= 104
- -104 <= nums[i] <= 104
- -104 <= val <= 104
- 最多调用 add 方法 104 次
- 题目数据保证，在查找第 k 大元素时，数组中至少有 k 个元素

---

### 题解 1

#### 基本流程

```
1. 创建堆并将所有元素入堆 该堆保存序列的最大值 (Go 中 sort.IntSlice 为最小堆)
2. 如果堆的长度 < k, 则将元素入堆
3. 如果新元素大于堆中最小值, 则将新元素入堆, 并将堆中最小值出堆
4. 堆顶元素即为第 K 大元素
```