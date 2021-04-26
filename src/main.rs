mod hello;

fn main() {
    let msg = hello::get_message();
    println!("{}", msg);
}
