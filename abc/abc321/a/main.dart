import 'dart:io';

void main() {
  String N = stdin.readLineSync()!;
  bool is321Like = true;
  
  for (int i = 0; i < N.length - 1; i++) {
    if (N[i] <= N[i + 1]) {
      is321Like = false;
      break;
    }
  }
  
  if (is321Like) {
    print("Yes");
  } else {
    print("No");
  }
}
