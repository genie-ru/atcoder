#include <iostream>
#include <string>

int main() {
    int n;
    std::cin >> n;

    std::string pi3 = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679";
    std::cout << pi3.substr(0, n + 2) << std::endl;

    return 0;
}
