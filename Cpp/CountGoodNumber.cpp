#include <iostream>
#include <unordered_map>
using namespace std;

bool is_good_number(int n) {
    unordered_map<char, char> rotation_map = {
        {'0', '0'}, {'1', '1'}, {'8', '8'}, {'2', '5'}, {'5', '2'}, {'6', '9'}, {'9', '6'}
    };
    string invalid_digits = "347";
    
    string str_n = to_string(n);
    string rotated = "";
    bool is_different = false;
    
    for (char digit : str_n) {
	    cout << digit << endl;
        if (invalid_digits.find(str_n) != string::npos) {
            return false;
        }
        //rotated += rotation_map[digit];
        if (rotation_map[digit] != digit) {
            is_different = true;
        }
    }
    
    return is_different;
}

int count_good_numbers(int K) {
    int count = 0;
    for (int i = 1; i <= K; ++i) {
        if (is_good_number(i)) {
            count++;
        }
    }
    return count;
}

int main() {
    int K;
    cin >> K;
    cout << count_good_numbers(K) << endl;
    return 0;
}

