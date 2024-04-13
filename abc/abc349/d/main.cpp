#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

typedef long long ll;

int main() {
    ll L, R;
    cin >> L >> R;

    vector<pair<ll, ll>> segments;
    
    while (L < R) {
        ll power = 1;
        while (power * 2 <= R - L) {
            power *= 2;
        }
        
        ll next_start = (L / power + 1) * power;
        if (next_start > R) next_start = R;

        segments.push_back({L, next_start});
        L = next_start;
    }

    cout << segments.size() << endl;
    for (auto& segment : segments) {
        cout << segment.first << " " << segment.second << endl;
    }

    return 0;
}
