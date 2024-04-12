import java.util.Scanner;
class Main{
  public static void main(String[] args){
    Scanner sc = new Scanner(System.in);
    int N = sc.nextInt();
    String S = sc.next();
    int ans = S.indexOf("A");
    ans = Math.max(ans,S.indexOf("B"));
    ans = Math.max(ans,S.indexOf("C"));
    System.out.println(ans+1);
  }
}
