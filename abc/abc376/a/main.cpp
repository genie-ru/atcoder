#include <bits/stdc++.h>
using namespace std;
#define rep(i, n) for (int i = 0; i < (n); ++i)

int main() {
    int n, c;
    cin >> n >> c;
    vector<int> t(n);
    rep(i, n) cin >> t[i];

    int ans = 0, pre = -999999;
    rep(i, n) {
        if (t[i]-pre < c) continue;
        pre = t[i];
        ans++;
    }
    cout << ans << endl;
    return 0;
} 