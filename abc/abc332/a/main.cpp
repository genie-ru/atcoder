#include <iostream>
using namespace std;

int main(){
    int N, S, K;
    cin >> N >> S >> K;
    int res = 0;
    for (int i = 0; i < N; i++) {
        int p, q;
        cin >> p >> q;
        res = res + p * q;
    }
    if (res >= S) {
        cout << res << endl;
    } else {
        cout << res + K << endl;
    }
    return 0;
}