#include <iostream>
#include <vector>
#include <string>
using namespace std;

string multiply(string num1, string num2) {
    if (num1 == "0" || num2 == "0") return "0"; // Edge case

    int n = num1.size(), m = num2.size();
    vector<int> result(n + m, 0); // Store intermediate results

    // Multiply each digit of num1 with num2
    for (int i = n - 1; i >= 0; i--) {
        for (int j = m - 1; j >= 0; j--) {
            int mul = (num1[i] - '0') * (num2[j] - '0'); // Convert chars to int and multiply
            int sum = mul + result[i + j + 1]; // Add to existing result

            result[i + j + 1] = sum % 10; // Store unit place
            result[i + j] += sum / 10;    // Carry to next position
        }
    }

    // Convert result array to string
    string product = "";
    for (int num : result) {
        if (!(product.empty() && num == 0)) { // Skip leading zeros
            product += (num + '0'); // Convert int to char
        }
    }

    return product.empty() ? "0" : product;
}

int main() {
    string num1 = "123", num2 = "456";
    cout << multiply(num1, num2) << endl; // Output: "56088"
    return 0;
}

