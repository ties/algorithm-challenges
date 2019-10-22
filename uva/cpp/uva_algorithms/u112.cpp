#include <iostream>
#include <sstream>
#include <tuple>
#include <algorithm>
#include <stack>
#include <utility>
#include <limits.h>

using namespace std;

#define MIN(s1, s2) (s1 < s2 ? s1 : s2)
#define MAX(s1, s2) (s1 > s2 ? s1 : s2)

bool canary(stack<int>* st) {
    if (st->top() == INT_MAX)
    {
        st->pop();
        return true;
    }
    return false;
}

void push_canary(stack<int>* st) {
    st->push(INT_MAX);
}

/**
 * Parse s-expressions like
 * 22 (5 (4 (11 (7 () ()) (2 () ()) ) ()) (8 (13 () ()) (4 () (1 () ()) ) ) )
 *
 */
bool sumTree(int goal) {
	stack<int>* expression = new stack<int>;
    
	expression->push(goal);
    
    bool found = false;
    int val = INT_MIN;
    
	while(expression->size() > 0) {
        // skip whitespace and stff
		while (cin.peek() == ' ' || cin.peek() == '\n' || cin.peek() == '\r')
			cin.ignore(1);
        
		switch(cin.peek())
		{
            case '(':
                // open brackets, push canary to make sure the bracket is not empty
                cin.ignore(1);
                
                push_canary(expression);
                break;
            case ')':
                cin.ignore(1);
                
                // if the brackets were empty, evaluate. Else, regular end; pop.
                if(canary(expression))
                {
                    found |= (val != INT_MIN) && expression->top() == 0;
                } else
                {
                    expression->pop();
                }
                
                // when the expression stack is empty, stop.
                if (expression->size() == 1)
                    return found;

                break;
            default:
                cin >> val;
                
                canary(expression);
                
                expression->push(expression->top() - val);
		}
	}
	return found;
}

/**
 * UVA "Tree Summing" problem <http://uva.onlinejudge.org/external/1/112.pdf>
 */
int main(int argc, const char * argv[])
{
	int targetNumber;
    
	while (cin >> targetNumber) {
		// parse the s-expression:
        
		if(sumTree(targetNumber))
		{
			cout << "yes" << endl;
		} else
		{
			cout << "no" << endl;
		}
    }
    
    return 0;
}

