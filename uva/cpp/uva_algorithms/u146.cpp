//============================================================================
// Name        : u146.cpp
// Author      :
// Version     :
// Copyright   : Your copyright notice
// Description : Hello World in C++, Ansi-style
//============================================================================

#include <iostream>
#include <stdio.h>
#include <math.h>
#include <algorithm>

#include <vector>
#include <list>
#include <map>
#include <set>
#include <string>

#include <queue>
#include <deque>
#include <stack>
#include <bitset>
#include <algorithm>
#include <functional>
#include <numeric>
#include <utility>
#include <sstream>
#include <iostream>
#include <fstream>
#include <iomanip>
#include <cstdio>
#include <cmath>
#include <cstdlib>
#include <ctime>
#include <cstring>
#include <assert.h>

using namespace std;

int main(int argc, const char * argv[])
{
    string line;

    while (cin >> line) {
        if (line.compare("#") == 0) {
            break;
        }

        std::vector<char> items(line.begin(), line.end());


        if(std::next_permutation(items.begin(), items.end())){
        	cout << std::string(items.begin(), items.end()) << endl;
        } else {
        	cout << "No Successor" << endl;
        }
    }

    return 0;
}
