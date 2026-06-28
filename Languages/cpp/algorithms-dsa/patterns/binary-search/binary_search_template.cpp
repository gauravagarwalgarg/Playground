/**
 * Binary Search Templates
 * 
 * Three core templates covering 99% of binary search problems.
 * Key decision: what does the loop invariant guarantee?
 * 
 * Compile: g++ -std=c++17 -o binary_search binary_search_template.cpp
 */

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

// =============================================================================
// TEMPLATE 1: Standard Binary Search
// Find exact target in sorted array. Returns index or -1.
// Loop invariant: target (if exists) is in [lo, hi]
// Use when: You need the exact position of a known value
// =============================================================================
int standard_search(const vector<int>& arr, int target) {
    int lo = 0, hi = static_cast<int>(arr.size()) - 1;
    
    while (lo <= hi) {  // NOTE: <= because both lo and hi are valid candidates
        int mid = lo + (hi - lo) / 2;  // avoid overflow vs (lo+hi)/2
        
        if (arr[mid] == target) {
            return mid;
        } else if (arr[mid] < target) {
            lo = mid + 1;  // mid is too small, exclude it
        } else {
            hi = mid - 1;  // mid is too large, exclude it
        }
    }
    return -1;  // not found
}

// =============================================================================
// TEMPLATE 2: Lower Bound (First position where arr[pos] >= target)
// Equivalent to std::lower_bound
// Loop invariant: answer is in [lo, hi)  hi is always a valid answer or n
// Use when: First occurrence, insertion point, "minimum X satisfying condition"
// =============================================================================
int lower_bound_search(const vector<int>& arr, int target) {
    int lo = 0, hi = static_cast<int>(arr.size());  // NOTE: hi = n (past-the-end)
    
    while (lo < hi) {  // NOTE: < not <=, terminates when lo == hi
        int mid = lo + (hi - lo) / 2;
        
        if (arr[mid] < target) {
            lo = mid + 1;  // mid is too small, first valid is after mid
        } else {
            hi = mid;      // mid might be the answer, don't skip it
        }
    }
    // lo == hi == first index where arr[index] >= target
    // If lo == arr.size(), target is larger than all elements
    return lo;
}

// =============================================================================
// TEMPLATE 3: Upper Bound (First position where arr[pos] > target)
// Equivalent to std::upper_bound
// Loop invariant: answer is in [lo, hi)
// Use when: Last occurrence (upper_bound - 1), count of target values
// =============================================================================
int upper_bound_search(const vector<int>& arr, int target) {
    int lo = 0, hi = static_cast<int>(arr.size());
    
    while (lo < hi) {
        int mid = lo + (hi - lo) / 2;
        
        if (arr[mid] <= target) {  // NOTE: <= instead of < (only difference from lower_bound)
            lo = mid + 1;  // mid is <= target, first strictly-greater is after mid
        } else {
            hi = mid;      // mid is > target, it might be the answer
        }
    }
    // lo == hi == first index where arr[index] > target
    return lo;
}

// =============================================================================
// TEMPLATE 4: Search on Answer (Binary Search the Result Space)
// Use when: "minimize the maximum", "can we achieve X?", capacity problems
// Example: Minimum capacity to ship packages within D days
// =============================================================================
bool canShipInDays(const vector<int>& weights, int capacity, int days) {
    int current_load = 0, days_needed = 1;
    for (int w : weights) {
        if (current_load + w > capacity) {
            days_needed++;
            current_load = 0;
        }
        current_load += w;
    }
    return days_needed <= days;
}

int shipWithinDays(const vector<int>& weights, int days) {
    // Search space: [max_single_weight, sum_of_all_weights]
    int lo = *max_element(weights.begin(), weights.end());
    int hi = 0;
    for (int w : weights) hi += w;
    
    while (lo < hi) {
        int mid = lo + (hi - lo) / 2;
        
        if (canShipInDays(weights, mid, days)) {
            hi = mid;      // mid capacity works, try smaller
        } else {
            lo = mid + 1;  // mid capacity too small, need more
        }
    }
    return lo;  // minimum capacity that works
}

// =============================================================================
int main() {
    vector<int> arr = {1, 2, 3, 4, 4, 4, 5, 6, 7, 8};
    
    // Standard search
    cout << "Standard search for 4: index " << standard_search(arr, 4) << endl;
    
    // Lower bound: first occurrence of 4
    cout << "Lower bound of 4: index " << lower_bound_search(arr, 4) << endl;
    
    // Upper bound: first element > 4
    cout << "Upper bound of 4: index " << upper_bound_search(arr, 4) << endl;
    
    // Count of 4s = upper_bound - lower_bound
    cout << "Count of 4: " << upper_bound_search(arr, 4) - lower_bound_search(arr, 4) << endl;
    
    // Search on answer: ship packages
    vector<int> weights = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    cout << "Min capacity for 5 days: " << shipWithinDays(weights, 5) << endl;
    
    return 0;
}
