# 目录

[TOC]



# 0.  背景说明



本规范有别于其他编程语言（比如C/Cpp/Java/Python等）的编程风格指南。



因为 Rust 有别于其他语言。因为 Rust 编译器内建安全编程模型，会强制要求并通过编译器静态检查来训练开发者也掌握这个安全编程模型，才能顺利产出代码。并且还配备各种Cargo插件来进一步保证Rust的编码质量。并且 Rust 语言的设计特点，几乎一切信息皆显式呈现给开发者，并且天生工程能力优秀，代码的正确性、健壮性、 可扩展性和可维护性具有天然优势。



而其他编程语言并没有Rust这种强制且统一的硬性保证，所以才需要编程风格指南去规范开发者的行为，从而保证程序的安全性、正确性和易维护性。

 

但任何一门语言，都会有其编码风格，Rust也不例外。所以本指南的目的是，通过阐述一些Rust编码约定来帮助开发者加强代码的一致性，让团队成员可以方便地维护代码。并且可以让 Rust 新手快速了解 Rust Style 和一些注意事项。



本指南分为两部分：



- Safe Rust 编码规范
- Unsafe Rust 编码规范



本指南除了笔者自己的编码实践经验之外，还参考了：

 

- 《法国网络安全局(ANSSI) Rust开发指南》

- Rust官方出品的：

     a)     《Rust API Guide》

     b)    《 Unsafe Code Guidelines》。



本指南仅为建议，并非强制要求，需要看具体情况酌情考虑。并且本指南为初稿，还待进一步完善。



# 1  Safe Rust编码规范

##  1.1  开发环境

**Rust** **工具链安装：**

使用 `Rustup` （https://github.com/rust-lang/rustup）。如需替代安装方式，为了保证安全，最好选择官方推荐的替代安装方式。



**Rust Edition** **说明**

Rust从2015开始，每三年发布一个 Edition 版次：

> 1. Rust 2015 edition （Rust 0.1.0 ~ Rust 1.0.0）
>
> 2. Rust 2018 edition （Rust 1.0.0 ~ Rust 1.31.0）
>
> 3. Rust 2021 edition (Rust 1.31.0 ~ ? )

以此类推。Edition是向前兼容的。Edition 和语义化版本是正交的，不冲突。

关于 Edition 更详细的内容可以查看：https://doc.rust-lang.org/edition-guide/



**稳定版、 开发版和测试版工具链**

Rust 工具链提供三种不同的发布渠道：

> 1. Nightly（开发版），每晚发布（release）一次。
>
> 2. Beta（测试版），每六周发布一次，基于Nightly版本。
>
> 3. Stable（稳定版），每六周发布一次，基于 beta版本。

注意：

> 1. 推荐使用 Stable Rust。
>
> 2. 在基于Nightly Rust 开发项目的时候，最好通过在项目中增加 rust-toolchain 文件来指定一个固定的版本，避免因为Nightly Rust 的频繁变更而导致项目编译问题。
>
> 3. 当在稳定版工作的时候，如果需要Nightly工具链，不需要整体上去切换工具链到Nightly，只需要再命令中指明Nightly就可以了。比如 `cargo +nightly fmt`。



**包管理器 Cargo**

Cargo 是 Rust 项目必不可少的包管理器，除此之外，它也是一种工作流：

> 1. 可以用Cargo创建一个项目（bin/lib）
>
> 2. 可以用它编译项目
>
> 3. 可以用它生产项目的文档（依据文档注释）
>
> 4. 可以用它运行单元测试（test）和基准测试（bench）
>
> 5. 可以用它下载和管理crate依赖
>
> 6. 可以用它分发软件包，默认分发到 crates.io 上面
>
> 7. 可以为它编写插件，使用子命令的方式，扩展它的功能。

 

Cargo 通过 Cargo.toml 配置文件来管理 crate。

Toml 配置文件是一种最小化且无歧义的文件格式，Rust社区最常用Toml。可以通过 toml.io 进一步了解 Toml 的细节。

值得说明的是，在配置文件中如果有 [profile.*] 这种配置，需要引起注意，因为这类配置决定了编译器的调用方式，比如：

> 1. debug-assertions ，决定了是否开启debug断言。
>
> 2. overflow-checks，决定了是否检查整数运算溢出。

关于Cargo的更多细节可以查看：https://doc.rust-lang.org/cargo/index.html



## 1.2  常用Cargo插件



**Clippy**

Clippy 是一个静态分析工具，它提供了很多检查，比如错误、 样式、 性能问题、 Unsafe UB问题等等。从1.29版本开始，Clippy可以用于 Stable Rust中。

可以通过 `rustup component add clippy` 来安装此 Cargo 插件。

细节参考：https://github.com/rust-lang/rust-clippy

Clippy 的全部 lint 检查建议列表： https://rust-lang.github.io/rust-clippy/master/



**Rustfmt**

Rustfmt 是一个根据风格指南原则来格式化代码的工具。

可以通过 Rustup 来安装它： `rustup component add rustfmt`

Rustfmt 依赖的社区维护的 Rust风格指南：https://github.com/rust-dev-tools/fmt-rfcs/tree/master/guide 

开发者也可以通过 `rustfmt.toml` 或 `.rustfmt.toml` 来定制团队统一的代码风格，比如：

```toml
# Set the maximum line width to 120
max_width = 120
# Maximum line length for single line if-else expressions
single_line_if_else_max_width = 40
```



**Rustfix**

从 Rust 2018 edition开始，Rustfix就被包含在 Rust 中。它可以用来修复编译器警告。

需要注意的是，在使用 cargo fix 进行自动修复警告的时候，需要开发者确认这个警告是否真的需要修复，并且要验证修复的是否正确。



**Cargo Edit**

Cargo Edit插件为Cargo扩展了三个命令：

> 1. Cargo add，在命令行增加新的依赖，而不需要去知道这个依赖的语义版本。
>
> 2. Cargo rm，在命令行删除一个指定依赖。
>
> 3. Cargo upgrade，在命令行升级一个指定依赖。

  Cargo-edit地址：https://github.com/killercup/cargo-edit



**Cargo Audit**

Cargo Audit 可以根据 Rust安全警报数据库（RestSec Advisory Database ）的漏洞数据，扫描crate以及它的所有依赖库，然后给出一份安全报告。

更多细节：https://github.com/RustSec/cargo-audit

Rust 安全警报数据库：https://rustsec.org/



**Cargo Outdated**

该插件可以检测依赖库是否有新版本可用。

更多细节：https://github.com/kbknapp/cargo-outdated



**Cargo Deny**

该插件可以检测依赖中的软件许可证（License），如果和开发者配置的不符合，则会拒绝使用该依赖。

更多细节：https://github.com/EmbarkStudios/cargo-deny

Cargo Deny Book： https://embarkstudios.github.io/cargo-deny/



## 1.3  Rust 编码风格指南

### 1.3.1    命名

Rust基本的命名约定被定义在 RFC 430 中：https://github.com/rust-lang/rfcs/blob/master/text/0430-finalizing-naming-conventions.md

 

简单来说，Rust 中 类型级别（Type Level）的构造要使用 `UpperCamelCase` 命名风格，而值级别（`Value Level`）的构造则使用 `snake_case` 的命名风格，而静态变量和常量则使用 `SCREAMING_SNAKE_CASE`的命名风格，生命周期参数则使用 `'lowercase`。

| **Item**                | **Convention**                                               |
| ----------------------- | ------------------------------------------------------------ |
| Crates                  | [unclear](https://github.com/rust-lang/api-guidelines/issues/29) |
| Modules                 | `snake_case`                                                 |
| Types                   | `UpperCamelCase`                                             |
| Traits                  | `UpperCamelCase`                                             |
| Enum variants           | `UpperCamelCase`                                             |
| Functions               | `snake_case`                                                 |
| Methods                 | `snake_case`                                                 |
| General constructors    | `new` or `with_more_details`                                 |
| Conversion constructors | `from_some_other_type`                                       |
| Macros                  | `snake_case!`                                                |
| Local variables         | `snake_case`                                                 |
| Statics                 | `SCREAMING_SNAKE_CASE`                                       |
| Constants               | `SCREAMING_SNAKE_CASE`                                       |
| Type parameters         | concise `UpperCamelCase`, usually single uppercase letter: `T` |
| Lifetimes               | short `lowercase`, usually a single letter: '`a`, `'de`，`'src` |
| Features                | [unclear](https://github.com/rust-lang/api-guidelines/issues/101) but see [C-FEATURE](https://rust-lang.github.io/api-guidelines/naming.html#c-feature) |



**Crate** **命名**

Crate 的名字最好不要使用 “-rs”或“rust”这样的后缀，因为这没有必要，因为 crate 本来就是 Rust语言的库。



**类型转换方法命名**

要遵循下面的约定：

| **Prefix** | **Cost**  | **Ownership**                                                |
| ---------- | --------- | ------------------------------------------------------------ |
| `as_`      | Free      | borrowed  -> borrowed                                        |
| `to_`      | Expensive | borrowed  -> borrowed   borrowed -> owned (non-Copy types)   owned -> owned (Copy types) |
| `into_`    | Variable  | owned  -> owned (non-Copy types)                             |

 

`as_`前缀，表示转换没有开销成本（free）。比如 `str::as_bytes()` ，从字符串切片转换为字节切片，这个过程是没有开销的，输入是一个引用，输出也是一个引用。

`to_`前缀，表示转换会有一些运行时开销（Expensive）。比如，`str::to_string()`，转换为 String 类型会有内存分配开销。

`into_`前缀，表示转换的开销是不确定的。比如 `String::into_bytes`，将`String`转为 Vec<u8>，这个转换是没有开销的，但是它会拿走String的所有权。而 BufWriter::into_iter，则需要对缓冲数据进行潜在的flush，这个过程会有expensive的开销。

在类型的抽象级别而言，通常，`as_`和`into_`前缀的方法会降低抽象级别，比如 `str::as_bytes` 和 `String::into_bytes` ，`都是从抽象的字符串切片和字符串转换为更基础的字节表现形式`。而`to_`前缀的方法，则表示转换双方的抽象层次还在同一层面，比如 `str::to_string`，字符串切片换为动态字符串类型，抽象层面保持一致。

当类型包装单个值，使其与更高级别的语义相关联的时候，应该通过 `into_inner` 方法提供对包装值的访问。比如`BufReader`、 `GzDecoder`、 `AtomicBool`之类的包装器。

如果转换方法名字中`mut`构成返回类型的一部分，则它应该在该类型中出现相同的外观呈现。比如：`Vec::as_mut_slice`返回的是 `&mut [T]` 切片。

可以在标准库文档中查看更多的命名来验证此约定。



**Getter** **方法命名**

Rust 虽然有OOP编码风格，但在实现类型 setter/getter 中的 getter 方法的时候，没必要加 `get_`前缀，`set_`前缀是可以加的。

比如：

```rust
pub struct S {
    first: First,
    second: Second,
}

impl S {
    // Not get_first.
    pub fn first(&self) -> &First {
        &self.first
    }

    // Not get_first_mut, get_mut_first, or mut_first.
    pub fn first_mut(&mut self) -> &mut First {
        &mut self.first
    }
}

```



可以查看标准库中的 `Cell::get` 、`Cursor::get_mut`、 `<[T]>::get_unchecked` 等包含get前缀的体会命名有什么意义。



**迭代器命名：**

遵循 `iter`/`iter_mut`/`into_iter` 风格。详细参加标准库迭代器相关文档。

迭代器类型名应该与生成它们的方法名相似：

比如：

```rust
Vec::iter  ->  Iter
Vec::iter_mut -> IterMut
Vec::into_iter -> IntoIter
BTreeMap::keys -> Keys
BtreeMap::Values -> Values
```



**Features** **命名**

不要在 cargo features 的名字中增加任何无意义的单词，比如： `use-abc`或`with-abc`，请直接用`abc`。比如：

```toml
# In Cargo.toml

[features]
default = ["std"]
std = []
```

```rust
// In lib.rs

#![cfg_attr(not(feature = "std"), no_std)]
```

不要使用`use-std`，直接使用`std`就可以了。



**命名词序要一致**

比如标准库中错误类型的命名：

```rust
JoinPathsError
ParseBoolError
ParseCharError
ParseFloatError
ParseIntError
RecvTimeoutError
StripPrefixError
```

上面类型命名的单词词序为`verb-object-error`这样的词序，非常统一，而不是突然改变了词序。

### 1.3.2    内存管理

#### 1.3.2.1 内存泄露

内存泄露不属于Rust内存安全的保证范围。所以，Rust 程序中也有可能发生内存泄露。

一般是没有正确调用析构函数而导致内存泄露，在Rust中下面这些情况是正常的：

> 1. 循环引用
>
> 2. 使用forget等函数主动跳过drop
>
> 3. 析构函数内部发生了panic
>
> 4. 程序中止（abort-on-panic 开启时发生panic）



**循环引用**

最有可能发生内存泄露的场景是“循环引用”。

在Safe Rust下，循环引用的场景一般是使用`Rc`和`Arc` 引用计数容器，要注意在循环引用的地方区分强引用和弱引用的区别。弱引用不拥有所有权，合理的使用弱引用，可以避免内存泄露的问题。

**Forget Drop**

标准库中实现了 `std::mem::forget` 方法，可以让类型“遗忘”自动析构。这个方法自身是Safe的，但是它的行为可能会导致内存泄露，所以不要滥用它。`forget`方法一般是用在 FFi的时候避免因为Rust的自动析构而导致传到C-ABI对面的指针出现异常。

建议使用设置`Clippy`的lint，让代码中拒绝使用forget：

```rust
#![deny(clippy::mem_forget)]
```

需要注意的是，标准库中还有一些方法的作用类似于 forget：

> 1. Box::leak
>
> 2. Box::into_raw
>
> 3. ManuallyDrop

在Safe Rust中一定不能通过上述方法来泄露资源。



**ManuallyDrop**

使用ManuallyDrop 包装的类型，必须提供方法可以实现自动或手动drop内部的类型。比如通过 `ManuallyDrop::into_inner` 方法提取出内部的类型能被自动析构，或者，通过`ManuallyDrop::drop`方法`(unsafe)`手动析构内部的值。



**创建裸指针**

下列类型都可以通过 `into_raw` 方法来得到裸指针：

> 1. `std::boxed::Box (or alloc::boxed::Box)`
>
> 2. `std::rc::Rc (or alloc::rc::Rc)`
>
> 3. `std::rc::Weak (or alloc::rc::Weak)`
>
> 4. `std::sync::Arc (or alloc::sync::Arc)`
>
> 5. `std::sync::Weak (or alloc::sync::Weak)`
>
> 6. `std::ffi::CString`
>
> 7. `std::ffi::OsString`

同时，也可以通过调用相应的 `from_raw` 将其转换为值，并允许其回收。比如：

```rust
let boxed = Box::new(String::from("Crab"));
let raw_ptr = unsafe { Box::into_raw(boxed) };
let _ = unsafe { Box::from_raw(raw_ptr) }; // will be freed
```



#### 1.3.2.2    未初始化内存

默认情况下，Rust会强制初始化所有值，以防止使用未初始化的内存。

除非，使用 `std::mem::uninitialized`（自1.38已弃用） 或者 `std::mem::MaybeUninit`（1.36中稳定），但最好不要使用它们，原因：

> 1. Drop 掉 未初始化内存
>
> 2. 初始化内存没有Drop



#### 1.3.2.3    安全零化内存（Zeroing Memory）

零化内存是非常有用的，比如在一些加密信息，使用后需要销毁，可以在内存中零化处理，或者在FFi中也非常有用。

想要安全零化内存不容易，推荐使用`zeroize`： https://crates.io/crates/zeroize



### 1.3.3    错误处理

Rust 中的错误处理，和以往的语言，比如C++、Java、Python等包含异常的语言是非常不同的。

Rust 中的错误处理更像是C语言，基于返回值的错误处理，但是 Rust 把表示错误的返回值纳入了类型系统，从而成就了Rust的错误处理机制。

错误处理需要注意的地方包括以下几点：

1. 不要忽略错误。也就是说，不要随便使用 `unwrap` 来处理 `Result`。

2. 使用后备值。如果有些错误没必要处理，那么可以使用后备值。比如 `unwrap_or` 方法，可以指定一个后备值。

3. 面对无法处理的错误请大胆终止程序。对于一些开发者无法处理的错误，可以使用 `panic!` 来引发线程崩溃从而终止程序。建议使用 `expect` 方法来自定义`panic`时的错误信息。

4. 如果要返回单个错误，请使用 “?”操作符。该操作符会自动match Result，并且自动返回 `Error`，向上传播错误。

5. 如果要返回多个错误，建议使用 `Error` 的trait对象。也推荐使用第三方库： `anyhow` 。

6. 处理 `Error` trait对象，使用 `downcast` 来向下转换为具体的错误类型进行匹配处理。

7. 如果是编写应用，建议使用 `Error` trait对象，如果是编写库，则建议返回自定义类型，方便下游处理。

8. 如果是自定义类型，需要为其实现 `From`来自动转换为统一的错误类型，方便使用“？”来自动向上传播错误。



### 1.3.4    类型系统

#### 1.3.4.1    标准库trait



**Drop trait**

为自定义类型实现 `Drop` trait 的时候，需要遵循下面规范：

1. 合理实现`Drop`。`Drop` 通常用于释放外部资源（网络连接、文件等）或内存。

   > a)  实现的`Drop`必须不能发生 panic。
   >
   > b)  不要有循环引用
   >
   > c)  某些敏感资源处理结束的时候，不要仅仅依赖`Drop`。（比如加密结束的时候擦除密钥）

2. 要包含文档注释。

3. 要经过同行评审。



**Send** **和 Sync trait**

需要注意的地方：

1. 原生指针没有实现`Send`和`Sync`，意味着原生指针不能跨线程共享和传递。

2. UnsafeCell（包括`Cell`和`RefCell`） 不是 `Sync`，因为它提供了内部可变性。

3. `Rc` 也没有实现`Send`和`Sync`。

`Sync` 和`Send` 是自动实现的trait，当复合类型中所有元素都实现了`Sync` 和 `Send`，那么整个复合类型也会自动实现`Sync`和`Send`。通常情况下，要避免手工实现`Sync`和`Send`。



**比较相关的trait** **（`ParitialEq`，`Eq`，`PartialOrd`，`Ord`）**

在实现标准库比较trait的时候，必须，遵守标准库文档中描述的不变式（invariant）。

Rust也提供了 `#[derive(…)]` 宏为结构体自动实现这些比较trait。要注意，要实现比较trait的结构体的字段顺序，会影响自动生成的实现。



#### 1.3.4.2    Trait 一致性规则

使用 trait 的时候，必须要满足 trait 一致性规则，即，**孤儿规则（orphans rule）**：

类型和trait，必须有一个是在本地crate内定义的。



#### 1.3.4.3   使用标准库的转换trait来实现相应的类型转换

标准库提供的转换 trait：

1. `From/Into`

2. `TryFrom`

3. `AsRef`

4. `AsMut`



#### 1.3.4.4    通过类型来传递意义，而不是简单的布尔值或 Option。

布尔值或 `Option`类型，传达的意义太宽泛。最好能包装一个类型来表明具体的意义。比如：

使用 ：

```rust
let w = Widget::new(Small, Round)
```

来替代：

```rust
let w = Widget::new(true, false)
```



### 1.3.5    宏

#### 1.3.5.1    宏语法要贴近Rust语法

Rust 宏可以让开发者定义自己的DSL，但是，在使用宏的时候，要尽可能贴近Rust的语法。这样可以增强可读性，让其他开发者在使用宏的时候，可以猜测出它的生成的代码。

比如：

```rust
// Prefer this...
bitflags! {
    struct S: u32 { /* ... */ }
}

// ...over no keyword...
bitflags! {
    S: u32 { /* ... */ }
}

// ...or some ad-hoc word.
bitflags! {
    flags S: u32 { /* ... */ }
}
```

另外还需要注意语义和逗号，比如：

```rust
// Ordinary constants use semicolons.
const A: u32 = 0b000001;
const B: u32 = 0b000010;

// So prefer this...
bitflags! {
    struct S: u32 {
        const C = 0b000100;
        const D = 0b001000;
    }
}

// ...over this.
bitflags! {
    struct S: u32 {
        const E = 0b010000,
        const F = 0b100000,
    }
}
```

#### 1.3.5.2    宏语法应该支持属性

比如：

```rust
bitflags! {
    struct Flags: u8 {
        #[cfg(windows)]
        const ControlCenter = 0b001;
        #[cfg(unix)]
        const Terminal = 0b010;
    }
}
=
```

#### 1.3.5.3    宏语法应该支持可见性声明

比如：

```rust
bitflags! {
    pub struct PublicFlags: u8 {
        const C = 0b0100;
        const D = 0b1000;
    }
}
```

#### 1.3.5.4    要注意宏的卫生性

声明宏，拥有比较严格的卫生性。

过程宏，卫生性则不如声明宏严格。

 

根据适合的场景去选择相应的宏。



### 1.3.6    设计模式

Rust 项目中也会经常使用到一些经典设计模式，比如**工厂模式、策略模式、迭代器模式、构建模式、适配器模式**等等。除此之外，Rust也有属于它自己的模式。



#### 1.3.6.1    NewType 模式

使用元组结构体对某个类型进行包装，就是NewType模式。

NewType模式的好处是，可以赋予多个相同类型不同的意义，并且可以静态检查。

比如：

```rust
struct Miles(pub f64);
struct Kilometers(pub f64);

impl Miles {
    fn to_kilometers(self) -> Kilometers { /* ... */ }
}
impl Kilometers {
    fn to_miles(self) -> Miles { /* ... */ }
}
```



#### 1.3.6.2    Deref 模式

```rust
struct Vec<T> {
    ...
}

impl<T> Deref for Vec<T> {
    type Target = [T];

    fn deref(&self) -> &[T] {
        ...
    }
}
```



#### 1.3.6.3    RAII 模式

Rust 借鉴了Cpp 的RAII来实现智能指针的自动化内存管理。RAII 同样也可以被扩展为一种模式来自动化管理其他资源。

比如：

```rust
struct Mutex<T> {
    // We keep a reference to our data: T here.
    ...
}

struct MutexGuard<'a, T: 'a> {
    data: &'a T,
    ...
}

// Locking the mutex is explicit.
impl<T> Mutex<T> {
    fn lock(&self) -> MutexGuard<T> {
        // Lock the underlying OS mutex.
        ...

        // MutexGuard keeps a reference to self
        MutexGuard { data: self, ... }
    }
}

// Destructor for unlocking the mutex.
impl<'a, T> Drop for MutexGuard<'a, T> {
    fn drop(&mut self) {
        // Unlock the underlying OS mutex.
        ...
    }
}

// Implementing Deref means we can treat MutexGuard like a pointer to T.
impl<'a, T> Deref for MutexGuard<'a, T> {
    type Target = T;

    fn deref(&self) -> &T {
        self.data
    }
}

fn main(x: Mutex<Foo>) {
    let xx = x.lock();
    xx.foo(); // foo is a method on Foo.
    // The borrow checker ensures we can't store a reference to the underlying
    // Foo which will outlive the guard xx.

    // x is unlocked when we exit this function and xx's destructor is executed.
}
```

Mutex实现了MutexGuard，利用Drop自动调用的时候，实现自动解锁。



## 1.4  代码优化指北

关于性能优化，从下面三点来描述：

> 1. 原则。对于性能优化，首先需要有一定的原则，按原则行事。
>
> 2. 技巧。 掌握一些性能优化的技巧。
>
> 3. 利用一些工具去发现性能瓶颈。



### 1.4.1    原则

- 不要盲目优化，最优先应该优化的是代码的可读性。
- 不要费心去优化一次性成本。 如果只是一次性成本，就没必要优化它。
- 真的需要优化的时候，要优先改进算法。



### 1.4.2    技巧

- 避免使用 trait 对象。 当然这一条也不是绝对的，trait对象虽然有开销，但性能也能满足大多数情况。除非对性能有极致的要求，可以使用 Enum 或 泛型 静态分发来替代。必须使用 trait 对象的场景除外。
- 使用基于栈的动态数据类型，比如smallvec，它可以在一定长度的情况下存在栈上，长度大于一个临界值就存到堆上。也有一些避免堆内存频繁分配的数据结构。
- 合理使用断言避免数组越界检查。
- 减少复制。尽可能降低代码中的复制开销。
- Rust 是可以自定义内存分配器的，所以可以自由替换符合你场景的内存分配器。
- 避免嵌套使用堆分配的容器，比如 嵌套使用动态数组 `Vec<Vec<_>>`
- 另外需要掌握一些编译器配置。比如 要使用 release 模式编译
- 使用 LTO
- 设置 codegen-units=1 ，codegen-units 叫做代码生成单元，Rust 编译器会把crate 生成的 LLVMIR进行分割，默认分割为16个单元，每个单元就叫 codegen-units，如果分割的太多，就不利于 Rust编译器使用内联优化一些函数调用，分割单元越大，才越容易判断需要内联的地方。但是这也有可能增大编译文件大小，需要大小和性能间寻找平衡。
- 设置panic=abort。可以缩减编译文件的大小。



### 1.4.3    工具

- clippy，前面介绍过了，可以通过静态分析帮你识别代码里的坏味道。
- perf 、 flamegraph、 Vtune 等工具，都是性能测试常用的一些工具
- Criterion，是 第三方库，可以为 rust项目编写基准测试。并且生成漂亮的报告，而且还可以检测项目的性能是否回退还是提升



# 2  Unsafe Rust 编码规范

Unsafe Rust 是Safe Rust的超集。 在Unsafe Rust 中照样会进行安全检查，只有下列五种违反静态语义内存安全保证的操作会跳过安全检查：

> 1. 解引用原生指针。
>
> 2. 读写可变的（mutable）或外部的（external）静态变量。
>
> 3. 访问 union 的字段，不论读或者写。
>
> 4. 调用 unsafe 的函数（包括内部函数和外部函数）
>
> 5. 实现 unsafe trait。

使用Unsafe Rust 的时候，要注意避免产生 UB（Undefined Behavior）未定义行为。



## 2.1  安全抽象

首先，在概念上应该正确认识 Unsafe 。 Unsafe Rust 是 Safe Rust 的超集。在Unsafe Rust 中也有编译器检查，只不过是在几种特殊的情况，编译器才无法检查，比如解引用原生指针，使用Union类型等。

 

其次，要理解Unsafe Rust 的职责范围。 Unsafe Rust 的职责是帮助开发者标识 Unsafe 边界， 如何划分还是开发者的责任。

 

开发者划分Unsafe 边界，就必须了解接口的安全边界，如果不了解，是无法划分的。

 

Rust 官方的 Unsafe 实践规范如下：

 

> 1. 调用unsafe 代码必须加上unsafe 块
>
> 2. 调用unsafe代码的函数，需要进行安全抽象。安全抽象的意思就是通过一些验证条件来消除unsafe 发生的情况。
>
> 3. 安全抽象的代码必须增加注释说明为什么安全，以及安全边界在哪。
>
> 4. 如果没有安全抽象，那么整个函数或trait 都必须被标志为unsafe。不能省略unsafe 关键字，并且要增加 注释 说明什么情况下有可能会发生 UB。
>
> 5. 了解 Unsafe Rust 实践中常犯的错误。之前有世界顶级学术期刊的一篇论文《理解真实Rust程序中的内存和线程安全实践》，通过调研五个开源项目，五个广泛使用的库，以及标准库中850个unsafe代码的用法，分析了170个bug，最终得出一些结论：
>
>    ​		a)     证明了我们上面所说的 unsafe 最佳实践的做法确实是最佳实践。平时编写 unsafe 代码最好采用这个最佳实践。
>
>    ​		b)    不管是内存问题还是线程安全问题，总的来说，都是因为开发者对生命周期的判断失误导致的。所以理解 Rust 中生命周期的概念很重要。

参考：https://cloud.tencent.com/developer/article/1655438



安全抽象的时候，需要注意 PhatomData 的使用：

![img](rust1.jpg)

这段代码里 `MappedMutexGuard<T, U>` 结构体中使用了`PhantomData`，来保证类型U是被`MappedMutexGuard`拥有的，否则，就会出现安全问题。

`PhantomeData`在Unsafe Rust中通常被用来给类型指派所有权，以及型变规则。

详细参考：https://doc.rust-lang.org/nomicon/

### 2.1.1 指针别名规则

不要随意更改指针的可变性，保持指针别名规则 (pointer aliasing rule)。在 Rust reference 14.3 节，对于未定义行为的描述，有这样两条规则：

- Breaking the [pointer aliasing rules](http://llvm.org/docs/LangRef.html#pointer-aliasing-rules). `&mut T` and `&T` follow LLVM’s scoped [noalias](http://llvm.org/docs/LangRef.html#noalias) model, except if the `&T` contains an [`UnsafeCell`](https://doc.rust-lang.org/std/cell/struct.UnsafeCell.html).
- Mutating immutable data. All data inside a [`const`](https://doc.rust-lang.org/reference/items/constant-items.html) item is immutable. Moreover, all data reached through a shared reference or data owned by an immutable binding is immutable, unless that data is contained within an [`UnsafeCell`](https://doc.rust-lang.org/std/cell/struct.UnsafeCell.html).

因此，不要随意把 `*const T` 转换为 `*mut T` 然后用这个指针对目标地址做修改，这是未定义行为。

如果确实有这样的需求，请把数据包裹到 `UnsafeCell` 里面，然后通过它的 `get()` 函数获取。

反面示例：

```rust
fn main() {
    let x = 10;
    unsafe {
        let p = &x as *const i32 as *mut i32; // Do NOT do this!
        *p = 20;
    }
    println!("{}", x);
}
```


### 2.1.2 内存对齐规则

注意对象的内存对齐alignment。内存不对齐的对象，不要把指针 `*const T` 转换为 `&T` 类型。

参考 [Rust reference type-layout](https://doc.rust-lang.org/reference/type-layout.html)

   > ***Warning:*** Dereferencing an unaligned pointer is [undefined behavior](https://doc.rust-lang.org/reference/behavior-considered-undefined.html) and it is possible to [safely create unaligned pointers to `packed` fields](https://github.com/rust-lang/rust/issues/27060). Like all ways to create undefined behavior in safe Rust, this is a bug.

### 2.1.3 避免数组越界

在 safe rust 里面，数组越界访问是不可能的。但是在 unsafe rust 里面，当你需要用裸指针管理一块连续内存的时候，就有可能发生越界行为。在这种情况下用户一定需要注意避免这种风险。反面示例如下，假设成员 p 是指向数组的指针，而 len 代表的是数组长度：

```rust
#[repr(C)]
pub struct User {
    pub id: i32,
    pub p: *const i32,
    pub len: usize,  // 危险!
}

impl User {
    pub fn method(&self) {
        for i in 0..len {
            println!("{}", unsafe { *self.p.offset(i) } );
        }
    }
}
```

这个例子的问题是，把指针和长度都作为 pub 成员给公开到外部了。这意味着任何用户都可以随意修改长度，那么调用 method 这个方法的时候就很可能内存越界了。数组长度应该作为一个私有成员保护起来，避免越界。如果要把修改 len 这个功能暴露成公开方法，那么它必须被 unsafe 修饰：

```rust
#[repr(C)]
pub struct User {
    pub id: i32,
    p: *const i32,
    len: usize,
}

impl User {
    pub fn method(&self) {
        for i in 0..len {
            println!("{}", unsafe { *self.p.offset(i) } );
        }
    }

    // 因为这个函数使用不当可能会造成安全问题，所以它需要用 unsafe 标记
    pub unsafe fn set_len(&mut self, l: usize) {
        self.len = l;
    }
}
```



## 2.2     整数溢出

 尽管Rust对可能的整数溢出进行了一些验证，但对整数执行算术运算时应采取预防措施。

特别要注意的是，使用debug或release编译配置文件会更改整数溢出行为。 在debug配置中，溢出导致程序终止（紧急），而在release配置中，结果为二进制补码的截断。

通过使用`Wrapping`泛型类型或对整数（`<op>`部分为`add`，`mul`，`sub`，`shr`等）的`overflowing_<op>`和`wrapping_<op>`操作，可以使此最后一个行为明确。

```rust
use std::num::Wrapping;

let x: u8 = 242;

println!("{}", x + 50);                      // panics in debug, prints 36 in release.
println!("{}", x.overflowing_add(50).0);     // always prints 36.
println!("{}", x.wrapping_add(50));          // always prints 36.
println!("{}", Wrapping(x) + Wrapping(50));  // always prints 36.

// always panics:
let (res, c) = x.overflowing_add(50);
if c { panic!("custom error"); }
else { println!("{}", res); }
```

## 2.3  FFi

Rust 可以通过C-ABI无缝与C语言打交道，但是两种边界本质上是不安全的。所以在使用 FFi 的时候，需要注意一些相关事项。



### 2.3.1    数据布局

为了保证一致的数据布局，Rust提供了 `#[repr(C)]` 属性来兼容C的内存布局。

```rust
#[repr(C)]
struct Data {
    a: u32,
    b: u16,
    c: u64,
}
#[repr(C, packed)]
struct PackedData {
    a: u32,
    b: u16,
    c: u64,
}
```

但注意有一些类型是不保证内存布局的，所以不合适用于和 C 语言交互：

> 1. 没有使用任何 `#[repr( )]` 属性修饰的自定义类型
> 
> 2. 动态大小类型 (dynamic sized type)
>
> 3. 指向动态大小类型对象的指针或引用 (fat pointers)
>
> 4. str 类型、tuple 元组、闭包类型


类型必须在 FFi 边界两端保持一致。

建议使用自动生成绑定的工具，如`Rust-bindgen`或`cbindgen`，可能有助于使C和Rust之间的类型保持一致。

> 注意：rust-bindgen的某些选项可能会导致危险的转换，尤其是rustified_enum



### 2.3.2    依赖平台的类型

当与外部(如C或c++)接口交互时，通常需要使用平台相关的类型，如C的`int`、`long`等。除了`std::ffi`(或`core::ffi`)中的c void外，标准库还在`std:os::raw`(或`core::os::raw`)中提供了可移植类型别名。

Libc crate 基本覆盖了所有的C标准库中的C兼容类型。

> 建议：Rust代码必须使用诸如标准库或libc crate所提供的可移植类型别名，而不是特定于平台的 类型。


### 2.3.3    字符串类型

1. 使用指向 `c_char` 类型的指针作为 FFI 接口。不要使用Rust的 `i8` `u8` 或者 `char` ，它们都不能对应C语言char类型。注意，即便是 `std::ffi` 模块中的 `CStr` `CString` 类型，也不可以直接作为 FFI 函数的接口。
2. FFI 接口中使用的字符串要符合 C 语言的约定，即使用 `\0` 结尾，且中间不要包含 `\0` 字符。
3. 注意Rust字符串和C字符串的编码区别。Rust 的 `str` `String` 类型要求字符串编码为 utf-8 编码，Rust 的 `char` 类型是 32bit。C 侧的字符串理论上是不限编码的。建议遵循 [utf8 everywhere](http://utf8everywhere.org/) 的准则，整个项目全部使用 utf8 编码，避免不必要的麻烦。

### 2.3.4    非健壮类型：引用、 函数指针和Enum

特定类型的*trap representation*是一种表示(位模式)，它尊重类型的表示约束(如大小和对齐)，但不表示该类型的有效值，并导致未定义的行为。简单地说，如果Rust变量被设置为这样一个无效的值，从简单的程序崩溃到任意的代码执行，任何事情都可能发生。当编写safe Rust时，这不会发生(除非通过Rust编译器中的错误)。然而，当编写不安全Rust，特别是在FFI中，这是非常容易的。

很多 Rust 类型都不健壮：

> 1. bool
>
> 2. 引用类型
>
> 3. 函数指针
>
> 4. Enum
>
> 5. 浮点数
>
> 6. 包含了上述类型的复合类型

所以 ，不要在Rust 中使用未经检查的上述不健壮类型的外部的值。要尽可能的对外部值进行检查。



**指针**：

尽管Rust编译器允许使用引用和指针，但在FFI中使用Rust的引用可能会破坏Rust的内存安全性。 由于它们的“不安全性”更为明确，因此在绑定到另一种语言时，与Rust引用相比，首选使用指针。

 

一方面，引用类型非常不健壮：它们仅允许指向有效内存对象的指针。 任何偏差都会导致不确定的行为。

与C绑定时，问题特别严重，因为C没有引用（在有效指针的意义上），并且编译器不提供任何安全保证。

 

当与 C++ 绑定时，Rust引用实际上可以绑定到 C++ 引用，即使带有引用的 C++ 中外部“ C”函数的实际ABI是“实现定义的”。 另外，应检查 C++ 代码是否与指针/引用混淆。

 

Rust引用可以合理地与其他C兼容语言一起使用，包括允许非空类型检查的C变体，例如 Microsoft SAL批注的代码。

 

另一方面，Rust的指针类型也可能导致未定义的行为，但更可验证，主要针对`std / core :: ptr :: null（）`（C的`（void *）0`），但在某些情况下还针对已知的有效内存 范围（特别是在嵌入式系统或内核级编程中）。 在FFI中使用Rust指针的另一个优点是，在不安全的块或函数内部清楚地标记了指向值的所有负载。

 

在Rust中，任何对外部指针解引用的操作，都必须验证其有效性，并且要检查指针是否非空。



**函数指针**：

跨越FFI边界的函数指针可能最终导致任意代码执行。

在安全Rust开发中，FFI边界上的任何函数指针类型都必须标记为extern(可能使用特定的ABI)和不安全的。

与普通指针相比，Rust中的函数指针与引用非常相似。 特别是，不能直接在Rust一侧检查函数指针的有效性。 但是，Rust提供了两种替代方法：

> 1. 使用Option包装的函数指针并检查是否为null
>
> 2. 转换为原始指针。

任何外部函数都必须在FFi边界上进行检查。



**Enum**

在FFi开发中，Rust代码不应该接受任何 Rust Enum类型的传入值。

### 2.3.5    不透明类型（Opaque Types）

当绑定外部不透明类型时，应该使用指向专用不透明类型的指针，而不是使用c void指针。

最好的创建不透明的类型的方法是这样的：
```rust
extern {
    type Foo;
}
extern "C" {
    fn foo(arg: *mut Foo);
}
```

这个功能在 RFC 1861 extern type 中描述。目前这个功能要求 `#![feature(extern_types)]`，如果这个功能还没有稳定，替代方案为：

```rust
#[repr(C)]
pub struct Foo {_private: [u8; 0]}
extern "C" {
    fn foo(arg: *mut Foo);
}
```

当与C或C ++接口时，应将在C / C ++中被视为不透明的Rust类型转换为不完整的结构类型（即声明为未定义），并提供专用的构造函数和 析构函数。

比如：

```rust
struct Opaque {
    // (...) details to be hidden
}

#[no_mangle]
pub unsafe extern "C" fn new_opaque() -> *mut Opaque {
    catch_unwind(|| // Catch panics, see below
        Box::into_raw(Box::new(Opaque {
            // (...) actual construction
        }))
    ).unwrap_or(std::ptr::null_mut())
}

#[no_mangle]
pub unsafe extern "C" fn destroy_opaque(o: *mut Opaque) {
    catch_unwind(||
        if !o.is_null() {
            drop(Box::from_raw(o))
        }
    ); // Only needed if Opaque or one of its subfield is Drop
}
```



### 2.3.6    内存和资源管理

编程语言以不同的方式处理内存。因此，当Rust和另一种语言之间传输数据时，知道哪一种语言负责回收该数据的内存空间是很重要的。对于其他类型的资源，如套接字或文件，也是如此。

Rust可利用变量的所有权和生存期，以在编译时确定是否以及何时应该释放内存。 由于Drop特性，开发者可以利用此系统来回收其他类型的资源，例如文件或网络访问。 将一些数据从Rust迁移到外语还意味着放弃与之相关的可能的回收。
 
在安全Rust开发中，任何用外部分配和释放的非敏感的外来数据块都应该封装在Drop类型中，这样就可以在Rust中通过自动调用外部语言释放例程来提供自动释放。


当某种类型的数据通过FFI边界非Copy传递时，必须确保：

1. 单一语言负责数据的分配和释放。谁分配，谁释放。

2. 如果一个对象的指针被传递给了 C 侧，需要在 Rust 侧确保此对象不会被移动。反面事例：

```rust
#[repr(C)]
struct Isolate {
    id: i32
}

impl Isolate {
    fn as_raw_ptr(&self) -> *const Isolate { self as *const Isolate }
    ...
}
```

类似这种API，就很容易误用。想象一下这个场景：

* 创建了一个 Isolate 实例
* 获取了 raw ptr，并传递给了 C 侧，C 侧将这个指针的值保存下来了
* 移动了这个 Isolate 实例
* C 侧依然在使用开始时候保存的那个指针的值解引用

解决方案：在Rust侧的API强制把对象分配到堆上，然后使用这个地址传递给C侧。或者使用新的 `Pin` 类型做 API，防止实例被移动。

3. 防止内存泄漏。严格来说，内存泄漏在Rust中不算内存安全问题。但是依然是一个严重的bug，我们应该尽量避免。
一般情况下，大家记着分配和释放要对应就够了。但是极端场景下，还是可能会发生不易察觉的内存泄漏，比如下例：

```rust
fn test(x: Box<i32>) {
    let raw_ptr = unsafe { Box::into_raw(boxed) };
    // ...
    do_smth();
    // 如果上面的函数调用发生了 panic，那么就会发生内存泄漏
    let _ = unsafe { Box::from_raw(raw_ptr) }; // will be freed
}
```

如果要保证即便在发生 panic 的情况下，释放内存依然可以正常执行，那么就需要使用 RAII 机制，把释放操作封装到一个类型的 `drop` 函数中。

4. 需要特别注意生命周期的正确性。因为此时对象的生命周期是手工管理的，必须由用户自己保证指针指向的内存没有被提前释放。

5. 避免拷贝和移动自引用类型。所谓自引用类型，就是如果一个类型中存在指针和引用，指向同一个实例的其它成员。

假设C里面有一个结构体：

```C
struct Snake {
	int weight;
    int* p;
}
```

而且这个成员 p 会指向同一个对象的 weight 成员。那么这种类型就叫自引用类型。这样的类型不要直接用于 FFI 的接口。请保证这样的类型在堆上分配，在 FFI 接口中使用指针。

### 2.3.7    线程安全

1. 如果C函数内部使用了未同步的全局变量，那么Rust方的调用者需要注意不要将这个函数用于多线程。这样的 extern 函数最好是用更安全的 API 包装起来，不要直接作为模块对外 API。

反面示例：

```C
int new_count() {
   static int count = 0;

   count++;
   return count;
}
```

对于这样的函数，在Rust那边封装的时候，不要直接暴露成模块对外的安全 API。因为它不是线程安全的。应该想办法保证这个函数只能在一个线程中使用。

2. FFI 对象类型，要特别注意它是不是 Sync / Send 的。定义FFI类型的时候要显式 impl Send / Sync，避免编译器误判。

反面示例：https://github.com/rust-lang/rust/issues/41622

`MutexGuard` 类型不是 Send 安全的，可是由于 auto trait 机制，编译器默认把它推理成了 Send 的，这就会造成线程安全问题。

### 2.3.8    panic 安全

参考 https://github.com/rust-lang/rfcs/blob/master/text/2945-c-unwind-abi.md 这份提案。在 `#![feature(c_unwind)]` 功能稳定以前，建议用户遵循以下规则：

1. 不允许Rust的panic被扩散到C函数中。如果一个Rust函数需要被C调用且可能发生panic，那么最好使用`catch_unwind` 把 panic 处理掉。或者使用 `-C panic=abort` 选项编译Rust代码。


2. 不要把C++ exception扩散到Rust函数中。


3. 在当前版本的Rust实现中，使用longjmp跨越Rust函数调用帧是安全的，但是此事还没有一个完整的RFC来规定下来。建议暂时尽量不要依赖 longjmp。



## 2.4  Unsafe 相关检测工具

### 2.4.1    Miri

Miri是 Rust编译器内部提供的一个工具，它可以检测 Unsafe Rust中的UB行为，但目前此功能是实验性的，准确率不是很高。

详细查看：https://github.com/rust-lang/miri

目前Rust标准库已经使用 Miri 发现了多个UB问题：

- Debug for `vec_deque::Iter` accessing uninitialized     memory
- `Vec::into_iter` doing an unaligned ZST     read
- `From<&[T]>` for `Rc` creating a not sufficiently aligned reference
- `BTreeMap` creating a shared reference pointing to a too small allocation
- `Vec::append` creating a dangling   reference
- Futures turning a shared  reference into a mutable one
- `str` turning a shared reference into a mutable one
- `rand` performing unaligned reads

在 `Readme` 文档里还有更多。

### 2.4.2    Cargo Geiger

一个第三方插件，可以统计当前crate以及其依赖中Unsafe代码信息。
https://github.com/rust-secure-code/cargo-geiger

# 3  参考资源列表

- https://anssi-fr.github.io/rust-guide/01_introduction.html
- https://rust-lang.github.io/api-guidelines/about.html
- https://rust-lang.github.io/unsafe-code-guidelines/introduction.html
- http://cliffle.com/p/dangerust/
- https://github.com/rust-dev-tools/fmt-rfcs/tree/master/guide
- https://github.com/rust-unofficial/patterns