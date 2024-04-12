#include <stdio.h>
#include <stdbool.h>

int main() {
    char N[100000];  // 文字列の最大長さに注意
    scanf("%s", N);
    char prev_digit = 'A';

    for (int i = 0; N[i] != '\0'; i++) {
        if (N[i] >= prev_digit) {
            printf("No\n");
            return 0;
        }
        prev_digit = N[i];
    }
    printf("Yes\n");
    return 0;
}