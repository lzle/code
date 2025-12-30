# effective-python


### 1、遵循 PEP8 风格指南

https://peps.python.org/pep-0008/


### 2、了解 bytes、str 和 unicode 的区别

Python3 中有两种表示字符序列的类型：bytes 和 str。 前者的实例包含原始的 8 为值；后者的实例
包含 Unicode 字符。

Python2 也有两种表示字符序列的类型：str 和 unicode。与 Python3 不同的是，str 的实例包含原始
的 8 位值；而 unicode 的实例，则包含 Unicode 字符。

在 Python3 中，需要编写接受 str 或 bytes，并总是返回 str 的方法：

``` python
def to_str(bytes_or_str):
    if isinstance(bytes_or_str, bytes):
        value = bytes_or_str.decode('utf-8')
    else:
        value = bytes_or_str
    return value
```

另外接受 str 或 bytes，并总是返回 bytes 的方法：

``` python
def to_str(bytes_or_str):
    if isinstance(bytes_or_str, str):
        value = bytes_or_str.encode('utf-8')
    else:
        value = bytes_or_str
    return value
```

在 Python2 中，需要编写接受 str 或 unicode，并总是返回 unicode 的方法：

``` python
def to_str(unicode_or_str):
    if isinstance(unicode_or_str, str):
        value = unicode_or_str.decode('utf-8')
    else:
        value = unicode_or_str
    return value
```

另外接受 str 或 unicode，并总是返回 str 的方法：

``` python
def to_str(unicode_or_str):
    if isinstance(unicode_or_str, unicode):
        value = unicode_or_str.encode('utf-8')
    else:
        value = unicode_or_str
    return value
```

### 3、使用辅助函数取代复杂的表达式

开发者很容易过度运用 Python 的语法特性，从而写出那种特步复杂并且难以理解的单行表达式。

```python
red = my_value.get('red',[''])
red = int(red[0] if red[0] else 0)
```

对于上面这个例子来说，它的清晰程度还是比不上跨越多行的完整 if/else 语句。但如果改成下面这种形式，
就会让感觉到：上面的紧缩的写法其实挺复杂的。

```python
green = my_value.get('green', [''])
if green[0]:
    green = int(green[0])
else:
    green = 0
```

现在应该把它总结成辅助函数，如果需要频繁的使用这种逻辑，那就更应该这么做。

```python
def get_first_int(values, key, default=0):
    found = values.get(key, [''])
    if found[0]:
        found = int(found[0])
    else:
        found = default
    return found
```

调用这个辅助函数时所使用的代码更加清晰

```python
green = get_first_int(my_value, 'green')
```

表达式如果变得比较复杂，那就应该考虑将其拆分成小块，并把这些逻辑移入辅助函数中。会令代码更加
易读。

### 4、了解切割序列的方法

对原列表进行切割后，会产生另一方全新的列表。系统依然维护这只想原列表中各个对象的引用。在切割后得到
的新列表上进行修改，不会影响原列表。

```python
a = ['a', 'b', 'c', 'd']
b = a[2:]
b[1] = 99
print('After', b)
print('No change', a)

>>>
After ['c', 99]
No change ['a', 'b', 'c', 'd']
```

在赋值时对左侧列表使用切割操作，会把该列表中处于指定范围内的对象替换为新值。两边的位数无需一致。

```python
a = ['a', 'b', 'c', 'd']
a[1:2] = [99, 100, 101]
print('After', a)

>>>
After ['a', 99, 100, 101, 'c', 'd']
```

如果对赋值操作的右侧的列表使用切片，切片索引都留空，会产生一份原列表的拷贝。

```python
b = a[:]
assert  b == a and b is not a
```

如果左侧列表切片未指定起止索引，系统会把右侧的**新值复制一份**，并用这份拷贝替换左侧列表的全部内容，
并不会重新分配列表。

```python
a = ['a', 'b', 'c', 'd']
b = a
a[:] =  [101, 102, 103]
assert a is b
print('After', a)

>>>
After [101, 102, 103]
```

### 5、单次切片操作内，不要同时指定 start、end 和 stride

```python
a = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h']
a[-2::2] # ['g', 'e', 'c', 'a']
```

在一个中括号里写上 3 个数字显得太拥挤，从而导致代码难以阅读。可以考虑先做步进式切片，再在结果上做切割。

```python
b = a[::2]  # ['a', 'c', 'e', 'g']
c = b[1:-1] # ['c', 'e']
```

上面这种先做步进切割，再做范围切割的办法，会多产生一份原数据的浅拷贝，执行第一次切割操作时，应该尽量缩减
切割后的列表尺寸。

既有 start 和 end, 又有 stride 的切割操作，可能会令人费解。

尽量使用 stride 为正数，且不带 start 或 end 索引的切割操作。尽量避免用负数做 stride。

在同一个切片操作呢，不要同时使用 start、end 和 stride。如果确实需要执行这种操作，那就考虑将其
拆解为两条赋值语句，其中一条做范围切割，另一条做步进切割，或考虑内置 itertools 模块中的 islice。


### 6、使用列表推导来取代 map 和 filter

```python
a = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
even_squares = [x**2 for x in a if x % 2 == 0]

>>>
[4, 16,36, 64, 100]
```

内置的 filter 函数和 map 结合起来，也能达到同样的效果，但是代码会写得非常难懂。

```python
alt = map(lambda x:x**2, filter(lambda x:x % 2 == 0, a))
assert even_squares == alt
```

字典和集合也支持推导表达式。