#include <stdio.h>
#include <limits.h>

int main() {
    int n, p, q;
    scanf("%d %d %d", &n, &p, &q);

    int d_min = INT_MAX, d;
    for (int i = 1; i <= n; i++) {
        scanf("%d", &d);
        if (d < d_min) {
            d_min = d;
        }
    }
    printf("%d\n", (p < (q + d_min)) ? p : (q + d_min));

    return 0;
}
