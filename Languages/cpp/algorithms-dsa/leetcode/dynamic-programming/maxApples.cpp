/* A grocery store is selling apples by count. 
They want to set the same price for each apple that they sell. 
They received total N apples of different weights from the farm.

To be able to set the same price for each apple, they need to select only those set of apples, where the weight difference between the largest and the smallest is not more than K
Given the range K, what is the maximum number of apples they can select from the N apples.

Input format:

Locked stub code in the editor reads the following input from stdin and passes to the function:
int find_max_apples(int max_size_difference, vector<int> apple_sizes) {

The first line contains the range K
The second line contains the number of apples N.
the sequence lines contains the size of apples, per one line

K, number of apple (N) , size of apple can be 1 to 10000

Output:
The function must return an integer specifying the maximum number of apples that can be selected from the input set of N apples  which are all within the wright range of maximum K among each other.
*/

#include <bits/stdc++.h>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);


/*
 * Complete the 'find_max_apples' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER max_size_difference
 *  2. INTEGER_ARRAY apple_sizes
 */

int find_max_apples(int max_size_difference, vector<int> apple_sizes) {
    sort(apple_sizes.begin(), apple_sizes.end());
    
    int max_apples = 0;
    int left = 0;
    
    for(int right = 0; right < apple_sizes.size(); right++) {
        while(apple_sizes[right] - apple_sizes[left] > max_size_difference) {
            left++;
        }
        max_apples = max(max_apples, right - left + 1);
    }
    
    return max_apples;
}
int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string max_size_difference_temp;
    getline(cin, max_size_difference_temp);

    int max_size_difference = stoi(ltrim(rtrim(max_size_difference_temp)));

    string apple_sizes_count_temp;
    getline(cin, apple_sizes_count_temp);

    int apple_sizes_count = stoi(ltrim(rtrim(apple_sizes_count_temp)));

    vector<int> apple_sizes(apple_sizes_count);

    for (int i = 0; i < apple_sizes_count; i++) {
        string apple_sizes_item_temp;
        getline(cin, apple_sizes_item_temp);

        int apple_sizes_item = stoi(ltrim(rtrim(apple_sizes_item_temp)));

        apple_sizes[i] = apple_sizes_item;
    }

    int result = find_max_apples(max_size_difference, apple_sizes);

    fout << result << "\n";

    fout.close();

    return 0;
}

string ltrim(const string &str) {
    string s(str);

    s.erase(
        s.begin(),
        find_if(s.begin(), s.end(), not1(ptr_fun<int, int>(isspace)))
    );

    return s;
}

string rtrim(const string &str) {
    string s(str);

    s.erase(
        find_if(s.rbegin(), s.rend(), not1(ptr_fun<int, int>(isspace))).base(),
        s.end()
    );

    return s;
}

