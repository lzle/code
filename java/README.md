 
## 目录

* [环境](#环境)
    * [安装](#安装)
    * [Hello World](#hello-world)
* [基础类型](#基础类型)
    * [整型](#整型)
    * [浮点型](#浮点型)
    * [布尔](#boolean-类型)
    * [枚举](#枚举)
    * [字符串](#字符串)
* [复合类型](#复合类型)
    * [数组](#数组)
    * [ArrayList](#ArrayList)
    * [LinkedList](#LinkedList)
    * [HashSet](#HashSet)
    * [HashMap](#HashMap)
* [变量与常量](#变量与常量)
    * [声明变量](#声明变量)
    * [初始化变量](#初始化变量)
    * [常量](#常量)
* [修饰符](#修饰符)
    * [default](#default)
    * [private](#private)
    * [public](#public)
    * [protected](#protected)
    * [static](#static)
    * [final](#final)
    * [abstract](#abstract)
    * [synchronized](#synchronized)
    * [volatile](#volatile)
* [流程控制](#流程控制)
    * [while](#while)
    * [do while](#do-while)
    * [for](#for)
    * [break](#break)
    * [continue](#continue)
    * [if else](#if-else)
    * [switch case](#switch-case)
* [对象与类](#对象与类)
    * [对象](#对象)
    * [类](#类)
        * [隐式参数和显式参数](#隐式参数和显式参数)
        * [基于类的访问权限](#基于类的访问权限)
        * [final 实例字段](#final-实例字段)
        * [方法参数](#方法参数)
        * [构造器](#构造器)
        * [类设计技巧](#类设计技巧)
* [继承](#继承)
    * [类、超类、子类](#类、超类、子类)
        * [定义子类](#定义子类)
        * [覆盖方法](#覆盖方法)
        * [子类构造器](#子类构造器)
        * [阻止继承：final 类和方法](#阻止继承final-类和方法)
        * [强制类型转换](#强制类型转换)
    * [Object：所有类的超类](#object所有类的超类)
    * [抽象类](#抽象类)
    * [枚举类](#枚举类)
* [接口](#接口)
* [单例](#单例)
* [性能分析](#性能分析)
    * [async-profiler](#async-profiler)
    * [arthas](#arthas)


## 环境
Java是一种广泛使用的编程语言，具有跨平台、面向对象、安全性高等特点。它最初由Sun Microsystems公司于1995年发布，现在属于Oracle公司。Java可以用于开发各种应用程序，包括桌面应用程序、网站后端、移动应用程序(Android)、嵌入式系统等。Java的核心理念是“一次编写，到处运行”，这得益于Java虚拟机（JVM）的架构，使得Java编写的程序可以在任何支持JVM的平台上运行而无需重新编译。Java是一种静态类型、编译型语言，同时也支持自动垃圾回收，减轻了内存管理的负担。

### 安装
Java的安装分为两个部分：JDK（Java Development Kit）和JRE（Java Runtime Environment）。JDK包含了Java的开发工具，如javac、java等，JRE包含了Java的运行环境，用于运行Java程序。在安装Java之前，需要先安装JDK，然后再安装JRE。以下是在CentOS上安装Java的步骤：

#### 方法 1：通过 yum 安装

查看 yum 库支持的 Java 版本

```bash
$ yum search java|grep jdk
```

选择版本安装 jdk

```bash
$ yum install java-1.8.0-openjdk
$ yum install java-1.8.0-openjdk-devel
```

验证是否安装成功

```bash
$ java -version
openjdk version "1.8.0_412"
OpenJDK Runtime Environment (build 1.8.0_412-b08)
OpenJDK 64-Bit Server VM (build 25.412-b08, mixed mode)
```

#### 方法 2：通过源码安装

从 [官网](https://www.oracle.com/cn/java/technologies/javase-downloads.html) 获取安装包, 使用 wget 命令下载

```bash
$ mkdir /usr/local
$ cd /usr/local
$ wget https://download.oracle.com/otn/java/jdk/8u202-b08/1961070e4c9b4e26a04e7f5a083f551e/jdk-8u202-linux-x64.tar.gz?AuthParam=1718551424_b71ec810b7f4719651ac8dfd4ad35cf9
$ tar -zxvf jdk-8u202-linux-x64.tar.gz
```

配置环境变量

```bash
$ vim /etc/profile
export JAVA_HOME=/usr/local/jdk1.8.0_202
export JRE_HOME=$JAVA_HOME/jre
export PATH=$JAVA_HOME/bin:$PATH
export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
$ source /etc/profile
```

验证是否安装成功

```bash
$ java -version
java version "1.8.0_202"
Java(TM) SE Runtime Environment (build 1.8.0_202-b08)
Java HotSpot(TM) 64-Bit Server VM (build 25.202-b08, mixed mode)
```

### Hello World

测试文件 HelloWorld.java

```java
public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello World");
    }
}
```

执行文件

```
$ javac HelloWorld.java
$ java HelloWorld
Hello World
```

## 基础类型

### 整型

整数型用于表示没有小数部分的数，可以是负数。Java 提供了 4 种整型，如表所示。

| 类型   | 存储需求 | 取值范围 |
|--------|----------|-----------|
| int    | 4 字节    | -2,147,483,648 ～ 2,147,483,647（略高于 20 亿） |
| short  | 2 字节    | -32,768 ～ 32,767 |
| long   | 8 字节    | -9,223,372,036,854,775,808 ～ 9,223,372,036,854,775,807 |
| byte   | 1 字节    | -128 ～ 127 |

在 Java 中，整型的范围与运行 Java 代码的机器无关。这就解决了软件从一个平台移植到
另一个平台时（或者甚至在同一个平台中不同操作系统之间移植时）让程序员头疼的主要问题。

长整型数值有一个后缀 L 或 l（如 40000000L）。十六进制数值有一个前缀 0x 或 0X（ 如 0xCAFE）

加上前缀 0b 或 0B 还可以写二进制数。例如，0b1001 就是 9。另外，可以为数字字面量加下画线，
如用 1_000_000表示 100 万。这些下画线只是为了让人更易读。Java 编译器会去除这些下画线。

### 浮点类型

浮点类型用于表示有小数部分的数值。在 Java 中有两种浮点类型，如表所示。

| 类型     | 存储需求 | 取值范围 |
|----------|----------|----------|
| float    | 4 字节    | 大约 ±3.40282347 × 10³⁸（6～7 位有效数字） |
| double   | 8 字节    | 大约 ±1.79769313486231570 × 10³⁰⁸（15 位有效数字） |

double 表示这种类型的数值精度是 float 类型的两倍（有人称之为双精度教（ double-precision））. 很多情况下，float 类型的精度（ 6 7 位有效数字）都不能满足需求。实际上，
只有很少的情况适合使用 float 类型，例如，所使用的库需要单精度数，或者需要存储大量单精度数时。

float 类型的数值有一个后缀 F 或 f（例如，3.14F）. 没有后缀 F 的浮点数值（如 3.14）总是默认为 double 类型。可选地，也可以在 double 数值后面添加后缀 D 或 d（例如，3.14D）。

所有的浮点数计算都遵循 IEEE 754 规范。具体来说，有 3 个特殊的浮点数值表示溢出和出错情况：
* 正无穷大
* 负无穷大
* NaN（不是一个数）

例如，一个正整数除以 0 的结果为正无穷大白 计算 `0/0` 或者负数的平方根结果为 NaN。

### Unicode

Java 使用 Unicode 字符集来表示字符。Unicode 是一种国际标准，旨在为世界上所有的字符提供唯一的编码。Java 中的字符类型是 `char`，它使用 16 位来存储一个字符。

### boolean 类型

boolean (布尔) 类型有两个值：false 和 true, 用来判定逻辑条件。整型值和布尔值之间不能进行相互转换。

### 枚举

自定义枚举类型。

```java
public enum Day {
    SUNDAY, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY
}
```

声明这种类型的变量。

```java
public class EnumExample {
    public static void main(String[] args) {
        Day today = Day.MONDAY;
        System.out.println("Today is: " + today);
    }
}
```

### 字符串

从概念上讲，Java 字符串就是 Unicode 字符串序列。Java 没有内置的字符串类型，而是标准 Java 类库中提供了一个预定义类，很自然地叫作 `String`。 
每个用双引号括起来的字符串都是 `String` 类的一个实例：

```java
String e = "";  // an empty string
String greeting = "Hello!";
```

#### 子串

`String` 类的 `substring` 方法可以从一个较大的字符串提取出一个子串。

```java
String greeting = "Hello, World!";
String s = greeting.substring(0,3);
System.out.println(s);

// 输出: Hel
```

#### 拼接

允许使用 `+` 号连接（拼接）两个字符串。

```java
String firstName = "John";
String lastName = "Doe";
String fullName = firstName + " " + lastName;
System.out.println(fullName);
// 输出: John Doe
```

如果需要把多个字符串放在一起，用一个界定符分隔，可以使用静态 `join` 方法：

```java
String[] words = {"Hello", "World", "Java"};
String sentence = String.join(" ", words);
System.out.println(sentence);
// 输出: Hello World Java
```

#### 字符串不可变

如果复制一个字符串变量，原始字符串和复制的字符串共享相同的字符序列。字符串是不可变的，这意味着一旦创建了一个字符串，就不能更改它的内容。
只能进行拼接、裁剪生成新的字符串。

#### 字符串比较

可以使用 `equals` 方法检测两个字符串是否相等。

```java
s.equals(t)
```

要想检测两个字符串是否相等，而不区分大小写，可以使用 `equalsIgnoreCase` 方法。

```java
"Hello".equalsIgnoreCase("hello"); 
```

不要使用 `==` 运算符检测两个字符串是否相等！这个运算符只能够确定两个字符串是否存放在同一个位置上。

#### 空串和 Null 串

空串 "" 是长度为 0 的字符串，`String` 变量还可以存放一个特殊的值，名为 `null`, 表示目前没有任何对象与该变量关联。

```java
if (str.length() = 0)

if (str.equals(""))

if (str == null)
```

有时要检查一个字符串既不是 `null` 也不是空串，这种情况下可以使用：

```java
if (str != null && !str.isEmpty())
```

#### 构建字符串

采用字符串拼接的方式，每次拼接时会产生新的 `String` 对象，性能比较低。使用 `StringBuilder` 来构建字符串可以避免这个问题。

```java
StringBuilder sb = new StringBuilder();
sb.append("Hello");
sb.append(" World");
String result = sb.toString();
System.out.println(result);
// 输出: Hello World
```

## 复合类型

### 数组

在 `Java` 中，数组是一种用于存储固定大小的同一类型元素的容器。数组可以存储基本数据类型或对象。

#### 声明:

```java
int[] numbers;      // 推荐方式
int numbers[];      // 也合法，但不推荐，C/C++风格
```

#### 初始化:

```java
int[] arr = new int[5];  // 创建一个长度为5的整型数组，默认值为0

int[] arr = {1, 2, 3, 4, 5};

int[] arr = new int[]{1, 2, 3, 4, 5};
```

#### 访问数组元素:

```java
for (int i = 0; i < arr.length; i++) {
    System.out.println(arr[i]);
}

// Java 5 之后可以使用增强 for 循环
for (int num : arr) {
    System.out.println(num);
}
```

#### 数组拷贝：

允许将一个数组变量拷贝到另一个数组变量。这时, 两个变量将引用同一个数组：

```java
int[] arr1 = {1, 2, 3};
int[] arr2 = arr1;  // arr2 引用 arr1 的数组
// 修改 arr2 的元素会影响 arr1
arr2[0] = 10;
System.out.println(arr1[0]);  // 输出 10
```

如果确实希望将一个数组的所有值拷贝到一个新的数组中，就要使用类的 `copyOf` 方法：

```java
int[] copiedLuckyNumbers = Arrays.copyOf(luckyNLjnberst, luckyNumbers.length);
```

第 2 个参数是新数组的长度。

#### 数组排序

要想对数值型数组进行排序，可以使用 `Arrays` 类中的 `sort` 方法：

```java
int[] a = new int[1000];
...
Arrays.sort(a)
```

### ArrayList

`ArrayList` 类是一个可以动态修改的数组列表，与普通数组的区别就是它是没有固定大小的限制，我们可以添加或删除元素。

`ArrayList` 继承了 `AbstractList` ，并实现了 `List` 接口。

```java
import java.util.ArrayList;


ArrayList<E> objectName =new ArrayList<>();

// Java 10 中，最好使用以下语法
var objectName = new ArrayList<E>();
```

**E**: 泛型数据类型，用于设置 `objectName` 的数据类型，只能为引用数据类型。
**objectName**: 对象名。

`ArrayList` 是一个数组队列，提供了相关的添加、删除、修改、遍历等功能。

#### 基础语法

下面是一个简单的 `ArrayList` 使用示例：

```java
import java.util.ArrayList;

public class Test {
   public static void main(String[] args) {
      ArrayList<String> sites = new ArrayList<String>(100); // 指定初始容量为 100

      // 添加元素
      sites.add("Google");
      sites.add("Runoob");
      sites.add("Taobao");
      sites.add(3,"Weibo");
      System.out.println(sites);

      // 获取元素
      System.out.println(sites.get(1));

      // 修改元素
      sites.set(2, "Wiki");
      System.out.println(sites);

      // 删除元素
      sites.remove(2);
      System.out.println(sites);

      // 获取 ArrayList 的大小
      System.out.println(sites.size());
   }
}

// [Google, Runoob, Taobao, Weibo]
// Runoob
// [Google, Runoob, Wiki, Weibo]
// [Google, Runoob, Weibo]
// 3
```

#### 迭代数组

我们可以使用 `for` 来迭代数组列表中的元素：

```java
import java.util.ArrayList;

public class RunoobTest {
    public static void main(String[] args) {
        ArrayList<String> sites = new ArrayList<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Weibo");
        for (int i = 0; i < sites.size(); i++) {
            System.out.println(sites.get(i));
        }
    }
}

// Google
// Runoob
// Taobao
// Weibo
```

也可以使用 `for-each` 来迭代元素：

```java
import java.util.ArrayList;

public class RunoobTest {
    public static void main(String[] args) {
        ArrayList<String> sites = new ArrayList<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Weibo");
        for (String i : sites) {
            System.out.println(i);
        }
    }
}

// Google
// Runoob
// Taobao
// Weibo
```

#### 引用类型

`ArrayList` 中的元素实际上是对象，在以上实例中，数组列表元素都是字符串 `String` 类型。

如果我们要存储其他类型，而 `<E>` 只能为引用数据类型，这时我们就需要使用到基本类型的包装类。

基本类型对应的包装类表如下：

| 基本类型 | 引用类型  |
|----------|-----------|
| boolean  | Boolean   |
| byte     | Byte      |
| short    | Short     |
| int      | Integer   |
| long     | Long      |
| float    | Float     |
| double   | Double    |
| char     | Character |


以下实例使用 `ArrayList` 存储数字(使用 `Integer` 类型):

```java
import java.util.ArrayList;

public class RunoobTest {
    public static void main(String[] args) {
        ArrayList<Integer> myNumbers = new ArrayList<Integer>();
        myNumbers.add(10);
        myNumbers.add(15);
        myNumbers.add(20);
        myNumbers.add(25);
        for (int i : myNumbers) {
            System.out.println(i);
        }
    }
}

// 10
// 15
// 20
// 25
```

#### ArrayList 排序

`Collections` 类也是一个非常有用的类，位于 `java.util` 包中，提供的 `sort()` 方法可以对字符或数字列表进行排序。

以下实例对字母进行排序：

```java
import java.util.ArrayList;
import java.util.Collections;  // 引入 Collections 类

public class RunoobTest {
    public static void main(String[] args) {
        ArrayList<String> sites = new ArrayList<String>();
        sites.add("Taobao");
        sites.add("Wiki");
        sites.add("Runoob");
        sites.add("Weibo");
        sites.add("Google");
        Collections.sort(sites);  // 字母排序
        for (String i : sites) {
            System.out.println(i);
        }
    }
}

// Google
// Runoob
// Taobao
// Weibo
// Wiki
```

[更多 API 方法](https://www.runoob.com/manual/jdk11api/java.base/java/util/ArrayList.html)


### LinkedList

链表（Linked list）是一种常见的基础数据结构，是一种线性表，但是并不会按线性的顺序存储数据，而是在每一个节点里存到下一个节点的地址。

Java LinkedList（链表） 类似于 ArrayList，是一种常用的数据容器。

与 ArrayList 相比，LinkedList 的增加和删除的操作效率更高，而查找和修改的操作效率较低。

语法格式如下：

```java
// 引入 LinkedList 类
import java.util.LinkedList; 

LinkedList<E> list = new LinkedList<E>();   // 普通创建方法
或者
LinkedList<E> list = new LinkedList(Collection<? extends E> c); // 使用集合创建链表
```

#### 基础语法

以下是一个简单的 LinkedList 使用示例：

```java
import java.util.LinkedList;

public class Test {
    public static void main(String[] args) {
        LinkedList<String> sites = new LinkedList<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");

        // 使用 addFirst() 在头部添加元素
        sites.addFirst("Wiki");
        System.out.println(sites);

        // 使用 addLast() 在尾部添加元素
        sites.addLast("Wiki");
        System.out.println(sites);

        // 使用 removeFirst() 移除头部元素
        sites.removeFirst();
        System.out.println(sites);

        // 使用 removeLast() 移除尾部元素
        sites.removeLast();
        System.out.println(sites);

        // 使用 getFirst() 获取头部元素
        System.out.println(sites.getFirst());

        // 使用 getLast() 获取尾部元素
        System.out.println(sites.getLast());
    }
}

// [Wiki, Google, Runoob, Taobao]
// [Wiki, Google, Runoob, Taobao, Wiki]
// [Google, Runoob, Taobao, Wiki]
// [Google, Runoob, Taobao]
// Google
// Taobao
```

#### 迭代链表

我们可以使用 for 来迭代链表中的元素：

```java
// 引入 LinkedList 类
import java.util.LinkedList;

public class RunoobTest {
    public static void main(String[] args) {
        LinkedList<String> sites = new LinkedList<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Weibo");
        for (int size = sites.size(), i = 0; i < size; i++) {
            System.out.println(sites.get(i));
        }
    }
}
```

也可以使用 for-each 来迭代元素：

```java
// 引入 LinkedList 类
import java.util.LinkedList;

public class RunoobTest {
    public static void main(String[] args) {
        LinkedList<String> sites = new LinkedList<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Weibo");
        for (String i : sites) {
            System.out.println(i);
        }
    }
}
```

[更多 API 方法](https://www.runoob.com/manual/jdk11api/java.base/java/util/LinkedList.html)

### HashSet

HashSet 基于 HashMap 来实现的，是一个不允许有重复元素的集合。

HashSet 允许有 null 值。

HashSet 是无序的，即不会记录插入的顺序。

HashSet 不是线程安全的， 如果多个线程尝试同时修改 HashSet，则最终结果是不确定的。 您必须在多线程访问时显式同步对 HashSet 的并发访问。

HashSet 中的元素实际上是对象，一些常见的基本类型可以使用它的包装类。

语法格式如下：

```java
import java.util.HashSet; // 引入 HashSet 类

HashSet<String> sites = new HashSet<String>();
```

#### 基础语法

以下是一个简单的 HashSet 使用示例：

```java
import java.util.HashSet;

public class Test {
    public static void main(String[] args) {
    HashSet<String> sites = new HashSet<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Zhihu");

        // 重复的元素不会被添加
        sites.add("Runoob");
        System.out.println(sites);

        // 检查元素是否存在
        System.out.println(sites.contains("Taobao"));

        // 删除元素，删除成功返回 true，否则为 false
        sites.remove("Taobao");
        System.out.println(sites);

        // 清空集合
        sites.clear();  
        System.out.println(sites);

        // 计算大小
        System.out.println("集合大小: " + sites.size());
    }
}

// [Google, Runoob, Zhihu, Taobao]
// true
// [Google, Runoob, Zhihu]
// []
// 集合大小: 0
```

#### 迭代 HashSet

我们可以使用 for-each 来迭代 HashSet 中的元素：

```java
// 引入 HashSet 类      
import java.util.HashSet;

public class RunoobTest {
    public static void main(String[] args) {
    HashSet<String> sites = new HashSet<String>();
        sites.add("Google");
        sites.add("Runoob");
        sites.add("Taobao");
        sites.add("Zhihu");
        sites.add("Runoob");     // 重复的元素不会被添加
        for (String i : sites) {
            System.out.println(i);
        }
    }
}

// Google
// Runoob
// Zhihu
// Taobao
```

[更多 API 方法](https://www.runoob.com/manual/jdk11api/java.base/java/util/HashSet.html)

#### HashMap

`HashMap` 是一个散列映射，它存储的内容是键值对映射。实现了 `Map` 接口，根据键的 `HashCode` 值存储数据，不支持线程同步。

声明语法如下：

```java
import java.util.HashMap;

HashMap<Integer, String> sites = new HashMap<Integer, String>();
```

#### 基础语法

常见的基础操作命令：

```java  
import java.util.HashMap;

public class Test {
    public static void main(String[] args) {
        HashMap<Integer, String> sites = new HashMap<Integer, String>();
        
        sites.put(1, "Google");
        sites.put(2, "Runoob");
        sites.put(3, "Taobao");
        sites.put(4, "Zhihu");
        System.out.println(sites);

        // 获取元素
        System.out.println(sites.get(3));

        // 删除元素
        sites.remove(4);
        System.out.println(sites);

        // 清空
        sites.clear();
        System.out.println(sites);

        // 获取大小
        System.out.println(sites.size());
    }
}

// {1=Google, 2=Runoob, 3=Taobao, 4=Zhihu}
// Taobao
// {1=Google, 2=Runoob, 3=Taobao}
// {}
// 0
```

#### 迭代 HashMap

我们可以使用 `for-each` 来迭代 `HashMap` 中的键值对：

```java
import java.util.HashMap;

public class Test {
    public static void main(String[] args) {
        HashMap<Integer, String> sites = new HashMap<Integer, String>();

        sites.put(1, "Google");
        sites.put(2, "Runoob");
        sites.put(3, "Taobao");
        sites.put(4, "Zhihu");

        // 返回所有 key 值
        for (Integer i : sites.keySet()) {
            System.out.println("key: " + i + " value: " + sites.get(i));
        }

        // 返回所有 value 值
        for(String value: sites.values()) {
          System.out.print(value + ", ");
        }

        System.out.println();

        // 输出所有键值对
        sites.forEach((key, value) -> {
            System.out.println(key + " " + value);
        });
    }
}
// key: 1 value: Google
// key: 2 value: Runoob
// key: 3 value: Taobao
// key: 4 value: Zhihu
// Google, Runoob, Taobao, Zhihu, 
// 1 Google
// 2 Runoob
// 3 Taobao
// 4 Zhihu
```

[更多 API 方法](https://www.runoob.com/manual/jdk11api/java.base/java/util/HashMap.html)

## 变量与常量

### 声明变量

声明一个变量时，先指定变量的类型, 然后是变量名。

```java
double salary;
int vacationDays;
long earthPofwlation;
boolean done;
```

可以在一行中声明多个变量。

```java
int i j;
```

### 初始化变量

声明一个变量之后，必须用赋值语句显式地初始化变量. 千万不要使用未初始化的变量的值。

```java
int vacatlonDays;
System.out.println(vacatlonDays);

// ERROR
// The local variable id may not have been initialized
```

下面两种初始化方式是等价的。

```java
int vacationDays;
vacationDays = 12;
// 也可以将变量的声明和初始化放在同一行中
int vacationDays = 12;
```

### 常量

用关键字 `final` 指示常量，常量只能被赋值一次。一旦赋值，就不能再更改了，通常常量名使用大写字母。

```java
public class Constants
{
    public static void main( String[] args)
    {
        final double CM_PER_INCH = 2.54;
        double paperWidth = 8.5;
        double paperHeight = 11;
        System.out.println( "Paper size in centimeters: "
            + paperWidth * CM_PER_INCH + " by " + paperHeight * CM_PER_INCH);
    } 
}
```

使用关键字 `static final` 设置一个类常量。

```java
public class Constants
{
    static final double CM_PER_INCH = 2.54;
    public static void main( String[] args)
    {
        double paperWidth = 8.5;
        double paperHeight = 11;
        System.out.println( "Paper size in centimeters: "
            + paperWidth * CM_PER_INCH + " by " + paperHeight * CM_PER_INCH);
    } 
}
```


## 修饰符
修饰符是 `Java` 中用于控制类、方法、变量等访问权限和行为的关键字。它们可以分为访问修饰符和非访问修饰符两大类。

### default

`default` 表示默认，什么也不写，在同一包内可见，不使用任何修饰符。使用对象：类、接口、变量、方法。

如果在类、变量、方法或构造函数的定义中没有指定任何访问修饰符，那么它们就默认具有默认访问修饰符。

默认访问修饰符的访问级别是包级别，即只能被同一包中的其他类访问。

如下例所示，变量和方法的声明可以不使用任何修饰符。

```java
// MyClass.java
 
class MyClass {  // 默认访问修饰符
 
    int x = 10;  // 默认访问修饰符
 
    void display() {  // 默认访问修饰符
        System.out.println("Value of x is: " + x);
    }
}
 
// MyOtherClass.java
 
class MyOtherClass {
    public static void main(String[] args) {
        MyClass obj = new MyClass();
        obj.display();  // 访问 MyClass 中的默认访问修饰符变量和方法
    }
}
```

以上实例中，`MyClass` 类和它的成员变量 x 和方法 `display()` 都使用默认访问修饰符进行了定义。`MyOtherClass` 类在同一包中，因此可以访问 `MyClass` 类和它的成员变量和方法。

`default` 与 `protected` 的区别在于，`protected` 访问修饰符允许子类访问父类的成员变量和方法，即使它们不在同一包中。而 `default` 访问修饰符只允许同一包中的类（子类）访问。


### private

在 `Java` 中，`private` 是最严格的访问控制修饰符，用于实现封装。它限定成员只能在 声明它的类内部被访问，外部包括子类和同包的其他类都无法访问。使用对象：变量、方法。

```java
public class Logger {
   private String format;
   public String getFormat() {
      return this.format;
   }
   public void setFormat(String format) {
      this.format = format;
   }
}
```

`private` 访问修饰符的使用主要用来隐藏类的实现细节和保护类的数据。

### public

在 `Java` 中，`public` 是访问权限最大的修饰符，表示可以被任何其他类访问，不受包和继承关系的限制。使用对象：类、接口、变量、方法

被声明为 `public` 的类、方法、构造方法和接口能够被任何其他类访问。

如果几个相互访问的 `public` 类分布在不同的包中，则需要导入相应 `public` 类所在的包。由于类的继承性，类所有的公有方法和变量都能被其子类继承。

```java
// 文件：people/Person.java
package people;

public class Person {
    public String name = "Alice";       // public 字段

    public void sayHello() {            // public 方法
        System.out.println("Hello, I'm " + name);
    }
}

// 文件：test/TestPerson.java
package test;

import people.Person;

public class TestPerson {
    public static void main(String[] args) {
        Person p = new Person();
        System.out.println(p.name);    // 合法访问 public 字段
        p.sayHello();                  // 合法访问 public 方法
    }
}
```

### protected

在 `Java` 中，`protected` 是一种访问修饰符，使用对象：变量、方法，用于控制类成员（属性、方法、构造器）的可见性。
它的访问范围介于 `private` 和 `public` 之间，具体如下：

`protected` 的访问范围：

1. **同一包内的类**：`protected` 成员可以被同一包内的其他类访问。

2. **子类**：`protected` 成员可以被任何子类访问，无论子类是否在同一包内。

3. **不同包的非子类**：`protected` 成员不能被不同包内的非子类访问。

测试目录结构

```
src/
├── a/
│   ├── Base.java
│   ├── A1.java
│   └── A2.java
└── b/
    ├── B1.java
    └── B2.java
```

测试文件 `Base.java`

```java
package a;

public class Base {
    protected String msg = "protected value";

    protected void sayHello() {
        System.out.println("Base says: " + msg);
    }
}
```

测试文件 `A1.java`, 同包非继承。

```java
package a;

public class A1 {
    public static void main(String[] args) {
        Base base = new Base();
        System.out.println(base.msg); 
        base.sayHello();   
        
        // sayHello(); // The method sayHello() is undefined for the type A1
        // System.out.println(msg); // msg cannot be resolved to a variable
    }
    
    // public void bark() {
    //     System.out.println(msg); // The method sayHello() is undefined for the type A1
    //     sayHello(); // msg cannot be resolved to a variable
    // }
}
```

测试文件 `A2.java`，同包继承。

```java
package a;

public class A2 extends Base {
    public static void main(String[] args) {
        A2 obj = new A2();
        System.out.println(obj.msg); 
        obj.sayHello();

        // System.out.println(msg); // Cannot make a static reference to the non-static field msg
        // sayHello(); // Cannot make a static reference to the non-static method sayHello() from the type A2

        Base base = new Base();
        System.out.println(base.msg);
        base.sayHello(); // OK
    }

    public void bark() {
        // 这里可以访问父类的 protected 方法
        System.out.println(msg);
        sayHello(); // OK
    }
}
```

测试文件 `B1.java`，不同包非继承。

```java
package b;

import a.Base;

public class B1 {
    public static void main(String[] args) {
        Base base = new Base(); 
        // System.out.println(base.msg); // The field Base.msg is not visible
        // base.sayHello();              // The method sayHello() from the type Base is not visible
    }
}
```

测试文件 `B2.java`，不同包继承。

```java
package b;

import a.Base;

public class B2 extends Base {
    public static void main(String[] args) {
        B2 obj = new B2();
        System.out.println(obj.msg); 
        obj.sayHello();              

        // System.out.println(msg); // Cannot make a static reference to the non-static field msg
        // sayHello(); // Cannot make a static reference to the non-static method sayHello() from the type A2

        // 无法通过父类对象访问
        // Base base = new Base();
        // System.out.println(base.msg); // The field Base.msg is not visible
        // base.sayHello();              // The method sayHello() from the type Base is not visible
    }

    public void bark() {
        // 这里可以访问父类的 protected 方法
        System.out.println(msg); // OK
        sayHello();              // OK
    }
}
```

### static

`static` 是一个非访问修饰符，用于表示类的成员（变量和方法）属于类本身，用来修饰类方法和类变量。

**静态变量**：`static` 关键字用来声明独立于对象的静态变量，无论一个类实例化多少对象，它的静态变量只有一份拷贝。
静态变量也被称为类变量。局部变量不能被声明为 `static` 变量。

**静态方法**：`static` 关键字用来声明独立于对象的静态方法。静态方法不能使用类的非静态变量。静态方法从参数列表得到数据，然后计算这些数据。

```java
public class InstanceCounter {
    private static int numInstances = 0;
    protected static int getCount() {
       return numInstances;
    }
  
    private static void addInstance() {
       numInstances++;
    }
  
    InstanceCounter() {
       InstanceCounter.addInstance();
    }
  
    public static void main(String[] arguments) {
       System.out.println("Starting with " +
       InstanceCounter.getCount() + " instances");
       for (int i = 0; i < 500; ++i){
         InstanceCounter.addInstance();
         new InstanceCounter();
      }
       System.out.println("Created " +
       InstanceCounter.getCount() + " instances");
    }
 }
```

输出结果：

```
Starting with 0 instances
Created 1000 instances
```

`static` 还可以声明静态代码块，静态代码块在类加载时执行一次，用于初始化类的静态变量。

```java
public class StaticExample {
    static int staticVariable;

    static {
        staticVariable = 10;
        System.out.println("Static block executed. Static variable initialized to: " + staticVariable);
    }

    public static void main(String[] args) {
        System.out.println("Main method executed. Static variable value: " + staticVariable);
    }
}
```

输出结果：

```
Static block executed. Static variable initialized to: 10
Main method executed. Static variable value: 10
```

### final

`final` 表示"最后的、最终的"含义，变量一旦赋值后，不能被重新赋值。被 `final` 修饰的实例变量必须显式指定初始值。
通常和 `static` 修饰符一起使用来创建类常量。

```java
public class Test{
  final int value = 10;
  // 下面是声明常量的实例
  public static final int BOXWIDTH = 6;
  static final String TITLE = "Manager";
 
  public void changeValue(){
     value = 12; //将输出一个错误
  }
}
```

`final` 可以修饰引用类型变量（如 Map），意味着：

* 变量本身的引用不能变（即不能执行 instances = new HashMap<>(); 这种操作）；

* 引用指向的对象的内容可以变（即 instances.put(...) 是允许的）；

```java
private static final Map<String, EventHeaderKey> instances = new HashMap<>();
```

父类中的 `final` 方法可以被子类继承，但是不能被子类重写。声明 `final` 方法的主要目的是防止该方法的内容被修改。

```java
public class Test{
    public final void changeName(){
       // 方法体
    }
}
```

### abstract

抽象类不能用来实例化对象，声明抽象类的唯一目的是为了将来对该类进行扩充。

一个类不能同时被 `abstract` 和 `final` 修饰。如果一个类包含抽象方法，那么该类一定要声明为抽象类，否则将出现编译错误。

抽象类可以包含抽象方法和非抽象方法。

```java
abstract class Caravan{
   private double price;
   private String model;
   private String year;
   public abstract void goFast(); //抽象方法
   public abstract void changeColor();
}
```

抽象方法是一种没有任何实现的方法，该方法的具体实现由子类提供。

抽象方法不能被声明成 `final` 和 `static`。

任何继承抽象类的子类必须实现父类的所有抽象方法，除非该子类也是抽象类。

如果一个类包含若干个抽象方法，那么该类必须声明为抽象类。抽象方法的声明以分号结尾。

```java
public abstract class SuperClass{
    abstract void m(); //抽象方法
}
 
class SubClass extends SuperClass{
     //实现抽象方法
      void m(){
          .........
      }
}
```

### synchronized

关键字 `synchronized` 声明的方法同一时间只能被一个线程访问。

Java 中的每个对象都有一个内部锁（intrinsic lock）。如果一个方法声明时有 synchronized 关键字，那么对象的锁将保护该方法的执行，也就要调用必须先获取锁。

```java
public synchronized void method() {
    // method body
}

等价于
public void method() {
    synchronized(this) {
        // method body
    }
}

等价于
public void method() {
    this.intrinsicLock.lock();
    try {
        // method body
    } finally {
        this.intrinsicLock.unlock();
    }
}
```

修饰符 `synchronized` 有三种使用方式：

1、修饰实例方法（锁住当前实例 this）。对象中多个 `synchronized` 函数共用一个锁，只能同一时刻执行一个代码块。

```java
public synchronized void method() { ... }
```

2、修饰静态方法（锁住类的 Class 对象）。锁定类的 Class 对象，同一类的静态方法使用的同一个锁，和对象内部锁不同。

```java
public static synchronized void staticMethod() { ... }


public class Example {
    public static synchronized void staticMethod() {
        // method body
    }
}

这等价于：
public static void staticMethod() {
    synchronized (Example.class) {
        // method body
    }
}
```

3、修饰代码块（锁住指定对象）。

```java
public class Example {

    private final Object lock = new Object();

    public void method() {
        synchronized (lock) {
            // method body
        }
    }
```

### volatile

`volatile` 是一个非访问修饰符，用于声明变量的可见性和禁止指令重排序。它确保变量的值在多个线程之间保持一致。

* 多处理器能够在寄存器或本地缓存中保存内存值，导致不同处理器上的线程看到同一内存位置有不同的值。

* 编译器可能改变指令执行的顺序以到达最大的吞吐量。

使用锁机制 `synchronized`、`Lock` 可以解决可见性问题，但会导致性能下降。

```java
private volatile boolean done;

public boolean isDone() {
    return done;
}

public void setDone(boolean done) {
    this.done = done;
}
```

关键字 `volatile` 为实例字段的同步访问提供了免锁机制，编辑器会插入适当的代码来禁止指令重排。
`Java` 内存模型会插入内存屏障（一个处理器指令，可以对 CPU 或编译器重排序做出约束）来确保以下两点：

* 写屏障（Write Barrier）：当一个 `volatile` 变量被写入时，写屏障确保在该屏障之前的所有变量的写入操作都提交到主内存。
* 读屏障（Read Barrier）：当读取一个 `volatile` 变量时，读屏障确保在该屏障之后的所有读操作都从主内存中读取。

注意：关键字 `volatile` 的变量不能提供原子性。

#### 实现单例模式的双重锁

下面是一个使用 "双重检查锁定"（double-checked locking）实现的单例模式（Singleton Pattern）的例子

```java
public class Penguin {
    private static volatile Penguin m_penguin = null;

    // 一个成员变量 money
    private int money = 10000;

    // 避免通过 new 初始化对象，构造方法应为 private
    private Penguin() {}

    public void beating() {
        System.out.println("打豆豆" + money);
    }

    public static Penguin getInstance() {
        if (m_penguin == null) {
            synchronized (Penguin.class) {
                if (m_penguin == null) {
                    m_penguin = new Penguin();
                }
            }
        }
        return m_penguin;
    }
}
```

使用 `volatile` 关键字是为了防止 `m_penguin = new Penguin()` 这一步被指令重排序。因为实际上，`new Penguin()` 这一行代码分为三个子步骤：

* 步骤 1：为 Penguin 对象分配足够的内存空间，伪代码 memory = allocate()。

* 步骤 2：调用 Penguin 的构造方法，初始化对象的成员变量，伪代码 ctorInstanc(memory)。

* 步骤 3：将内存地址赋值给 m_penguin 变量，使其指向新创建的对象，伪代码 instance = memory。

如果不使用 volatile 关键字，JVM 可能会对这三个子步骤进行指令重排。

* 为 Penguin 对象分配内存
* 将对象赋值给引用 m_penguin
* 调用构造方法初始化成员变量


## 流程控制

### while

`while` 循环是 `Java` 中的一种循环控制结构，用于在满足特定条件时重复执行一段代码。它的基本语法如下：

```java
while( 布尔表达式 ) {
  //循环内容
}
```

`while` 循环的执行过程如下：

```java
public class Test {
   public static void main(String[] args) {
      int x = 10;
      while( x < 20 ) {
         System.out.print("value of x : " + x );
         x++;
         System.out.print("\n");
      }
   }
}
```

### do while

`do while` 循环是 `Java` 中的一种循环控制结构，与 `while` 循环类似，但它至少会执行一次循环体。它的基本语法如下：

```java
do {
  //循环内容
} while( 布尔表达式 );
```

`do while` 循环的执行过程如下：

```java
public class Test {
   public static void main(String[] args){
      int x = 10;
 
      do{
         System.out.print("value of x : " + x );
         x++;
         System.out.print("\n");
      }while( x < 10 );
   }
}

// 输出结果
// value of x : 10
```

### for

`for` 循环是 `Java` 中的一种循环控制结构，用于在满足特定条件时重复执行一段代码。它的基本语法如下：

```java
for (初始化; 布尔表达式; 更新) {
  //循环内容
}
```

`for` 循环的执行过程如下：

```java
public class Test {
   public static void main(String[] args) {
 
      for(int x = 10; x < 20; x = x+1) {
         System.out.print("value of x : " + x );
         System.out.print("\n");
      }
   }
}
```

### break

`break` 主要用在循环语句或者 `switch` 语句中，用来跳出整个语句块。

### continue

`continue` 主要用在循环语句中，用来跳过当前循环的剩余部分，直接进入下一次循环。

### if else

`if else` 语句是 `Java` 中的一种条件控制结构。

```java
public class Test {
   public static void main(String args[]){
      int x = 30;
      if( x == 10 ){
         System.out.print("Value of X is 10");
      }else if( x == 20 ){
         System.out.print("Value of X is 20");
      }else if( x == 30 ){
         System.out.print("Value of X is 30");
      }else{
         System.out.print("这是 else 语句");
      }
   }
}
```

### switch case

`switch case` 执行时，一定会先进行匹配，匹配成功返回当前 `case` 的值，再根据是否有 `break`，判断是否继续输出。

```java
public class Test {
   public static void main(String args[]){
      //char grade = args[0].charAt(0);
      char grade = 'B';
 
      switch(grade)
      {
         case 'A' :
            System.out.println("优秀"); 
            break;
         case 'B' :
         case 'C' :
            System.out.println("良好");
         case 'D' :
            System.out.println("及格");
            break;
         case 'F' :
            System.out.println("你需要再努力努力");
            break;
         default :
            System.out.println("未知等级");
      }
      System.out.println("你的等级是 " + grade);
   }
}

// 良好
// 及格
// 你的等级是 B
```

如果当前匹配成功的 `case` 语句块没有 `break` 语句，则从当前 `case` 开始，后续所有 `case` 的值都会输出。

```java
public class Test {
   public static void main(String args[]){
      int i = 2;
      switch(i){
         case 0:
            System.out.println("0");
         case 1:
            System.out.println("1");
         case 2:
            System.out.println("2");
         default:
            System.out.println("default");
      }
   }
}

// 2
// default
```

如果都不匹配走 `default`，如果没有 `default` 语句，则不执行任何语句。


## 对象与类

类 (class) 指定了如何构造对象。由一个类构造 (construct) 对象的过程称为创建这个类的一个实例 ( instance )。

### 对象

要想使用对象，首先必须构造对象，并指定其初始状态。然后对对象应用方法。

在 `Java` 中，任何对象变量的值都是一个引用，指向存储在另外一个地方的某个对象。`new` 操作符的返回值也是一个引用。

```java
Date startTime = new Date();
```

表达式 `new Date()` 构造了一个 `Date` 类型的对象, 它的值是新创建对象的一个引用，再将这个引用存储在 `startTime` 变量中。

可以显式地将对象变量设置为 `null`, 指示这个对象变量目前没有引用任何对象。

```java
Date startTime = null;
```

### 类

常见的类定义示例 `Employee.java` 如下：

```java
public class Employee {
    private String name;
    private double salary;
    private LocalDate hireDate;

    public Employee(String n, double s, int year, int month, int day) {
        name = n;
        salary = s;
        hireDate = LocalDate.of(year, month, day);

    }
    public String getName() {
        return name;
    }
    // more methods
    // ......
}
```

源文件名是 `Employee.java`，文件名必须与 `public` 类名相同。一个源文件中只能有一个公共类，但可以有任意数目的非公共类。

示例中有 3 个实例字段，关键字 private 确保这些字段只能在 `Employee` 类的内部访问，任何其他类的方法都不能读写这些字段。

```java
private String name;
private double salary;
private LocalDate hireDate;
```

#### 隐式参数和显式参数

方法会操作对象并访问它们的实例字段。

```java
public void raiseSalary(double byPercent) {
    double raise = salary * byPercent / 100;
    salary += raise;
}
```

方法 `raiseSalary` 有两个参数，一个是显式参数 `byPercent`，另一个是隐式参数 `this`，指向调用该方法的对象。

可以如下改写 `raiseSalary` 方法，显示地使用 `this` 关键字来引用隐式参数：

```java
public void raiseSalary(double byPercent) {
    double raise = this.salary * byPercent / 100; // 显式使用 this
    this.salary += raise; // 显式使用 this
}
```

#### 基于类的访问权限

方法可以调用这个方法的对象的私有数据，一个类的方法可以访问这个类的所有对象的私有数据。

```java
{
    public boolean equals(Employee other) {
        return name.equals(other.name)
    }
}

// 调用
harry.equals(boss);
```

`boss` 是 `Employee` 类型的对象，`harry` 也是 `Employee` 类型的对象，`equals` 方法可以访问 `boss` 的私有数据。

#### final 实例字段

在 `Java` 中，`final` 关键字可以用于修饰实例字段，这样的字段必须在构造对象时初始化，并且之后不能再进行修改。

```java
class Employee {
    private final String name; // final 实例字段
}
```

#### 方法参数

在 `Java` 中，方法参数可以是基本类型或引用类型。对于基本类型(数值型和布尔型)，参数传递的是值的副本；对于引用类型，传递的是对象的引用。

#### 构造器

下面看看 `Employee` 类的构造器：

```java
public Employee(String n, double s, int year, int month, int day) {
    name = n;
    salary = s;
    hireDate = LocalDate.of(year, month, day);
}
```

构造器具有下面特点：

* 构造器的名称必须与类名相同。
* 每个类可以有多个构造器，允许不同的参数列表。
* 构造器可以有0个、1个或多个参数。
* 构造器没有返回值类型，甚至没有 `void`。
* 构造器总是结合 `new` 关键字使用，当创建一个对象时，构造器会被自动调用。

注意：不要引入与实例字段同名的局部变量，构造器将不会正确设置实例字段。

```java
public Employee(String n, double s, ....) {
    String name = n;    // ERROR
    double salary = s;  // ERROR
    ...
}
```

**默认字段初始化**

如果没有在构造器中显式地为一个字段设置初始值，就会将它自动设置为默认值，数值将设置为 0 ，布尔值将设置为 false，引用类型将设置为 null。

**无参数构造器**

以下是 `Employee` 类的无参数构造器示例：

```java
public Employee() {
    name = "Unknown";
    salary = 0;
    hireDate = LocalDate.now();
}
```

如果没有构造器，会为你提供一个无参数构造器，构造器将为所有的实例字段设置默认值。

#### 类设计技巧

1、一定要保证数据私有。

2、一定要初始化数据。

3、不要在类中使用过多的基本类型（多个基本类型可以单独封装为类）。

4、不是所有的字段都需要单独的字段访问器和更改器

5、分解有过多职责的类。

6、类名和方法名要能够体现它们的职责。

7、优先使用不可变的类。


## 继承

### 类、超类、子类

#### 定义子类

在 `Java` 中通过 `extends` 关键字可以申明一个类是从另外一个类继承而来的：

```java
public class Manager extends Employee 
{
    added methods and fields
}
```

特性:

* 子类可以继承父类的所有字段（不能直接访问private字段）、非private方法;

* 子类可以拥有自己的字段和方法，可以对父类进行扩展;

* 子类可以重写父类的方法;

#### 覆盖方法

在子类中可以覆盖父类的方法，但是如何在方法里调用父类的同名方法呢？

```java
public double getSalary() {
    return super.getSalary() + bonus; // 调用父类的 getSalary 方法
}
```

注：`super` 不是一个对象的引用，只是一个指示编译器调用超类方法的特殊关键字。

#### 子类构造器

子类的构造器必须调用父类的构造器，通常通过 `super` 关键字来实现。

```java
public Manager(String n, double s, int year, int month, int day) {
    super(n, s, year, month, day); // 调用父类的构造器
    bonus = 0;
}
```

由于 `Manager` 类的构造器不能访问 `Employee` 类的私有字段，因此必须通过 `super` 调用父类的构造器来初始化继承的字段。
使用 `super` 调用构造器的语句必须是子类构造器的第一个语句。


#### 阻止继承：final 类和方法

有时候，我们可能希望阻止人们定义某个类的子类，不允许扩展的类被称为 `final` 类。

```java
public final class Executive extends Manager {
    ...
}
```

也可以将类中的某个特定方法声明为 `final`。如果这样做，那么所有子类都不能覆盖这个方法，例如：

```java
public class Employee {
    public final double getSalary() {
        return salary;
    }
}
```

字段也可以被声明为 `final`，构造对象之后就允许改变了。

#### 强制类型转换

在 `Java` 中，子类对象可以被赋值给父类引用。

```java
var staff = new Employee

staff [0] = new Manager("Carl Cracker", 80000, 1987, 12, 15);
staff [1] = new Employee("Harry Hacker"，50000, 1989, 10, 1);
staff [2] = new Employee("Tony Tester", 40000, 1998, 3. 15);
```

如果 `staff[0]` 要用到 `Manager` 的方法特性，需要还原对象原本类型，进行强制转换。

```java
Manager boss = (Manager) staff[0];
```

如果 `staff[0]` 不是 `Manager` 类型的对象，那么强制转换将会失败，抛出 `ClassCastException` 异常。

```java
if (staff[i] instanceof Manager) {
    Manager boss = (Manager) staff[i];
}
```

在将超类强制转换成子类之前，应该使用 `instanceof` 进行检查。

在 `Java16` 及更高版本中，可以使用 `instanceof` 进行类型检查和强制转换的简化写法：

```java
if (staff[i] instanceof Manager boss) {
    boss.setBonus(5000);
}
```

### Object：所有类的超类

可以使用 `object` 类型的变量引用任何类型的对象。

```java
Object obj = new Employee("Alice", 50000, 2020, 1, 1);
```
在 `Java` 中，只有基本类型不是对象，例如：数值、字符和布尔类型的值都不是对象。

`Object` 类中有一些通用方法，例如 `toString()`、`equals()` 和 `hashCode()`，所有类都继承了这些方法。

### 抽象类

### 枚举类

枚举类（`enum`）是 `Java` 中的一种特殊类型的类，用于定义一组常量。枚举类可以包含字段、方法和构造器。

```java
public enum Size {
    SMALL, MEDIUM, LARGE, EXTRA_LARGE
}
```

所有的枚举类型都是抽象类 `Enum` 的子类中 它们继承了这个类的许多方法。其中，最有用的一个是 `toString`, 这个方法会返回枚举常量名。
例如，`Size.SMALL.toString()` 将返回字符串 "SMALL"。

方法 `toString` 的逆方法是 `valueOf`，它可以将字符串转换为枚举常量。

```java
Size s = Enum.valueOf(Size.class, "SMALL");
```

将 `s` 设置为 `Size.SMALL`，如果没有找到匹配的枚举常量，将抛出 `IllegalArgumentException` 异常。

每个枚举类型都有一个 `values` 方法，它返回一个包含所有枚举常量的数组。

```java
Size [] values = Size.values();
for (Size size : values) {
    System.out.println(size);
}
```

下面是枚举类的真实使用示例。

```java
public enum EventHeaderKey {
    DATA_KEY("x-log-key"),
    DATA_TIME("x-log-time"),
    DATA_TTL("x-log-ttl"),
    DATA_REGION("x-log-region"),
    DATA_LOGSTORE_NAME("x-log-logstore-name"),
    DATA_TIMEZONE("x-log-timezone");

    private final String value;

    EventHeaderKey(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }
}

public class TestEnum {
    public static void main(String[] args) {
        EventHeaderKey key = EventHeaderKey.DATA_KEY;
        System.out.println(key.getValue()); // 输出：x-log-key

        // 遍历所有枚举值
        for (EventHeaderKey eventHeaderKey : EventHeaderKey.values()) {
            System.out.println(eventHeaderKey + ": " + eventHeaderKey.getValue());
        }
    }
}
```

## 接口

在 `Java` 中，接口（interface） 是一种抽象类型，用于定义类必须实现的一组方法。接口是 `Java` 实现多态和多继承的关键机制。

要点：

* 接口中的方法默认是 public abstract（即公开的、抽象的）。

* 接口中的变量默认是 public static final（即常量）。

* 类使用 implements 关键字来实现接口。

* 一个类可以实现多个接口，实现接口是支持 多继承 的方式

```java
// 定义接口
interface Animal {
    void speak(); // 接口中的方法默认是 public abstract
    String getType();
}

// 实现接口
class Dog implements Animal {
    @Override
    public void speak() {
        System.out.println("Dog barks.");
    }

    @Override
    public String getType() {
        return "Dog";
    }
}

class Cat implements Animal {
    @Override
    public void speak() {
        System.out.println("Cat meows.");
    }

    @Override
    public String getType() {
        return "Cat";
    }
}

// 测试类
public class TestInterface {
    public static void main(String[] args) {
        Animal a1 = new Dog();
        Animal a2 = new Cat();

        a1.speak(); // Dog barks.
        a2.speak(); // Cat meows.
        System.out.println(a1.getType()); // Dog
    }
}
```

## 枚举

在 `Java` 中，enum（枚举）是一种特殊的类，用于定义常量集合。它是一种类型安全的方式来表示固定集合的值。

```java
enum Day {
    MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY
}

public class TestEnum {
    public static void main(String[] args) {
        Day today = Day.MONDAY;
        System.out.println(today); // 输出：MONDAY

        // 使用 switch 语句
        switch (today) {
            case MONDAY:
                System.out.println("Start of the week");
                break;
            case FRIDAY:
                System.out.println("Almost weekend");
                break;
            default:
                System.out.println("Midweek");
        }
    }
}
```

遍历所有枚举值

```java
for (Day day : Day.values()) {
    System.out.println(day);
}  
```

## 单例

常见的 `Java` 单例模式有下面几种实现方式：

1、静态内部类方式

```java
public class Singleton {
    private Singleton() {}

    public static Singleton getInstance() {
        return Holder.INSTANCE;
    }

    private static class Holder {
        private static final Singleton INSTANCE = new Singleton();
    }
}
```

特点：

* 懒加载：第一次调用 getInstance() 时才初始化
* 线程安全：由 JVM 保证类加载的线程安全性
* 性能好：无锁，无额外同步
* 推荐用法：高效 + 简洁

2、双重检查锁（Double-Checked Locking，DCL）

```java
public class Singleton {
    private static volatile Singleton instance;

    private Singleton() {}

    public static Singleton getInstance() {
        if (instance == null) {                 // 第一次检查（无锁）
            synchronized (Singleton.class) {
                if (instance == null) {         // 第二次检查（有锁）
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }
}
```

特点：

* 懒加载：是
* 线程安全：使用 volatile 防止指令重排
* 复杂性高：要理解内存模型、指令重排
* 性能一般：虽然锁只在初始化时进入，但代码复杂，容易出错
* Java 1.5 之后才推荐使用（在 Java 1.5 之前，Java 内存模型（JMM）不保证 volatile 能正确禁止指令重排，导致 DCL 实现可能不安全。）

3、枚举单例

```java
public enum Singleton {
    INSTANCE;

    public void doSomething() {
        System.out.println("Doing something...");
    }
}

// 使用方式：
Singleton.INSTANCE.doSomething();
```

特点：

* 不是懒加载：类加载时立即初始化
* 线程安全：由 JVM 枚举机制保证
* 防反射、防序列化攻击：天然安全
* 适合用于不可延迟初始化的单例

`JVM` 会为每个枚举值创建唯一的实例，类似于：

```java
public final class Singleton extends Enum<Singleton> {
    public static final Singleton INSTANCE = new Singleton();
    private Singleton() {}
}
```

`JVM` 不允许通过 `new` 创建枚举对象。

## 性能分析

### async-profiler

开源工具 `async-profiler` 是一个高性能的 Java 采样分析器，支持 CPU、内存和锁等多种分析模式。它可以生成火焰图（Flame Graph）来可视化性能数据。

下载工具：

```bash
$ wget https://github.com/async-profiler/async-profiler/releases/download/v4.1/async-profiler-4.1-linux-x64.tar.gz
$ tar zxvf async-profiler-4.1-linux-x64.tar.gz
```

获取要分析的 `Java` 进程的 `PID`：

```bash
$  jps -v
36561 Application -Xms56G -Xmx56G -XX:+UseG1GC -XX:ParallelGCThreads=32 -XX:MaxGCPauseMillis=150 -XX:InitiatingHeapOccupancyPercent=60 -javaagent:/usr/local/bls/flume/lib/jmx-prometheus-javaagent-0.16.1.jar=18017:/usr/local/bls/flume/conf/kafka-sink-jmx.yaml -Dflume.monitoring.type=http -Dflume.monitoring.port=41414 -Dlog4j.configurationFile=/usr/local/bls/current/conf.d/kafka-sink/log4j2.xml -Dkafka-sink.conf.path=/usr/local/bls/current/conf.d/conf.yaml -Djava.library.path=
12809 Jps -Dapplication.home=/usr/java/jdk1.8.0_181-amd64 -Xms8m
17551 QuorumPeerMain -Dzookeeper.log.dir=. -Dzookeeper.root.logger=INFO,CONSOLE -Dzookeeper.DigestAuthenticationProvider.superDigest=super:Zu5Tckgnn822Oi3gy2jMA7auDdE= -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.local.only=false

$ pgrep -a java
17551 java -Dzookeeper.log.dir=. -Dzookeeper.root.logger=INFO,CONSOLE -cp /usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../build/classes:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../build/lib/*.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../lib/slf4j-log4j12-1.6.1.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../lib/slf4j-api-1.6.1.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../lib/netty-3.10.5.Final.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../lib/log4j-1.2.16.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../lib/jline-0.9.94.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../zookeeper-3.4.10.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../src/java/lib/*.jar:/usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../conf: -Dzookeeper.DigestAuthenticationProvider.superDigest=super:Zu5Tckgnn822Oi3gy2jMA7auDdE= -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.local.only=false org.apache.zookeeper.server.quorum.QuorumPeerMain /usr/local/bls/zk4bls/zookeeper-3.4.10/bin/../conf/zoo.cfg
36561 /usr/bin/java -Xms56G -Xmx56G -XX:+UseG1GC -XX:ParallelGCThreads=32 -XX:MaxGCPauseMillis=150 -XX:InitiatingHeapOccupancyPercent=60 -javaagent:/usr/local/bls/flume/lib/jmx-prometheus-javaagent-0.16.1.jar=18017:/usr/local/bls/flume/conf/kafka-sink-jmx.yaml -Dflume.monitoring.type=http -Dflume.monitoring.port=41414 -Dlog4j.configurationFile=/usr/local/bls/current/conf.d/kafka-sink/log4j2.xml -Dkafka-sink.conf.path=/usr/local/bls/current/conf.d/conf.yaml -cp /usr/local/bls/flume/conf:/var/lib/flume-ng:/usr/local/bls/flume/lib/*:/usr/local/bls/flume/plugins.d/kafka-sink/lib/*:/usr/local/bls/flume/plugins.d/kafka-sink/libext/*:/lib/* -Djava.library.path= org.apache.flume.node.Application -n kafka-sink -f /usr/local/bls/current/conf.d/kafka-sink/kafka-sink.properties --no-reload-conf
```

执行命令生成 30s `CPU` 火焰图：

```bash
$ cd async-profiler-4.1-linux-x64/bin/
$ ./asprof -d 30 -f flamegraph.html 36561
$ put_file2 flamegraph.html
```

生成的 `flamegraph.html` 文件可以在浏览器中打开，查看火焰图。

<div align=center><img src="images/flamegraph_cpu.png" width=500></div>

执行命令生成 30s `Mem` 火焰图：

```bash
$ ./asprof -d 30 -e alloc -f flamegraph.html 36561
$ put_file2 flamegraph.html
```

直接在浏览器中打开。

<div align=center><img src="images/flamegraph_alloc.png" width=500></div>

更多参数使用，可以通过 `./asprof -h` 查看帮助文档。官网地址 [async-profiler](https://github.com/async-profiler/async-profiler)。

### arthas

工具 `arthas` 是 Alibaba 开源的 `Java` 诊断工具。


官网地址 [arthas](https://arthas.aliyun.com/en/doc/quick-start.html)
