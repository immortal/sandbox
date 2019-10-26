use libc;
use std::io;

fn main() {
    match unsafe { libc::close(0) } {
        -1 => panic!("could not close"),
        _ => {
            let mut input = String::new();
            io::stdin()
                .read_line(&mut input)
                .expect("failed to read line");
            println!("stdin: {}", input.trim());
        }
    }
}
