// ========== Rust 常用属性（Attributes）大全 ==========

// ========== 1. Lint 控制属性 ==========

// #[allow(...)] - 允许特定的 lint 警告
#[allow(dead_code)]
enum UnusedEnum {
    Variant1,
}

// #[warn(...)] - 提升警告级别
#[warn(dead_code)]
fn warn_function() {}

// #[deny(...)] - 将警告变为错误
// #[deny(dead_code)]
// fn deny_function() {}  // 未使用会编译失败

// #[forbid(...)] - 最严格，不能后续覆盖
// #[forbid(dead_code)]

// ========== 2. derive 属性 - 自动实现 trait ==========

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
enum Status {
    Active,
    Inactive,
}

// ========== 3. cfg 条件编译属性 ==========

#[cfg(target_os = "linux")]
fn linux_function() {
    println!("只在 Linux 上编译");
}

#[cfg(target_os = "windows")]
fn windows_function() {
    println!("只在 Windows 上编译");
}

#[cfg(debug_assertions)]
fn debug_only() {
    println!("只在 debug 模式下编译");
}

#[cfg(test)]
mod tests {
    // 测试模块
}

// ========== 4. test 测试属性 ==========

#[test]
fn test_example() {
    assert_eq!(2 + 2, 4);
}

#[test]
#[should_panic]
fn test_panic() {
    panic!("这个测试应该 panic");
}

#[test]
#[ignore]
fn ignored_test() {
    // 这个测试会被忽略
}

// ========== 5. inline 内联优化属性 ==========

#[inline]
fn small_function(x: i32) -> i32 {
    x * 2
}

#[inline(always)]
fn always_inline() {
    // 总是内联
}

#[inline(never)]
fn never_inline() {
    // 从不内联
}

// ========== 6. must_use 属性 ==========

#[must_use]
fn important_result() -> i32 {
    42
}

#[must_use = "这个结果很重要，不要忽略"]
fn critical_result() -> i32 {
    100
}

// ========== 7. deprecated 废弃属性 ==========

#[deprecated]
fn old_function() {
    println!("这个函数已废弃");
}

#[deprecated(note = "请使用 new_function() 代替")]
fn deprecated_with_note() {}

// ========== 8. doc 文档属性 ==========

/// 这是文档注释（使用 ///）
/// 
/// # 示例
/// ```
/// let x = example_function();
/// ```
#[doc = "这是文档属性"]
fn example_function() {}

// ========== 9. repr 表示属性（内存布局）==========

#[repr(C)]
struct CStruct {
    x: i32,
    y: i32,
}

#[repr(u8)]
enum Number {
    Zero,
    One,
    Two,
}

#[repr(transparent)]
struct Wrapper(i32);

// ========== 10. no_std 属性 ==========

// #![no_std]  // 在 crate 根使用，禁用标准库

// ========== 11. panic 处理属性 ==========

// #[panic_handler]  // 只在 no_std 环境中使用，标准库已有实现
// fn panic_handler(_info: &std::panic::PanicInfo) -> ! {
//     loop {}
// }

// ========== 12. 其他常用属性 ==========

#[cold]
fn unlikely_path() {
    // 标记为冷路径（不常执行）
}

#[track_caller]
fn track_caller_example() {
    // 跟踪调用者位置
}

#[unsafe(no_mangle)]
pub extern "C" fn exported_function() {
    // 不进行名称修饰，用于 FFI（需要 unsafe 包装）
}

// ========== 13. 模块级属性 ==========

// #![allow(dead_code)]  // 注意：! 表示应用到整个模块/crate（只能在文件顶部使用）

// ========== 14. 宏相关属性 ==========

#[macro_export]
macro_rules! my_macro {
    () => {
        println!("宏定义");
    };
}

#[macro_use]
mod macros {
    // 导出宏
}

// ========== 主函数 ==========

fn main() {
    println!("=== Rust 常用属性示例 ===\n");
    
    // 使用 derive 的结构体
    let p1 = Point { x: 10, y: 20 };
    let p2 = p1.clone();
    println!("点: {:?}", p1);
    println!("点相等: {}\n", p1 == p2);
    
    // must_use 示例
    let _result = important_result();  // 如果不用会有警告
    let result = critical_result();
    println!("重要结果: {}\n", result);
    
    // 条件编译
    #[cfg(target_os = "linux")]
    {
        println!("当前是 Linux 系统");
    }
    
    #[cfg(not(target_os = "linux"))]
    {
        println!("当前不是 Linux 系统");
    }
    
    println!("\n=== 属性分类总结 ===");
    println!("1. Lint 控制: allow, warn, deny, forbid");
    println!("2. 自动实现: derive");
    println!("3. 条件编译: cfg");
    println!("4. 测试: test, should_panic, ignore");
    println!("5. 优化: inline, cold");
    println!("6. 文档: doc");
    println!("7. 废弃: deprecated");
    println!("8. 内存布局: repr");
    println!("9. 其他: must_use, no_mangle, track_caller");
}
