use nix::unistd::daemon;
use std::process;
use std::{thread, time};

fn main() {
    if let Ok(_) = daemon(false, false) {
        println!(" my pid {}", process::id());
        let mut count = 0u32;
        loop {
            count += 1;
            print!("{} ", count);
            if count == 300 {
                println!("OK, that's enough");
                // Exit this loop
                break;
            }
            thread::sleep(time::Duration::new(1, 0));
        }
    }
}
