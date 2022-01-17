# Map声明、初始化和 make

- map 是一种特殊的数据结构：一种元素对（pair）的**无序集合**，给定 key，对应的 value 可以迅速定位。
  - key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key 
    - 含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的
    - 指针和接口类型也可以作为key
  - value 可以是任意类型的
- map 是**引用类型**，可以使用如下声明：`var mapname map[keytype]valuetype`，如`var map1 map[string]int`
- 在声明的时候不需要知道 map 的长度，map 是可以动态增长的。
- 未初始化的 map 的值是 nil。
- 常用的 `len(map1)` 方法可以获得 map 中的 pair 数目
- 使用make创建map
  - map 的初始化：`var map1 = make(map[keytype]valuetype)`。
  - 或者`map1 := make(map[keytype]valuetype)`。
  - **不要使用 new，永远用 make 来构造 map**
    - 如果你错误地使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址。 如`mapCreated := new(map[string]float32)`，接下来当我们调用：`mapCreated["key1"] = 4.5` 的时候，编译器会报错：`invalid operation: mapCreated["key1"] (index of type *map[string]float32).`
  - map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制。但是你也可以选择标明 map 的初始容量 `capacity`，就像这样：`make(map[keytype]valuetype, cap)`
    - 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1

# Map的key与value

- 测试 map1 中是否存在 key1：`val1, isPresent = map1[key1]`

  - isPresent 返回一个 bool 值：如果 key1 存在于 map1，val1 就是 key1 对应的 value 值，并且 isPresent为true；如果 key1 不存在，val1 就是一个空值，并且 isPresent 会返回 false。

  - 如果只是想判断key是否存在而不关心值是多少：`_, ok := map1[key1]`

  - 常见场景：与if合用：

    - ```go
      if _, ok := map1[key1]; ok {
      	// ...
      }
      ```

- 从 map1 中删除 key： `delete(map1, key1)` 

  - 如果 key1 不存在，该操作不会产生错误。

## 搭配for range

```
for key, value := range map1 {
	...
}
```

## Map是无序的

map 不是按照 key 的顺序排列的，也不是按照 value 的序排列的。map的本质是散列表，而map的增长扩容会导致重新进行散列。

## 排序

如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序（使用 sort 包，详见第 7.6.6 节），然后可以使用切片的 for-range 方法打印出所有的 key 和 value。

```go
// the telephone alphabet:
package main
import (
	"fmt"
	"sort"
)

var (
	barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
							"delta": 87, "echo": 56, "foxtrot": 12,
							"golf": 34, "hotel": 16, "indio": 87,
							"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("unsorted:")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	keys := make([]string, len(barVal))
	i := 0
	for k, _ := range barVal {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println()
	fmt.Println("sorted:")
	for _, k := range keys {
		fmt.Printf("Key: %v, Value: %v / ", k, barVal[k])
	}
}
```

## 多值map

使用 `map[int][]string` 类型