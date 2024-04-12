#include <iostream>
#include <algorithm> // min関数のために必要
using namespace std;

int main() {
    int K, G, M;
    cin >> K >> G >> M;
    
    int x = 0, y = 0;

    for (int i = 0; i < K; i++) {
        if (x == G) {
            x = 0;
        } 
        else if (y == 0) {
            y = M;
        }
        else {
            int z = min(y, G - x);
            x += z;
            y -= z;
        }
    }

    cout << x << " " << y << endl;
    return 0;
}
