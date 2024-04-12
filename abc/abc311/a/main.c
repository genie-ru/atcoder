#include<stdio.h>
int main(void){
  int n;scanf("%d",&n);
  char s[105];scanf("%s",s);
  int cou=1;
  for(int i=0;i<n;i++){
   if(s[i]=='A'){
     if(cou%2){cou*=2;};
   }
    else if(s[i]=='B'){
      if(cou%3){cou*=3;};
    }
    else if(s[i]=='C'){
      if(cou%5){cou*=5;};
    };
    if(cou%30==0){printf("%d",i+1);break;};
  }
  return 0;
}