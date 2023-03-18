fn main() {
    let pointer = 0x00007000 as *mut String;
    unsafe { pointer.as_mut().unwrap().insert_str(10, "Henlo"); }
}
