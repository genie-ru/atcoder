import 'dart:io';

void main() {
  stdin.readLineSync().toString();
  final str = stdin.readLineSync().toString();
  final l = str.split('');
  final state = List<bool>.generate(3, (index) => false);
  var counter = 0;
  for (final e in l) {
    if (e == 'A') {
      state[0] = true;
    } else if (e == 'B') {
      state[1] = true;
    } else if (e == 'C') {
      state[2] = true;
    }
    counter++;
    if (state.every((element) => element)) {
      print(counter);
      return;
    }
  }
}
