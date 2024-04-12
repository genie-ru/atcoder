#include <stdio.h>

int main(void) {
    char a, b, str[2][4] = { "No", "Yes" };
    scanf("%hhd %hhd", &a, &b);
    printf("%s", str[a == b - 1 && b % 3 != 1]);
}