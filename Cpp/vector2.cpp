#include <iostream>
#include <vector>

using namespace std;

int main()
{
  vector<vector<int>> pascal;
  pascal.push_back(vector<int>(1, 1));

  for(auto it: pascal)
  {
      std::cout << it << std::endl;
    for(auto i: it)
    {
      std::cout << i << std::endl;
    }
  }
}
