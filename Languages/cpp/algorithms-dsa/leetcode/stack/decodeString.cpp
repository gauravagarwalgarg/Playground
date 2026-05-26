#include <iostream>
#include <stack>
#include <string>

std::string decodeString(std::string s) {
    std::stack<int> countStack;  
    std::stack<std::string> stringStack;  
    std::string currentString = "";  
    int count = 0;

    for (char ch : s) {
        if (isdigit(ch)) {
            count = count * 10 + (ch - '0');  // Build multi-digit numbers
            std::cout << "Digit encountered: " << ch << ", count = " << count << std::endl;
        } else if (ch == '[') {
            countStack.push(count);  // Push repeat count
            stringStack.push(currentString);  // Store the previous string
            std::cout << "Encountered '[', pushing to stack -> count: " << count << ", storing previous currentString: '" << currentString << "'" << std::endl;
            
            currentString = "";  // Reset for new part inside brackets
            count = 0;  // Reset count
        } else if (ch == ']') {
            int repeatCount = countStack.top();
            countStack.pop();
            std::string decodedString = stringStack.top();  // Get previous stored string
            stringStack.pop();
            
            std::cout << "Encountered ']', repeatCount: " << repeatCount 
                      << ", previous stored string: '" << decodedString 
                      << "', currentString before repeat: '" << currentString << "'" << std::endl;

            // Append repeated `currentString` to `decodedString`
            for (int i = 0; i < repeatCount; i++) {
                decodedString += currentString;
            }

            currentString = decodedString;  // Update `currentString` after decoding
            std::cout << "After decoding, currentString: '" << currentString << "'" << std::endl;
        } else {
            currentString += ch;  // Append normal characters
            std::cout << "Character appended: " << ch << ", currentString: '" << currentString << "'" << std::endl;
        }
    }
    
    return currentString;
}

int main() {
    std::string input = "3[a2[c]]";
    std::cout << "Final decoded string: " << decodeString(input) << std::endl;
    return 0;
}

