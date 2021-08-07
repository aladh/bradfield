#include <stdio.h>

int main() {
    float celsius, fahr;
    int lower, upper, step;

    lower = 0;
    upper = 150;
    step = 10;

    celsius = lower;

    printf("%7s%13s\n", "Celsius", "Fahrenheit");

    while (celsius <= upper) {
        fahr = (celsius / (5.0 / 9.0)) + 32;
        printf("%3.0f%16.1f\n", celsius, fahr);
        celsius = celsius + step;
    }
}