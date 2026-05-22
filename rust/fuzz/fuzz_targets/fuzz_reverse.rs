#![no_main]

use libfuzzer_sys::fuzz_target;
use scorecard_example::reverse;

fuzz_target!(|data: &str| {
    let reversed = reverse(data);
    let double_reversed = reverse(&reversed);
    assert_eq!(data, double_reversed);
});
