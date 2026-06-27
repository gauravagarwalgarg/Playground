/**
 * Backtracking Templates
 * 
 * Three fundamental backtracking patterns: Subsets, Permutations, Combinations.
 * Each demonstrates the "choose → explore → unchoose" framework.
 * 
 * Compile: g++ -std=c++17 -o backtrack backtrack_template.cpp
 */

#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

// =============================================================================
// TEMPLATE 1: Subsets (Power Set)
// Generate all subsets of the array.
// Decision at each element: include it or skip it.
// Time: O(2^n), Space: O(n) recursion depth
// =============================================================================
class Subsets {
public:
    vector<vector<int>> result;
    
    vector<vector<int>> subsets(vector<int>& nums) {
        vector<int> current;
        backtrack(nums, 0, current);
        return result;
    }
    
private:
    void backtrack(vector<int>& nums, int start, vector<int>& current) {
        // Every state is a valid subset record it
        result.push_back(current);
        
        for (int i = start; i < (int)nums.size(); i++) {
            // CHOOSE: include nums[i]
            current.push_back(nums[i]);
            
            // EXPLORE: recurse with i+1 (no element reuse)
            backtrack(nums, i + 1, current);
            
            // UNCHOOSE: remove nums[i] (backtrack)
            current.pop_back();
        }
    }
};

// =============================================================================
// TEMPLATE 1b: Subsets with Duplicates
// When input has duplicates, sort first and skip consecutive same elements.
// =============================================================================
class SubsetsWithDup {
public:
    vector<vector<int>> result;
    
    vector<vector<int>> subsetsWithDup(vector<int>& nums) {
        sort(nums.begin(), nums.end());  // MUST sort for dedup
        vector<int> current;
        backtrack(nums, 0, current);
        return result;
    }
    
private:
    void backtrack(vector<int>& nums, int start, vector<int>& current) {
        result.push_back(current);
        
        for (int i = start; i < (int)nums.size(); i++) {
            // PRUNE: skip duplicates at the same recursion level
            if (i > start && nums[i] == nums[i-1]) continue;
            
            current.push_back(nums[i]);
            backtrack(nums, i + 1, current);
            current.pop_back();
        }
    }
};

// =============================================================================
// TEMPLATE 2: Permutations
// Generate all orderings of the array.
// Use a "used" array to track which elements are placed.
// Time: O(n!), Space: O(n)
// =============================================================================
class Permutations {
public:
    vector<vector<int>> result;
    
    vector<vector<int>> permute(vector<int>& nums) {
        vector<int> current;
        vector<bool> used(nums.size(), false);
        backtrack(nums, current, used);
        return result;
    }
    
private:
    void backtrack(vector<int>& nums, vector<int>& current, vector<bool>& used) {
        // Complete permutation when all elements are placed
        if (current.size() == nums.size()) {
            result.push_back(current);
            return;
        }
        
        for (int i = 0; i < (int)nums.size(); i++) {
            if (used[i]) continue;  // skip already-used elements
            
            // CHOOSE
            used[i] = true;
            current.push_back(nums[i]);
            
            // EXPLORE
            backtrack(nums, current, used);
            
            // UNCHOOSE
            current.pop_back();
            used[i] = false;
        }
    }
};

// =============================================================================
// TEMPLATE 3: Combinations (Choose k from n)
// Generate all combinations of size k from [1..n].
// Like subsets but only record when size == k.
// Time: O(C(n,k)), Space: O(k)
// =============================================================================
class Combinations {
public:
    vector<vector<int>> result;
    
    vector<vector<int>> combine(int n, int k) {
        vector<int> current;
        backtrack(n, k, 1, current);
        return result;
    }
    
private:
    void backtrack(int n, int k, int start, vector<int>& current) {
        if ((int)current.size() == k) {
            result.push_back(current);
            return;
        }
        
        // PRUNING: need (k - current.size()) more elements
        // so i can go at most to n - (k - current.size()) + 1
        int remaining = k - current.size();
        
        for (int i = start; i <= n - remaining + 1; i++) {
            current.push_back(i);
            backtrack(n, k, i + 1, current);
            current.pop_back();
        }
    }
};

// =============================================================================
// TEMPLATE 4: Combination Sum (Reuse Allowed)
// Find all combinations that sum to target. Elements can be reused.
// Key difference: recurse with i (not i+1) to allow reuse.
// =============================================================================
class CombinationSum {
public:
    vector<vector<int>> result;
    
    vector<vector<int>> combinationSum(vector<int>& candidates, int target) {
        sort(candidates.begin(), candidates.end());  // sort for pruning
        vector<int> current;
        backtrack(candidates, target, 0, current);
        return result;
    }
    
private:
    void backtrack(vector<int>& candidates, int remaining, int start, vector<int>& current) {
        if (remaining == 0) {
            result.push_back(current);
            return;
        }
        
        for (int i = start; i < (int)candidates.size(); i++) {
            if (candidates[i] > remaining) break;  // PRUNE: too large
            
            current.push_back(candidates[i]);
            // Pass i (not i+1) because we CAN reuse the same element
            backtrack(candidates, remaining - candidates[i], i, current);
            current.pop_back();
        }
    }
};

// =============================================================================
void printResult(const string& label, const vector<vector<int>>& result) {
    cout << label << ":" << endl;
    for (const auto& v : result) {
        cout << "  [";
        for (int i = 0; i < (int)v.size(); i++) {
            cout << v[i] << (i < (int)v.size()-1 ? "," : "");
        }
        cout << "]" << endl;
    }
}

int main() {
    // Subsets
    vector<int> nums1 = {1, 2, 3};
    Subsets s;
    printResult("Subsets of [1,2,3]", s.subsets(nums1));
    
    // Subsets with duplicates
    vector<int> nums2 = {1, 2, 2};
    SubsetsWithDup sd;
    printResult("Subsets of [1,2,2] (no dup)", sd.subsetsWithDup(nums2));
    
    // Permutations
    vector<int> nums3 = {1, 2, 3};
    Permutations p;
    printResult("Permutations of [1,2,3]", p.permute(nums3));
    
    // Combinations: choose 2 from 4
    Combinations c;
    printResult("Combinations C(4,2)", c.combine(4, 2));
    
    // Combination Sum: target = 7
    vector<int> candidates = {2, 3, 6, 7};
    CombinationSum cs;
    printResult("Combination Sum [2,3,6,7] target=7", cs.combinationSum(candidates, 7));
    
    return 0;
}
