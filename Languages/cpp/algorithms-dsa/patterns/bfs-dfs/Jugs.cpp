#include <bits/stdc++.h>
#include <exception>

using namespace std;

string ltrim(const string &);
string rtrim(const string &);



#include <queue>
/*
 * Complete the 'balance_water' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY elements as parameter.
 */

struct WaterState{
    vector<int> water;
    int steps;
};

int balance_water(vector<int> elements) {
    int n = 4; 
    vector<int> capacities(elements.begin(), elements.begin() + n);
    vector<int> initial(elements.begin() + n, elements.begin() + 2 * n);
    vector<int> target(elements.begin() + n * 2, elements.begin() + 3 * n);
    
    queue<WaterState> q;
    set<vector<int>> visited;
    
    q.push({initial, 0});
    
    visited.insert(initial);
    
    while(!q.empty()) {
        WaterState current = q.front();
        q.pop();
        
        if(current.water == target) {
            return current.steps;
        }
        
        for(int i = 0; i < n; i++) {
            for(int j = 0; j < n; j++) {
                if(i != j) {
                    vector<int> nextWater = current.water;
                    int transfer = min(nextWater[i], capacities[j] - nextWater[j]);
                    
                    nextWater[i] -= transfer;
                    nextWater[j] += transfer;
                    
                    if(visited.find(nextWater) == visited.end()) {
                        visited.insert(nextWater);
                        q.push({nextWater, current.steps + 1});
                    }
                }
            }
        }
    }
    
    return -1;
}   
int main()
{
    ofstream fout(getenv("OUTPUT_PATH"));

    string elements_count_temp;
    getline(cin, elements_count_temp);

    int elements_count = stoi(ltrim(rtrim(elements_count_temp)));

    vector<int> elements(elements_count);

    for (int i = 0; i < elements_count; i++) {
        string elements_item_temp;
        getline(cin, elements_item_temp);

        int elements_item = stoi(ltrim(rtrim(elements_item_temp)));

        elements[i] = elements_item;
    }

    int result = balance_water(elements);

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

