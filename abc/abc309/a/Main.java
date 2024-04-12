interface Main{
  static void main(String[]g){
    var sc=new java.util.Scanner(System.in);
    int a=sc.nextInt();
    System.out.println(sc.nextInt()-a<2&&a%3>0?"Yes":"No");
  }
}