#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

string longestCommonPrefix(vector<string>& strs) {
    if (strs.empty()) return "";

    cout << "Sorted strings: ";
    for (const string& str : strs) {
        cout << str << " ";
    }
    cout << endl;


    // Sort the strings lexicographically
    sort(strs.begin(), strs.end());

    cout << "Sorted strings: ";
    for (const string& str : strs) {
        cout << str << " ";
    }
    cout << endl;

    string first = strs[0], last = strs.back();
    int i = 0;
    
    // Compare the first and last string character by character
    while (i < first.size() && first[i] == last[i]) {
        i++;
    }

    return first.substr(0, i);
}

// Example usage
int main() {
    vector<string> words = {"flower", "flow", "flight"};
    cout << "Longest Common Prefix: " << longestCommonPrefix(words) << endl;
    return 0;
}

