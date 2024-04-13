#include <iostream>
#include <string>
using namespace std;

bool isAirportCode(const string& S, const string& T) {
    int n = S.length();
    string uppercase_S = S;  
    for (char& c : uppercase_S) {
        c = toupper(c);
    }

    for (int i = 0; i < n; ++i) {
        if (uppercase_S[i] == T[0]) {  
            for (int j = i + 1; j < n; ++j) {
                if (uppercase_S[j] == T[1]) {  
                    for (int k = j + 1; k < n; ++k) {
                        if (uppercase_S[k] == T[2]) {  
                            return true;
                        }
                    }
                }
            }
        }
    }


    if (T[2] == 'X') {
        for (int i = 0; i < n; ++i) {
            if (uppercase_S[i] == T[0]) {  
                for (int j = i + 1; j < n; ++j) {
                    if (uppercase_S[j] == T[1]) {  
                        return true;
                    }
                }
            }
        }
    }

    return false;
}

int main() {
    string S, T;
    cin >> S >> T;

    if (isAirportCode(S, T)) {
        cout << "Yes\n";
    } else {
        cout << "No\n";
    }

    return 0;
}
