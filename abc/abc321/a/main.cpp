#include <iostream>
#include <string>
using namespace std;

int main()
{
    string N;
    cin >> N;
    char prev_digit = 'A';

    for (char digit : N) {
        if (digit >= prev_digit) {
            cout << "No" << endl;
            return 0;
        }
        prev_digit = digit;
    }
    cout << "Yes" << endl;
    return 0;
}