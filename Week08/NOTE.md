## 位运算
### 表达符
* 左移 `<<` 0011 << 1 =  0110
* 右异动 `>>` 0110 >> 1 = 0011
* 或运算 `|`  0011 | 1011 = 1011
* 与运算 `&` 0011 & 1011 = 0011
* 取反 `~` ~0011 = 1100
* 异或 `^` 0011 ^ 1011 = 1000
### 实战操作
* 判断奇数偶数
x % 2 == 1 等于与 x&1 == 1
x % 2 == 0 等同于 x&1 == 0 
* x >> 1 等同于 x /2
* X= X & (X-1) 清零最低位的1
* X &-x 得到最低位的1
* X & ~X = 0

## 布隆过滤器
一个很长的二进制向量和一系列随机映射函数，布隆过滤器可以用来检索一个元素是否存在一个集合中
优点：空间效率和时间效率远远超过一般算法
缺点：有一定的误识别率和删除困难
### 使用场景
* 可以作为DB前面的一层缓存来判断元素是否在库里，在就查询，不在就直接返回，避免一次查库。
一般对于缓存穿透，可以使用这个来进行解决。
* 垃圾邮件识别，判重
* IP黑名单

## LRU Cache
hash table + double linked
O(1)查询
O(1)修改/更新
对于get操作，首先判断key是否存在，如果不存在，返回-1；如果key存在，则key对应的节点是最近被使用的节点，
通过hash表定位到其在双向链表中的位置，将其移动到双向链表的头部，最后返回该节点的值
对于put操作，首先判断key是否存在，不存在则创建一个新的节点，在双向链表的头添加节点，并加入hash表中，
如果双向链表的节点超出容量了，就删除尾部节点，并删除对应hash表中的项。

## 排序
### 比较类排序
* 交换排序：冒泡排除、快速排序
* 插入排序：简单插入排序、希尔排序
* 选择排序: 简单选择排序、堆排序
* 归并排序：二路归并排序、多路归并排序
nlogN
堆排序、快速排序、归并排序

### 非比较类排序
* 计数排序
* 桶排序
* 基数排序

https://www.cnblogs.com/onepixel/p/7674659.html
### 初级排序
* 选择排序:每次找最小值，然后放到待排序数组的起始位置
```php
function selectSort($arr){
    $len = count($arr);
    for ($i=0;$i<$len-1;$i++){
        $minIndex = $i;
        for ($j=$i+1;$j<$len-1;$j++){
            if ($arr[$j] < $arr[$minIndex]){
                $minIndex = $j;
            }
        }
        list($arr[$minIndex],$arr[i]) = [$arr[i],$arr[$minIndex]];
    }
    return $arr;
}
```
* 插入排序：从前往后构建有序序列，对于未排序的元素，在已排序的序列中从后往前扫描，找到位置插入
```php
function insertSort($arr){
    $len = count($arr);
    for ($i=1;$i<$len-1;$i++){
        for ($j=$i;$j>=0;$j--){
            if ($arr[j] < $arr[j-1]){
                list($arr[j],$arr[j-1]) = [$arr[j-1],$arr[j]];
            }else{
                break;
            }
        }
    }
}
```
* 冒泡排序：循环嵌套，每次查看相邻元素，逆序就交换
```php
function BubbleSort($arr){
    $len = count($arr);
    for ($i= 0;$i<$len-1;$i++){
        for ($j = 0; $j<$len-1-$i;$j++){
            if($arr[$j]> $arr[$j+1]){
                list($arr[$j],$arr[$j+1]) = [$arr[$j+1],$arr[$j]]
            }
        }
    }
}
```
### 高级排序
* 快速排序：在数组中取一个标杆，将小元素放在左边，大元素放在右侧，然后依次对右边和右边的子数组继续进行快排
```php
function quickSort(&$arr, $begin, $end){
    if ($end <= $begin){
        return true;
    }
    $pivot =  partition($arr, $begin, $end);
    quickSort($arr, $begin, $pivot-1);
    quickSort($arr, $pivot+1, $end);
}
function partition(&$arr, $begin, $end){
    $pivot = $end;
    $index = $pivot+1;
    for($i=$index;$i<=$end;$i++){
        if($arr[$i]<$arr[$pivot]){
            list($arr[i],$arr[$index]) = [$arr[$index],$arr[$i]];
            $index++;
        }
    }
    list($arr[$pivot],$arr[$index-1]) = [$arr[$index-1],$arr[$pivot]];
    return $index-1
}

```
* 归并排序：将数组一分为2 ，将两个子序列分别调用归并排序
```php
function mergeSort(&$arr,$left,$right){
    if ($right <= $left) {
        return true;
    }
    $mid = $left - ($left-$right) >> 1;
    mergeSort($arr,$left,$mid)
    mergeSort($arr,$mid+1,$right)
    merge($arr,$left,$mid,$right)
}

function merge(&$arr,$left,$mid,$right){
    $tmp = [];
    $i = $left;
    $j = $mid+1;
    while($i<=$mid && $j <= right){
        $tmp[] = $arr[$i]
        if ($arr[i] <= $arr[$j]){
            $i++;    
        }else{
            $j++;
        }
    }
    while ($i<=$mid){
        $tmp[] = $arr[$i] 
    }
    while($j <= $right){
        $tmp[] = $arr[$j]
    }
    for ($p=0;$p<count($tmp);$p++){
        $arr[$left+$p] = $tmp[$p]
    }
}
```
* 堆排序：插入是O(logN) 取最大值最小值都是O(1),从数组元素中依次建立小顶堆，取堆顶元素并删除。
### 特殊排序
* 计数排序


