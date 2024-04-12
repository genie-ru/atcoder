#include <iostream>
#include <climits>  // INT_MAXを使用するために必要
using namespace std;

int main(void)
{
    int n, p, q;
    cin >> n >> p >> q;

    int d_min = INT_MAX, d;
    for(int i = 1; i <= n; i++){
        cin >> d;
        d_min = min(d, d_min);
    }
    cout << min(p, q+d_min) << endl;
    
    return 0;
}
