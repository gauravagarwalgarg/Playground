use std::collections::HashMap;
use std::collections::HashSet;

/// LeetCode 1: Two Sum
/// Given an array of integers nums and an integer target,
/// return indices of the two numbers such that they add up to target.
/// Time: O(n), Space: O(n)
fn two_sum(nums: &[i32], target: i32) -> Vec<usize> {
    let mut seen: HashMap<i32, usize> = HashMap::new();
    for (i, &num) in nums.iter().enumerate() {
        let complement = target - num;
        if let Some(&j) = seen.get(&complement) {
            return vec![j, i];
        }
        seen.insert(num, i);
    }
    vec![]
}

/// LeetCode 217: Contains Duplicate
/// Given an integer array nums, return true if any value appears at least twice.
/// Time: O(n), Space: O(n)
fn contains_duplicate(nums: &[i32]) -> bool {
    let mut seen: HashSet<i32> = HashSet::new();
    for &num in nums {
        if !seen.insert(num) {
            return true;
        }
    }
    false
}

/// LeetCode 121: Best Time to Buy and Sell Stock
/// Find the maximum profit from one buy-sell transaction.
/// Time: O(n), Space: O(1)
fn max_profit(prices: &[i32]) -> i32 {
    let mut min_price = i32::MAX;
    let mut profit = 0;

    for &price in prices {
        min_price = min_price.min(price);
        profit = profit.max(price - min_price);
    }
    profit
}

/// LeetCode 70: Climbing Stairs
/// Time: O(n), Space: O(1)
fn climbing_stairs(n: u32) -> u64 {
    if n <= 2 {
        return n as u64;
    }
    let mut prev: u64 = 1;
    let mut curr: u64 = 2;
    for _ in 3..=n {
        let next = prev + curr;
        prev = curr;
        curr = next;
    }
    curr
}

/// LeetCode 11: Container With Most Water
/// Time: O(n), Space: O(1)
fn max_area(height: &[i32]) -> i32 {
    let mut left = 0;
    let mut right = height.len() - 1;
    let mut max_water = 0;

    while left < right {
        let width = (right - left) as i32;
        let h = height[left].min(height[right]);
        max_water = max_water.max(width * h);

        if height[left] < height[right] {
            left += 1;
        } else {
            right -= 1;
        }
    }
    max_water
}

fn main() {
    println!("Running Playground Rust...");
    println!("two_sum([2,7,11,15], 9) = {:?}", two_sum(&[2, 7, 11, 15], 9));
    println!("contains_duplicate([1,2,3,1]) = {}", contains_duplicate(&[1, 2, 3, 1]));
    println!("max_profit([7,1,5,3,6,4]) = {}", max_profit(&[7, 1, 5, 3, 6, 4]));
    println!("climbing_stairs(10) = {}", climbing_stairs(10));
    println!("max_area([1,8,6,2,5,4,8,3,7]) = {}", max_area(&[1, 8, 6, 2, 5, 4, 8, 3, 7]));
    println!("\nRun `cargo test` to execute all tests.");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_two_sum() {
        assert_eq!(two_sum(&[2, 7, 11, 15], 9), vec![0, 1]);
        assert_eq!(two_sum(&[3, 2, 4], 6), vec![1, 2]);
        assert_eq!(two_sum(&[3, 3], 6), vec![0, 1]);
        assert_eq!(two_sum(&[1], 2), Vec::<usize>::new());
    }

    #[test]
    fn test_contains_duplicate() {
        assert_eq!(contains_duplicate(&[1, 2, 3, 1]), true);
        assert_eq!(contains_duplicate(&[1, 2, 3, 4]), false);
        assert_eq!(contains_duplicate(&[]), false);
    }

    #[test]
    fn test_max_profit() {
        assert_eq!(max_profit(&[7, 1, 5, 3, 6, 4]), 5);
        assert_eq!(max_profit(&[7, 6, 4, 3, 1]), 0);
        assert_eq!(max_profit(&[1, 2]), 1);
    }

    #[test]
    fn test_climbing_stairs() {
        assert_eq!(climbing_stairs(1), 1);
        assert_eq!(climbing_stairs(2), 2);
        assert_eq!(climbing_stairs(3), 3);
        assert_eq!(climbing_stairs(5), 8);
        assert_eq!(climbing_stairs(10), 89);
    }

    #[test]
    fn test_max_area() {
        assert_eq!(max_area(&[1, 8, 6, 2, 5, 4, 8, 3, 7]), 49);
        assert_eq!(max_area(&[1, 1]), 1);
        assert_eq!(max_area(&[4, 3, 2, 1, 4]), 16);
    }
}
