/*
 * LeetCode 20 - Valid Parentheses
 * Topic: Stack
 * Difficulty: Easy
 *
 * Push opening brackets onto stack, pop and match for closing brackets.
 * Time: O(n)
 * Space: O(n)
 */
#include <iostream>
#include <string>
#include <stack>
#include <cassert>
using namespace std;

bool isValid(string s) {
    stack<char> st;
    for (char c : s) {
        if (c == '(' || c == '{' || c == '[') st.push(c);
        else {
            if (st.empty()) return false;
            char top = st.top(); st.pop();
            if ((c == ')' && top != '(') ||
                (c == '}' && top != '{') ||
                (c == ']' && top != '[')) return false;
        }
    }
    return st.empty();
}

int main() {
    assert(isValid("()") == true);
    assert(isValid("()[]{}") == true);
    assert(isValid("(]") == false);
    assert(isValid("([)]") == false);
    assert(isValid("{[]}") == true);

    cout << "All tests passed!" << endl;
    return 0;
}
