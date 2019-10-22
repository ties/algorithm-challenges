#include <iostream>

using namespace std;

#define MAX(s1, s2) (s1 > s2 ? s1 : s2)
#define MIN(s1, s2) (s1 < s2 ? s1 : s2)

/**
 * UVA "hashmat the brave warrior" <http://uva.onlinejudge.org/external/100/10055.pdf>
 */
int main(int argc, const char * argv[])
{
    int s1, s2;

    while (cin >> s1 >> s2) {
    	cout << max(s1, s2) - min(s1, s2) << endl;
    }
    return 0;
}
