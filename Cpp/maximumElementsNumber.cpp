// Find the maximum elements in an given array that would sum upto a given array

#include <bits/stdc++.h>
#include <unordered_map>
#include <vector>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);



/*
 * Complete the 'find_max_elements' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY array as parameter.
 */
int find_max_elements(vector<int> array) {
    if(array.size() < 2) return 0;
    
    int target = array[0];
    //cout << target << endl;
    
    vector<int> new_array(array.begin() + 1, array.end());
    sort(new_array.begin(), new_array.end());
    
    // int sum = 0, count = 0;

    // for(int num : new_array) {
    //     if(sum + num <= target) {
    //         sum += num;
    //         count++;
    //     } else {
    //         break;
    //     }
        
    // }
    
    // return count;
    
    int n = new_array.size();
    vector<int> dp(target + 1, -1);  
    dp[0] = 0;  // Base case: sum of 0 is achieved with 0 elements

    for (int num : new_array) {
        if(num > target) continue;
        
        for (int j = target; j >= num; j--) {  // Iterate in reverse to avoid overwriting
            if (dp[j - num] != -1) {
                dp[j] = max(dp[j], dp[j - num] + 1);
            }
        }
    }

    return (dp[target] == -1) ? 0 : dp[target];
}

    //int n = new_array.size();
    
    //sort(new_array.begin(), new_array.end());
    
    // vector<int> dp(target + 1, -1);  
    // dp[0] = 0;  // Base case: sum of 0 is achieved with 0 elements

    // for (int num : new_array) {
    //     for (int j = target; j >= num; j--) {  // Iterate in reverse
    //         if (dp[j - num] != -1) {
    //             dp[j] = max(dp[j], dp[j - num] + 1);
    //         }
    //     }
    // }

    //return (dp[target] == -1) ? 0 : dp[target];

int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string array_count_temp;
    getline(cin, array_count_temp);

    int array_count = stoi(ltrim(rtrim(array_count_temp)));

    vector<int> array(array_count);

    for (int i = 0; i < array_count; i++) {
        string array_item_temp;
        getline(cin, array_item_temp);

        int array_item = stoi(ltrim(rtrim(array_item_temp)));

        array[i] = array_item;
    }

    int result = find_max_elements(array);

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

