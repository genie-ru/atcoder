import 'dart:io';

void main() {
  final inputs = stdin.readLineSync()?.split(" ")?.map(int.parse);
  final a = inputs?.first;
  final b = inputs?.last;

  if (a?.remainder(3) == 0) {
    print('No');
  } else if (b != null && a != null && b - a == 1) {
    print('Yes');
  } else {
    print('No');
  }
}
