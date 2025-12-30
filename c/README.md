# The C Programming Language

## 目录

* [导言](#导言)
    * [1.1 入门](#11-入门)
    * [1.2 变量和算术表达式](#12-变量和算术表达式)
    * [1.3 for 语句](#13-for语句)
    * [1.4 符号常量](#14-符号常量)
    * [1.5 字符输入/输出](#15-字符输入/输出)
    * [1.6 数组](#16-数组)
    * [1.7 函数](#17-函数)
    * [1.8 参数-传值调用](#18-参数-传值调用)
    * [1.9 字符数组](#19-字符数组)
    * [1.10 外部变量与作用域](#110-外部变量与作用域)

* [类型、运算符与表达式](#类型运算符与表达式)
    * [2.1 变量名](#21-变量名)
    * [2.2 数据类型及长度](#22-数据类型及长度)
    * [2.3 常量](#23-常量)
    

## 导言


### **1.1 入门**


编写下列程序，然后打印 "hello, world"。

```c
#include <stdio.h>

main()
{
    printf("hello, world\n");
}
```

以 ".c" 作为文件的扩展名，例如 hello.c ,然后通过下列命令进行编译：

```c
cc hello.c
```

生成可执行文件 a.out, 执行后输出结果：

```c
./a.out

hello, world
```

需要注意的是，printf 函数永远不会自动换行，上面的程序也可以改成下面的形式：

```c
#include <stdio.h>

main()
{
    printf("hello, ");
    printf("world");
    printf("\n");
}
```


### **1.2 变量和算术表达式**

C 语言只提供了下列几种基本数据类型：

```C
char    字符型，占用一个字节。
int     整型，通常反映了所有机器中整数的最自然长度。
float   单精度浮点型
double  双精度浮点型
```

此外，还可以在基本数据类型前面加一些限定符，`short` 与 `long` 两个限定符用于限定整数。

```C
short int sh;
long int counter;
```

编译器根据硬件特性自主选择合适的类型长度，但要遵循下列限制：

`short` 与 `int` 类型至少为 16 位，而 `long` 类型至少为 32 位，并且 `short` 类型不得长与 `int`类型，
而 `int` 类型不得长于 `long` 类型。


类型限定符 `signed` 与 `unsigned` 可用于限定 `char` 类型或者任何整型。

```C
unsigned char   /* 取值范围0~255 */
signed char     /* 取值范围-128~127 */
```

下面使用公式 C° = (5/9)(°F-32) 打印华氏温度与摄氏温度对照表：

```C
#include <stdio.h>

int main()
{
    float fahr, celsius;
    int lower, upper, step;

    lower = 0;     /* 温度表的下限 */
    upper = 300;   /* 温度表的上限 */
    step = 20;     /* 步长 */

    fahr = lower;
    while (fahr <= upper) {
        celsius = (5.0/9.0) * (fahr-32.0);
        printf("%3.0f %6.1f\n", fahr, celsius);
        fahr = fahr + step;
    }
}

  0  -17.8
 20   -6.7
 40    4.4
 60   15.6
 80   26.7
100   37.8
120   48.9
140   60.0
160   71.1
180   82.2
200   93.3
220  104.4
240  115.6
260  126.7
280  137.8
300  148.9
```

`printf` 中的转化说明 %3.0f 表明待打印的浮点数（即fahr）至少占3个字符宽，且
不带小数点和小数部分；%6.1f 表明另一个待打印的数（celsius）至少占6个字符宽，
且小数点后面有1位数字。

```shell
%d      按照十进制整型数打印
%6d     按照十进制整型数打印, 至少6个字符宽
%f      按照浮点数打印
%6f     按照浮点数打印, 至少6个字符宽
%.2f    按照浮点数打印, 小数点后有两位小数
%6.2f   按照浮点数打印, 至少6个字符宽, 小数点后有两位小数
```


### **1.3 for 语句**

与 `while` 语句相比， `for` 语句的操作更加直观。 `for` 语句更适合初始化
和增加步长都是单条语句并且逻辑相关的情形，因为它将循环控制语句集中放在一起，
比 `while` 语句更紧凑。

```C
#include <stdio.h>

int main()
{
    int fahr;
    for (fahr = 0; fahr <= 300; fahr = fahr + 20) {
        printf("%3d %6.1f\n", fahr, (5.0/9.0) * (fahr-32.0));
    }
}
```

在实际编程中，可以选择 `while` 与 `for` 中任意一种循环语句，主要看哪种语句更清晰。

循环的第三部分 `fahr = fahr + 20` 在函数体执行后再执行。


### **1.4 符号常量**

`#define `指令可以把符号名（或符号常量）定义为一个特定的字符串。

```shell
#define 名字 替换文本
```

替换文本可以是任何字符序列，而不仅限于数字。

```C
#include <stdio.h>

#define LOWER 0
#define UPPER 300
#define STEP  20

int main()
{
    int fahr;
    for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP) {
        printf("%3d %6.1f\n", fahr, (5.0/9.0) * (fahr-32.0));
    }
}
```

符号常量相对于 20、300 等"幻数"更加清晰明了，符号常量通常是大写字母，
便于与普通变量区分。另外注意指令行的末尾没有分号。


### **1.5 字符输入/输出**

标准库提供了一次读/写一个字符的函数，其中最简单的是 `getchar()`、`putchar()`。
每次调用时，getchar 函数从文本流中读入下一个输入字符，并将结果返回。

```C
c = getchar()
```

这个字符通常是由键盘输入的，每次调用 `putchar` 会打印一个字符，例如：

```C
getchar(c)
```

关于从文件中输入字符的方法，在第7章讨论。

#### 文件复制

借助 `getchar` 与 `putchar` 函数可以实现字符串复制的功能。

```C
#include <stdio.h>

int main()
{
    int c;
    c = getchar();
    while (c != EOF) {
        putchar(c);
        c = getchar();
    }
}
```

`char` 类型专门用于存储这种字符型数据，当然任何整型 `int` 也可以用于存储字符型数据。
这些在字符型数据在机器内部都是以位模式存储的。

`EOF` 是一个特殊字符，这个特殊值与任何实际字符都不同，表示文本结束。这里之所以不把 c 声明成 `char` 类型，
是因为他必须足够大，除了可以存储任何字符外还可以存储文件结束符 `EOF`。

`EOF` 定义在头文件 `<stdio.h>` 中，是一个整型数。

上面代码也可以改成下面这样，代码更加紧凑，里面的括号不能省。

```C
#include <stdio.h>

int main()
{
    int c;
    while ((c = getchar())!= EOF) {
        putchar(c);
    }
}
```

#### 字符计数

统计输入的字符数量。

```C
#include <stdio.h>

int main()
{
    long nc;

    nc = 0;
    while (getchar() != EOF) {
        ++nc;
     }
     printf("%ld\n", nc);
}
```

也可以用 `for` 来实现：

```C
#include <stdio.h>

int main()
{
    double nc;
    for (nc = 0; getchar() != EOF; ++nc)
        ;
    printf("%.0f\n", nc);
}
```

#### 行计数

```C
#include <stdio.h>

int main()
{
    int c, n;

    n = 0;
    while ((c = getchar()) != EOF) {
        if (c == '\n')
            ++n;
    }
    printf("%d\n", n);
}
```

单引号 `'\n''` 是一个字符常量，表示一个整型值，改值等于此字符在机器字符集中对应的数值。
例如 `'\A''` 在ASCII字符集中对应的数值为 65,使用字符常量可以表达意义更清晰。

#### 单词计数

下面这段代码是 `UNIX` 系统中 `wc` 程序的骨干部分， 包含了单词统计，行数统计，字符统计。

```C
#include <stdio.h>

#define IN 1
#define OUT 0

int main()
{
    int c, nl, nw, nc, state;

    state = OUT;
    nl = nw = nc = 0;

    while ((c = getchar()) != EOF) {
        nc++;
        if (c == '\n')
            ++nl;
        if (c == ' ' || c == '\n' || c == '\t')
            state = OUT;
        else if (state == OUT)
            state = IN;
            ++nw;
    }
    printf("%d %d %d\n", nl, nw, nc);
}
```


### **1.6 数组**

统计文本中数字出现的频次，可以使用一个数组存放各个数字出现的次数，这比使用10个独立的变量更方便。

```C
#include <stdio.h>


int main()
{
    int c, i, nwhite, nother;
    int ndigit[10];

    nwhite = nother = 0;
    for (i = 0; i < 10; ++i)
        ndigit[i] = 0;

    while ((c = getchar()) != EOF)
        if (c >= '0' && c <= '9')
             ++ndigit[c-'0'];
        else if (c == ' ' || c == '\n' || c == '\t')
             ++nwhite;
        else
            ++nother;

    printf("digits =");
    for (i = 0; i < 10; ++i)
        printf(" %d", ndigit[i]);
    printf(" white space = %d, other = %d\n", nwhite, nother);
}
```


### **1.7 函数**

编写程序实现整数的次幂，使用函数来实现。

```C
#include <stdio.h>

int power(int m, int n);

int main()
{
    int i;

    for (i = 0; i < 10; ++i)
        printf("%d %d %d\n", i, power(2,i), power(-3,i));
    return 0;
}

int power(int base, int n)
{
    int i, p;

    p = 1;
    for (i = 1; i <= n; ++i)
        p = p * base;
    return p;
}
```

通常把函数定义中圆括号列表中出现的变量称为形式参数，而把函数调用中与形式参数对应的值称为实际参数。

在 `main` 函数中末尾有一个 `return` 语句，一般来说，返回值为 0 表示正常终止，返回值为非 0 表示出现异常情况或出错结束条件。

上面程序中可以注意到，在 `main` 函数之前要声明语句：

```C
int power(int m, int n);
```

函数原型和函数声明中的参数名不要求相同，事实上函数原型的参数是可选的，这样上面的函数原型也可以写成以下形式。

```C
int power(int, int);
```


### **1.8 参数-传值调用**

在 C 语言中，所有函数参数都是"通过值"传递的，被传的参数值存在临时变量中，而不是存放在原来的变量中。

函数的参数作为临时变量存在，对其进行修改不会影响到调用函数的原始参数值，如果要修改调用函数的原始变量值，
则需要传入变量地址。

把数组名作为参数时，传递给函数的值是数组起始元素的位置或地址，并不会复制元素数组本身。


### **1.9 字符数组**

使用字符数组统计文本中最长的行，并打印出来。

```C
#include <stdio.h>
#define MAXLINE 1000

int getline1(char s[], int lim);
void copy(char to[], char from[]);

/* 打印最长输入行 */
int main()
{
    int len;
    int max;
    char line[MAXLINE];
    char longest[MAXLINE];

    max = 0;
    while ((len = getline1(line, MAXLINE)) > 0)
        if (len > max) {
            max = len;
            copy(longest, line);
        }

    if (max > 0)
        printf("%s", longest);
    return 0;
}

/* 将一行读入到s中，并返回长度 */
int getline1(char s[], int lim)
{
    int c, i;

    for (i=0; i<lim-1 && (c = getchar()) != EOF && c != EOF && c != '\n'; ++i)
        s[i] = c;
    if (c == '\n') {
        s[i] = c;
        ++i;
    }
    s[i] = '\0';
    return i;
}

/* copy 函数 */
void copy(char to[], char from[])
{
    int i;

    i =  0;
    while((to[i] = from[i]) != '\0')
        i++;
}
```

`copy` 函数的返回值类型为 `void`，显式说明该函数不需要返回值。


### **1.10 外部变量与作用域**

除了自动变量外，还可以定义位于所有函数外部的变量，外部变量只能定义一次，定义后编译程序将为它分配存储单元。
在每个需要访问外部变量的函数中，必须声明相应的外部变量，此时说明其类型。声明时可以使用 `extern` 语句显示说明，
也可以通过上下文隐式声明。

```C
#include <stdio.h>
#define MAXLINE 1000

int max;
char line[MAXLINE];
char longest[MAXLINE];

int getline1();
void copy();

/* 打印最长输入行 */
int main()
{
    int len;
    extern int max;
    extern char longest[];

    max = 0;
    while ((len = getline1()) > 0)
        if (len > max) {
            max = len;
            copy();
        }

    if (max > 0)
        printf("%s", longest);
    return 0;
}

/* 将一行读入到s中，并返回长度 */
int getline1()
{
    int c, i;
    extern char line[];

    for (i=0; i<MAXLINE-1 && (c = getchar()) != EOF && c != EOF && c != '\n'; ++i)
        line[i] = c;
    if (c == '\n') {
        line[i] = c;
        ++i;
    }
    line[i] = '\0';
    return i;
}


/* copy 函数 */
void copy()
{
    int i;
    extern char line[];
    extern char longest[];

    i =  0;
    while((longest[i] = line[i]) != '\0')
        i++;
}
```

如果外部变量的定义出现在使用它的函数之前，那么这个函数就没有必要使用 `extern` 声明，因此，
上面函数中的 `extern` 声明都是多余的。

某个变量在 file1 中定义，在 file2 和 file3 文件中使用，那么在文件 file2 和 file3 中就需要使用 `extern` 声明来建立该变量与定义之间的联系。
通常会把变量和函数的 `extern` 声明放到一个单独的文件中(头文件)，并在每个源文件的开头使用 `#include` 语句把所要用的头文件包含进来。




















## 类型、运算符与表达式


### **2.1 变量名**

对变量的命名与符号常量的命名存在一些限制条件。

命名的名字是有`字母`和`数字`组成的序列，第一个字符必须为字母。

较长的变量名使用下划线"_"分隔，提高可读性。

在传统的 C 语言中，`普通变量`使用小写字母，`符号常量`全部使用大写字母。

注：变量名要尽量从字面上表达变量的用途，这样做不容易引起混淆。局部变量一般使用较短的变量名（尤其是循环
控制变量），外边变量使用较长的名字。


### **2.2 数据类型及长度**

C 语言只提供了下列几种基本数据类型：

```C
char    字符型，占用一个字节。
int     整型，通常反映了所有机器中整数的最自然长度。
float   单精度浮点型
double  双精度浮点型
```

此外，还可以在基本数据类型前面加一些限定符，`short` 与 `long` 两个限定符用于限定整数。

```C
short int sh;
long int counter;
```

编译器根据硬件特性自主选择合适的类型长度，但要遵循下列限制：

`short` 与 `int` 类型至少为 16 位，而 `long` 类型至少为 32 位，并且 `short` 类型不得长与 `int`类型，
而 `int` 类型不得长于 `long` 类型。


类型限定符 `signed` 与 `unsigned` 可用于限定 `char` 类型或者任何整型。

```C
unsigned char   /* 取值范围0~255 */
signed char     /* 取值范围-128~127 */
```


### **2.3 常量**

一个字常量是一个整数，书写✍️时将一个字符括在单引号中，如'x'。字符在机器字符集中的数值就是字符常量的值。
例如，在 ASCII 字符集中，字符 `'0'` 的值为48，与数值 0 没有关系。

`字符常量`增加了程序的易读性，一般用来与其他字符进行比较，也可以像其他整数一样参与数值运算。

某些字符可以通过转义字符序列表示为字符和字符串常量，`'\ooo'` 与 `'\xhh'` 分别表示八进制数字和十六进制数字。

```C
#define VTAB '\013'
#define BELL '\007'

#define VTAB '\xb'
#define BELL '\x7'
```

字符常量 `'\0'` 表示值为 0 的字符，也就是空字符（null）。通常使用 `'\0'` 的形式代替 0，以强调
某些表达式的字符属性，但其数字值为0。

`常量表达式`是仅仅包含常量的表达式，这种表达式在**编译时求值**，**不再运行时求值**。

`字符串常量`也叫字符串字面值，使用双引号括起来的 0 个或多个字符组成的字符序列。

```
"I am a string" 或 ""
```

从技术角度，字符串常量就是字符数组。字符串的内部使用一个空字符串 `'\0'` 作为字符串的结尾。
存储的物理单元要比双引号中的字符数多一个。
