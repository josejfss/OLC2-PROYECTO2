fn main() {
    let a: i64 = 25 - 25;
    let b: i64 = 1 + 1 / 2 * 10 / a;
    println!("{}", b);

    let x: i64 = 10 + 2 - a;
    let arr: [i64; 2] = [1,2];
    println!("{}", arr[x * 2 + 10 - a * 20]);
}
