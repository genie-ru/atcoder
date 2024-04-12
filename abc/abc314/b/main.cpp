#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main(void)
{
    int n, x, c[101];
    vector<int> a[101];

    cin >> n;
    for(int i = 1; i <= n; i++) {
        cin >> c[i];
        a[i].resize(c[i]);
        for(int j = 0; j < c [i]; j++) cin >> a[i][j];
    }
    cin >> x;

    vector<int> vec;
    for(int)
}