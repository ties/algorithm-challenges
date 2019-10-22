#include <iostream>
#include <sstream>
#include <trie>
#include <algorithm>

using namespace std;

#define MIN(s1, s2) (s1 < s2 ? s1 : s2)
#define MAX(s1, s2) (s1 > s2 ? s1 : s2)

/**
 * UVA "List of Conquests" problem <http://uva.onlinejudge.org/external/104/10420.pdf>
 */
int main(int argc, const char * argv[])
{
	int num_conquests;
	string country;

	cin >> num_conquests;

	std::map<string, int> conquests;

	std::cin.ignore(); // eat 1 char (= \n)
	while (num_conquests > 0 && cin >> country) {
		// ignore until eol
		std::cin.ignore(256, '\n');

    	if(conquests.find(country) == conquests.end()) {
    		conquests[country] = 1;
    	} else {
    		conquests[country] += 1;
    	}

    	num_conquests--;
    }

	// map *should* be sorted in key order - internal guarantee
	for_each(conquests.begin(), conquests.end(), [] (std::map<string, int>::value_type ent) {
		cout << ent.first << " " << ent.second << endl;
	});

    return 0;
}
