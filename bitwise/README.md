### 二进制中负数的计算
负数以正数的补码表示,int是有符号的整形数,最左端的1位是符号位
- 原码：一个整数按照绝对值的大小转化成二进制的数
- 反码：将二进制数按位取反
- 补码：反码加 1
正数的补码和原码、反码一样，负数的补码就是反码+1。

| 十进制  | 原码       | 反码      | 补码       |
| ------:| ---------:| ---------:| ---------:|
| 2      | 0000 0010 | 0000 0010 | 0000 0010 |
| -2     | 1000 0010 | 1111 1101 | 1111 1110 |

### 位运算：
- 按位与（ AND）	a & b	对于每一个比特位，只有两个操作数相应的比特位都是1时，结果才为1，否则为0。
- 按位或（OR）	a | b	对于每一个比特位，当两个操作数相应的比特位至少有一个1时，结果为1，否则为0。
- 按位异或（XOR）	a ^ b	对于每一个比特位，当两个操作数相应的比特位有且只有一个1时，结果为1，否则为0。
- 按位非（NOT）	~ a	反转操作数的比特位，即0变成1，1变成0。
- 左移（Left shift）  ``	a << b ``	将 a 的二进制形式向左移 b (< 32) 比特位，右边用0填充。
- 有符号右移 ``	a >> b ``	将 a 的二进制表示向右移 b (< 32) 位，丢弃被移出的位。
- 无符号右移 ``	a >>> b	`` 将 a 的二进制表示向右移 b (< 32) 位，丢弃被移出的位，并使用 0 在左侧填充。

### 异或运算有以下三个性质:
任何数和 0 做异或运算，结果仍然是原来的数，即 a ^ 0=a^0=a。  
任何数和其自身做异或运算，结果是 0，即 a^a=0。  
异或运算满足交换律和结合律，即:  
 a ^ b ^ a=b ^ a ^ a=b ^ (a ^ a)=b^0=b^b^a=b^a^a=b^(a^a)=b^0=b

 ### 应用实例
 * 按位与运算通常用来对某些位清0或保留某些位。例如把a 的高八位清 0 ， 保留低八位， 可作 `` a&255 `` 运算 ( 255 的二进制数为0000000011111111)。 
 * X=2^N与任何小于2^N的整数进行按位与的结果都是0：`` X&(X-1)=0 ``
 * 判断奇偶  `` a & 1 == 0 ``
 * 交换符号  `` ~a + 1 ``, 取反+1
 * -1 & a = a ( 111 000 001)
 
