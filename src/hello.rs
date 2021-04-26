pub fn get_message() -> String {
    return "Hello,World!".to_string();
}

#[test]
fn test_get_message(){
    let want = "Hello,World!";
    let got = get_message();
    assert_eq!(want, got);
}