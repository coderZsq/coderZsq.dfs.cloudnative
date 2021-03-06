## LeetCode 50 Pow(x, n) [中等]

### 题目描述

- https://leetcode-cn.com/problems/powx-n/

实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，xn）。

![](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/binarysearchtree_improved.png)

示例 1：

```
输入：x = 2.00000, n = 10
输出：1024.00000
```

示例 2：

```
输入：x = 2.10000, n = 3
输出：9.26100
```

示例 3:

```
输入：x = 2.00000, n = -2
输出：0.25000
解释：2^-2 = 1/2^2 = 1/4 = 0.25
```

提示：

-100.0 < x < 100.0
-2^31 <= n <= 2^31-1
-10^4 <= x^n <= 10^4

---

### 题解 1 (暴力)

### 基本流程

```
1. 判断指数是否为负数
2. 如果为负数则需要转换
3. 进行逐个累乘
```

### 题解 2 (分治 + 递归)

#### 基本流程

```
设: x = 2, n = 8, 即 x^n 为 2^8
分治: 2^16 == 2^4*2^4 == 2^2*2^2*2^2*2^2

1. 如果指数为 0, 返回 1, 因为所有数的 0 次幂均为 1
2. 如果指数为负数, 则返回倒数
3. 如果指数为偶数, 则进行分治
4. 如果指数为奇数, 先将其变为偶数PoW 乘以 x
```

### 题解 3 (分治 + 迭代)

### 基本流程

```
1. 先判断指数是否为负数
2. 如果为负数, 进行转变为正数
3. 循环取出 指数 n 的最低位 进行累乘法
4. x 进行累乘法, 指数进行移位
5. 返回累乘结果
```
